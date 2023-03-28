package database

import (
	"context"
	"github.com/juniorcarrillo/support-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"math/rand"
	"strconv"
	"time"
)

func (db *DB) SaveTicket(input *model.AddTicket) *model.Ticket {

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
		Title:       input.Title,
		Owner:       input.Owner,
		Status:      "Pentiente",
		Technical:   technical.ID,
		Equipment:   input.Equipment,
		Description: input.Description,
		Timestamp:   timestamp,
	})
	if err != nil {
		panic(err)
	}

	// Retorno
	return &model.Ticket{
		ID:          resTickets.InsertedID.(primitive.ObjectID).Hex(),
		Status:      "Pendiente",
		Equipment:   input.Equipment,
		Technical:   technical.ID,
		Timestamp:   timestamp,
		Description: input.Description,
		Title:       input.Title,
		Tv:          &tv,
	}
}

func (db *DB) FindByIDTicket(ID string) *model.Ticket {

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
		ID:          ticket.ID,
		Equipment:   ticket.Equipment,
		Technical:   ticket.Technical,
		Status:      ticket.Status,
		Timestamp:   ticket.Timestamp,
		Description: ticket.Description,
		Title:       ticket.Title,
		Tv:          &tv,
	}
}

func (db *DB) FindByTicket(ATT, VAL string) []*model.Tickets {

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

func (db *DB) AllTicket() []*model.Tickets {

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
		ID:          ticket.ID,
		Equipment:   ticket.Equipment,
		Technical:   ticket.Technical,
		Status:      ticket.Status,
		Timestamp:   ticket.Timestamp,
		Description: ticket.Description,
		Title:       ticket.Title,
		Tv:          &tv,
	}
}
