package com_cliente

import (
	context "context"
	"log"
	"os"
	"strconv"
)

type ServerCliente struct {
	placeholder int
}

func (c *ServerCliente) SubirLibro(ctx context.Context, in *Libro) (*EstadoSubida, error) {
	val := strconv.FormatUint(in.GetChunkActual(), 10)
	f, err := os.OpenFile(in.GetTitulo()+"_"+val, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Printf("No se pude crear chunk: %s", err.Error())
		return &EstadoSubida{
			Estado: false,
			Msg:    err.Error(),
		}, nil
	}
	defer f.Close()
	log.Println("xd")
	_, err = f.Write(in.GetChunk())
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

func (c *ServerCliente) PedirChunk(ctx context.Context, in *SolicitudChunk) (*Chunk, error) {
	val := strconv.FormatUint(in.GetNChunk(), 10)
	buf := make([]byte, 250*1000)
	data, err := os.Open(in.GetTitulo() + "_" + val)
	if err != nil {
		return &Chunk{
			Estado: false,
			Data:   buf[:1],
		}, nil
	}
	defer data.Close()
	n, err := data.Read(buf)
	return &Chunk{
		Estado: true,
		Data:   buf[:n],
	}, nil
}
