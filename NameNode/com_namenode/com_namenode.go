package com_namenode

import (
	context "context"
	"fmt"
	"os"
)

type ServerNamenode struct {
	placeholder int
}

func (c *ServerNamenode) EscribirLog(ctx context.Context, in *Log) (*EstadoEscritura, error) {
	f, err := os.OpenFile("./LOG.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

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
