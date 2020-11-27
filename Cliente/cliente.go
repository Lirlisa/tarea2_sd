package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"./ClienteName"
)

var tipo int

func main() {

	for {
		for {
			fmt.Println("\nIngrese comportamiento a seguir (1/2):\n1 Cliente Uploader\n2 Cliente Downloader")
			fmt.Scanln(&tipo)
			if tipo == 1 || tipo == 2 {
				break
			}
		}

		if tipo == 1 {
			ClienteUploader()
		} else {
			ClienteDownloader()
		}
	}

}

func ClienteUploader() {

	var libro string
	nodoAzar := rand.Intn(3)

	fmt.Println("\nIngrese nombre de libro a subir (FORMATO: NOMBRE.pdf)")
	fmt.Scanln(&libro)

	file, err := os.Open(libro)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 250000 // 1 MB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)
		////////////////////////////////////////////////////////////////////////REEMPLAZAR PARA ENVIAR A DATANODE
		// write to disk
		fileName := "bigfile_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

		fmt.Println("Split to : ", fileName)
		fmt.Println(nodoAzar)
	}

}

func ClienteDownloader() {

	fmt.Println("ClienteDownloader")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := ClienteName.NewNameServiceClient(conn)

	//solicitar titulos
	response, err := c.SolicitudCliente(context.Background(), &ClienteName.Message{Body: ""})
	if err != nil {
		log.Fatalf("Error al llamar SolicitudCliente: %s", err)
	}

	if len(response.Body) == 0 {
		fmt.Println("\nNO HAY LIBROS DISPONIBLES\n")
		return
	}
	var titulos []string = strings.Split(response.Body, " ")
	fmt.Println("\nLIBROS DISPONIBLES\n")
	for t := 0; t < len(titulos); t++ {
		fmt.Println(titulos[t])
	}

	var libro string
	fmt.Println("\nIngrese nombre de libro a solicitar(Igual al mostrado anteriormente, sin espacios)")
	fmt.Scanln(&libro)

	//solicitar ubicaciones
	response, err = c.SolicitudCliente(context.Background(), &ClienteName.Message{Body: libro})
	if err != nil {
		log.Fatalf("Error al llamar SolicitudCliente: %s", err)
	}
	ubicaciones := response.Body
	log.Printf("Response from server: %s", ubicaciones)
	var ubicacioneslista []string = strings.Split(ubicaciones, " ")
	log.Printf("Response from server: %s", ubicacioneslista)
	///////////////////////////////////COMENZAR A PEDIR CHUNKS

}
