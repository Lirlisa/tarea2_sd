package com_namenode

import (
	context "context"
	"log"
	"os"
	sync "sync"

	"../../estructuras"
)

//ServerNamenode estructura para crear el servidor
type ServerNamenode struct {
	placeholder int
}

//EscribirLog para escribir en el log
func (c *ServerNamenode) EscribirLog(ctx context.Context, in *Log) (*EstadoEscritura, error) {
	var candado sync.Mutex
	candado.Lock()
	if estructuras.Ocupado {
		return &EstadoEscritura{
			Estado: false,
			Msg:    "toy ocupao",
		}, nil
	}
	estructuras.Ocupado = true
	candado.Unlock()
	f, err := os.OpenFile("LOG.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("error: ", err.Error())
	}
	defer f.Close()
	candado.Lock()
	defer candado.Unlock()
	estructuras.Locaciones[in.GetTitulo()] = in.GetTexto()
	if _, err = f.WriteString(in.GetTexto()); err != nil {
		estructuras.Ocupado = true
		return &EstadoEscritura{
			Estado: false,
			Msg:    err.Error(),
		}, nil
	}
	estructuras.Ocupado = true
	return &EstadoEscritura{
		Estado: true,
		Msg:    "",
	}, nil

}

//Request permite saber si se puede escribir en el log
func (c *ServerNamenode) Request(ctx context.Context, in *Id) (*Disponibilidad, error) {
	if !estructuras.Ocupado {
		return &Disponibilidad{
			Disponible: true,
		}, nil
	}
	return &Disponibilidad{
		Disponible: false,
	}, nil
}
