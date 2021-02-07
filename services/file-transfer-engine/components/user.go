package filetransferengine

import "time"

type Users []User

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"user_name"`
	Password    string    `json:"password"`
	IsAdmin     bool      `json:"is_admin"`
	IsActive    bool      `json:"is_activ"`
	UserCreated time.Time `json:"user_created"`
}
