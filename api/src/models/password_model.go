package models

// Representa o formato da requisição de alteração de senha
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
