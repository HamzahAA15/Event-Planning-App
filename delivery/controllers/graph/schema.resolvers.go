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
	password := input.Password
	user.Password, _ = entities.HashPassword(password)
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
	// dataLogin := ctx.Value("EchoContextKey")
	// if dataLogin == nil {
	// 	return nil, errors.New("unauthorized")
	// } else {
	// 	convId := ctx.Value("EchoContextKey")
	// 	fmt.Println("id user", convId)
	// }
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
	// dataLogin := ctx.Value("EchoContextKey")
	// if dataLogin == nil {
	// 	return nil, errors.New("unauthorized")
	// } else {
	// 	convId := ctx.Value("EchoContextKey")
	// 	fmt.Println("id user", convId)
	// }
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

	hashedPass, _ := entities.HashPassword(password)
	token, user, err := r.authRepo.Login(email, hashedPass)
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
	hasil.User = &penampung
	return &hasil, nil
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	// dataLogin := ctx.Value("EchoContextKey")
	// if dataLogin == nil {
	// 	return nil, errors.New("unauthorized")
	// } else {
	// 	convId := ctx.Value("EchoContextKey")
	// 	fmt.Println("id user", convId)
	// }

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
