package main

import (
	component "maxzilla/services/filetransferengine/components"
	"net/http"
)

func main() {
	s, err := component.NewService(component.DbUserName, component.DbPassword)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/login", s.Login)
	http.HandleFunc("/info", s.GetUserInfo)
	http.HandleFunc("/logout", s.LogOut)

	_ = http.ListenAndServe(":8085", nil)
}
