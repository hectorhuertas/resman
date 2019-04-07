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
	size      int64
	status    string
	origin    gitOrigin
}

type Git struct {
	name string
}

func (g *Git) Name() string {
	return g.name
}

func New(name string) (Git, error) {
	g := Git{
		name: name,
	}
	return g, nil
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
		ret.WriteString(g.origin.owner)
		ret.WriteString("/")
		ret.WriteString(g.name)
		ret.WriteString(" ")
		ret.WriteString(g.status)
		ret.WriteString(" ")
		ret.WriteString(g.origin.host)
		ret.WriteString(" ")
		switch {
		case g.size > 1024*1024*1024:
			ret.WriteString(fmt.Sprintf("%.1fGi", float64(g.size)/1024/1024/1024))
		case g.size > 1024*1024:
			ret.WriteString(fmt.Sprintf("%.1fMi", float64(g.size)/1024/1024))
		case g.size > 1024:
			ret.WriteString(fmt.Sprintf("%.1fKi", float64(g.size)/1024))
		default:
			fmt.Printf("%d", g.size)
			ret.WriteString(fmt.Sprintf("%v", g.size))
		}
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
		gits = append(gits, git{name: path.Base(p), localPath: p, status: status, origin: getOrigin(p), size: getSize(p)})
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

type gitOrigin struct {
	owner       string
	host        string
	writeAccess bool
}

func getOrigin(path string) gitOrigin {
	var origin gitOrigin
	rawOriginBytes, _ := exec.Command("git", "-C", path, "config", "--get", "remote.origin.url").Output()
	rawOrigin := string(rawOriginBytes)
	origin.writeAccess = rawOrigin[0:3] == "git"

	if origin.writeAccess {
		atIndex := strings.Index(rawOrigin, "@")
		colonIndex := strings.Index(rawOrigin, ":")
		origin.host = rawOrigin[atIndex+1 : colonIndex]
		slashIndex := strings.Index(rawOrigin, "/")
		origin.owner = rawOrigin[colonIndex+1 : slashIndex]
	} else {
		doubleSlashIndex := strings.Index(rawOrigin, "/")
		noProtocol := rawOrigin[doubleSlashIndex+2:]
		firstSlashIndex := strings.Index(noProtocol, "/")
		origin.host = noProtocol[:firstSlashIndex]
		noHost := noProtocol[firstSlashIndex+1:]
		secondSlashIndex := strings.Index(noHost, "/")
		origin.owner = noHost[:secondSlashIndex]
	}
	return origin
}

func getSize(path string) int64 {
	var size int64
	filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size
}

type RepositoryX interface {
	Create(g Git) error
}
