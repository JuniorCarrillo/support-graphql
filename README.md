El siguiente es un instructivo adicional al que ya presenta en sí la plataforma GraphQL de manera predeterminada en la interface primaria de ejecución, la cual contiene documentación especifica y detallada de ejecución, y en cuyo caso se debe poder por parte de quien emplemente el sistema en caso de ser necesario.

# SupportGraphQL v1
Esta es una API GraphQL para el uso como Back End, el cual se encuentra desarrollado en Go y MongoDB Atlas en plataformas que requieran soporte de tickets, clientes y componentes electrónicos. Recomendado para tiendas de electrónicos que requieren soporte a clientes con implementación de CORS ajustable.

Está plataforma corre sobre el puerto `:8080` el cual es modificable añadiendo la variable del entorno `port` al OS que lo ejecute o modificando el archivo `server.go` en la raíz del sistema, así como la modificación de uso en modo CORS actualmente predefinido.

### Estructura del sistema
Aquí se muestra la organización y componentes de estruturales, arquitectura y organización especifica del repositorio.

#### Archivos del repositorio
Dentro de este repositorio se encuentran un total de 10 archivos y 4 directorios. Los cuales se muestran en el siguiente mapa estructural:
```
.
|-- README.md
|-- database
|   |-- message.go
|   |-- ticket.go
|   |-- tv.go
|   `-- user.go
|-- go.mod
|-- graph
|   |-- generated
|   |   `-- generated.go
|   |-- model
|   |   `-- models_gen.go
|   |-- resolver.go
|   |-- schema.graphqls
|   `-- schema.resolvers.go
`-- server.go
```

#### Librerías del sistema
Está plaforma fué desarrollada en `go 1.15` con una configuración `modules`, haciendo uso de los siguientes componentes o librerías:
```
require (
	github.com/99designs/gqlgen v0.13.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/cors v1.7.0
	github.com/vektah/gqlparser v1.3.1 // indirect
	github.com/vektah/gqlparser/v2 v2.1.0
	go.mongodb.org/mongo-driver v1.4.3
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
)
```

#### Schema de la API GraphQL
Dentro de la información se implementa el siguiente esquema de peticiones y mutaciones:
```
# Output
type Login {
  _id:    ID!
  status: Boolean!
  message: String!
  user:   User!
}

type User {
  _id: ID!
  name: String!
  address: String!
  phone: String!
  email: String!
  role: String!
}

type UserAuth {
  _id: ID!
  name: String!
  address: String!
  phone: String!
  email: String!
  role: String!
  password: String!
}

type Tv {
  _id: ID!
  type: String!
  model: String!
  brand: String!
  serial: String!
  owner: String!
  user: User!
}

type Tvs {
  _id: ID!
  type: String!
  model: String!
  brand: String!
  serial: String!
  owner: String!
}

type Ticket {
  _id: ID!
  title: String!
  status: String!
  equipment: String!
  technical: String!
  timestamp: String!
  description: String!
  tv: Tv!
}

type Tickets {
  _id: ID!
  title: String!
  owner: String!
  status: String!
  technical: String!
  equipment: String!
  timestamp: String!
  description: String!
}

type Message {
  _id: ID!
  author: String!
  ticket: String!
  content: String!
  timestamp: String!
}

# Input
input NewUser {
  name: String!
  address: String!
  phone: String!
  password: String!
  email: String!
  role: String!
}

input NewTv {
  type: String!
  model: String!
  brand: String!
  serial: String!
  owner: String!
}

input NewTicket {
  title: String!
  owner: String!
  status: String!
  technical: String!
  equipment: String!
  description: String!
  timestamp: String!
}

input AddTicket {
  title: String!
  owner: String!
  equipment: String!
  description: String!
}

input NewMessage {
  author: String!
  ticket: String!
  content: String!
  timestamp: String!
}

input AddMessage {
  author: String!
  ticket: String!
  content: String!
}

# Querys
type Query {
  login(email: String!, password: String!): Login!
  user(_id: String!): User!
  users: [User!]!
  usersBy(att: String!, val: String!): [User!]!
  tv(_id: String!): Tv!
  tvs: [Tvs!]!
  tvsBy(att: String!, val: String!): [Tv!]!
  ticket(_id: String!): Ticket!
  tickets: [Tickets!]!
  ticketsBy(att: String!, val: String!): [Tickets!]!
  message(_id: String!): Message!
  messages: [Message!]!
  messagesBy(att: String!, val: String!): [Message!]!
}

