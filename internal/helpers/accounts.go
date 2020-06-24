package helpers

import (
	"errors"
	"pizza-api/pkg/types"
	"strings"

	"github.com/badoux/checkmail"
)

func validateEmail(email string) error {
	e := strings.ToLower(strings.TrimSpace(email))

	if !strings.Contains(email, "@") {
		return errors.New("Email is invalid")
	}

	// Does email exist?
	err := checkmail.ValidateHost(e)
	if smtpErr, ok := err.(checkmail.SmtpError); ok && err != nil {
		if smtpErr.Code() == "550" {
			return errors.New("Email does not exist")
		}
	}

	// Does email have a valid format?
	if err = checkmail.ValidateFormat(email); err != nil {
		return errors.New("Email has an invalid format")
	}

	// Does email have a resolvable host?
	if err = checkmail.ValidateHost(email); err != nil {
		if err.Error() == "unresolvable host" {
			return errors.New("Email has an unresolvable host")
		}
	}
	return nil
}

// ValidateCreateAccountInput : validates CreateAccountInput
func ValidateCreateAccountInput(in *types.CreateAccountInput) error {
	if in.Email == nil {
		return errors.New("Specify an Email")
	}
	if err := validateEmail(*in.Email); err != nil {
		return err
	}

	if in.FirstName == nil {
		return errors.New("Specify a FirstName")
	} else if !IsOnlyAlphabetic(*in.FirstName) {
		return errors.New("FirstName must be alphabetic")
	}

	if in.LastName == nil {
		return errors.New("Specify a LastName")
	} else if !IsOnlyAlphabetic(*in.LastName) {
		return errors.New("LastName must be alphabetic")
	}

	if in.Password == nil {
		return errors.New("Specify a Password")
	} else if !IsOnlyAlphabetic(*in.Password) { // TODO : let's validate for stronger passwords, maybe learn some stuff about security, huh bud?
		return errors.New("Password must be alphabetic")
	}

	return nil
}
