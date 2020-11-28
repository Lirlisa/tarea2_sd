package com_datanode

import (
	"golang.org/x/net/context"
)

type ServerDatanode struct {
	placeholder int
}

func (s *ServerDatanode) Disponible(ctx context.Context, in *Empty) (*Disponibilidad, error) {
	return &Disponibilidad{
		Data: true,
	}, nil
}

func (s *ServerDatanode) SubirArchivo(Interacciones_SubirArchivoServer) error {

}
