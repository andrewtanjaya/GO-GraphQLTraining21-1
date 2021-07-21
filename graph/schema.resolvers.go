package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Go-Graphql/graph/generated"
	"Go-Graphql/graph/model"
	"context"
	"github.com/pkg/errors"
)

func (r *mutationResolver) CreateFood(ctx context.Context, input model.NewFood) (*model.Food, error) {
	if input.Price < 1000 {
		return nil, errors.New("price must be more than 1000")
	}
	food := model.Food{
		Name: input.Name,
		Description: input.Description,
		Price: input.Price,
	}
	_, err := r.DB.Model(&food).Insert()
	if err != nil {
		return nil, errors.New("insert new food error")
	}
	return &food, nil
}

func (r *mutationResolver) UpdateFood(ctx context.Context, id string, input model.NewFood) (*model.Food, error) {
	var food model.Food
	if input.Price < 1000 {
		return nil, errors.New("price must be more than 1000")
	}

	err := r.DB.Model(&food).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("food with id provided not found")
	}

	food.Price = input.Price
	food.Description = input.Description
	food.Name = input.Name

	_ , error := r.DB.Model(&food).Where("id = ? ", id).Update()
	if error != nil {
		return nil, errors.New("failed update food")
	}
	return &food, nil

}

func (r *mutationResolver) DeleteFood(ctx context.Context, id string) (bool, error) {
	var food model.Food
	err := r.DB.Model(&food).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("food with id provided not found")
	}
	_ , error := r.DB.Model(&food).Where("id = ? ", id).Delete()
	if error != nil {
		return false, errors.New("failed delete food")
	}
	return true, nil
}

func (r *queryResolver) Foods(ctx context.Context) ([]*model.Food, error) {
	var foods []*model.Food
	err := r.DB.Model(&foods).Select()
	if err!= nil {
		return nil, errors.New("get foods failed")
	}
	return foods, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
