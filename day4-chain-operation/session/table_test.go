package session

import "testing"

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age int
}

func TestSession_CreateTable(t *testing.T){
	s := NewSession().Model(&User{})
	table := s.RefTable()
	s.Model(&Session{})
	if table.Name != "User" || s.RefTable().Name != "Session"{
		t.Fatal("Failed to change model")
	}
}