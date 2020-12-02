package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"./com_cliente"
	"./com_datanode"
	"./com_namenode"
	"./estructuras"
	"google.golang.org/grpc"
)

func obtenerVecinos(nombre string) *[2]string {
	buf, err := os.Open(nombre)
	if err != nil {
		log.Fatalf("Se ha producido un error al buscar vecinos: %s", err.Error())
	}
	defer buf.Close()
	arr := new([2]string)
	lector := bufio.NewScanner(buf)
	var i int
	for lector.Scan() {
		arr[i] = lector.Text()
		i++
	}
	return arr
}

func initServer(listener net.Listener, canal chan *grpc.Server) {
	servidor := com_datanode.ServerDatanode{}
	grpcServer := grpc.NewServer()
	com_datanode.RegisterInteraccionesServer(grpcServer, &servidor)
	canal <- grpcServer
	if err := grpcServer.Serve(listener); err != nil {
		log.Printf("Terminada ejecución. Motivo: %s", err.Error())
	}
}

func main() {
	var wait sync.WaitGroup

	//iniciar servidor datanodes
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Se ha producido un error al iniciar el servidor: %s", err.Error())
	}

	listener_clientes, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Se ha producido un error al iniciar el servidor: %s", err.Error())
	}
	go func(listener net.Listener) {
		servidor := com_cliente.ServerCliente{}
		grpcServer := grpc.NewServer()
		com_cliente.RegisterInteraccionesServer(grpcServer, &servidor)
		if err := grpcServer.Serve(listener); err != nil {
			log.Printf("Terminada ejecución. Motivo: %s", err.Error())
		}
	}(listener_clientes)

	// dejar el servidor grpc en otro hilo
	canalServer := make(chan *grpc.Server)
	wait.Add(1)
	go func(canal chan *grpc.Server) {
		initServer(listener, canal)
		wait.Done()
	}(canalServer)

	_ = <-canalServer

	time.Sleep(time.Second) //pausa para darle tiempo a los resagados

	yo := obtenerVecinos("yo.txt")[0]           //nombre de este nodo
	vecinos := obtenerVecinos("vecinos.txt")    //nombres de los vecinos (dist45, dist46, etc.)
	conexiones := make([](*grpc.ClientConn), 2) //las conexioenes con cada vecino en el mismo orden del arreglo vecinos
	activos := make([]bool, 2)                  //para llevar cuenta de lso vecinos activos

	//establecer conexion con vecinos activos
	canalVecinos := new(bool)
	*canalVecinos = false
	clientes := make([](com_datanode.InteraccionesClient), 2) //clientes grpc
	go func(conexiones *[](*grpc.ClientConn), activos *[]bool, vecinos *[2]string, canal *bool, clientes *[](com_datanode.InteraccionesClient)) {
		var i int
		var ctx context.Context
		for {
			ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
			if !(*activos)[i] {
				(*conexiones)[i], err = grpc.DialContext(
					ctx,
					vecinos[i]+":9000", grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Printf("Nodo se pudo establecer conexión con %s\n", vecinos[i])
					(*activos)[i] = false
				} else {
					(*activos)[i] = true
					(*clientes)[i] = com_datanode.NewInteraccionesClient((*conexiones)[i])
					log.Printf("Conectado con %s", vecinos[i])
					defer (*conexiones)[i].Close()
				}
			}
			if *canal {
				break
			}
			i = (i + 1) % 2
			time.Sleep(time.Millisecond)
		}
	}(&conexiones, &activos, vecinos, canalVecinos, &clientes)

	for i := 0; i < 2; i++ {
		clientes[i] = com_datanode.NewInteraccionesClient(conexiones[i])

	}

	var conexionNamenode *grpc.ClientConn

	conexionNamenode, err = grpc.Dial("dist45:9000", grpc.WithInsecure())
	if err != nil {
		log.Printf("no se pudo conectar con namenode... Reintentando...")
	} else {
		log.Printf("Conectado con namenode")
	}
	clienteNamenode := com_namenode.NewInteraccionesClient(conexionNamenode)

	go func(vecinos *[2]string, clientes *[]com_datanode.InteraccionesClient, clienteNamenode *com_namenode.InteraccionesClient, yo string) {
		for {
			if len(estructuras.ColaParaEnvios) > 0 {
				elem := estructuras.Pop(&estructuras.ColaParaEnvios)
				libro := *estructuras.AlmacenLibros[elem]
				total := libro.ChunksTotales
				titulo := libro.Titulo
				paraMandar := make([]uint64, 0)
				for i := 0; i < 2; i++ {
					ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
					respuesta, err := (*clientes)[i].Disponible(ctx, &com_datanode.Empty{})
					if err != nil {
						log.Printf("Nodo %s no está disponible.", vecinos[i])
					} else if !respuesta.Estado {
						log.Printf("Nodo %s no está disponible.", vecinos[i])
					} else {
						paraMandar = append(paraMandar, uint64(i))
					}
				}
				paraMandar = append(paraMandar, uint64(2))
				var contador uint64
				var k uint64
				var buf []byte
				var data *os.File
				var n int
				var titulo2 string
				paraLogear := titulo + " " + strconv.FormatUint(total, 10) + "\n"
				//fmt.Println(paraMandar)
				for contador < total {
					titulo2 = titulo + "_" + strconv.FormatUint(contador+1, 10)
					if paraMandar[k] == 2 {
						paraLogear += titulo2 + " " + yo + "\n"
						contador++
						k = (k + 1) % uint64(len(paraMandar))
						continue
					}
					data, err = os.Open(titulo2)
					if err != nil {
						log.Fatalf("Error al leer archivo: %s", err.Error())
					}
					buf = make([]byte, 250*1000)
					n, _ = data.Read(buf)
					respuesta, err := (*clientes)[paraMandar[k]].SubirArchivo(context.Background(), &com_datanode.Chunk{
						Data:   buf[:n],
						Nombre: titulo2,
					})
					data.Close()
					if err != nil {
						log.Printf("%s FRACASADO por nodo %s", vecinos[paraMandar[k]])
					} else if !respuesta.Estado {
						log.Printf("%s FRACASADO por nodo %s. Motivo: %s", vecinos[paraMandar[k]], respuesta.Msg)
					} else {
						log.Printf("%s RECIBIDO con exito por nodo %s", titulo2, vecinos[paraMandar[k]])
						err = os.Remove(titulo2)
						if err != nil {
							log.Printf("No se pudo eliminar %s", titulo2)
						}
					}
					paraLogear += titulo2 + " " + vecinos[paraMandar[k]] + "\n"
					contador++
					k = (k + 1) % uint64(len(paraMandar))
				}
				res, err := (*clienteNamenode).EscribirLog(context.Background(), &com_namenode.Log{
					Titulo: titulo,
					Texto:  paraLogear,
				})
				if err != nil {
					log.Printf("No se pudo escribir en log")
				} else if !res.Estado {
					log.Printf("No se pudo escribir en log debido a %s", res.Msg)
				} else {
					log.Printf("Escrito en log con éxito")
				}

			}
		}
	}(vecinos, &clientes, &clienteNamenode, yo)
	//
	// data, err := os.Open(yo + ".txt")
	// if err != nil {
	// 	log.Fatalf("Error al leer archivo: %s", err.Error())
	// }
	// defer data.Close()
	// buf := make([]byte, 250*1000)
	// //generar clientes grpc
	// var contador, i int
	// for {
	// 	var enviado [2]bool
	// 	if activos[i] && !enviado[i] {
	// 		clientes[i] = com_datanode.NewInteraccionesClient(conexiones[i])
	// 		n, err := data.Read(buf)
	// 		respuesta, err := clientes[i].SubirArchivo(context.Background(), &com_datanode.Chunk{
	// 			Nombre: yo + ".txt",
	// 			Data:   buf[:n],
	// 		})
	// 		if err != nil {
	// 			log.Panicf("Ha ocurrido un error: %s", err.Error())
	// 		}
	// 		log.Printf("Estado de envío: %s", respuesta.Estado)
	// 		contador++
	// 	}
	// 	if contador == 2 {
	// 		*canalVecinos = true
	// 		ser.GracefulStop()
	// 		break
	// 	}
	// 	i = (i + 1) % 2
	// }

	wait.Wait()
}
