package ListaTienda

import (
	"fmt"
)

type nodot struct {
	Nombre              string
	Descripcion         string
	Contacto            string
	Calificacion        int
	anterior, siguiente *nodot
}
type Listat struct {
	primero, ultimo *nodot
	tamaño          int
}

func NuevaLista() *Listat {
	return &Listat{nil, nil, 0}
}

func (list *Listat) Insertar(nombre_ string, descripcion_ string, contacto_ string, calificacion_ int) {
	nuevo := &nodot{nombre_, descripcion_, contacto_, calificacion_, nil, nil}

	if list.primero == nil {
		list.primero = nuevo
		list.ultimo = nuevo

	} else {
		list.ultimo.siguiente = nuevo
		nuevo.anterior = list.ultimo
		list.ultimo = nuevo
	}
	list.tamaño++
}

func (list *Listat) Buscar(nombre_ string, calificacion_ int) *nodot {
	aux := list.primero
	for aux != nil {
		if (aux.Nombre == nombre_) && (aux.Calificacion == calificacion_) {
			return aux
		} else {
			aux = aux.siguiente
		}
	} // fin del for
	return nil // no encuentra el nodo por lo que regresa nulo
}
func (list *Listat) Imprimir() {
	aux := list.primero
	for aux != nil {
		fmt.Print("nombre: ", aux.Nombre)
		fmt.Print(" des: ", aux.Descripcion)
		fmt.Print(" contacto: ", aux.Contacto)
		fmt.Println(" calificacion : ", aux.Calificacion)

		aux = aux.siguiente

	}

}
func (list *Listat) ImprimirGraph() string {
	var cadena_aux string = ""
	if list.primero != nil {
		if list.primero == list.ultimo {
			cadena_aux = "\" " + list.primero.Nombre + "\\n Contacto: " + list.primero.Contacto + "\" "
		} else {
			aux := list.primero
			for aux != nil {
				cadena_aux = cadena_aux + "\" " + aux.Nombre + "\\n Contacto: " + aux.Contacto + "\"-> "

				aux = aux.siguiente
			}

			aux = list.ultimo.anterior
			for aux != list.primero {
				cadena_aux = cadena_aux + "\" " + aux.Nombre + "\\n Contacto: " + aux.Contacto + "\" ->"

				aux = aux.anterior
			}
			cadena_aux = cadena_aux + "\" " + aux.Nombre + "\\n Contacto: " + aux.Contacto + "\"; "

		}

	}
	//fmt.Println(cadena_aux)
	return cadena_aux
}
func GetAscii(nombre_ string) int {
	longitud := len(nombre_)
	valor := 0
	for i := 0; i < longitud; i++ {
		letra := nombre_[i]
		valor = valor + int(letra)

	}
	return valor

}
func (list *Listat) Eliminar(nombre_ string) {
	if list.primero.Nombre == nombre_ {
		// si vamos a eliminar al primero de la lista verificamos si hay mas nodos
		if list.tamaño == 1 {
			list.primero = nil
			list.ultimo = nil
			list.tamaño--
		} else {
			//eliminamos el primero de la lista y colocamos al siguien como nuevo inicio
			list.primero = list.primero.siguiente
			list.tamaño--

		}
	} else if list.ultimo.Nombre == nombre_ {
		// hacemos que el ultimo nodo sea el penultimo
		list.ultimo = list.ultimo.anterior
		list.ultimo.siguiente = nil
		list.tamaño--

	} else {
		// ahora buscamos si el nodo esta en medio de nuestra lista
		aux := list.primero
		for aux != nil {
			if aux.Nombre == nombre_ {
				// eliminamos el nodo aux y cambiamos los punteros
				aux.anterior.siguiente = aux.siguiente
				aux.siguiente.anterior = aux.anterior
				list.tamaño--
			} // fin if

			aux = aux.siguiente

		} // fin for
	} // fin if anidado
}

func (list *Listat) Ordenar() {
	for j := 0; j < list.tamaño; j++ {
		aux := list.primero
		for i := 0; i < list.tamaño-1; i++ {
			nombre1 := aux.Nombre
			nombre2 := aux.siguiente.Nombre

			if GetAscii(nombre1) > GetAscii(nombre2) {

				nom := aux.Nombre
				des := aux.Descripcion
				con := aux.Contacto
				cali := aux.Calificacion

				aux.Nombre = aux.siguiente.Nombre
				aux.Descripcion = aux.siguiente.Descripcion
				aux.Contacto = aux.siguiente.Contacto
				aux.Calificacion = aux.siguiente.Calificacion

				aux.siguiente.Nombre = nom
				aux.siguiente.Descripcion = des
				aux.siguiente.Contacto = con
				aux.siguiente.Calificacion = cali

			} // fin del if
			aux = aux.siguiente

		} //fin del for i

	} // fin del for j

}
