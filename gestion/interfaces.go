// Dylan Alcivar, Ihair Llamuca y Mateo Vivas
package gestion

import "sistema-libros/models"

type GestorInventario interface {
	Cargar() error
	Guardar() error
	Agregar(l models.Libro) error
	Listar() []models.Libro
	Buscar(titulo string) (models.Libro, error)
	Actualizar(tituloOriginal string, nuevoLibro models.Libro) error
	Eliminar(titulo string) error
}
