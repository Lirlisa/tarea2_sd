package com_datanode

import (
	"log"
	"os"

	"golang.org/x/net/context"
)

type ServerDatanode struct {
	placeholder int
}

func (s *ServerDatanode) Disponible(ctx context.Context, in *Empty) (*Disponibilidad, error) {
	/*
		test test test
	*/
	return &Disponibilidad{
		Estado: true,
	}, nil
}

func (c *ServerDatanode) SubirArchivo(ctx context.Context, in *Chunk) (*EstadoSubida, error) {
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
