package models

type User struct {
	Username   	string
	Password 	string
	Key         string
	Token       string
}

type InsUser struct {
	Username   	string
	EncUsername []byte
	Password 	[]byte
	Key         []byte
	Token       string
}

