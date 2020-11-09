package database

import (
	"context"
	"github.com/juniorcarrillo/SupportGraphQL/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"strconv"
	"time"
)

// Estructura
type MessageDB struct {
	client *mongo.Client
}

// Instancia
func Message() *MessageDB {

	// La URI de MongoDB se recomienda guardar como variable de entorno
	// se deje esta como ejemplo y para testeo. Pero por seguridad se recomienda
	// no colocar de esta forma por medida de seguridad en producción.
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://SoyJrCarrillo:CEjp249@cluster0.nkkan.gcp.mongodb.net/SupportGraphQL?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &MessageDB{client: client}

}

// Crear un Mensaje
func (db *MessageDB) Save(input *model.AddMessage) *model.Message {

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
func (db *MessageDB) FindByID(ID string) *model.Message {

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
func (db *MessageDB) All() []*model.Message {

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
func (db *MessageDB) FindBy(ATT, VAL string) []*model.Message {

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
