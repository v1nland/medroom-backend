# MEDROOM BACKEND

A simple RESTful API compatible with heroku using:

1. **Gin Framework**
2. **Gorm**
3. **Go 1.13**

## Installation & Run

```bash
# Download this project
$ git clone github.com/v1nland/medroom-backend
```

To run locally, generate .env file and modify its variables

```bash
cp .env.example .env
```

Install dependencies

```bash
$ dep ensure
```

Build swagger

```bash
$ swag init
```

## API

#### /estudiantes

-   `GET` : Get all estudiantes
-   `POST` : Create a new estudiante

#### /estudiantes/:uuid

-   `GET` : Get a estudiante
-   `PUT` : Update a estudiante
-   `DELETE` : Delete a estudiante
