package github

type Reposer interface {
	repos() ([]string, error)
}

func Repos(r Reposer) ([]string, error) {
	return r.repos()
}

type Org struct {
	Name string
}

func (o Org) repos() ([]string, error) {
	return []string{"org1", "org2"}, nil
}

type User struct {
	Name string
}

func (u User) repos() ([]string, error) {
	return []string{"user1", "user2"}, nil
}
