package database

import (
	"context"
	"github.com/juniorcarrillo/SupportGraphQL/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

// Estructura de la base de datos
type DB struct {
	client *mongo.Client
}

// Instancia general
func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://SoyJrCarrillo:CEjp249@cluster0.nkkan.gcp.mongodb.net/SupportGraphQL?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{client: client}
}


// Login
func (db* DB) Login(email, password string) *model.Login {

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	res := collection.FindOne(ctx, bson.M{"email": email})
	user := model.UserAuth{}
	_ = res.Decode(&user)

	// Verificación de contraseña
	status := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil

	// Retorno
	return &model.Login{
		ID: 	user.ID,
		Status: status,
		User:	&model.User{
			ID:      user.ID,
			Name:    user.Name,
			Address: user.Address,
			Phone:   user.Phone,
			Email:   user.Email,
			Role:    user.Role,
		},
	}
}

/**
* Gestión de usuarios
**/

// Crear un usuario
func (db* DB) SaveUser(input *model.NewUser) *model.User {

	// Instancia de guardado
	collection := db.client.Database("SupportGraphQL").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Codificar la contraseña
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		panic(err)
	}

	// Guardar
	res, err := collection.InsertOne(ctx, &model.NewUser{
		Name:     input.Name,
		Address:  input.Address,
		Phone:    input.Phone,
		Password: string(password),
		Email:    input.Email,
		Role:     input.Role,
	})
	if err != nil {
		panic(err)
	}

	// Retorno
	return &model.User{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
		Name: input.Name,
		Address: input.Address,
		Phone: input.Phone,
		Email: input.Email,
		Role: input.Role,
	}
}

// Buscar un usuario
func (db *DB) FindUserByID(ID string) *model.User {

	// Codificar _id en la petición
	UserID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	res := collection.FindOne(ctx, bson.M{"_id": UserID})
	user := model.User{}
	_ = res.Decode(&user)
	return &user
}

// Buscar un usuario por
func (db *DB) FindUserBy(ATT, VAL string) []*model.User {

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	cur, err := collection.Find(ctx, bson.M{ATT: VAL})
	if err != nil {
		panic(err)
	}
	var users []*model.User
	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	// Retorno
	return users
}

// Listar todos los usuarios
func (db *DB) AllUsers() []*model.User {

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	var users []*model.User
	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	// Retorno
	return users
}


/**
* Gestión de televisores
**/

// Crear un TV
func (db* DB) SaveTV(input *model.NewTv) *model.Tv {

	// Codificar _id en la petición
	UserID, err := primitive.ObjectIDFromHex(input.Owner)
	if err != nil {
		panic(err)
	}

	// Instancia de búsqueda
	colTv := db.client.Database("SupportGraphQL").Collection("tvs")
	colUser := db.client.Database("SupportGraphQL").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	res := colUser.FindOne(ctx, bson.M{"_id": UserID})
	user := model.User{}
	_ = res.Decode(&user)

	// Guardar
	resTv, err := colTv.InsertOne(ctx, &input)
	if err != nil {
		panic(err)
	}

	// Retorno
	return &model.Tv{
		ID: resTv.InsertedID.(primitive.ObjectID).Hex(),
		Type: input.Type,
		Model: input.Model,
		Brand: input.Brand,
		Serial: input.Serial,
		Owner: input.Owner,
		User: &user,
	}
}

// Buscar un tv por _id
func (db *DB) FindTvByID(ID string) *model.Tv {

	// Codificar _id en la petición
	TvID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}

	// Instancia de búsqueda
	colTv := db.client.Database("SupportGraphQL").Collection("tvs")
	colUser := db.client.Database("SupportGraphQL").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	resTv := colTv.FindOne(ctx, bson.M{"_id": TvID})
	tv := model.Tv{}
	_ = resTv.Decode(&tv)

	// Cecodificar _id del dueño del tv
	OwnerID, err := primitive.ObjectIDFromHex(tv.Owner)
	if err != nil {
		panic(err)
	}

	// Extraer
	resUser := colUser.FindOne(ctx, bson.M{"_id": OwnerID})
	user := model.User{}
	_ = resUser.Decode(&user)

	// Retorno
	return &model.Tv{
		ID:     tv.ID,
		Type:   tv.Type,
		Model:  tv.Model,
		Brand:  tv.Brand,
		Serial: tv.Serial,
		Owner:  tv.Owner,
		User:   &user,
	}
}