# Mutations
type Mutation {
  createUser(input: NewUser!): User!
  createTv(input: NewTv!): Tv!
  createTicket(input: AddTicket!): Ticket!
  createMessage(input: AddMessage!): Message!
  updateTicket(_id: String!, att: String!, val: String!): Ticket!
}
```

## Interacciones disponibles
Las interacciones realizadas en el sistema se pueden realizar por medio de Postman o desde la interface home de la API, en este caso todas las interacciones se realizan directamente desde el home de la API, con la intención de que sea más flexible y requiera menos herramientas su funcionamiento. El funcionamiento de este sistema posee las siguientes interacciones disponibles:

### Mutations o mutaciones
Para la carga o guardado y, actualización de información dentro del sistema en la plataforma:

##### Carga o guardado de información
En estas mutaciones se encuentran las interacciones que ejecutan la creación de los elementos en la plataforma.

---
**NOTA**

Dentro de la creación del usuario la contraseña se encripta y se guarda de esa manera en la base de datos, luego de eso para realizar la autenticación se debe compara mediante un sistema de autenticación simple de usuario y correo electrónico.

---
```
mutation AddUser {
  createUser(
    input: {
      name: "Junior Carrillo"
      address: "Medellin"
      phone: "3003328389"
      password: "123456"
      email: "soyjrcarrillo@gmail.com"
      role: "user"
    }
  ) {
    _id
    name
  }
}

mutation AddTV {
  createTv(
    input: {
      type: "LCD"
      model: "DEMO-1"
      brand: "LG"
      serial: "000-000-000-000"
      owner: "5fa583653f3a3896431173cc"
    }
  ) {
    _id
    serial
    owner
    user {
      name
      email
    }
  }
}

mutation AddTicket {
  createTicket(
    input: {
      title: "El tv se apaga luego de unos minutos"
      owner: "5fa583653f3a3896431173cc"
      equipment: "5fa5bf6c0b9407de5ba89f1c"
      description: "Hola, me gustaría notificar una falla en mi equipo, tengo un problma que al encender tras unos minutos se apaga solo"
    }
  ) {
    _id
    equipment
    technical
    timestamp
    description
    status
    title
    tv {
      _id
      model
      serial
    }
  }
}

mutation AddMessage {
  createMessage(
    input: {
      author: "5fa583653f3a3896431173cc"
      ticket: "5fa64d6042c421ac24bc1381"
      content: "Hola, esto es una prueba"
    }
  ) {
    _id
    author
    ticket
    content
    timestamp
  }
}
```
##### Actualizar
Dentro del sistema solo se puede actualizar la información de los tickets, para lo cual se requiere enviar:
- `_id`: Identificador del ticket, este se utiliza como identificación y puede ser usado para extraer toda la información relacionada con este ticket.
- `att`: Atributo del campo que se requiere cambiar.
- `val`: Valor del atributo que se desea cambiar o actualizar.
```
mutation Update {
  updateTicket(
    _id: "5fa686d9e828d2235f06c6f6"
    att: "status"
    val: "Pendiente"
  ) {
    _id
    title
  }
}
```

### Querys o peticiones
En este punto se muestran las peticiones que se pueden realizar al sistema:

##### Peticiones únicas
Las peticiones únicas son las que se ejecutan directamente y solo esperan un resultado especifico, por ejemplo: solicitud de un usuario, ticket, televisor o mensaje.
```
query User {
  user(_id: "5fa57bb09e60d49b64cc60fe") {
    _id
    name
    email
    address
    role
  }
}

query Tv {
  tv(_id: "5fa5bf6c0b9407de5ba89f1c") {
    _id
    serial
    brand
    owner
    user {
      name
    }
  }
}

query Ticket {
  ticket(_id: "5fa64d6042c421ac24bc1381") {
    _id
    equipment
    technical
    timestamp
    description
    title
    tv {
      _id
      model
      serial
    }
  }
}

query Message {
  message(_id: "5fa66b387c1f99c6895a0b69") {
    _id
    author
    ticket
    content
    timestamp
  }
}
```
##### Peticiones multiples
Estas son peticiones que al ser ejecutadas, reciben dos o más elementos.
```
query AllUsers {
  users {
    _id
    name
    email
    address
    role
  }
}

query UsersBy {
  usersBy(att: "role", val: "user") {
    _id
    name
    email
    address
    role
  }
}

query AllTvs {
  tvs {
    _id
    type
    model
    brand
    serial
    owner
  }
}

