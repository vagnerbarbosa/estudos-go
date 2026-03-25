package models

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"required,min=3"`
	CPF  string `json:"cpf" validate:"required,len=11,numeric"`
	RG   string `json:"rg" validate:"required,len=9,numeric"`
}

func ValidateAluno(aluno *Aluno) error {
	validate := validator.New()
	return validate.Struct(aluno)
}
