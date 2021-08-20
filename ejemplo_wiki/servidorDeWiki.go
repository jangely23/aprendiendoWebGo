package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Pagina struct {
	Titulo string
	Cuerpo []byte
}

func main() {
	//Creamos y guardamos una página para que el cliente la pida
	pag1 := &Pagina{Titulo: "Ejemplo", Cuerpo: []byte(
		"¡Hola personita! Este es el cuerpo de tu página.")}
	pag1.guardar()
	http.HandleFunc("/view/", manejadorMostrarPagina)
	fmt.Println("El servidor se encuentra en ejecución")
	http.ListenAndServe(":8080", nil)
}

//Manejador de peticiones
func manejadorMostrarPagina(w http.ResponseWriter, r *http.Request) {
	titulo := r.URL.Path
	pagina, _ := cargarPagina(titulo)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", pagina.Titulo[len("/view/"):], pagina.Cuerpo)
	fmt.Println("¡Página servida!")
}

//Método para guardar página
func (p *Pagina) guardar() error {
	nombre := p.Titulo + ".txt"
	return ioutil.WriteFile("./view/"+nombre, p.Cuerpo, 0600)
}

//Método para cargar página
func cargarPagina(titulo string) (*Pagina, error) {
	nombre_archivo := titulo + ".txt"
	fmt.Println("El cliente ha pedido:" + nombre_archivo)
	cuerpo, err := ioutil.ReadFile("." + nombre_archivo)
	if err != nil {
		return nil, err
	}
	return &Pagina{Titulo: titulo, Cuerpo: cuerpo}, nil
}
