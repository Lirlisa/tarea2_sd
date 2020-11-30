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

	rand.Seed(time.Now().UnixNano())

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
	var dir string
	nodoAzar := rand.Intn(3)

	fmt.Println("\nIngrese nombre de libro a subir (FORMATO: NOMBRE.pdf)")
	fmt.Scanln(&libro)

	file, err := os.Open(libro)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	if nodoAzar == 0 {
		dir = "dist46:9001" ////////////////////////////////////////////////////////////////////////////
		fmt.Println(dir)
	} else if nodoAzar == 1 {
		dir = "dist47:9001"
		fmt.Println(dir)
	} else {
		dir = "dist48:9001"
		fmt.Println(dir)
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(dir, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

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
		/*////////////////////////////////////////////////////////////////////////REEMPLAZAR PARA ENVIAR A DATANODE
		// write to disk
		fileName := "bigfile_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)*/
		c := com_cliente.NewInteraccionesClient(conn)

		respuesta, err = c.SubirLibro(context.Background(), &com_cliente.Libro{
			Titulo: strings.TrimRight(libro, ".pdf"),
			TotalChunks: totalPartsNum,
			ChunkActual: i+1,
			Data:   partBuffer})
		if err != nil {
			log.Fatalf("Error al subir chunk: %s", err)
		}
		if respuesta.Estado==false{
			log.Fatalf("Error al subir chunk: %s", respuesta.Msg)
		}

	}

}

func ClienteDownloader() {

	fmt.Println("ClienteDownloader")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
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
	if len(ubicaciones) == 0 {
		return
	}
	var ubicacioneslista []string = strings.Split(ubicaciones, " ")
	log.Printf("Response from server: %s", ubicacioneslista)
	///////////////////////////////////COMENZAR A PEDIR CHUNKS

	newFileName := "Descarga" + libro + ".pdf"
	_, err = os.Create(newFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//set the newFileName file to APPEND MODE!!
	// open files r and w

	file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var writePosition int64 = 0

	for d := 0; d < len(ubicacioneslista); d++ {

		fmt.Println(ubicacioneslista[d])
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(ubicacioneslista[d], grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c := com_cliente.NewInteraccionesClient(conn)

		respuesta, err = c.PedirChunk(context.Background(), &com_cliente.Libro{
			Titulo: libro,
			NChunck: uint64(d+1))
		if err != nil {
			log.Fatalf("Error al pedir chunk: %s", err)
		}

		if respuesta.Estado==true{
			_, err := file.Write(response.Data)

			if err != nil {
					fmt.Println(err)
					os.Exit(1)
			}

			file.Sync() //flush to disk
		} else{
			log.Fatalf("Error al pedir chunk %d",d+1)
		}


	}
	file.Close()
}

