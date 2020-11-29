package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"./com_namenode"
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

func main() {
	//iniciar servidor
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Se ha producido un error al iniciar el servidor: %s", err.Error())
	}
	time.Sleep(time.Second) //pausa para darle tiempo a los resagados

	yo = obtenerVecinos("yo.txt")[0] //nombre de este nodo
	vecinos := obtenerVecinos("vecinos.txt") //nombres de los vecinos (dist45, dist46, etc.)
	var conexiones [2](*grpc.ClientConn) //las conexioenes con cada vecino en el mismo orden del arreglo vecinos
	activos := [2](bool) //para llevar cuenta de lso vecinos activos

	//establecer conexion con vecinos activos
	for i := range vecinos {
		conexiones[i], err = grpc.Dial(vecinos[i]+":9000", grpc.WithInsecure())
		if err != nil {
			log.Println("Nodo se pudo establecer conexión con %s", vecinos[i])
			activos[i] = false
		} else {
			activos[i] = true
			defer conexiones[i].Close()
		}
	}

	clientes := [2](com_datanode.InteraccionesClient) //clientes grpc
	data, err := ioutil.ReadFile(yo+".txt")
	if err != nil {
		log.Fatalf("Error al leer archivo: %s", err.Error())
	}
	//generar clientes grpc
	for i := range vecinos {
		if activo[i] {
			clientes[i] = com_namenode.NewInteraccionesClient(conexiones[i])
			respuesta, err := cliente[i].SubirArchivo(context.Background(), &com_datanode.Chunk{
				Nombre: yo+".txt",
				Data: data,
			})
			log.Printf("Estado de envío: %s", respuesta.Estado)
		}
	}
}
