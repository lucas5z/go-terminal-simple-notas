package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/lucas5z/ternimal/funcionalida"
	"github.com/lucas5z/ternimal/structura"
)

var n []structura.Notas

func main() {
	open()
}

// EN ESTA FUNCION abro el archivo veo si hay contenido y lo guardo
func open() {
	//creamos una slice de notas

	//abrimos el archivo
	file, err := os.OpenFile("lista.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//vemos si hay algo dentro del archivo

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		by, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(by, &n)
		if err != nil {
			panic(err)
		}
	} else {
		n = []structura.Notas{}

	}

	if len(os.Args) < 2 {

		ImprimirUser()
		return
	}
	switch os.Args[1] {
	case "list":
		funcionalida.ListaNotas(n)
	case "add":
		escrito := bufio.NewReader(os.Stdin)
		fmt.Println("escriba el curso ")
		name1, _ := escrito.ReadString('\n')

		fmt.Println("escriba la candidad de credito ")
		name2, _ := escrito.ReadString('\n')
		name2 = strings.TrimSpace(name2)
		int_mane2, err := strconv.Atoi(name2)
		if err != nil {
			panic(err)
		}

		fmt.Println("escriba su nota")
		name3, _ := escrito.ReadString('\n')
		name3 = strings.TrimSpace(name3)
		float_name3, err := strconv.ParseFloat(name3, 32)
		if err != nil {
			fmt.Println("Error al convertir la nota:", err)
			return
		}
		//llamo la funcion de devuelve un []Notas
		n = funcionalida.CrearNota(n, name1, int_mane2, float32(float_name3))
		funcionalida.AñadirNotas(file, n)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("proporciona un id")
		}
		a := os.Args[2]
		str_a, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("ese no es un id")
		}
		n = funcionalida.Actulalizado(n, str_a)
		funcionalida.AñadirNotas(file, n)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("proporciona un id")
		}
		a := os.Args[2]
		str_a, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("ese no es un id")
		}
		n = funcionalida.Eliminar(n, str_a)
		funcionalida.AñadirNotas(file, n)
	default:
		ImprimirUser()
	}

}

func ImprimirUser() {
	fmt.Println("USO: NOTAS-SEMESTRE [list|add|complete|delete]")
}
