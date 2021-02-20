package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./ListaTienda"
	vectores "./Vectores"

	"github.com/gorilla/mux"
)

var dat Datos
var matriz vectores.Matriz

func main() {

	request()

}

func request() {
	myrouter := mux.NewRouter().StrictSlash(true)
	//esta es una pagina de home por defecto
	myrouter.HandleFunc("/", homePage)
	// creamos el primer endpoint de carga de datos
	myrouter.HandleFunc("/cargartienda", cargartienda).Methods("POST")
	myrouter.HandleFunc("/Eliminar", Eliminar).Methods("POST")
	myrouter.HandleFunc("/getArreglo", Graficar).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", myrouter))

}
func Graficar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "La grafica ya esta dibujada")

}
func Eliminar(w http.ResponseWriter, r *http.Request) {
	//metodo para eliminar desde un json
	body, _ := ioutil.ReadAll(r.Body)
	var elim Del
	json.Unmarshal(body, &elim)

	//fmt.Println(elim.Nombre, " ", elim.Categoria, " ", elim.Calificacion)
	w.Header().Set("Content-Type", "application/json")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "EDD_VirtualMall")
	fmt.Fprintln(w, "Juan Pablo Orellana")
	fmt.Fprintln(w, "201314881")

}
func cargartienda(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	// hacer estructura auxiliar para obtener datos de json
	json.Unmarshal(body, &dat)
	//	fmt.Println(dat.Datos[0].Indice)
	Ingresar(dat)
	vectores.Linealizacion(matriz)

	// ahora que tenemos los datos en json vamos a guardarlos en nuestras estructuras
	//fmt.Println("estos indies: ", ma.Indice[0].)

	//vectores.Linealizacion(ma)
	w.Header().Set("Content-Type", "application/json")
	//json.Encoder(w).Encode(re)

}
func Ingresar(datos Datos) {
	matriz = vectores.GetMatriz()
	matriz.Indice = make([]vectores.Indice, len(datos.Datos))
	//recorremos los indices
	for a := 0; a < len(datos.Datos); a++ {
		matriz.Indice[a].Depto = make([]vectores.Departamento, len(datos.Datos[a].Departamentos))
		matriz.Indice[a].Letra = datos.Datos[a].Indice
		//	fmt.Println("el indice ", matriz.Indice[a], " tiene estos deptos ", len(matriz.Indice[a].Depto))
		//recorremos las clasificaciones
		for b := 0; b < len(datos.Datos[a].Departamentos); b++ {
			matriz.Indice[a].Depto[b].Clasi = make([]vectores.Clasificacion, 5)
			matriz.Indice[a].Depto[b].Nombre = datos.Datos[a].Departamentos[b].Nombre
			for z := 1; z <= 5; z++ {
				matriz.Indice[a].Depto[b].Clasi[z-1].Clasi = z
				matriz.Indice[a].Depto[b].Clasi[z-1].Tiendas = ListaTienda.NuevaLista()

			} // asignacion de los 5 espacios de clasificacion
			for c := 0; c < len(datos.Datos[a].Departamentos[b].Tiendas); c++ {
				//aqui insertamos tienda por tienda
				for z := 0; z < 5; z++ {
					if datos.Datos[a].Departamentos[b].Tiendas[c].Calificacion == matriz.Indice[a].Depto[b].Clasi[z].Clasi {
						aux := datos.Datos[a].Departamentos[b].Tiendas[c]

						matriz.Indice[a].Depto[b].Clasi[z].Tiendas.Insertar(aux.Nombre, aux.Descripcion, aux.Contacto, aux.Calificacion)

					}

				}

			}

		}
	}

}
func ImprimirM() {
	for a := 0; a < len(matriz.Indice); a++ {
		fmt.Println("                                           Indice: ", matriz.Indice[a].Letra)
		for b := 0; b < len(matriz.Indice[a].Depto); b++ {
			fmt.Println("               Departamento:", matriz.Indice[a].Depto[b].Nombre)
			for c := 0; c < len(matriz.Indice[a].Depto[b].Clasi); c++ {
				fmt.Println("       Clasificacion: ", matriz.Indice[a].Depto[b].Clasi[c].Clasi)
				matriz.Indice[a].Depto[b].Clasi[c].Tiendas.Imprimir()

			}
		}

	}
}

type Tienda struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
}
type Departamento struct {
	Nombre  string
	Tiendas []Tienda
}
type Dato struct {
	Indice        string
	Departamentos []Departamento
}
type Datos struct {
	Datos []Dato
}
type Del struct {
	Nombre       string
	Categoria    string
	Calificacion int
}
