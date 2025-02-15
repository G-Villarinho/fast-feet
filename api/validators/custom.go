package validators

import (
	"github.com/go-playground/validator/v10"
)

const (
	CPFTag = "cpf"
)

func SetupCustomValidations(validator *validator.Validate) error {
	if err := validator.RegisterValidation(CPFTag, cpfValidator); err != nil {
		return err
	}

	return nil
}

func cpfValidator(fl validator.FieldLevel) bool {
	cpf := fl.Field().String()

	digits := []int{}
	for _, r := range cpf {
		if r >= '0' && r <= '9' {
			digits = append(digits, int(r-'0'))
		}
	}

	if len(digits) != 11 {
		return false
	}

	allEqual := true
	for i := 1; i < len(digits); i++ {
		if digits[i] != digits[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	var firstCheck int
	if sum%11 < 2 {
		firstCheck = 0
	} else {
		firstCheck = 11 - (sum % 11)
	}
	if digits[9] != firstCheck {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += digits[i] * (11 - i)
	}
	var secondCheck int
	if sum%11 < 2 {
		secondCheck = 0
	} else {
		secondCheck = 11 - (sum % 11)
	}
	if digits[10] != secondCheck {
		return false
	}

	return true
}
