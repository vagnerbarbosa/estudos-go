package models

type Aluno struct {
	ID   uint   `json:"id"`
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Alunos = []Aluno{}
