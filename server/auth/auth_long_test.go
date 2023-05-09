package auth_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/mnrva-dev/owltier.com/server/auth"
	"github.com/mnrva-dev/owltier.com/server/config"
	"github.com/mnrva-dev/owltier.com/server/db"
	"github.com/mnrva-dev/owltier.com/server/token"
)

// TODO also write unit tests

type userdata struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Name         string `json:"name"`
}

var (
	testuser = &db.UserSchema{
		Username: "test",
		Password: "testpassword1234!!",
		Email:    "test@example.com",
	}
	permatestuser = &db.UserSchema{
		Id:       "0a85b0ae-a577-4be3-8609-d443d50f6939",
		Username: "user",
		Password: "password1234!!",
		Email:    "user@example.com",
		Refresh:  "",
	}
)

func runTestServer() *httptest.Server {
	return httptest.NewServer(auth.BuildRouter())
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestRegister(t *testing.T) {
	ts := runTestServer()
	defer ts.Close()
	data := url.Values{}
	data.Set("email", testuser.Email)
	if config.UsernamesEnabled() {
		data.Set("username", testuser.Username)
	}
	data.Set("password", testuser.Password)
	resp, err := http.PostForm(fmt.Sprintf("%s/register", ts.URL), data)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Could not read body")
	}
	if resp.StatusCode/100 != 2 {
		t.Errorf("Expected status in 200-299 range, got %d", resp.StatusCode)
		fmt.Println(string(body))
		t.FailNow()
	}
	AccessToken := strings.Split(string(body), "\n")[0]
	RefreshToken := strings.Split(string(body), "\n")[1]
	if AccessToken == "" || RefreshToken == "" {
		t.Errorf("Expected access and refresh token to be set")
	}
	if AccessToken == RefreshToken {
		t.Errorf("Expected access and refresh tokens to be different")
	}
	_, err = token.ValidateAccess(AccessToken)
	if err != nil {
		t.Errorf("Expected valid access token, got %v", err)
	}
	_, err = token.ValidateRefresh(RefreshToken)
	if err != nil {
		t.Errorf("Expected valid refresh token, got %v", err)
	}
}

func TestLogin(t *testing.T) {
	ts := runTestServer()
	defer ts.Close()
	data := url.Values{}
	data.Set("email", testuser.Email)
	if config.UsernamesEnabled() {
		data.Set("username", testuser.Username)
	}
	data.Set("password", testuser.Password)
	resp, err := http.PostForm(fmt.Sprintf("%s/login", ts.URL), data)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Could not read body")
	}
	if resp.StatusCode/100 != 2 {
		t.Errorf("Expected status in 200-299 range, got %d", resp.StatusCode)
		fmt.Println(string(body))
		t.FailNow()
	}
	AccessToken := strings.Split(string(body), "\n")[0]
	RefreshToken := strings.Split(string(body), "\n")[1]
	if AccessToken == "" || RefreshToken == "" {
		t.Errorf("Expected access and refresh token to be set")
	}
	if AccessToken == RefreshToken {
		t.Errorf("Expected access and refresh tokens to be different")
	}
	_, err = token.ValidateAccess(AccessToken)
	if err != nil {
		t.Errorf("Expected valid access token, got %v", err)
	}
	_, err = token.ValidateRefresh(RefreshToken)
	if err != nil {
		t.Errorf("Expected valid refresh token, got %v", err)
	}
}

func TestRefresh(t *testing.T) {
	ts := runTestServer()
	defer ts.Close()
	data := url.Values{}
	data.Set("email", testuser.Email)
	data.Set("password", testuser.Password)
	resp, err := http.PostForm(fmt.Sprintf("%s/login", ts.URL), data)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Could not read body")
	}
	if resp.StatusCode/100 != 2 {
		t.Error("Failed to login")
		t.FailNow()
	}
	RefreshToken := strings.Split(string(body), "\n")[1]
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/token/refresh", ts.URL), nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+RefreshToken)
	resp, err = http.DefaultClient.Do(req)
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Could not read body")
	}
	if resp.StatusCode/100 != 2 {
		t.Errorf("%s\n", string(body))
		t.Fatalf("Expected token to be validated but got status %d", resp.StatusCode)
	}
	AccessToken := strings.Split(string(body), "\n")[0]
	RefreshToken = strings.Split(string(body), "\n")[1]
	if AccessToken == "" || RefreshToken == "" {
		t.Errorf("Expected access and refresh token to be set")
	}
	if AccessToken == RefreshToken {
		t.Errorf("Expected access and refresh tokens to be different")
	}
	_, err = token.ValidateAccess(AccessToken)
	if err != nil {
		t.Errorf("Expected valid access token, got %v", err)
	}
	_, err = token.ValidateRefresh(RefreshToken)
	if err != nil {
		t.Errorf("Expected valid refresh token, got %v", err)
	}
}

func TestValidateAccess(t *testing.T) {
	ts := runTestServer()
	defer ts.Close()
	AccessToken := token.GenerateAccess(testuser)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/token/validate", ts.URL), nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+AccessToken)
	resp, err := http.DefaultClient.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Could not read body")
	}
	if err != nil {
		t.Error(string(body))
		t.Fatal(err)
	}
	if resp.StatusCode/100 != 2 {
		t.Error(string(body))
		t.Errorf("Expected token to be validated but got status %d", resp.StatusCode)
	}
}

func TestDeleteAccount(t *testing.T) {
	ts := runTestServer()
	defer ts.Close()
	data := url.Values{}
	data.Set("email", testuser.Email)
	data.Set("password", testuser.Password)
	resp, err := http.PostForm(fmt.Sprintf("%s/login", ts.URL), data)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Could not read body")
	}
	if resp.StatusCode/100 != 2 {
		t.Error("Failed to login")
		t.FailNow()
	}
	AccessToken := strings.Split(string(body), "\n")[0]
	data = url.Values{}
	data.Set("password", testuser.Password)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/delete", ts.URL), strings.NewReader(data.Encode()))
	req.Header.Add("Authorization", "Bearer "+AccessToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Could not read body")
	}
	// fmt.Println(string(body))
	if resp.StatusCode/100 != 2 {
		t.Errorf("Expected status in 200-299 range, got %d", resp.StatusCode)
		fmt.Println(string(body))
		t.FailNow()
	}
}
