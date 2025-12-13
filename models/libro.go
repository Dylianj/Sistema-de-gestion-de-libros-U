// Dylan Alcivar, Ihair Llamuca y Mateo Vivas
package models

//se mantienen los campos públicos para que el json pueda acceder
type Libro struct {
	Titulo    string `json:"titulo"`
	Autor     string `json:"autor"`
	Categoria string `json:"categoria"`
	Anio      int    `json:"anio"`
}