// Buscar tvs por usuario
func (db *DB) FindTvsBy(ATT, VAL string) []*model.Tv {

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("tvs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	cur, err := collection.Find(ctx, bson.M{ATT: VAL})
	if err != nil {
		panic(err)
	}
	var tvs []*model.Tv
	for cur.Next(ctx) {
		var tv *model.Tv
		err := cur.Decode(&tv)
		if err != nil {
			panic(err)
		}
		tvs = append(tvs, tv)
	}

	// Retorno
	return tvs
}

// Listar los tvs
func (db *DB) AllTvs() []*model.Tvs {

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("tvs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	var tvs []*model.Tvs
	for cur.Next(ctx) {
		var tv *model.Tvs
		err := cur.Decode(&tv)
		if err != nil {
			panic(err)
		}
		tvs = append(tvs, tv)
	}

	// Mostrar
	return tvs
}


/**
* Gestión de tickets
**/

// Crear un Ticket
func (db* DB) SaveTicket(input *model.AddTicket) *model.Ticket {

	// Codificar _id
	TvID, err := primitive.ObjectIDFromHex(input.Equipment)
	if err != nil {
		panic(err)
	}

	// Instancia de guardado
	colTickets := db.client.Database("SupportGraphQL").Collection("tickets")
	colUsers := db.client.Database("SupportGraphQL").Collection("users")
	colTvs := db.client.Database("SupportGraphQL").Collection("tvs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer TV
	resTvs := colTvs.FindOne(ctx, bson.M{"_id": TvID})
	tv := model.Tv{}
	_ = resTvs.Decode(&tv)

	// Extraer Tecnicos
	cur, err := colUsers.Find(ctx, bson.M{"role": "technical"})
	if err != nil {
		panic(err)
	}
	var users []*model.User
	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	// Tecnico asignado al azar
	technical := users[rand.Intn(len(users))]

	// Timestamp
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Guardar
	resTickets, err := colTickets.InsertOne(ctx, &model.NewTicket{
		Title:       	input.Title,
		Owner:       	input.Owner,
		Status: 		"Pentiente",
		Technical: 		technical.ID,
		Equipment:  	input.Equipment,
		Description: 	input.Description,
		Timestamp: 		timestamp,
	})
	if err != nil {
		panic(err)
	}

	// Retorno
	return &model.Ticket{
		ID:        		resTickets.InsertedID.(primitive.ObjectID).Hex(),
		Status: 		"Pendiente",
		Equipment: 		input.Equipment,
		Technical: 		technical.ID,
		Timestamp: 		timestamp,
		Description:    input.Description,
		Title:      	input.Title,
		Tv:				&tv,
	}
}

// Buscar por _id
func (db *DB) FindTicketByID(ID string) *model.Ticket {

	// Codificar _id en la petición
	TicketID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}

	// Instancia de búsqueda
	colTicket := db.client.Database("SupportGraphQL").Collection("tickets")
	colTv := db.client.Database("SupportGraphQL").Collection("tvs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer tv
	resTicket := colTicket.FindOne(ctx, bson.M{"_id": TicketID})
	ticket := model.Ticket{}
	_ = resTicket.Decode(&ticket)

	// Codificar _id del duseño en la petición
	EquipmentID, err := primitive.ObjectIDFromHex(ticket.Equipment)
	if err != nil {
		panic(err)
	}

	// Extraer usuario
	resTv := colTv.FindOne(ctx, bson.M{"_id": EquipmentID})
	tv := model.Tv{}
	_ = resTv.Decode(&tv)

	// Retorno
	return &model.Ticket{
		ID:        		ticket.ID,
		Equipment: 		ticket.Equipment,
		Technical: 		ticket.Technical,
		Status: 		ticket.Status,
		Timestamp: 		ticket.Timestamp,
		Description:    ticket.Description,
		Title:      	ticket.Title,
		Tv:				&tv,
	}
}

// Listar todos tickets
func (db *DB) AllTickets() []*model.Tickets {

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("tickets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	var tickets []*model.Tickets
	for cur.Next(ctx) {
		var ticket *model.Tickets
		err := cur.Decode(&ticket)
		if err != nil {
			panic(err)
		}
		tickets = append(tickets, ticket)
	}

	// Retorno
	return tickets
}

// Buscar tickets por
func (db *DB) FindTicketsBy(ATT, VAL string) []*model.Tickets {

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("tickets")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	cur, err := collection.Find(ctx, bson.M{ATT: VAL})
	if err != nil {
		panic(err)
	}
	var tickets []*model.Tickets
	for cur.Next(ctx) {
		var ticket *model.Tickets
		err := cur.Decode(&ticket)
		if err != nil {
			panic(err)
		}
		tickets = append(tickets, ticket)
	}

	// Retorno
	return tickets
}

// Actualizar tickets
func (db *DB) UpdateTicket(ID, ATT, VAL string) *model.Ticket {

	// Codificar _id en la petición
	TicketID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}

	// Instancia de búsqueda
	colTicket := db.client.Database("SupportGraphQL").Collection("tickets")
	colTv := db.client.Database("SupportGraphQL").Collection("tvs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer tv
	resTicket := colTicket.FindOne(ctx, bson.M{"_id": TicketID})
	ticket := model.Ticket{}
	_ = resTicket.Decode(&ticket)

	// Guardar
	_, err = colTicket.UpdateOne(ctx, bson.M{"_id": TicketID}, bson.D{
		{"$set", bson.D{{ATT, VAL}}},
	})
	if err != nil {
		panic(err)
	}

	// Codificar _id del duseño en la petición
	EquipmentID, err := primitive.ObjectIDFromHex(ticket.Equipment)
	if err != nil {
		panic(err)
	}

	// Extraer usuario
	resTv := colTv.FindOne(ctx, bson.M{"_id": EquipmentID})
	tv := model.Tv{}
	_ = resTv.Decode(&tv)

	// Retorno
	return &model.Ticket{
		ID:        		ticket.ID,
		Equipment: 		ticket.Equipment,
		Technical: 		ticket.Technical,
		Status: 		ticket.Status,
		Timestamp: 		ticket.Timestamp,
		Description:    ticket.Description,
		Title:      	ticket.Title,
		Tv:				&tv,
	}
}

/**
* Gestión de mensajes
**/

// Crear un Mensaje
func (db* DB) SaveMessage(input *model.AddMessage) *model.Message {

	// Instancia de guardado
	collection := db.client.Database("SupportGraphQL").Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Timestamp
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// Guardar
	resTickets, err := collection.InsertOne(ctx, &model.NewMessage{
		Author:     input.Author,
		Ticket:     input.Ticket,
		Content: 	input.Content,
		Timestamp: 	timestamp,
	})
	if err != nil {
		panic(err)
	}

	// Retorno
	return &model.Message{
		ID:        	resTickets.InsertedID.(primitive.ObjectID).Hex(),
		Author: 	input.Author,
		Ticket: 	input.Ticket,
		Content: 	input.Content,
		Timestamp: 	timestamp,
	}
}

// Buscar por _id
func (db *DB) FindMessageByID(ID string) *model.Message {

	// Codificar _id en la petición
	MessageID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		panic(err)
	}

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	res := collection.FindOne(ctx, bson.M{"_id": MessageID})
	message := model.Message{}
	_ = res.Decode(&message)

	// Retorno
	return &message
}

// Listar todos mensajes
func (db *DB) AllMessages() []*model.Message {

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	var messages []*model.Message
	for cur.Next(ctx) {
		var message *model.Message
		err := cur.Decode(&message)
		if err != nil {
			panic(err)
		}
		messages = append(messages, message)
	}

	// Retorno
	return messages
}

// Buscar mensajes por
func (db *DB) FindMessagesBy(ATT, VAL string) []*model.Message {

	// Instancia de búsqueda
	collection := db.client.Database("SupportGraphQL").Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Extraer
	cur, err := collection.Find(ctx, bson.M{ATT: VAL})
	if err != nil {
		panic(err)
	}
	var messages []*model.Message
	for cur.Next(ctx) {
		var message *model.Message
		err := cur.Decode(&message)
		if err != nil {
			panic(err)
		}
		messages = append(messages, message)
	}

	// Retorno
	return messages
}
