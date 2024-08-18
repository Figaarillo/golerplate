[Español](./README.es.md) | [Inglés](./README.md)

# GOlerplate

GOlerplate es una plantilla para aplicaciones web en Golang. Esto significa que es punto de partida para aplicaciones web en Golang, diseñado para facilitar el desarrollo con una arquitectura limpia y modular.

## Tecnologías

- [gorilla/mux](https://github.com/gorilla/mux): HTTP router para construir servidores web con 🦍.
- [gorm](https://gorm.io): Librería de ORM para Golang.
- [Docker](https://www.docker.com) & [Docker Compose](https://docs.docker.com/compose/): Para la contenedoraización.
- [Swagger](https://swagger.io): Para la documentación de la API.

## Características

- **API REST**: Implementación de CRUD para entidades.
- **Tests**: Pruebas unitarias e integradas, para validar el correcto funcionamiento de la aplicación.
- **Documentació**n: Documentación de la API generada por Swagger.
- **Arquitectura**: Implementación de principios de arquitectura limpia y hexagonal.
- **Dockerización**: Dockerización de la aplicación y la base de datos, para ejecutar la aplicación de agnosticamente.
- **Automatización**: Automatización ciertas tareas mediante comandos personalizados, como ejecutar la aplicación y la base de datos en Docker, o ejecutar pruebas unitarios.

## Primeros pasos

### Requisitos previos

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/doc/install) (opcional si deseas utilizar el proyecto de forma local)

### Clona el repositorio

1. Clonar el repositorio:

```sh
git clone https://github.com/Figaarillo/golerplate.git
cd golerplate
```

2. Setear las variables de entorno. Copia el archivo `.env.example` a `.env`

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

- Ejecutar la compilación de la aplicación localmente y la base de datos con Docker

```sh
make run.build
```

### ¿Cómo ejecutar los test?

#### Test Unitarios

- Para ejecutar todos los tests unitarios:

```sh
make test.unit
```
- Si desea ver el coverage de los tests unitarios:

```sh
make test.unit.cover
```

-  Para ejecutar un solo test unitario:

```sh
make test.unit.[entity_name]
```

- Por ejemplo, para ejecutar el test unitario de categoría:

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

### Documentación

La documentación se generará usando [Swagger](https://swagger.io).

Para generar la documentación:

```sh
make docs
```

Y ahora puedes ver la generada en la ruta *http://localhost:8080/api/docs/*

## Estructura del proyecto

Para la realización de este proyecto, decidí utilizar una arquitectura hexagonal basada en los principios de la Clean Architecture.
Bajo esta estructura, el proyecto se compone de tres capas principales por cada entidad: Dominio, Aplicación y Infraestructura.

La capa de Dominio contiene las entidades de dominio, las excepciones de dominio, interfaces de los repositorios y toda la lógica relacionada con el dominio. La capa de Aplicación contiene los casos de uso de la aplicación y la capa de Infraestructura contiene los componentes externos de la aplicación, como el servidor HTTP y la base de datos.

```sh
.
├── cmd
│   └── api                # Entrada principal de la aplicación
├── docs                   # Documentación generada por Swagger
├── internal
│   ├── application
│   │   └── usecase        # Casos de uso de la aplicación
│   ├── bootstrap          # Inicialización de cada entidad
│   ├── domain
│   │   ├── entity         # Definición de entidades del dominio y pruebas unitarios
│   │   ├── exception      # Manejo de excepciones del dominio
│   │   └── repository     # Interfaces de repositorios
│   ├── infrastructure
│   │   ├── handler        # Manejadores HTTP
│   │   ├── middleware     # Middlewares HTTP
│   │   ├── repository     # Implementaciones de repositorios
│   │   └── router         # Definición de rutas
│   ├── setup              # Configuración inicial
│   ├── share
│   │   ├── config         # Configuración compartida
│   │   ├── exception      # Manejo de excepciones compartidas
│   │   ├── utils          # Utilidades compartidas
│   │   └── validation     # Validaciones compartidas
│   └── test               # Pruebas integración
└── scripts                # Scripts para automatización
```

## To-do

- [ ] Agregar manejo de autenticación  integrando OAuth2.
- [ ] Mejorar la performance al momento de compilar la aplicaicón con Docker.
- [ ] Mejorar la documentación de Swagger.
- [ ] Mejorar los test unitarios agregando más casos de prueba.
- [ ] Mejorar el manejo de errores, agregando más excepciones, mensajes de error y etc.
- [ ] Agregar goroutines para mejorar la eficiencia de la aplicación.
- [ ] Agregar manejo de logs.
- [ ] Integrar otros Frameworks de servidor y de Base de Datos

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
