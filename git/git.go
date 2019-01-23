package git

import "github.com/shurcooL/githubv4"
import "golang.org/x/oauth2"
import (
	"context"
	"fmt"
	"os"
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

func Get() gitList {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var query struct {
		Viewer struct {
			Login     githubv4.String
			CreatedAt githubv4.DateTime
		}
	}

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
	}
	fmt.Println("    Login:", query.Viewer.Login)
	fmt.Println("CreatedAt:", query.Viewer.CreatedAt)
	fmt.Println("whatever")

	var gits []git
	gits = append(gits, git{Name: "resman", localPath: "/Users/hh/xdev/resman"})
	gits = append(gits, git{Name: "exercism", localPath: "/Users/hh/xdev/exercism"})
	return gits

}
