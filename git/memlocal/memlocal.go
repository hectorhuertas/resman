package memlocal

import (
	"errors"
	"math/rand"

	"github.com/hectorhuertas/resman/git"
)

type Store struct {
	locals []git.Local
}

func New() *Store {
	s := &Store{}
	return s
}

func (s *Store) Add(l git.Local) error {
	if l.Location == "" {
		return errors.New("add error: invalid location")
	}

	_, err := s.Get(l.Location)
	if err == nil {
		return errors.New("add error: local already exists")
	}

	s.locals = append(s.locals, l)
	return nil
}

func (s *Store) Get(loc string) (git.Local, error) {
	for _, l := range s.locals {
		if l.Location == loc {
			return l, nil
		}
	}

	return git.Local{}, errors.New("get error: not found")
}

func (s *Store) GetAll() []git.Local {
	return s.locals
}

func (s *Store) Delete(loc string) {
	for i, l := range s.locals {
		if l.Location == loc {
			s.locals = append(s.locals[:i], s.locals[i+1:]...)
			return
		}
	}
}

// Generate 30 chars long [a-zA-Z] non-crypto random strings
func (s *Store) GenerateLocation() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 30)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
