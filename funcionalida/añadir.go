package funcionalida

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/lucas5z/ternimal/structura"
)

func AñadirNotas(file *os.File, n []structura.Notas) {
	by, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(file)
	_, err = w.Write(by)
	if err != nil {
		panic(err)
	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}

}
