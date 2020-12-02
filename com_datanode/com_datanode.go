package com_datanode

import (
	"log"
	"os"

	"../estructuras"
	"golang.org/x/net/context"
)

//ServerDatanode representa el servidor para datanode
type ServerDatanode struct {
	placeholder int
}

//Disponible permite saber si un nodo está disponible
func (s *ServerDatanode) Disponible(ctx context.Context, in *Empty) (*Disponibilidad, error) {
	return &Disponibilidad{
		Estado: true,
	}, nil
}

//SubirArchivo le manda un archivo a los otros datanodes
func (s *ServerDatanode) SubirArchivo(ctx context.Context, in *Chunk) (*EstadoSubida, error) {
	f, err := os.OpenFile(in.GetNombre(), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Printf("No se pude crear chunk: %s", err.Error())
		return &EstadoSubida{
			Estado: false,
			Msg:    err.Error(),
		}, nil
	}
	defer f.Close()

	_, err = f.Write(in.GetData())
	if err != nil {
		log.Printf("Se recibió: %s", in.GetData())
		return &EstadoSubida{
			Estado: false,
			Msg:    err.Error(),
		}, nil
	}

	return &EstadoSubida{
		Estado: true,
		Msg:    "",
	}, nil
}

//Request para solicitar acceso al log
func (s *ServerDatanode) Request(ctx context.Context, in *Id) (*Disponibilidad, error) {
	//las id más pequeñas tienen prioridad
	if !estructuras.Ocupado || in.GetId() < estructuras.MiID {
		return &Disponibilidad{
			Estado: true,
		}, nil
	}
	return &Disponibilidad{
		Estado: false,
	}, nil
}
