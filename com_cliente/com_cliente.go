package com_cliente

import (
	context "context"
	"log"
	"os"
	"strconv"
	"sync"

	"../estructuras"
)

//ServerCliente elemento que represetará el servidor
type ServerCliente struct {
	placeholder int
}

//SubirLibro función que maneja la subida de libros
func (c *ServerCliente) SubirLibro(ctx context.Context, in *Libro) (*EstadoSubida, error) {
	var candado sync.Mutex
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
	_, err = f.Write(in.GetChunk())
	if err != nil {
		return &EstadoSubida{
			Estado: false,
			Msg:    err.Error(),
		}, nil
	}

	candado.Lock()
	if _, ok := estructuras.AlmacenLibros[in.GetTitulo()]; ok {
		estructuras.AlmacenLibros[in.GetTitulo()].ChunksRecibidos++
		if estructuras.AlmacenLibros[in.GetTitulo()].ChunksTotales == estructuras.AlmacenLibros[in.GetTitulo()].ChunksRecibidos {
			estructuras.Push(&estructuras.ColaParaEnvios, in.GetTitulo())
		}
	} else {
		estructuras.AlmacenLibros[in.GetTitulo()] = new(estructuras.LibrosGuardados)
		estructuras.AlmacenLibros[in.GetTitulo()].Titulo = in.GetTitulo()
		estructuras.AlmacenLibros[in.GetTitulo()].ChunksTotales = in.GetTotalChunks()
		estructuras.AlmacenLibros[in.GetTitulo()].ChunksRecibidos = 1
		estructuras.AlmacenLibros[in.GetTitulo()].Repartido = false
	}
	candado.Unlock()
	log.Printf("Creado chunk de %s", in.GetTitulo())
	return &EstadoSubida{
		Estado: true,
		Msg:    "",
	}, nil
}

//PedirChunk solicita los chunks a los datanodes
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
