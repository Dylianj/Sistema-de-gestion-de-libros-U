# Sistema de Gestión de Libros Electrónicos en Go

Este proyecto es un sistema simple de gestión de libros electrónicos desarrollado en **Go (Golang)**.  
Permite registrar, listar, buscar, actualizar y eliminar libros, guardando la información en un archivo JSON.

---

## Objetivo

Aplicar los fundamentos de programación en Go —estructuras, slices, funciones y manejo de archivos— para crear un sistema modular, funcional y mantenible.

---

## Estructura del Proyecto

sistema-libros/
│
├── main.go # Punto de entrada, contiene el menú principal
├── go.mod # Módulo Go
│
├── gestion/
│ └── gestion.go # Lógica CRUD (crear, leer, actualizar, eliminar)
│
└── models/
└── libro.go # Definición de la estructura Libro

---

## Funcionalidades

-  Agregar un nuevo libro  
-  Listar todos los libros  
-  Buscar libros por título  
-  Actualizar datos de un libro  
-  Eliminar libros  
-  Guardar y cargar datos desde un archivo `datos.json`

---

## Paquetes Utilizados

Solo se emplean paquetes estándar de Go:

- `fmt` — Entrada y salida por consola  
- `bufio` — Lectura controlada de texto  
- `encoding/json` — Manejo del formato JSON  
- `os` — Lectura y escritura de archivos  
- `strconv` — Conversión de cadenas a números  
- `strings` — Manipulación de texto  

No se usan paquetes externos.

---

## Ejecución del Programa

1. Clona el repositorio:
   ```bash
   git clone https://github.com/tuusuario/sistema-libros-go.git
   cd sistema-libros-go
