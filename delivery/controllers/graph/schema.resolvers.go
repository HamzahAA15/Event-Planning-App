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

func (r *mutationResolver) CreateParticipant(ctx context.Context, eventID int) (*model.SuccessResponse, error) {
	// dataLogin := ctx.Value("EchoContextKey")
	// if dataLogin == nil {
	// 	return nil, errors.New("unauthorized")
	// } else {
	// 	convId := ctx.Value("EchoContextKey")
	// 	fmt.Println("id user", convId)
	// }
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteParticipant(ctx context.Context, eventID int) (*model.SuccessResponse, error) {
	// dataLogin := ctx.Value("EchoContextKey")
	// if dataLogin == nil {
	// 	return nil, errors.New("unauthorized")
	// } else {
	// 	convId := ctx.Value("EchoContextKey")
	// 	fmt.Println("id user", convId)
	// }
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, eventID int, comment string) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	}
	loginId := dataLogin.(int)
	var commentData entities.Comment
	commentData.UserId = loginId
	commentData.EventId = eventID
	commentData.Comment = comment

	err := r.commentRepo.CreateComment(commentData)
	if err != nil {
		return nil, err
	}

	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil membuat comment"
	return &response, nil
}

func (r *mutationResolver) EditComment(ctx context.Context, commentID int, eventID int, comment string) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	}
	loginId := dataLogin.(int)
	var editComment entities.Comment
	editComment.Id = commentID
	editComment.UserId = loginId
	editComment.EventId = eventID
	editComment.Comment = comment

	err := r.commentRepo.EditComment(editComment)
	if err != nil {
		return nil, err
	}

	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil merubah comment"
	return &response, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, commentID int, eventID int) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	} else {
		convId := ctx.Value("EchoContextKey")
		fmt.Println("id user", convId)
	}
	err := r.commentRepo.DeleteComment(eventID, commentID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil menghapus comment"
	return &response, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.SuccessResponse, error) {
	var user entities.User
	user.Name = input.Name
	user.Email = input.Email
	password := input.Password
	user.Password, _ = entities.EncryptPassword(password)
	fmt.Println(user)
	// ini function buat bikin user dengan sebuah inputan entitites.User
	err := r.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil membuat user"
	return &response, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	}
	userID := dataLogin.(int)
	fmt.Println("id = ", userID)
	err := r.userRepo.DeleteUser(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil menghapus user"
	return &response, nil
}

func (r *mutationResolver) EditUser(ctx context.Context, edit model.EditUser) (*model.SuccessResponse, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	}
	userID := dataLogin.(int)
	var user entities.User
	user.Name = *edit.Name
	user.Email = *edit.Email
	user.Password, _ = entities.EncryptPassword(*edit.Password)

	err := r.userRepo.EditUser(user, userID)
	if err != nil {
		return nil, err
	}
	var response model.SuccessResponse
	response.Code = 200
	response.Message = "berhasil mengedit user"
	return &response, nil
}

func (r *queryResolver) Login(ctx context.Context, email string, password string) (*model.LoginResponse, error) {
	hashedPassword, err_checkdata := r.authRepo.GetEncryptPassword(email)
	if err_checkdata != nil {
		return nil, errors.New("email tidak ditemukan")
	}
	err_compare := entities.ComparePassword(hashedPassword, password)
	if err_compare != nil {
		return nil, errors.New("password salah")
	}
	token, user, err := r.authRepo.Login(email)
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
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
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

func (r *queryResolver) GetUser(ctx context.Context, userID int) (*model.User, error) {
	// dataLogin := ctx.Value("EchoContextKey")
	// if dataLogin == nil {
	// 	return nil, errors.New("unauthorized")
	// } else {
	// 	convId := ctx.Value("EchoContextKey")
	// 	fmt.Println("id user", convId)
	// }
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetParticipants(ctx context.Context, eventID int) ([]*model.User, error) {
	dataLogin := ctx.Value("EchoContextKey")
	if dataLogin == nil {
		return nil, errors.New("unauthorized")
	}

	responseData, err := r.participantRepo.GetParticipants(eventID)

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

func (r *queryResolver) GetComments(ctx context.Context, eventID int) ([]*model.Comment, error) {
	responseData, err := r.commentRepo.GetComments(eventID)
	if err != nil {
		return nil, err
	}

	commentResponseData := []*model.Comment{}

	for _, v := range responseData {
		var user model.User
		user.ID = &v.User.Id
		user.Name = v.User.Name
		user.Email = v.User.Email
		commentResponseData = append(commentResponseData, &model.Comment{ID: v.Id, User: &user, Comment: v.Comment})
	}
	return commentResponseData, nil
}

func (r *queryResolver) GetComment(ctx context.Context, commentID int) (*model.Comment, error) {
	responseData, err := r.commentRepo.GetComment(commentID)
	if err != nil {
		return nil, err
	}

	var commentResponseData model.Comment
	commentResponseData.ID = responseData.Id
	commentResponseData.Comment = responseData.Comment
	var user model.User
	user.ID = &responseData.User.Id
	user.Name = responseData.User.Name
	user.Email = responseData.User.Email
	commentResponseData.User = &user

	return &commentResponseData, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
