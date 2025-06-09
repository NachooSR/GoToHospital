package validations

import (
	"regexp"
	"unicode"
)

var Regex_correo = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

//True si esta vacio
func EmptyField(field string) bool {
	return field == ""
}

func ValidadUsername(email string) bool {
    return Regex_correo.MatchString(email)
}

func ValidarPassword(password string) bool {
	var (
		hasMinLen = len(password) >= 8
		hasUpper  = false
		hasLower  = false
		hasNumber = false
		hasSpecial = false
	)
		
	
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
            hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}