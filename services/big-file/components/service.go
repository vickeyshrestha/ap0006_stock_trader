package bigfile

import (
	"encoding/json"
	"fmt"
	guuid "github.com/google/uuid"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"strings"
	"time"
)

var key = []byte("super-secret-key")

type Services struct {
	store  Store
	cookie *sessions.CookieStore
}

func NewService(userName, password string) (Service, error) {
	s, err := NewStore(userName, password)
	if err != nil {
		return nil, err
	}
	return &Services{
		store:  s,
		cookie: sessions.NewCookieStore(key),
	}, nil
}

func (s *Services) Login(writer http.ResponseWriter, request *http.Request) {
	var credentials User
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = writer.Write([]byte(fmt.Sprintf("error while decoding json: %s", err)))
		return
	}

	if credentials.Username == "" || credentials.Password == "" {
		_, _ = writer.Write([]byte("missing field"))
		return
	}

	// get the user info and validate authentication
	user, err := s.store.GetUser(credentials.Username)
	if err != nil {
		_, _ = writer.Write([]byte(fmt.Sprintf("error retrieving user record: %s", err)))
		return
	}
	if !strings.EqualFold(user.Password, credentials.Password) {
		_, _ = writer.Write([]byte("password not matched"))
		return
	}

	session, _ := s.cookie.Get(request, CommonCookieName)
	session.Values["authenticated"] = true
	err = session.Save(request, writer)
	if err != nil {
		_, _ = writer.Write([]byte("failed generating user session"))
		return
	}
	_, _ = writer.Write([]byte(fmt.Sprintf("user %s logged in success", credentials.Username)))
}

func (s *Services) GetUserInfo(writer http.ResponseWriter, request *http.Request) {
	session, _ := s.cookie.Get(request, CommonCookieName)

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(writer, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(writer, "The cake is a lie!")
}

func (s *Services) LogOut(writer http.ResponseWriter, request *http.Request) {
	session, _ := s.cookie.Get(request, CommonCookieName)

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(request, writer)
}

func (s *Services) ViewContents() {

}

func (s *Services) UploadContents() {

}

func (s *Services) DownloadContents() {

}

func (s *Services) RemoveContents() {

}

/*
	requires POST request with following JSON structure:
	{
		"user_name":"",
		"password":"",
		"is_admin":"",
	}

*/
func (s *Services) AddUserAccount(writer http.ResponseWriter, request *http.Request) (User, error) {
	decoder := json.NewDecoder(request.Body)
	var newUser User
	err := decoder.Decode(&newUser)
	if err != nil {
		return User{}, err
	}
	user, err := s.store.InsertUser(guuid.New().String(), newUser.Username, newUser.Password, newUser.IsAdmin, true, time.Now())
	if err != nil {
		return User{}, err
	}
	return user, nil
}

/*
	This is just a soft delete and should deactivate user, not delete it.
	Requires POST request with following JSON structure:
	{
		"user_name":""
	}
*/
func (s *Services) CloseUserAccount(writer http.ResponseWriter, request *http.Request) error {
	decoder := json.NewDecoder(request.Body)
	var newUser User
	err := decoder.Decode(&newUser)
	if err != nil {
		return err
	}
	err = s.store.UpdateUserActive(newUser.Username, false)
	if err != nil {
		return err
	}
	return nil
}

func (s *Services) ExitApplication() {
	os.Exit(1)
}
