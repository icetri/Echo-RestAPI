package types

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Users struct {
	Users []User `json:"users"`
}
