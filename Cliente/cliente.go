package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"./ClienteName"
	"./com_cliente"
)

var tipo int

func main() {

	//semilla para aleatoriedad
	rand.Seed(time.Now().UnixNano())

	//Solicitud por pantalla para el comportamiento
	for {
		for {
			fmt.Println("\nIngrese comportamiento a seguir (1/2):\n1 Cliente Uploader\n2 Cliente Downloader")
			fmt.Scanln(&tipo)
			if tipo == 1 || tipo == 2 {
				break
			}
		}

		//Se ejecuta el comportamiento solicitado
		if tipo == 1 {
			ClienteUploader()
		} else {
			ClienteDownloader()
		}
	}

}

//Comportamiento de cliente uploader
func ClienteUploader() {

	//libro a enviar
	var libro string
	//variables para enviar a un nodo al azar
	var dir string
	nodoAzar := rand.Intn(3)

	//se solicita el nombre del libro a subir
	fmt.Println("\nIngrese nombre de libro a subir (FORMATO: NOMBRE.pdf)")
	fmt.Scanln(&libro)

	//se abre el archivo
	file, err := os.Open(libro)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	//a partir del valor al azar se define la dirección del nodo
	if nodoAzar == 0 {
		dir = "dist46:9001"
		fmt.Println(dir)
	} else if nodoAzar == 1 {
		dir = "dist47:9001"
		fmt.Println(dir)
	} else {
		dir = "dist48:9001"
		fmt.Println(dir)
	}

	//Se establece comunicación
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(dir, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	//datos necesarios para los chunk
	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 250000

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	//creación y envio de chunks
	for i := uint64(0); i < totalPartsNum; i++ {

		//creacion chunk
		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		c := com_cliente.NewInteraccionesClient(conn)

		//envio de chunk
		respuesta, err := c.SubirLibro(context.Background(), &com_cliente.Libro{
			Titulo:      strings.TrimRight(libro, ".pdf"),
			TotalChunks: totalPartsNum,
			ChunkActual: i + 1,
			Chunk:       partBuffer})
		if err != nil {
			log.Fatalf("No se pudo envíar chunk: %s", err)
		} else if respuesta.Estado == false {
			log.Fatalf("Error externo al subir chunk: %s", respuesta.Msg)
		} else {
			log.Printf("Chunk entregado con éxito")
		}

	}

}

//comportamiento cliente downloader
func ClienteDownloader() {

	//se establece comunicación con NameNode
	fmt.Println("ClienteDownloader")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := ClienteName.NewNameServiceClient(conn)

	//se solicitan titulos
	response, err := c.SolicitudCliente(context.Background(), &ClienteName.Message{Body: ""})
	if err != nil {
		log.Fatalf("Error al llamar SolicitudCliente: %s", err)
	}
	//si no hay titulos se termina la ejecucion de cliente downloader
	if len(response.Body) == 0 {
		fmt.Println("\nNO HAY LIBROS DISPONIBLES\n")
		return
	}

	//se muestran los libros disponibles
	//var titulos []string = strings.Split(response.Body, " ")
	fmt.Println("\nLIBROS DISPONIBLES\n")
	fmt.Print(response.Body)

	//se solicita el ingreso del libro a descargar
	var libro string
	fmt.Println("\nIngrese nombre de libro a solicitar(Igual al mostrado anteriormente, sin espacios)")
	fmt.Scanln(&libro)
	libro = strings.TrimSpace(libro)
	//se solicitan ubicaciones de los chunks
	response, err = c.SolicitudCliente(context.Background(), &ClienteName.Message{Body: libro})
	if err != nil {
		log.Fatalf("Error al llamar SolicitudCliente: %s", err)
	}
	ubicaciones := response.Body
	log.Printf("Response from server: %s", ubicaciones)
	if len(ubicaciones) == 0 {
		return
	}
	var ubicacioneslista []string = strings.Split(ubicaciones, " ")
	log.Printf("Response from server: %s", ubicacioneslista)
	//COMENZAR A PEDIR CHUNKS

	//creación del archivo de la descarga
	newFileName := "Descarga" + libro + ".pdf"
	_, err = os.Create(newFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//se abre el archivo de la descarga para comenzar a escribir

	file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Se solicitan los chunks de acurdo a la direccion indicada y se escriben en archivo
	for d := 0; d < len(ubicacioneslista); d++ {

		//Peticiones de chunks
		fmt.Println(ubicacioneslista[d])
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(ubicacioneslista[d], grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c := com_cliente.NewInteraccionesClient(conn)

		respuesta, err := c.PedirChunk(context.Background(), &com_cliente.SolicitudChunk{
			Titulo: libro,
			NChunk: uint64(d + 1),
		})
		if err != nil {
			log.Fatalf("Error al pedir chunk: %s", err)
		}

		//escritura en archivo
		if respuesta.Estado {
			_, err := file.Write(respuesta.Data)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			file.Sync() //flush to disk
		} else {
			log.Fatalf("Error al pedir chunk %d", d+1)
		}

	}
	file.Close()
}
