package estructuras

//Locaciones guarda las ubicaciones de los chunks del libro
var Locaciones map[string]string = make(map[string]string)

//Ocupado para saber si namenode está escribiendo en log
var Ocupado bool = false
