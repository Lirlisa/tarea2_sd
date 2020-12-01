package main

import (
	"fmt"
	"log"
	"net"

	"sync"

	"./ClienteName"
	"./com_namenode"
	"google.golang.org/grpc"
)

func main() {

	/*_, err := os.Create("./LOG.txt")
	if err != nil {
		log.Fatal(err)
	}*/

	//variable de sincronización de go routines
	var wait sync.WaitGroup

	//se abren los puertos
	listenerData, err1 := net.Listen("tcp", ":9000")
	listenerCliente, err2 := net.Listen("tcp", ":9001")
	if err1 != nil {
		log.Fatalf("Se ha producido un error: %s", err1)
	}
	if err2 != nil {
		log.Fatalf("Se ha producido un error: %s", err2)
	}

	fmt.Println("Iniciado servidor en escucha en puerto 9000 y 9001")

	//se ejecuta cada listener en rutinas distintas
	wait.Add(2)
	go func() {
		escucharCliente(listenerCliente)
		wait.Done()
	}()
	go func() {
		escucharData(listenerData)
		wait.Done()
	}()

	wait.Wait()
}

func escucharCliente(listener net.Listener) {
	//función encargada del servidor para escuchar al cliente, termina cuando ocurra un error es el servidor grpc
	//o en el listener

	servidorCliente := ClienteName.Server{}

	grpcServer := grpc.NewServer()

	ClienteName.RegisterNameServiceServer(grpcServer, &servidorCliente)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func escucharData(listener net.Listener) {
	//función encargada del servidor para escuchar a los DataNode, termina cuando ocurra un error es el servidor grpc
	//o en el listener

	servidorData := com_namenode.ServerNamenode{}

	grpcServer := grpc.NewServer()

	com_namenode.RegisterInteraccionesServer(grpcServer, &servidorData)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
