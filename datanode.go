package main

import (
	"context"
	"fmt"
	"log"

	"./com_namenode"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Se ha producido un error: %s", err)
	}
	defer conn.Close()

	cliente := com_namenode.NewInteraccionesClient(conn)
	respuesta, err := cliente.EscribirLog(context.Background(), &com_namenode{
		Texto: "xd",
	})
	if err != nil {
		log.Fatalf("Se ha producido un error %s", err)
	}

	fmt.Println(respuesta.Estado)
}
