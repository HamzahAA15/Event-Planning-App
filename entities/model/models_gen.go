// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID      int    `json:"id"`
	Comment string `json:"comment"`
	User    *User  `json:"user"`
}

type EditEvent struct {
	CategoryID  *int    `json:"categoryId"`
	Title       *string `json:"title"`
	Host        *string `json:"host"`
	Date        *string `json:"date"`
	Location    *string `json:"location"`
	Description *string `json:"description"`
	ImageURL    *string `json:"imageUrl"`
}

type EditUser struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Image    *string `json:"image"`
}

type Event struct {
	ID          *int    `json:"id"`
	UserID      int     `json:"userId"`
	CategoryID  int     `json:"categoryId"`
	Title       string  `json:"title"`
	Host        string  `json:"host"`
	Date        string  `json:"date"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	ImageURL    *string `json:"imageUrl"`
}

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
	User    *User  `json:"User"`
}

type NewEvent struct {
	UserID      int     `json:"userId"`
	CategoryID  int     `json:"categoryId"`
	Title       string  `json:"title"`
	Host        string  `json:"host"`
	Date        string  `json:"date"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	ImageURL    *string `json:"imageUrl"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type User struct {
	ID       *int    `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Image    *string `json:"image"`
}
