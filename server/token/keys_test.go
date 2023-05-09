package token_test

import (
	"testing"

	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/token"
)

var ts string

// TODO Write actual unit tests

func TestMain(m *testing.M) {
	ts = token.GenerateAccess(&db.UserSchema{
		Id:       "id_1234",
		Username: "myusername",
		Email:    "user@example.com",
		Scope:    "admin",
		Policies: []string{"policy1", "policy2"},
	})
	m.Run()
}
func TestAccessIdentity(t *testing.T) {
	if ts == "" {
		t.Fatal("No token was created")
	}
	c, err := token.ValidateAccess(ts)
	if err != nil {
		t.Errorf("Expected valid token, got %v", err)
	}
	if c.Email != "user@example.com" || c.Id != "id_1234" || c.Username != "myusername" {
		t.Errorf("Unexpected identity, got %v", c)
	}
}
func TestAccessScope(t *testing.T) {
	c, err := token.ValidateAccess(ts)
	if err != nil {
		t.Errorf("Expected valid token, got %v", err)
	}
	if c.Scope != "admin" {
		t.Errorf("Unexpected scope, expected %v got %v", "admin", c.Scope)
	}
	otherts := token.GenerateAccess(&db.UserSchema{
		Id:       "id_1234",
		Username: "myusername",
		Email:    "user@example.com",
		Policies: []string{"policy1", "policy2"},
	})
	c, err = token.ValidateAccess(otherts)
	if err != nil {
		t.Errorf("Expected valid token, got %v", err)
	}
	if c.Scope != "default" {
		t.Errorf("Unexpected scope, expected %v got %v", "default", c.Scope)
	}
}
func TestAccessPolicies(t *testing.T) {
	ts := token.GenerateAccess(&db.UserSchema{
		Id:       "id_1234",
		Username: "myusername",
		Email:    "user@example.com",
		Scope:    "admin",
		Policies: []string{"policy1", "policy2"},
	})
	if ts == "" {
		t.Error("No token was created")
	}
	c, err := token.ValidateAccess(ts)
	if err != nil {
		t.Errorf("Expected valid token, got %v", err)
	}
	if len(c.Policies) != 2 || c.Policies[0] != "policy1" || c.Policies[1] != "policy2" {
		t.Errorf("Unexpected policies, got %v", c.Policies)
	}
}

func TestRefresh(t *testing.T) {
	ts := token.GenerateRefresh(&db.UserSchema{
		Id: "12345",
	})
	if ts == "" {
		t.Error("No token was created")
	}
	c, err := token.ValidateRefresh(ts)
	if err != nil {
		t.Errorf("Expected valid token, got %v", err)
	}
	if c.Id != "12345" {
		t.Errorf("Unexpected identity, got %v", c)
	}
}
