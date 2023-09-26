package structura

type Notas struct {
	Id         int     `json:"id"`
	Curso      string  `json:"curso"`
	Cretidos   int     `json:"creditos"`
	Nota       float32 `json:"nota"`
	Completado bool    `json:"completado"`
}
