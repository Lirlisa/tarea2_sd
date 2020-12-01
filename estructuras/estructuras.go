package estructuras

//LibrosGuardados struct que mantiane los datos de los libros recibidos
type LibrosGuardados struct {
	Titulo          string
	ChunksTotales   uint64
	ChunksRecibidos uint64 //ya que se recibe de a 1, es para saber cuántos han llegado
	Repartido       bool   //para saber si ya se repartió entre los nodos
}

//AlmacenLibros variable global que lleva cuenta de los libros recibidos indexados por nombre
var AlmacenLibros map[string]*LibrosGuardados = make(map[string]*LibrosGuardados)
