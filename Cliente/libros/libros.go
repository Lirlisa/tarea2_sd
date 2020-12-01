package libros

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Subir(ctx context.Context, in *Chunk) (*Chunk, error) {
	log.Printf("Receive message body from client: %s", string(in.Data))
	return &Chunk{Titulo: "Hello From the Server!"}, nil
}
