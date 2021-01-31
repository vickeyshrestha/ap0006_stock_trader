package bigfile

import (
	"database/sql"
	"net/http"
	"time"
)

type Store interface {
	CreateTable(tableName string) (sql.Result, error)
	InsertUser(id, userName, password string, isAdmin, isActive bool, createdAt time.Time) (User, error)
	GetUser(userName string) (User, error)
	QueryUsers() (Users, error)
	UpdateUserActive(userName string, isActive bool) error
}

type Service interface {
	Login(http.ResponseWriter, *http.Request)
	GetUserInfo(http.ResponseWriter, *http.Request)
	LogOut(http.ResponseWriter, *http.Request)
	ViewContents()
	UploadContents()
	DownloadContents()
	RemoveContents()

	// Admin rights
	AddUserAccount(http.ResponseWriter, *http.Request) (User, error)
	CloseUserAccount(http.ResponseWriter, *http.Request) error
	ExitApplication()
}
