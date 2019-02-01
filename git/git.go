package git

import "github.com/hectorhuertas/resman/git/github"
import (
	"fmt"
	"strings"
)

type git struct {
	Name      string
	localPath string
}

func (g git) String() string {
	return fmt.Sprintf("%s", g.Name)
}

type gitList []git

func (gs gitList) String() string {
	var ret strings.Builder
	for _, g := range gs {
		ret.WriteString(g.Name)
		ret.WriteString("\n")
	}
	return ret.String()
}

func MockGet() gitList {
	var gits []git
	gits = append(gits, git{Name: "resman", localPath: "/Users/hh/xdev/resman"})
	gits = append(gits, git{Name: "exercism", localPath: "/Users/hh/xdev/exercism"})
	return gits
}

func Get(root string, exclude string) gitList {
	fmt.Println(root)
	fmt.Println(exclude)
	u := github.User{"Pepe"}
	fmt.Println(github.Repos(u))
	o := github.Org{"uw-labs"}
	fmt.Println(github.Repos(o))

	var gits []git
	gits = append(gits, git{Name: "resman", localPath: "/Users/hh/xdev/resman"})
	gits = append(gits, git{Name: "exercism", localPath: "/Users/hh/xdev/exercism"})
	return gits

}
