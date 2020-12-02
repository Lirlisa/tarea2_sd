package ClienteName

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"sync"

	"../../estructuras"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SolicitudCliente(ctx context.Context, in *Message) (*Message, error) {
	mensaje := in.Body
	if mensaje == "" {
		log.Printf("Solicitud titulos disponibles")
		return &Message{Body: titulos()}, nil
	}
	log.Printf("Solicitud ubicaciones de: %s", mensaje)
	return &Message{Body: ubicaciones(mensaje)}, nil
}

func ubicaciones(libro string) string {
	var candado sync.Mutex
	candado.Lock()
	retorno := estructuras.Locaciones[libro]
	candado.Unlock()
	return retorno
}

func titulos() string {
	archivo, err := os.Open("./LOG.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	partes := 0
	i := 1
	datos := ""
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()

		if partes != 0 {
			if i < partes {
				i = i + 1
				continue
			} else if i == partes {
				i = 1
				partes = 0
				continue
			}

		}

		var linealista []string = strings.Split(linea, " ")
		partes, err = strconv.Atoi(linealista[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		datos = datos + linealista[0] + " "
	}

	return strings.TrimSpace(datos)
}
