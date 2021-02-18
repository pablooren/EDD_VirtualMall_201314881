package vectores

import (
	"fmt"

	"../ListaTienda"
)

type Clasificacion struct {
	clasi   int
	tiendas *ListaTienda.Listat
}
type Departamento struct {
	clasi  [100]Clasificacion
	nombre string
}
type Indice struct {
	letra string
	depto [100]Departamento
}
type Matriz struct {
	indice [100]Indice
}

func Imprimir() {
	ls := ListaTienda.NuevaLista()
	ls.Insertar("Juan", "sexo", "123", 1)
	ls.Insertar("Juana", "sexo", "123", 2)
	ls.Insertar("Juanito", "sexo", "123", 6)
	ls.Insertar("Juano", "sexo", "123", 3)
	ls.Insertar("Juanita", "sexo", "123", 4)
	ls.Insertar("Juanote", "sexo", "123", 5)
	ls.Ordenar()
	ls.Imprimir()
	fmt.Println("eliminamos a Juanito")
	ls.Eliminar("Juanito")
	ls.Imprimir()

}
