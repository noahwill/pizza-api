package helpers

import (
	"errors"
	"pizza-api/pkg/types"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gofrs/uuid"
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

// ValidateCreateAccountInput : validates CreateAccountInput and constructs an account object to create
func ValidateCreateAccountInput(in *types.CreateAccountInput) (*types.Account, error) {
	account := &types.Account{}

	if in.Address == nil {
		return account, errors.New("Specify an Address")
	}

	address, err := validateAddress(in.Address)
	if err != nil {
		return account, err
	}

	if in.Email == nil {
		return account, errors.New("Specify an Email")
	}
	if err := validateEmail(*in.Email); err != nil {
		return account, err
	}

	if in.FirstName == nil {
		return account, errors.New("Specify a FirstName")
	} else if !alphabetic(strings.TrimSpace(*in.FirstName)) {
		return account, errors.New("FirstName must be alphabetic")
	}

	if in.LastName == nil {
		return account, errors.New("Specify a LastName")
	} else if !alphabetic(strings.TrimSpace(*in.LastName)) {
		return account, errors.New("LastName must be alphabetic")
	}

	if in.Password == nil {
		return account, errors.New("Specify a Password")
	} else if !alphabetic(strings.TrimSpace(*in.Password)) { // TODO : let's validate for stronger passwords, maybe learn some stuff about security, huh bud?
		return account, errors.New("Password must be alphabetic")
	}

	uu, _ := uuid.NewV4()
	account = &types.Account{
		Active:      true,
		Address:     *address,
		CreatedAt:   time.Now().Unix(),
		Email:       strings.ToLower(strings.TrimSpace(*in.Email)),
		FirstName:   strings.TrimSpace(*in.FirstName),
		LastName:    strings.TrimSpace(*in.LastName),
		LastUpdated: time.Now().Unix(),
		Orders:      []string{},
		Password:    strings.TrimSpace(*in.Password),
		UUID:        uu.String(),
	}

	return account, nil
}

// ValidateUpdateAccountInput : validates UpdateAccountInput and updates the given account object accordingly
func ValidateUpdateAccountInput(in *types.UpdateAccountInput, account *types.Account) (*types.Account, error) {
	if in.Active != nil {
		account.Active = *in.Active
	}

	if in.Address != nil {
		address, err := validateAddress(in.Address)
		if err != nil {
			return account, err
		}

		account.Address = *address
	}

	if in.Email != nil {
		if err := validateEmail(*in.Email); err != nil {
			return account, err
		}

		account.Email = strings.ToLower(strings.TrimSpace(*in.Email))
	}

	if in.FirstName != nil {
		if !alphabetic(strings.TrimSpace(*in.FirstName)) {
			return account, errors.New("FirstName must be alphabetic")
		}

		account.FirstName = strings.TrimSpace(*in.FirstName)
	}

	if in.LastName != nil {
		if !alphabetic(strings.TrimSpace(*in.LastName)) {
			return account, errors.New("LastName must be alphabetic")
		}

		account.LastName = strings.TrimSpace(*in.LastName)
	}

	if in.Order != nil {
		// TODO : search for order in orders DB to validate it exists
	}

	if in.Password != nil {
		if !alphabetic(strings.TrimSpace(*in.Password)) {
			return account, errors.New("Password must be alphabetic")
		}

		account.Password = strings.TrimSpace(*in.Password)
	}

	return account, nil
}
