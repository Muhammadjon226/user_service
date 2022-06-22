package models

//ID ...
type ID struct {
	ID string `json:"id"`
}

//User ...
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

//ListUserRequest ...
type ListUserRequest struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}

//ListUserResponse ...
type ListUserResponse struct {
	Users []*User `json:"users"`
	Count int64   `json:"count"`
}
