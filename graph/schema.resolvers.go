package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/juniorcarrillo/support-graphql/database"
	"github.com/juniorcarrillo/support-graphql/graph/generated"
	"github.com/juniorcarrillo/support-graphql/graph/model"
	_ "go.mongodb.org/mongo-driver/mongo"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return db.SaveUser(&input), nil
}

func (r *mutationResolver) CreateTv(ctx context.Context, input model.NewTv) (*model.Tv, error) {
	return db.SaveTv(&input), nil
}

func (r *mutationResolver) CreateTicket(ctx context.Context, input model.AddTicket) (*model.Ticket, error) {
	return db.SaveTicket(&input), nil
}

func (r *mutationResolver) CreateMessage(ctx context.Context, input model.AddMessage) (*model.Message, error) {
	return db.SaveMessage(&input), nil
}

func (r *mutationResolver) UpdateTicket(ctx context.Context, id string, att string, val string) (*model.Ticket, error) {
	return db.UpdateTicket(id, att, val), nil
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*model.Login, error) {
	return db.LoginUser(email, password), nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return db.FindByIDUser(id), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return db.AllUser(), nil
}

func (r *queryResolver) UsersBy(ctx context.Context, att string, val string) ([]*model.User, error) {
	return db.FindByUser(att, val), nil
}

func (r *queryResolver) Tv(ctx context.Context, id string) (*model.Tv, error) {
	return db.FindByIDTv(id), nil
}

func (r *queryResolver) Tvs(ctx context.Context) ([]*model.Tvs, error) {
	return db.AllTv(), nil
}

func (r *queryResolver) TvsBy(ctx context.Context, att string, val string) ([]*model.Tv, error) {
	return db.FindByTv(att, val), nil
}

func (r *queryResolver) Ticket(ctx context.Context, id string) (*model.Ticket, error) {
	return db.FindByIDTicket(id), nil
}

func (r *queryResolver) Tickets(ctx context.Context) ([]*model.Tickets, error) {
	return db.AllTicket(), nil
}

func (r *queryResolver) TicketsBy(ctx context.Context, att string, val string) ([]*model.Tickets, error) {
	return db.FindByTicket(att, val), nil
}

func (r *queryResolver) Message(ctx context.Context, id string) (*model.Message, error) {
	return db.FindByIDMessage(id), nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return db.AllMessage(), nil
}

func (r *queryResolver) MessagesBy(ctx context.Context, att string, val string) ([]*model.Message, error) {
	return db.FindByMessage(att, val), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Db()
