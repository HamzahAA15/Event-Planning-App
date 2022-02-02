package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"sirclo/entities"
	"sirclo/entities/model"
	"sirclo/util/graph/generated"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.SuccessResponse, error) {
	var user entities.User
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	fmt.Println(user)
	err := r.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil membuat user"
	return &response, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}
	fmt.Println("id = ", id)
	err := r.userRepo.DeleteUser(id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil menghapus user"
	return &response, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, edit model.NewUser) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}
	var user entities.User
	user.Name = edit.Name
	user.Email = edit.Email
	user.Password = edit.Password
	err := r.userRepo.EditUser(user, id)
	if err != nil {
		return nil, err
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil mengedit user"
	return &response, nil
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*model.LoginResponse, error) {
	token, user, err := r.authRepo.Login(email, password)
	if err != nil {
		return nil, errors.New("yang bener inputnya cuk")
	}
	var hasil model.LoginResponse
	hasil.Code = 200
	hasil.Message = "selamat anda berhasil login"
	hasil.Token = token
	var penampung model.User
	penampung.ID = &user.Id
	penampung.Name = user.Name
	penampung.Email = user.Email
	penampung.Password = user.Password
	hasil.User = &penampung
	return &hasil, nil
}

func (r *queryResolver) GetUser(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
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
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) RentBook(ctx context.Context, id *int) (*model.SuccessResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) ReturnBook(ctx context.Context, id *int) (*model.SuccessResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.SuccessResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (*model.SuccessResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) UpdateBook(ctx context.Context, id int, edit model.NewBook) (*model.SuccessResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}

	responseData, err := r.userRepo.GetUsers()

	if err != nil {
		return nil, errors.New("not found")
	}

	userResponseData := []*model.User{}

	for _, v := range responseData {
		theId := int(v.Id)
		userResponseData = append(userResponseData, &model.User{ID: &theId, Name: v.Name, Email: v.Email, Password: v.Password})
	}

	return userResponseData, nil
}
