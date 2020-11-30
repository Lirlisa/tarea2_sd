package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"./com_datanode"
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

func initServer(listener net.Listener) {
	servidor := com_datanode.ServerDatanode{}
	grpcServer := grpc.NewServer()
	com_datanode.RegisterInteraccionesServer(grpcServer, &servidor)
	if err := grpcServer.Serve(listener); err != nil {
		log.Printf("Terminada ejecución. Motivo: %s", err.Error())
	}
}

func main() {
	var wait sync.WaitGroup

	//iniciar servidor
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Se ha producido un error al iniciar el servidor: %s", err.Error())
	}

	// dejar el servidor grpc en otro hilo
	wait.Add(1)
	go func() {
		initServer(listener)
		wait.Done()
	}()

	time.Sleep(time.Second) //pausa para darle tiempo a los resagados

	yo := obtenerVecinos("yo.txt")[0]           //nombre de este nodo
	vecinos := obtenerVecinos("vecinos.txt")    //nombres de los vecinos (dist45, dist46, etc.)
	conexiones := make([](*grpc.ClientConn), 2) //las conexioenes con cada vecino en el mismo orden del arreglo vecinos
	activos := make([]bool, 2)                  //para llevar cuenta de lso vecinos activos

	//establecer conexion con vecinos activos
	go func(conexiones [](*grpc.ClientConn), activos []bool, vecinos *[2]string) {
		var i int
		var ctx context.Context
		var contador int
		for {
			ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
			if !activos[i] {
				conexiones[i], err = grpc.DialContext(
					ctx,
					vecinos[i]+":9000", grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Println("Nodo se pudo establecer conexión con %s", vecinos[i])
					activos[i] = false
				} else {
					activos[i] = true
					contador++
					defer conexiones[i].Close()
				}
			}
			if contador == 2 {
				break
			}
			i = (i + 1) % 2
			time.Sleep(time.Millisecond)
		}
	}(conexiones, activos, vecinos)

	var clientes [2](com_datanode.InteraccionesClient) //clientes grpc
	data, err := os.Open(yo + ".txt")
	if err != nil {
		log.Fatalf("Error al leer archivo: %s", err.Error())
	}
	defer data.Close()
	buf := make([]byte, 250*1000)
	//generar clientes grpc
	var contador, i int
	for {
		if activos[i] {
			clientes[i] = com_datanode.NewInteraccionesClient(conexiones[i])
			n, err := data.Read(buf)
			respuesta, err := clientes[i].SubirArchivo(context.Background(), &com_datanode.Chunk{
				Nombre: yo + ".txt",
				Data:   buf[:n],
			})
			if err != nil {
				log.Panicf("Ha ocurrido un error: %s", err.Error())
			}
			log.Printf("Estado de envío: %s", respuesta.Estado)
			contador++
		}
		if contador == 2 {
			break
		}
		i = (i + 1) % 2
	}

	wait.Wait()
}
