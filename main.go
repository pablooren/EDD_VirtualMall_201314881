package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"./ListaTienda"
	vectores "./Vectores"

	"github.com/gorilla/mux"
)

var dat Datos
var matriz vectores.Matriz
var vector vectores.VectorL

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
	myrouter.HandleFunc("/id", id).Methods("GET")
	myrouter.HandleFunc("/guardar", GuardarDatos).Methods("GET")
	myrouter.HandleFunc("/TiendaEspecifica", GetTiendaE).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", myrouter))

}
func GetTiendaE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	var elim Del
	json.Unmarshal(body, &elim)
	resultado := vectores.BuscarT(elim.Categoria, elim.Nombre, elim.Calificacion, vector)
	json.NewEncoder(w).Encode(resultado)

}

func GuardarDatos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matriz)

}
func Graficar(w http.ResponseWriter, r *http.Request) {
	vector = vectores.Linealizacion(matriz)
	vectores.Graficar(vector)
	fmt.Fprintln(w, "La grafica ya esta dibujada")

}
func id(w http.ResponseWriter, r *http.Request) {
	var resultado ListaTienda.ResL
	vars := r.URL.Query()
	valor := vars["id"]
	strinvalor := valor[0]

	valorr, error := strconv.ParseInt(strinvalor, 10, 64)
	fmt.Println("Esto es id :", valorr)
	if error != nil {
	}
	resultado = vectores.BusquedaL(int(valorr), vector)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultado)

}
func Eliminar(w http.ResponseWriter, r *http.Request) {
	//metodo para eliminar desde un json
	body, _ := ioutil.ReadAll(r.Body)
	var elim Del
	json.Unmarshal(body, &elim)
	// aqui toda la logica para eliminar desde el vector linealizado
	//Verificamos que la letra de la tienda sea la correcta
	for a := 0; a < len(matriz.Indice); a++ {
		if matriz.Indice[a].Letra == string(elim.Nombre[0]) {
			// si si es que la letra si existe por lo que puede estar dentro de un departamento
			for b := 0; b < len(matriz.Indice[a].Depto); b++ {
				if matriz.Indice[a].Depto[b].Nombre == elim.Categoria {
					// si si pueda que exista una tienda con ese nombre
					aux := matriz.Indice[a].Depto[b].Clasi[elim.Calificacion-1].Tiendas.Eliminar(elim.Nombre)
					fmt.Fprintf(w, aux)

				} else {
					//	fmt.Fprintf(w, "La tienda no existe 1")
				}
			}

		} else {
			//	fmt.Fprintf(w, "La tienda no existe 2")
		}
	}

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
	vector = vectores.Linealizacion(matriz)

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
