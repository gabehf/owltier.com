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
	"github.com/mnrva-dev/owltier.com/server/db"
)

// TODO also write unit tests

type userdata struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	testuser = &db.UserSchema{
		Username: "test",
		Password: "testpassword1234!!",
	}
	permatestuser = &db.UserSchema{
		Username: "user",
		Password: "password1234!!",
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
	data.Set("username", testuser.Username)
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
}

func TestLogin(t *testing.T) {
	ts := runTestServer()
	defer ts.Close()
	data := url.Values{}
	data.Set("username", testuser.Username)
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
}

func TestDeleteAccount(t *testing.T) {
	ts := runTestServer()
	defer ts.Close()
	data := url.Values{}
	data.Set("username", testuser.Username)
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
	// TODO: Fix this so that it uses session token
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
