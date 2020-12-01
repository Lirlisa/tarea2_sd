package ClienteName

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
				datos = datos + linea + " "
				i = i + 1
				continue
			} else if i == partes {
				datos = datos + linea
				i = i + 1
				break
			}

		}
		if strings.HasPrefix(linea, libro+" ") {
			partes, err = strconv.Atoi(strings.TrimLeft(linea, libro+" "))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			continue

		}

	}
	var datoslista []string = strings.Split(datos, " ")
	respuesta := ""
	if partes != 0 {
		for i = 1; i <= partes; i++ {
			aux := strconv.Itoa(i)
			for j := 0; j < len(datoslista); j++ {
				if aux == datoslista[j] {
					respuesta = respuesta + datoslista[j+1] + " "
					break
				}
			}
		}
	}
	return strings.TrimSpace(respuesta)
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
