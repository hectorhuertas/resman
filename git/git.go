package git

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type git struct {
	name      string
	localPath string
	size      string
	status    string
}

func (g git) String() string {
	return fmt.Sprintf("%s", g.name)
}

type gitList []git

func (gs gitList) String() string {
	var ret strings.Builder
	for i, g := range gs {
		if i != 0 {
			ret.WriteString("\n")
		}
		ret.WriteString(g.name)
		ret.WriteString(" ")
		ret.WriteString(g.status)
	}
	return ret.String()
}

func Scan(root string, exclude []string) gitList {
	paths, _ := findGitDirs(root, exclude)

	//u := github.User{"Pepe"}
	//fmt.Println(github.Repos(u))
	//o := github.Org{"uw-labs"}
	//fmt.Println(github.Repos(o))

	var gits []git
	for _, p := range paths {
		out, _ := exec.Command("git", "-C", p, "status", "--porcelain").Output()
		status := ""
		if string(out) != "" {
			status = "\u2757"
		}
		gits = append(gits, git{name: path.Base(p), localPath: p, status: status})

	}
	return gits

}

func findGitDirs(root string, exclude []string) (ret []string, err error) {
	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		for _, e := range exclude {
			if path == e {
				return filepath.SkipDir
			}
		}

		gitPath := path + "/.git"
		_, err = os.Stat(gitPath)
		if err == nil {
			ret = append(ret, path)
			return filepath.SkipDir
		}
		return nil
	})
	return ret, nil
}
