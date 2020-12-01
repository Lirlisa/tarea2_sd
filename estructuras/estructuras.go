package estructuras

import "sync"

//LibrosGuardados struct que mantiane los datos de los libros recibidos
type LibrosGuardados struct {
	Titulo          string
	ChunksTotales   uint64
	ChunksRecibidos uint64 //ya que se recibe de a 1, es para saber cuántos han llegado
	Repartido       bool   //para saber si ya se repartió entre los nodos
}

//AlmacenLibros variable global que lleva cuenta de los libros recibidos indexados por nombre
var AlmacenLibros map[string]*LibrosGuardados = make(map[string]*LibrosGuardados)

//Push para introducir elementos a la cola
func Push(q *[]string, elemento string) {
	var candado sync.Mutex
	candado.Lock()
	*q = append(*q, elemento)
	candado.Unlock()
}

//Pop saca un elemento de la cola
func Pop(q *[]string) string {
	var candado sync.Mutex
	candado.Lock()
	elem := (*q)[0]
	if len(*q) == 1 {
		*q = nil
	} else {
		*q = (*q)[1:]
	}
	candado.Unlock()
	return elem
}

//ColaParaEnvios Cuando un libro esté listo para repartir, su titulo se agregará acá
var ColaParaEnvios []string = make([]string, 0)
