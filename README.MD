# MEDROOM BACKEND

A simple RESTful API compatible with heroku and swaggo using:

1. **Gin Framework**
2. **Gorm**
3. **Go 1.15**

## Installation & Run

```bash
# Download this project
$ git clone github.com/v1nland/medroom-backend
```

To run locally, generate .env file and modify its variables

```bash
cp .env.example .env
```

Build swagger

```bash
$ swag init
```

Run and install dependencies

```bash
$ go run *.go
```

## API

#### /estudiantes

-   `GET` : Get all estudiantes
-   `POST` : Create a new estudiante

#### /estudiantes/:uuid

-   `GET` : Get a estudiante
-   `PUT` : Update a estudiante
-   `DELETE` : Delete a estudiante
