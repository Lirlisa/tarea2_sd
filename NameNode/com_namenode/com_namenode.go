package com_namenode

import (
	context "context"
	"log"
	"os"
	sync "sync"

	"../../estructuras"
)

//ServerNamenode representa al servidor
type ServerNamenode struct {
	placeholder int
}

//EscribirLog se encarga de escribir en el archivo LOG lo enviado por DataNode
func (c *ServerNamenode) EscribirLog(ctx context.Context, in *Log) (*EstadoEscritura, error) {
	var candado sync.Mutex
	f, err := os.OpenFile("LOG.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("error: ", err.Error())
	}
	defer f.Close()
	candado.Lock()
	estructuras.Locaciones[in.GetTitulo()] = in.GetTexto()
	candado.Unlock()
	if _, err = f.WriteString(in.GetTexto()); err != nil {
		return &EstadoEscritura{
			Estado: false,
			Msg:    err.Error(),
		}, nil
	} else {
		return &EstadoEscritura{
			Estado: true,
			Msg:    "",
		}, nil
	}
}
