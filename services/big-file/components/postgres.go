package bigfile

import (
	"database/sql"
	"fmt"
	"time"
)

type mySqlDb struct {
	db *sql.DB
}

func NewStore(userName, passWord string) (Store, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", userName, passWord, DbName))
	if err != nil {
		return nil, err
	}
	return &mySqlDb{db: db}, nil
}

func (s *mySqlDb) CreateTable(tableName string) (sql.Result, error) {
	stmt, err := s.db.Prepare(fmt.Sprintf("CREATE Table %s (id int NOT NULL, user_name varchar(50), password varchar(30), is_admin boolean, is_active boolean, created_at datetime,PRIMARY KEY (id));", tableName))
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec()
	return result, nil
}

func (s *mySqlDb) InsertUser(id, userName, password string, isAdmin, isActive bool, createdAt time.Time) (User, error) {
	stmt, err := s.db.Prepare(fmt.Sprintf("INSERT %s SET id=?,user_name=?,password=?,is_admin=?,is_active=?,created_at=?", DbTable))
	if err != nil {
		return User{}, err
	}
	_, err = stmt.Exec(id, userName, password, isAdmin, isActive, createdAt)
	if err != nil {
		return User{}, err
	}
	return User{
		Username: userName,
		IsAdmin:  false,
	}, nil
}

func (s *mySqlDb) QueryUsers() (Users, error) {
	rows, err := s.db.Query(fmt.Sprintf("SELECT * from %s", DbTable))
	if err != nil {
		return Users{}, err
	}
	user := User{}
	users := Users{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.IsAdmin, &user.IsActive, &user.UserCreated)
		users = append(users, user)
	}
	return users, nil
}

func (s *mySqlDb) GetUser(userName string) (User, error) {
	rows, err := s.db.Query(fmt.Sprintf("SELECT * FROM %s WHERE user_name=?", DbTable), userName)
	if err != nil {
		return User{}, err
	}
	user := User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.IsAdmin, &user.IsActive, &user.UserCreated)
	}
	return user, nil
}

func (s *mySqlDb) UpdateUserActive(userName string, isActive bool) error {
	stmt, err := s.db.Prepare(fmt.Sprintf("update %s set is_active=? where user_name=?", DbTable))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(isActive, userName)
	return err
}
