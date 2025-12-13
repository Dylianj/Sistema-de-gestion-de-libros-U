// Dylan Alcivar, Ihair Llamuca y Mateo Vivas
package gestion

import (
	"encoding/json"
	"errors"
	"os"
	"sistema-libros/models"
	"strings"
)

// gestion.go implementa la interfaz GestorInventario.
// encapsulado: 'libros' es privado, es decir, no se puede modificar directamente
type Sistema struct {
	libros  []models.Libro
	archivo string
}

// NuevoSistema es el constructor para inicializar el sistema encapsulado.
func NuevoSistema(archivo string) *Sistema {
	return &Sistema{
		libros:  []models.Libro{},
		archivo: archivo,
	}
}

// Cargar archivo json
// manejo de errores explicito
func (s *Sistema) Cargar() error {
	data, err := os.ReadFile(s.archivo)
	if err != nil {
		if os.IsNotExist(err) {
			s.libros = []models.Libro{}
			return nil
		}
		return err //retornamos para que lo maneje el main
	}
	return json.Unmarshal(data, &s.libros)
}

// escribe en el archivo json
func (s *Sistema) Guardar() error {
	data, err := json.MarshalIndent(s.libros, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.archivo, data, 0644)
}

// agrega y valida si se puede añadir un libro
func (s *Sistema) Agregar(l models.Libro) error {
	if l.Titulo == "" || l.Autor == "" {
		return errors.New("el título y el autor son obligatorios")
	}
	s.libros = append(s.libros, l)
	return nil
}

// devuelve una copia de los libros para seguir encapsulamiento
func (s *Sistema) Listar() []models.Libro {
	return s.libros
}

// encuentra un libro por su titulo, regresa error si no existe
func (s *Sistema) Buscar(titulo string) (models.Libro, error) {
	for _, l := range s.libros {
		if strings.EqualFold(l.Titulo, titulo) {
			return l, nil
		}
	}
	// regresa un error personalizado
	return models.Libro{}, errors.New("libro no encontrado")
}

// busca un libro por su titulo y cambia sus datos.
func (s *Sistema) Actualizar(tituloOriginal string, nuevoLibro models.Libro) error {
	for i, l := range s.libros {
		if strings.EqualFold(l.Titulo, tituloOriginal) {
			// Encontramos el libro
			s.libros[i] = nuevoLibro
			return nil
		}
	}
	return errors.New("libro no encontrado para actualizar")
}

// busca un libro y lo borra de la lista
func (s *Sistema) Eliminar(titulo string) error {
	for i, l := range s.libros {
		if strings.EqualFold(l.Titulo, titulo) {
			// borrado de un slice en Go
			// se une la parte anterior al indice [:i] con la parte posterior [i+1:]
			s.libros = append(s.libros[:i], s.libros[i+1:]...)
			return nil
		}
	}
	return errors.New("libro no encontrado para eliminar")
}
