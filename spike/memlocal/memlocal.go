package memlocal

import "github.com/hectorhuertas/resman/git"

type store struct {
	gits []git.Git
}

func New() store {
	return store{}
}

func (s *store) Add(g git.Git) error {
	s.gits = append(s.gits, g)
	return nil
}

func (s *store) Scan() []string {
	return []string{"location_1", "location_2"}
}

func (s *store) GetName(loc string) string {
	return "name"
}
