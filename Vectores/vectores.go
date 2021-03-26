package vectores

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"../ListaTienda"
)

type Clasificacion struct {
	Clasi   int
	Tiendas *ListaTienda.Listat
}
type Linealizar struct {
	indice, depto string
	clasi         int
	list          *ListaTienda.Listat
}
type VectorL struct {
	nodo []Linealizar
}
type Departamento struct {
	Clasi  []Clasificacion
	Nombre string
}
type Indice struct {
	Letra string
	Depto []Departamento
}
type Matriz struct {
	Indice []Indice
}

func GetMatriz() Matriz {
	var nuevo Matriz
	return nuevo
}

func Linealizacion(matriz Matriz) VectorL {
	var vector VectorL
	tamaño := len(matriz.Indice) * len(matriz.Indice[0].Depto) * 5
	vector.nodo = make([]Linealizar, tamaño)

	contador := 0
	// vamos a recorrer los indices
	for i := 0; i < len(matriz.Indice); i++ {
		for j := 0; j < len(matriz.Indice[i].Depto); j++ {
			for k := 0; k < 5; k++ {
				vector.nodo[contador].indice = matriz.Indice[i].Letra
				vector.nodo[contador].depto = matriz.Indice[i].Depto[j].Nombre
				vector.nodo[contador].clasi = matriz.Indice[i].Depto[j].Clasi[k].Clasi
				vector.nodo[contador].list = matriz.Indice[i].Depto[j].Clasi[k].Tiendas
				vector.nodo[contador].list.Ordenar()

				contador++

			}
		}
	}

	//for a := 0; a < len(vector.nodo); a++ {
	//	fmt.Println(vector.nodo[a].indice, ": ", vector.nodo[a].depto, ": ", vector.nodo[a].clasi)
	//	vector.nodo[a].list.Imprimir()

	//}
	//	Graficar(vector)
	// toda la logica para linealizar
	return vector
}
func GetFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDWR, 0775)
	if err != nil {
		log.Fatal(err)

	}
	return file
}

func Graficar(vector VectorL) {
	os.Create("vectores/grafica_vector.dot")
	graphdot := GetFile("vectores/grafica_vector.dot")

	fmt.Fprintf(graphdot, "Digraph G{\n")
	fmt.Fprintf(graphdot, "rankdir = TB;\n ")
	fmt.Fprintf(graphdot, "node [shape = record];\n ")
	fmt.Fprintf(graphdot, "\n ")
	fmt.Fprintf(graphdot, "\n ")
	fmt.Fprintf(graphdot, "label = \"Linealizacion\";\n ")
	fmt.Fprintf(graphdot, "\n ")
	fmt.Fprintf(graphdot, "// creamos el vector  \n ")
	var text_aux string = ""
	fmt.Fprintf(graphdot, "Vector [label=\"  ")
	for a := 0; a < len(vector.nodo); a++ {
		text_aux = "<" + vector.nodo[a].indice + strconv.Itoa(a) + ">" + "Indice: " + vector.nodo[a].indice + " \\n Departamento: " + vector.nodo[a].depto + "\\n Calificacion: " + strconv.Itoa(vector.nodo[a].clasi) + "|"

		fmt.Fprintf(graphdot, text_aux)
		if (a+1)%5 == 0 {
			fmt.Fprintf(graphdot, "\n")

		}

	}

	fmt.Fprintf(graphdot, " \" ]; \n")
	for a := 0; a < len(vector.nodo); a++ {
		text_aux = vector.nodo[a].list.ImprimirGraph()
		if text_aux != "" {
			pos := "Vector:" + vector.nodo[a].indice + strconv.Itoa(a) + "->"

			fmt.Fprintf(graphdot, pos)
			fmt.Fprint(graphdot, text_aux)
			fmt.Fprintf(graphdot, "\n")

		}

	}
	// aqui van los enlaces con los nodos
	fmt.Fprintf(graphdot, "}")
	graphdot.Close()
	exec.Command("C:\\Program Files\\Graphviz\\bin\\dot", "-Tpdf", "Vectores/grafica_vector.dot", "-o", "Vectores/grafica.pdf").Output()

}
func BusquedaL(posicion_ int, vector VectorL) ListaTienda.ResL {

	//fmt.Println("esto es:", vector.nodo[24].list.BuscarL())

	return vector.nodo[posicion_].list.BuscarL()
}
func BuscarT(departamento_ string, nombre_ string, calificacion_ int, vector VectorL) NodoT {
	var aux NodoT
	for a := 0; a < len(vector.nodo); a++ {
		if string(nombre_[0]) == vector.nodo[a].indice {
			//quiere decir que si existe el indice con la tienda
			if departamento_ == vector.nodo[a].depto {
				//quiere decir que el departamento si existe
				if calificacion_ == vector.nodo[a].clasi {
					aux2 := vector.nodo[a].list.BuscarL()
					for c := 0; c < len(aux2.Datos); c++ {
						if aux2.Datos[c].Nombre == nombre_ {
							aux.nombre = aux2.Datos[c].Nombre
							aux.descripcion = aux2.Datos[c].Descripcion
							aux.contacto = aux2.Datos[c].Contacto
							aux.calificacion = vector.nodo[a].clasi
						}
					}

				}
			}
		}
	}
	return aux
}

type NodoT struct {
	nombre, descripcion, contacto string
	calificacion                  int
}
