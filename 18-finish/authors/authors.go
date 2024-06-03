package authors

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type authorType struct {
	firstName string
	lastName  string
	bio       string
	birthDate string
}

var authorsArr []*authorType
var m sync.Mutex

func NewAuthor(fn, ln, bio, bd string) *authorType {
	return &authorType{
		firstName: fn,
		lastName:  ln,
		bio:       bio,
		birthDate: bd,
	}
}

func CreateAuthor(fn, ln, bio, bd string) error {
	_, err := time.Parse("2006-01-02", bd)
	if err != nil {
		return errors.New("Ошибка при разборе даты")
	}
	m.Lock()
	a := NewAuthor(fn, ln, bio, bd)
	authorsArr = append(authorsArr, a)
	m.Unlock()
	return nil
}

func GetAllAuthors() []string {
	r := make([]string, len(authorsArr))
	for i, _ := range authorsArr {
		r[i], _ = GetAuthorByIndex(i)
	}
	return r
}

func GetAuthorByIndex(i int) (string, error) {
	if i < 0 || i >= len(authorsArr) {
		return "", errors.New("Индекс вне диапазона")
	}
	r := fmt.Sprintf("%d %s %s %s", i, authorsArr[i].firstName, authorsArr[i].lastName, authorsArr[i].birthDate, authorsArr[i].bio)
	return r, nil
}

func UpdateAuthorByIndex(i int, fn, ln, bio, bd string) error {
	if i < 0 || i >= len(authorsArr) {
		return errors.New("Индекс вне диапазона")
	}
	_, err := time.Parse("2006-01-02", bd)
	if err != nil {
		return errors.New("Ошибка при разборе даты")
	}
	m.Lock()
	authorsArr[i].firstName = fn
	authorsArr[i].lastName = ln
	authorsArr[i].birthDate = bd
	authorsArr[i].bio = bio
	m.Unlock()
	return nil
}

func DeleteAuthorByIndex(i int) error {
	if i < 0 || i >= len(authorsArr) {
		return errors.New("Индекс вне диапазона")
	}
	authorsArr = append(authorsArr[:i], authorsArr[i+1:]...)
	return nil
}

type AuthorCRUD interface {
	CreateAuthor(string, string, string, string) error
	GetAllAuthors() []string
	GetAuthorByIndex(int) (string, error)
	UpdateAuthorByIndex(int) error
	DeleteAuthorByIndex(int) error
}
