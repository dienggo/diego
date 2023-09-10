package auth

import (
	"encoding/json"
	"net/http"
)

const USERAUTHKEY = "x-user-auth-key-diego-x"

// SetAuth : setter auth data
func SetAuth(user any, r *http.Request) error {
	marshal, err := json.Marshal(user)
	if err != nil {
		return err
	}
	r.Header.Set(USERAUTHKEY, string(marshal))
	return nil
}

// User : getter auth data
func User[T comparable](r *http.Request) (error, T) {
	user := new(T)
	su := r.Header.Get(USERAUTHKEY)
	err := json.Unmarshal([]byte(su), &user)
	if err != nil {
		return err, *user
	}
	return nil, *user
}
