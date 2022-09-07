package home

import (
	"net/url"
)

type loginData struct {
	Form url.Values
	Err  error
}
