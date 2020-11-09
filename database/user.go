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
	"time"
)

// Estructura
type UserDB struct {
	client *mongo.Client
}

// Instancia
func User() *UserDB {

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
	return &UserDB{client: client}

}

// Login
func (db* UserDB) Login(email, password string) *model.Login {

	// Instancias de notificaciones
	var message = ""

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
	if status {
		message = "Bienvenido " + user.Name
	} else {
		message = "La contraseña o el correo electrónico no se encuentran asociados a ningún usuario o están errados"
	}

	// Retorno
	if status {
		return &model.Login{
			ID: 	user.ID,
			Status: status,
			Message: message,
			User:	&model.User{
				ID:      user.ID,
				Name:    user.Name,
				Address: user.Address,
				Phone:   user.Phone,
				Email:   user.Email,
				Role:    user.Role,
			},
		}
	} else {
		return &model.Login{
			ID: 		"N/A",
			Status: 	status,
			Message: 	message,
			User:		&model.User{
				ID:      "N/A",
				Name:    "N/A",
				Address: "N/A",
				Phone:   "N/A",
				Email:   "N/A",
				Role:    "N/A",
			},
		}
	}
}

// Crear un usuario
func (db* UserDB) Save(input *model.NewUser) *model.User {

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
func (db *UserDB) FindByID(ID string) *model.User {

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
func (db *UserDB) FindBy(ATT, VAL string) []*model.User {

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
func (db *UserDB) All() []*model.User {

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
