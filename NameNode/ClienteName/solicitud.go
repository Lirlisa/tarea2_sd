package ClienteName

import (
	"log"
	"strings"

	"sync"

	"../../estructuras"
	"golang.org/x/net/context"
)

//Server representa al servidor
type Server struct {
}

//SolicitudCliente procesa las solucitudes de los clientes
func (s *Server) SolicitudCliente(ctx context.Context, in *Message) (*Message, error) {
	mensaje := in.Body
	// si el mensaje es vacio se están solicitando los titulos disponibles
	if mensaje == "" {
		log.Printf("Solicitud titulos disponibles")
		return &Message{Body: titulos()}, nil
	}
	//De lo contrario se ve si existe indormación sobre el titulo solicitado
	log.Printf("Solicitud ubicaciones de: %s", mensaje)
	return &Message{Body: ubicaciones(mensaje)}, nil
}

//función encargada de devolver las ubicaciones de los chunks
func ubicaciones(libro string) string {
	var candado sync.Mutex
	candado.Lock()
	retorno := estructuras.Locaciones[libro]
	candado.Unlock()
	retorno = strings.Trim(retorno, "\n ")
	trozos := strings.Split(retorno, "\n")
	trozos = trozos[1:]
	var paraMandar string
	for _, k := range trozos {
		aux := strings.Split(k, " ")
		paraMandar += aux[1] + " "
	}
	return strings.Trim(paraMandar, " ")
}

//función encargada de devolver los titulos disponibles
func titulos() string {
	var retorno string
	var candado sync.Mutex
	candado.Lock()
	for k := range estructuras.Locaciones {
		retorno += k + "\n"
	}
	candado.Unlock()
	return retorno
}
