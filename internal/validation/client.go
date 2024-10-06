package validation

import (
	"errors"
	"regexp"

	"github.com/CriticalNoob02/store-control-be/internal/model"
)

var ClientsFilterList = []string{"name", "cpf", "number"}

func ClientValidation(user model.Client) error {
	if user.Name == "" {
		return errors.New("campo de Nome e de preenchimento obrigatorio")
	}

	if user.Cpf == "" {
		return errors.New("campo de Cpf e de preenchimento obrigatorio")
	}

	phonePattern := `^\+\d{2}\s\d{2}\s\d{5}-\d{4}$`
	matched, _ := regexp.MatchString(phonePattern, user.Number)
	if !matched {
		return errors.New("número de telefone informado está no formato inválido")
	}
	return nil
}
