package auth

import (
	"encoding/base64"
	"errors"
	"net/http"
	"net/mail"
	"regexp"
	"strings"

	"github.com/mnrva-dev/owltier.com/server/config"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

const (
	minPasswordEntropy = 60
)

type RequestForm struct {
	Username    string
	Password    string
	Email       string
	Redirect    bool
	RedirectUrl string
}

func (h *RequestForm) validate() error {
	if h.Username == "" && config.UsernamesEnabled() {
		return errors.New("username is required")
	}
	if h.Password == "" {
		return errors.New("password is required")
	}
	if h.Email == "" {
		return errors.New("email is required")
	}
	if _, err := mail.ParseAddress(h.Email); err != nil {
		return errors.New("email address is not valid")
	}
	if h.Redirect {
		if h.RedirectUrl == "" {
			return errors.New("redirect url is required")
		}
		var url []byte
		_, err := base64.NewDecoder(base64.StdEncoding, strings.NewReader(h.RedirectUrl)).Read(url)
		if err != nil {
			return err
		}
		h.RedirectUrl = string(url)
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
	h.Email = strings.TrimSpace(r.FormValue("email"))
	h.Redirect = r.FormValue("redirect") != "" || strings.ToLower(r.FormValue("redirect")) == "false"
	h.RedirectUrl = strings.TrimSpace(r.FormValue("redirect_url"))

	return h.validate()
}
