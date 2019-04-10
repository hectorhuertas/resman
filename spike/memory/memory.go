package memory

import (
	"github.com/hectorhuertas/resman/git"
)

type Repository struct {
	gits []git.Git
}

func (r *Repository) Create(g git.Git) error {
	r.gits = append(r.gits, g)
	return nil
}

func (r *Repository) Get(s string) (git.Git, error) {
	for _, g := range r.gits {
		if g.Name() == s {
			return g, nil
		}
	}
	return git.Git{}, nil
}

func (r *Repository) Find(opt ...git.FindOptions) ([]git.Git, error) {
	return r.gits, nil
}

func (r *Repository) GetAll() ([]git.Git, error) {
	return r.gits, nil
}

func New() Repository {
	return Repository{}
}
