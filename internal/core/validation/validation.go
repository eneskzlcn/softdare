package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/rs/xid"
)

func IsValidXID(xID string) error {
	_, err := xid.FromString(xID)
	return err
}

func IsValidEmail(email string) error {
	return validation.Validate(email, is.Email)
}

func IsValidUsername(username string) error {
	return validation.Validate(username, validation.Length(2, 12))
}

func IsValidContent(content string) error {
	return validation.Validate(content, validation.Length(2, 1000))
}
