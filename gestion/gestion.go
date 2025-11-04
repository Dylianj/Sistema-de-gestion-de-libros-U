package gestion

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sistema-libros/models"
	"strconv"
	"strings"
)

var Libros []models.Libro
var nada = "a"

func CargarLibros() {
	archivo, err := os.ReadFile("datos.json")
	if err != nil {
		fmt.Println("No se encontró el archivo, se creará uno nuevo.")
		Libros = []models.Libro{}
		return
	}
	json.Unmarshal(archivo, &Libros)
}

func GuardarLibros() {
	datos, err := json.MarshalIndent(Libros, "", "  ")
	if err != nil {
		fmt.Println("Error al guardar:", err)
		return
	}
	os.WriteFile("datos.json", datos, 0644)
}

func AgregarLibro() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Scanln(&nada)
	fmt.Print("Ingrese el título: ")
	titulo, _ := reader.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	fmt.Print("Ingrese el autor: ")
	autor, _ := reader.ReadString('\n')
	autor = strings.TrimSpace(autor)

	fmt.Print("Ingrese la categoría: ")
	categoria, _ := reader.ReadString('\n')
	categoria = strings.TrimSpace(categoria)

	fmt.Print("Ingrese el año de publicación: ")
	anioStr, _ := reader.ReadString('\n')
	anioStr = strings.TrimSpace(anioStr)
	anio, _ := strconv.Atoi(anioStr)

	libro := models.Libro{Titulo: titulo, Autor: autor, Categoria: categoria, Anio: anio}
	Libros = append(Libros, libro)
	fmt.Println("Libro agregado exitosamente.")
}

func ListarLibros() {
	if len(Libros) == 0 {
		fmt.Println("No hay libros registrados.")
		return
	}

	fmt.Println("\nLista de libros:")
	for i, l := range Libros {
		fmt.Printf("%d. %s (%d) - %s [%s]\n", i+1, l.Titulo, l.Anio, l.Autor, l.Categoria)
	}
}

func BuscarLibro() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanln(&nada)
	fmt.Print("Ingrese el título del libro a buscar: ")
	titulo, _ := reader.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	for _, l := range Libros {
		if strings.EqualFold(l.Titulo, titulo) {
			fmt.Printf("Encontrado: %s (%d) - %s [%s]\n", l.Titulo, l.Anio, l.Autor, l.Categoria)
			return
		}
	}
	fmt.Println("Libro no encontrado.")
}

func ActualizarLibro() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanln(&nada)
	fmt.Print("Ingrese el título del libro a actualizar: ")
	titulo, _ := reader.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	for i, l := range Libros {
		if strings.EqualFold(l.Titulo, titulo) {
			fmt.Print("Nuevo autor: ")
			autor, _ := reader.ReadString('\n')
			Libros[i].Autor = strings.TrimSpace(autor)

			fmt.Print("Nueva categoría: ")
			cat, _ := reader.ReadString('\n')
			Libros[i].Categoria = strings.TrimSpace(cat)

			fmt.Print("Nuevo año: ")
			anioStr, _ := reader.ReadString('\n')
			anioStr = strings.TrimSpace(anioStr)
			anio, _ := strconv.Atoi(anioStr)
			Libros[i].Anio = anio

			fmt.Println("Libro actualizado correctamente.")
			return
		}
	}
	fmt.Println("Libro no encontrado.")
}

func EliminarLibro() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Scanln(&nada)
	fmt.Print("Ingrese el título del libro a eliminar: ")
	titulo, _ := reader.ReadString('\n')
	titulo = strings.TrimSpace(titulo)

	for i, l := range Libros {
		if strings.EqualFold(l.Titulo, titulo) {
			Libros = append(Libros[:i], Libros[i+1:]...)
			fmt.Println("Libro eliminado correctamente.")
			return
		}
	}
	fmt.Println("Libro no encontrado.")
}
