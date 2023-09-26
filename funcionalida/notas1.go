package funcionalida

import (
	"fmt"

	"github.com/lucas5z/ternimal/structura"
)

func ListaNotas(n []structura.Notas) {
	if len(n) == 0 {
		fmt.Println("la lista de nota esta vacia")
		return
	}
	//var salida string
	for _, conte := range n {
		estado := " "
		if conte.Completado {
			estado = "✓"
		}
		fmt.Printf("%s [%d]-> %-20s |   %-2d  |  %.2f\n", estado, conte.Id, conte.Curso, conte.Cretidos, conte.Nota)
		//salida += fmt.Sprintf("%s [%d]-> %s->%d->%.2f\n", estado, conte.Id, conte.Curso, conte.Cretidos, conte.Nota)
	}
	//fmt.Println(salida)
}

// creando tarea con los pareametros
func CrearNota(n []structura.Notas, c string, credi int, nt float32) []structura.Notas {
	/* nota := structura.Notas{

		Curso      string
		Cretidos   int
		Nota       float32
		Completado bool
	} */
	i := my_id(n)
	nota := structura.Notas{
		Id:         i,
		Curso:      c,
		Cretidos:   credi,
		Nota:       nt,
		Completado: false,
	}
	n = append(n, nota)
	return n
}

// creando el id
func my_id(n []structura.Notas) int {
	if len(n) == 0 {
		return 1
	}
	return n[len(n)-1].Id + 1
}

//actualizar

func Actulalizado(n []structura.Notas, id int) []structura.Notas {
	for i, conte := range n {
		if conte.Id == id {
			n[i].Completado = true
			return n
		}
	}
	return n
}
func Eliminar(n []structura.Notas, id int) []structura.Notas {
	for i, conte := range n {
		if conte.Id == id {
			n = append(n[:i], n[i+1:]...)
			return n
		}
	}
	return n
}