query TvsBy {
  tvsBy(att: "type", val: "LCD") {
    _id
    serial
    brand
  }
}

query AllTickets {
  tickets {
    _id
    title
    description
    equipment
    technical
    timestamp
  }
}

query TicketsBy {
  ticketsBy(att: "technical", val: "5fa62e9c07ec161f0fa3d39e") {
    _id
    title
    description
    timestamp
  }
}


query AllMessages {
  messages {
    _id
    author
    ticket
    content
    timestamp
  }
}

query MessagesBy {
  messagesBy(att: "ticket", val: "5fa64d6042c421ac24bc1381") {
    _id
    author
    ticket
    content
    timestamp
  }
}
```
### Autenticación de usuarios
Para el proceso de autenticación se debe realizar un query solicitando la información, para ello se envía `email` y `password`, en lo que se espera recibir un estatus de si el logueo es efectivo, la identificación del usuario, y su información general para el uso de restricciones en el Front End para diferentes roles. La estructura enviada debe ser similar a está:
```
query Auth {
  login(email: "soyjrcarrillo@gmail.com", password: "12345") {
    _id
    status
    user {
      email
      address
      role
    }
  }
}
```

**RESPUESTA EN CASO DE ERROR**

Cuando la contraseña o el correo no son válidos, o no esta registrado en el sistema
```
{
  "data": {
    "login": {
      "_id": "N/A",
      "status": false,
      "message": "La contraseña o el correo electrónico no se encuentran asociados a ningún usuario o están errados",
      "user": {
        "email": "N/A",
        "address": "N/A",
        "role": "N/A"
      }
    }
  }
}
```
**RESPUESTA EN CASO DE ÉXITO**

Cuando el usuario y la contraseña coinciden y están registrados en el sistema.
```
{
  "data": {
    "login": {
      "_id": "5fa583653f3a3896431173cc",
      "status": true,
      "message": "Bienvenido Junior Carrillo",
      "user": {
        "email": "soyjrcarrillo@gmail.com",
        "address": "Medellin",
        "role": "user"
      }
    }
  }
}
```

## Organización de la base de datos
En este proyecto se utiliza MongoDB Atlas como ejemplo practico, de la manera en la que es implementado solo funciona de forma educativa, si requiere un despliegue en producción es necesario optimizar la seguridad guardando la `MONGODB_URI` como una variable de entorno en el sistema.

#### Arquitectura de datos
```
|-- messages (Colección)
|   |-- _id: Identificación de item
|   |-- author: _id del autor del mensaje (String)
|   |-- ticket: _id del ticket padre del mensaje (String)
|   |-- content: contenido del mensaje (String)
|   |-- timestamp: Timestamp en formato Unix (String)
|-- tickets (Colección)
|   |-- _id: Identificación de item
|   |-- title: Título del ticket (String)
|   |-- owner: _id del dueño del ticket (String)
|   |-- status: nombre del estatus del ticket (String)
|   |-- technical: _id del usuario tecnico encargado (String)
|   |-- equipment: _id del televisor (String)
|   |-- description: Descripción del ticket (String)
|   |-- timestamp: Timestamp en formato Unix (String)
|-- tvs (Colección)
|   |-- _id: Identificación de item
|   |-- type: Tipo de televisor cargado (String)
|   |-- model: Modelo del televisor (String)
|   |-- brand: Marca del televisor (String)
|   |-- serial: Serial del televisor (String)
|   |-- owner: _id del usuario dueño del televisor (String)
|-- users (Colección)
|   |-- _id: Identificación de item
|   |-- name: Nombre del usuario (String)
|   |-- address: Dirección del usuario (String)
|   |-- phone: Nro. de teléfono del usuario (String)
|   |-- password: Contraseña del usuario (String [Bajo codificación de autenticación en Go])
|   |-- email: Correo electrónico del usuario (String)
|   |-- role: Rol del usuario (String)
```

### Más información
Esté proyecto fue enteramente desarrollado por Junior Carrillo bajo licencia MIT, es un aporte a la comunidad de Golang Venezuela. Puedes solicitar más información desde:
- +57 300 3328389
- soyjrcarrillo@gmail.com
- [Telegram](https://t.me/imjrcarrillo)
- [Facebook](https://fb.com/soyjrcarrillo)
- [Twitter](https://twitter.com/soyjrcarrillo)
- [Instagram](https://instagram.com/soyjrcarrillo)
