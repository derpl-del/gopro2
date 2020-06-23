package logincode

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

//StoreKey a
var StoreKey *sessions.CookieStore

//Sea a
const Sea = "welcome1"

//NewCookieStore Create Cookie
func NewCookieStore() *sessions.CookieStore {
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	StoreKey := sessions.NewCookieStore(authKey, encryptionKey)
	StoreKey.Options.Path = "/"
	StoreKey.Options.MaxAge = 86400 * 7
	StoreKey.Options.HttpOnly = true

	return StoreKey
}

//SetSession for session
func SetSession(input string, w http.ResponseWriter, r *http.Request) {
	session, _ := StoreKey.Get(r, Sea)
	session.Values["auth"] = true
	session.Values["username"] = input
	session.Save(r, w)
}

//GetUserName Session
func GetUserName(r *http.Request, input string) (bool, string) {
	session, _ := StoreKey.Get(r, Sea)
	if len(session.Values) == 0 {
		return false, ""
	}
	return session.Values["auth"].(bool), fmt.Sprintf("%v", session.Values["username"])
}

//ClearSession Session
func ClearSession(w http.ResponseWriter, r *http.Request, input string) {
	session, _ := StoreKey.Get(r, Sea)
	session.Options.MaxAge = -1
	session.Save(r, w)
}

//CreateStore Session
func CreateStore() {
	StoreKey = NewCookieStore()
}
