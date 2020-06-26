package valid

import (
	"errors"
	"pizza-api/pkg/types"
	"regexp"
	"strings"
)

var (
	// AlphaNumeric regex
	alphaNumeric = regexp.MustCompile(`^[A-Za-z0-9" "]+$`).MatchString
	// Alphabetic regex
	alphabetic = regexp.MustCompile(`^[a-zA-Z" "]+$`).MatchString
	// Numeric rexex
	numeric = regexp.MustCompile(`^[0-9\-]+$`).MatchString
	// AlphaNumericOcto regex
	alphaNumericOcto = regexp.MustCompile(`^[A-Za-z0-9#" "]+$`).MatchString
)

func validateAddress(address *types.Address) (*types.Address, error) {
	if address.StreetAddress == "" {
		return address, errors.New("Specify a StreetAddress")
	} else if !alphaNumeric(strings.TrimSpace(address.StreetAddress)) {
		return address, errors.New("StreetAddress contains invalid characters")
	}

	// ExtendedAddress is optional - this is like an appartment number
	extended := ""
	if address.ExtendedAddress != "" {
		extended = strings.TrimSpace(address.ExtendedAddress)
		if !alphaNumericOcto(extended) {
			return address, errors.New("ExtendedAddress contains invalid characters")
		}
	}

	if address.Locality == "" {
		return address, errors.New("Specify a Locality")
	} else if !alphabetic(strings.TrimSpace(address.Locality)) {
		return address, errors.New("Locality contains invalid characters")
	}

	if address.Region == "" {
		return address, errors.New("Specify a Region")
	} else if !alphabetic(strings.TrimSpace(address.Region)) {
		return address, errors.New("Region contains invalid characters")
	}

	if address.PostalCode == "" {
		return address, errors.New("Specify a PostalCode")
	} else if !numeric(strings.TrimSpace(address.PostalCode)) {
		return address, errors.New("PostalCode contains invalid characters")
	}

	trimAddress := types.Address{
		StreetAddress:   strings.TrimSpace(address.StreetAddress),
		ExtendedAddress: extended,
		Locality:        strings.TrimSpace(address.Locality),
		Region:          strings.TrimSpace(address.Region),
		PostalCode:      strings.TrimSpace(address.PostalCode),
	}

	return &trimAddress, nil
}
