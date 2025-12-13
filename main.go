// Dylan Alcivar, Ihair Llamuca y Mateo Vivas
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sistema-libros/gestion"
	"sistema-libros/models"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

const carpetaPersistencia = "persistencia"

func main() {
	err := os.MkdirAll(carpetaPersistencia, 0755)
	if err != nil {
		panic("Error crítico: " + err.Error())
	}
	rutaDB := filepath.Join(carpetaPersistencia, "datos.json")     //ruta de la DB
	rutaLog := filepath.Join(carpetaPersistencia, "auditoria.txt") // ruta de los logs

	canalAuditoria := make(chan string)
	go iniciarLogger(canalAuditoria, rutaLog)

	var inventario gestion.GestorInventario = gestion.NuevoSistema(rutaDB)

	if err := inventario.Cargar(); err != nil {
		fmt.Println("Iniciando base de datos nueva...")
	}

	for {
		fmt.Println("\nGestión de Libros")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Buscar libro")
		fmt.Println("4. Actualizar libro")
		fmt.Println("5. Eliminar libro")
		fmt.Println("6. Guardar y Salir")
		fmt.Print("Seleccione una opción: ")

		scanner.Scan()
		opcionStr := scanner.Text()
		opcion, _ := strconv.Atoi(opcionStr)

		switch opcion {
		case 1:
			l := leerLibro()
			err := inventario.Agregar(l)
			if err != nil {
				fmt.Println("Error:", err)
				canalAuditoria <- "Error al agregar: " + err.Error()
			} else {
				fmt.Println("¡Libro agregado!")
				canalAuditoria <- "Libro agregado: " + l.Titulo
			}

		case 2:
			libros := inventario.Listar()
			if len(libros) == 0 {
				fmt.Println("\nNo hay libros registrados.")
			} else {
				fmt.Println("\n--- Listado ---")
				for i, l := range libros {
					fmt.Printf("%d. %s | %s | %d | %s\n", i+1, l.Titulo, l.Autor, l.Anio, l.Categoria)
				}
				canalAuditoria <- "Listado consultado."
			}

		case 3:
			fmt.Print("Título a buscar: ")
			scanner.Scan()
			titulo := scanner.Text()
			libro, err := inventario.Buscar(titulo)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("ENCONTRADO: %s (%s)\n", libro.Titulo, libro.Autor)
				canalAuditoria <- "Busqueda exitosa: " + titulo
			}

		case 4:
			fmt.Print("Ingrese el título del libro a modificar: ")
			scanner.Scan()
			tituloOriginal := scanner.Text()

			// verificamos primero si existe el libro
			_, err := inventario.Buscar(tituloOriginal)
			if err != nil {
				fmt.Println("El libro no existe, no se puede editar.")
			} else {
				fmt.Println("--- Ingrese los NUEVOS datos ---")
				nuevoLibro := leerLibro()

				err := inventario.Actualizar(tituloOriginal, nuevoLibro)
				if err != nil {
					fmt.Println("Error al actualizar:", err)
					canalAuditoria <- "Error al actualizar: " + tituloOriginal
				} else {
					fmt.Println("¡Libro actualizado correctamente!")
					canalAuditoria <- "Libro actualizado: " + tituloOriginal + " -> " + nuevoLibro.Titulo
				}
			}

		case 5:
			fmt.Print("Ingrese el título del libro a eliminar: ")
			scanner.Scan()
			titulo := scanner.Text()

			err := inventario.Eliminar(titulo)
			if err != nil {
				fmt.Println("Error:", err)
				canalAuditoria <- "Intento fallido de eliminar: " + titulo
			} else {
				fmt.Println("¡Libro eliminado correctamente!")
				canalAuditoria <- "Libro eliminado: " + titulo
			}

		case 6:
			if err := inventario.Guardar(); err != nil {
				fmt.Println("Error al guardar:", err)
			}
			canalAuditoria <- "Sesión finalizada."
			fmt.Println("Guardando y saliendo...")
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func iniciarLogger(canal chan string, rutaArchivo string) {
	file, err := os.OpenFile(rutaArchivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	logger := log.New(file, "[LOG] ", log.Ldate|log.Ltime)
	for mensaje := range canal {
		logger.Println(mensaje)
	}
}

func leerLibro() models.Libro {
	fmt.Print("Título: ")
	scanner.Scan()
	t := scanner.Text()
	fmt.Print("Autor: ")
	scanner.Scan()
	a := scanner.Text()
	fmt.Print("Categoría: ")
	scanner.Scan()
	c := scanner.Text()
	fmt.Print("Año: ")
	scanner.Scan()
	anioStr := scanner.Text()
	anio, _ := strconv.Atoi(anioStr)
	return models.Libro{Titulo: t, Autor: a, Categoria: c, Anio: anio}
}
