package graph

import (
	"sirclo/repository/auth"
	"sirclo/repository/event"
	"sirclo/repository/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authRepo  auth.Auth
	userRepo  user.User
	eventRepo event.Event
}

func NewResolver(ur user.User, ar auth.Auth, er event.Event) *Resolver {
	return &Resolver{
		authRepo:  ar,
		userRepo:  ur,
		eventRepo: er,
	}
}
