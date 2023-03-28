package database

import (
	"context"
	"github.com/juniorcarrillo/support-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func (db *DB) SaveTv(input *model.NewTv) *model.Tv {

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
		ID:     resTv.InsertedID.(primitive.ObjectID).Hex(),
		Type:   input.Type,
		Model:  input.Model,
		Brand:  input.Brand,
		Serial: input.Serial,
		Owner:  input.Owner,
		User:   &user,
	}
}

// Buscar un tv por _id
func (db *DB) FindByIDTv(ID string) *model.Tv {

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
func (db *DB) FindByTv(ATT, VAL string) []*model.Tv {

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
func (db *DB) AllTv() []*model.Tvs {

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
