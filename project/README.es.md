[Español](./project/README.es.md) | [Inglés](../README.md)

# GOlerplate

GOlerplate es una plantilla para aplicaciones web en Golang. Esto significa que es punto de partida para aplicaciones web en Golang, diseñado para facilitar el desarrollo con una arquitectura limpia y modular.

## Tecnologías

- **Lenguaje**: Go
- **Frameworks y Librerías**: Gorilla Mux, GORM
- **Contenedores**: Docker, Docker Compose
- **Documentación**: Swagger

## Comenzando

### Requisitos previos

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/doc/install) (opcional si deseas utilizar el proyecto de forma local)

### Bajate el repositorio

1. Clona el repositorio:

    ```sh
    git clone https://github.com/Figaarillo/golerplate.git
    cd golerplate
    ```

2. Copia el archivo `.env.example` a `.env`:

    ```sh
    cp .env.example .env
    ```

### ¿Cómo usar?

- Ejecutar el servidor y la base de datos a traves de Docker

```sh
make docker.run
```

- Ejecutar el servidor localmente y la base de datos con Docker

```sh
make run
```

### ¿Cómo ejecutar los test?

#### Test Unitarios

Para ejecutar todos los tests unitarios:

```sh
make test.unit
```

Para ejecutar un solo test unitario:

```sh
make test.unit.[entity_name]
```

Por ejemplo, para ejecutar el test unitario de categoría:

```sh
make test.unit.category
```

#### Test de integración

Para ejecutar todos los tests de integración:

```sh
make test.e2e
```

Para ejecutar un solo test de integración:

```sh
make test.e2e.[entity_name]
```

Por ejemplo, para ejecutar el test de integración de categoría:

```sh
make test.e2e.category
```

## Estructura del proyecto

```sh
.
├── cmd
│   └── api                # Entrada principal de la aplicación
├── docs                   # Documentación generada por Swagger
├── internal
│   ├── application
│   │   └── usecase        # Casos de uso de la aplicación
│   ├── domain
│   │   ├── entity         # Definición de entidades del dominio
│   │   ├── exception      # Manejo de excepciones del dominio
│   │   └── repository     # Interfaces de repositorios
│   ├── infrastructure
│   │   ├── handler        # Manejadores de HTTP
│   │   ├── middleware     # Middlewares HTTP
│   │   ├── repository     # Implementaciones de repositorios
│   │   └── router         # Definición de rutas
│   ├── setup              # Configuración inicial
│   ├── share
│   │   ├── config         # Configuración compartida
│   │   ├── exception      # Manejo de excepciones compartidas
│   │   ├── utils          # Utilidades compartidas
│   │   └── validation     # Validaciones compartidas
│   └── test               # Pruebas unitarias y de integración
└── scripts                # Scripts para automatización
```

## Características

- **API REST**: Implementación de CRUD para entidades.
- **Pruebas**: Pruebas unitarias e integradas.
- **Documentació**n: Documentación de API con Swagger.
- **Arquitectura**: Implementa principios de arquitectura limpia y hexagonal.

<!--## Contribuir-->
<!---->
<!--Si deseas contribuir al proyecto, por favor, sigue estos pasos:-->
<!---->
<!--1. Haz un fork del repositorio.-->
<!--2. Crea una rama (git checkout -b feature/nueva-feature).-->
<!--3. Realiza tus cambios (git commit -am 'Agrega nueva feature').-->
<!--4. Haz push a la rama (git push origin feature/nueva-feature).-->
<!--4. Crea un nuevo Pull Request.-->

## License

This project is licensed under the MIT License.
