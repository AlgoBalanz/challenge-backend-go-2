# Balanz - Desafío de Backend en Go

¡Bienvenido al Desafío de Backend en Go! Este desafío está diseñado para evaluar tus habilidades en programación y desarrollo de backend utilizando el lenguaje de programación Go (Golang). En este desafío, se espera que diseñes y desarrolles dos endpoints que interactúen con una base de datos SQLite y procesen datos de un archivo CSV.

## Pautas Generales

- Utiliza el framework Gin Gonic para implementar los endpoints de manera eficiente.
- Utiliza GORM como ORM para interactuar con la base de datos SQLite.
- Asegúrate de manejar los errores de manera adecuada en tu código.

## Endpoint 1: Parseo y Persistencia de Datos desde un Archivo CSV

Diseña un endpoint que acepte un archivo CSV en el cuerpo de la solicitud. El archivo CSV contiene información de cursos con los siguientes campos:

- Name
- Price
- Category
- StartDate
- EndDate

El endpoint deberá parsear los datos del archivo CSV y almacenarlos en una tabla de una base de datos SQLite. Asegurate de validar que la fecha de inicio sea menor que la fecha de finalización antes de persistir el registro.

```
POST /upload-courses
Headers:
Content-Type: multipart/form-data
Body:
file: <archivo_csv>
```

## Endpoint 2: Obtener Cursos

Diseña otro endpoint que devuelva todos los productos almacenados en la base de datos SQLite.

```
GET /courses
```

## Notas Adicionales

- Utiliza la estructura `Course` definida en el código para representar los datos de los cursos.
- Asegurate de que el ID de cada registro de `Course` sea un UUID generado automáticamente.
