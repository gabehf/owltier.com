package db_test

import (
	"testing"

	"github.com/mnrva-dev/owltier.com/server/db"
)

func TestCreate(t *testing.T) {
	u := &db.UserSchema{
		Username: "myuser",
		Email:    "me@example.com",
		Password: "secret123",
	}

	err := db.Create(u)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestRead(t *testing.T) {
	u := &db.UserSchema{
		Username: "myuser",
	}
	err := db.Fetch(u, u)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if u.Email != "me@example.com" {
		t.Fatalf("Expected correct email, got %v", u.Email)
	}
}

func TestGsiRead(t *testing.T) {
	u := &db.UserSchema{
		Email: "me@example.com",
	}
	err := db.FetchByGsi(u, u)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if u.Username != "myuser" {
		t.Fatalf("Expected correct username, got %v", u.Email)
	}
}

func TestDelete(t *testing.T) {
	u := &db.UserSchema{
		Username: "myuser",
		Email:    "me@example.com",
		Password: "secret123",
	}
	err := db.Delete(u)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	err = db.Fetch(u, u)
	if err == nil {
		t.Fatalf("Expected item to be deleted, but item exists")
	}
}
