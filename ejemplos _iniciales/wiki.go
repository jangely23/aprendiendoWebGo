package main

import (
	"fmt"
	"io/ioutil"
)

type Pagina struct {
	Titulo string
	Cuerpo []byte //permite trabajar correctamente co ioutil.WriteFile
}

func main() {
	//crear y guardar la pagína en disco.
	pag1 := &Pagina{Titulo: "Prueba_1_wiki", Cuerpo: []byte("Este es el cuerpo de prueba 1 wiki")}
	pag1.guardar()

	//cargar la pagína guarda
	pag2, err := cargarPagina("Prueba_1_wiki")

	if err != nil {
		fmt.Println("Ha ocurrido un error: ", err)
		return
	}

	fmt.Println(string(pag2.Cuerpo))
}

//guardar almacena las paginas en le disco
func (p *Pagina) guardar() error {
	nombre := p.Titulo + ".txt"
	return ioutil.WriteFile(nombre, p.Cuerpo, 0600) //escribe un slice de archivos y retorna un error, el 0600 son los permisos del archivo
}

//cargarPagina permite leer el contenido de una pagina
func cargarPagina(titulo string) (*Pagina, error) {
	nombre_archivo := titulo + ".txt"
	cuerpo, err := ioutil.ReadFile(nombre_archivo)

	if err != nil {
		return nil, err
	}

	return &Pagina{Titulo: titulo, Cuerpo: cuerpo}, nil
}
