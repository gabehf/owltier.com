package auth

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/mnrva-dev/owltier.com/server/config"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

const (
	minPasswordEntropy = 60
)

type RequestForm struct {
	Username string
	Password string
}

func (h *RequestForm) validate() error {
	if h.Username == "" && config.UsernamesEnabled() {
		return errors.New("username is required")
	}
	if h.Password == "" {
		return errors.New("password is required")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9-_]{3,24}$`).MatchString(h.Username) && config.UsernamesEnabled() {
		return errors.New("username is not valid")
	}
	return passwordvalidator.Validate(h.Password, minPasswordEntropy)
}

func (h *RequestForm) Parse(r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	if config.UsernamesEnabled() {
		h.Username = strings.TrimSpace(r.FormValue("username"))
	}
	h.Password = strings.TrimSpace(r.FormValue("password"))
	// truncate extremely long passwords
	if len(h.Password) > 128 {
		h.Password = h.Password[:128]
	}

	return h.validate()
}
