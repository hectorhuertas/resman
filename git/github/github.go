package github

type Reposer interface {
	repos() []string
}

func Repos(r Reposer) []string {
	return r.repos()
}

type Org struct {
	Name string
}

func (o Org) repos() []string {
	return []string{"org1", "org2"}
}

type User struct {
	Name string
}

func (u User) repos() []string {
	return []string{"user1", "user2"}
}
