package main

import (
	"fmt"
	"sistema-libros/gestion"
)

func main() {
	gestion.CargarLibros()

	for {
		fmt.Println("\n===== Sistema de Gestión de Libros Electrónicos =====")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Listar libros")
		fmt.Println("3. Buscar libro")
		fmt.Println("4. Actualizar libro")
		fmt.Println("5. Eliminar libro")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			gestion.AgregarLibro()
		case 2:
			gestion.ListarLibros()
		case 3:
			gestion.BuscarLibro()
		case 4:
			gestion.ActualizarLibro()
		case 5:
			gestion.EliminarLibro()
		case 6:
			gestion.GuardarLibros()
			fmt.Println("Cambios guardados. Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción inválida, intente de nuevo.")
		}
	}
}
