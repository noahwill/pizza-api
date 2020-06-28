package valid

import (
	"errors"
	"pizza-api/pkg/types"
	"pizza-api/utils"
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

// GetAccountByID : gets an account from the accounts table by ID
func GetAccountByID(accountID string) (*types.Account, error) {
	var account types.Account
	if err := utils.Config.AccountsTableConn.Get("UUID", accountID).One(&account); err != nil {
		return &account, err
	}
	return &account, nil
}

// ValidateCreateAccountInput : validates CreateAccountInput and constructs an account object to create
func ValidateCreateAccountInput(in *types.CreateAccountInput) (*types.Account, error) {
	account := &types.Account{}

	// Validate and set Address
	if in.Address == nil {
		return account, errors.New("Specify an Address")
	}
	address, err := validateAddress(in.Address)
	if err != nil {
		return account, err
	}

	// Validate Email
	if in.Email == nil {
		return account, errors.New("Specify an Email")
	}
	if err := validateEmail(*in.Email); err != nil {
		return account, err
	}

	// Validate FirstName
	if in.FirstName == nil {
		return account, errors.New("Specify a FirstName")
	} else if !alphabetic(strings.TrimSpace(*in.FirstName)) {
		return account, errors.New("FirstName must be alphabetic")
	}

	// Validate LastName
	if in.LastName == nil {
		return account, errors.New("Specify a LastName")
	} else if !alphabetic(strings.TrimSpace(*in.LastName)) {
		return account, errors.New("LastName must be alphabetic")
	}

	// Validate Password
	if in.Password == nil {
		return account, errors.New("Specify a Password")
		// TODO : let's validate for stronger passwords, maybe learn some stuff about security, huh bud?
	} else if !alphaNumeric(strings.TrimSpace(*in.Password)) {
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
	// Update	 Active
	if in.Active != nil {
		account.Active = *in.Active
	}

	// Validate and Update Address
	if in.Address != nil {
		address, err := validateAddress(in.Address)
		if err != nil {
			return account, err
		}
		account.Address = *address
	}

	// Validate and Update Email
	if in.Email != nil {
		if err := validateEmail(*in.Email); err != nil {
			return account, err
		}
		account.Email = strings.ToLower(strings.TrimSpace(*in.Email))
	}

	// Validate and Update FirstName
	if in.FirstName != nil {
		if !alphabetic(strings.TrimSpace(*in.FirstName)) {
			return account, errors.New("FirstName must be alphabetic")
		}
		account.FirstName = strings.TrimSpace(*in.FirstName)
	}

	// Validate and Update LastName
	if in.LastName != nil {
		if !alphabetic(strings.TrimSpace(*in.LastName)) {
			return account, errors.New("LastName must be alphabetic")
		}
		account.LastName = strings.TrimSpace(*in.LastName)
	}

	// Validate and Update Password
	if in.Password != nil {
		if !alphabetic(strings.TrimSpace(*in.Password)) {
			return account, errors.New("Password must be alphabetic")
		}
		account.Password = strings.TrimSpace(*in.Password)
	}

	account.LastUpdated = time.Now().Unix()
	return account, nil
}
