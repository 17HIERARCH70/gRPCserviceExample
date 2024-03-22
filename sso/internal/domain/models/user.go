package models

type User struct {
	ID       int64
	email    string
	PassHash []byte
	Email    interface{}
}
