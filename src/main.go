package main

import "fmt"

// import "log"
import "os"
import "path/filepath"
import "os/exec"

func main() {
	CODEPATH := os.Getenv("HOME") + "/xdev"
	fmt.Printf("CODEPATH: %s\n", CODEPATH)
	// // cmdX := binary
	// // args := []string{"status",""}
	// // output, err := exec.Command(cmdX, args...).Output()
	// cmd := exec.Command("/usr/bin/git", "status")
	// output, err := cmd.Output()
	// // fmt.Println(cmdX, output, "\n", args)
	reposPaths := []string{}
	filepath.Walk(CODEPATH, getGitRepos(&reposPaths))

	bob := unmodifiedRepos(reposPaths)
	fmt.Printf("Clean: %v\n", bob)

	jar := modifiedRepos(reposPaths)
	fmt.Printf("Dirty: %v\n", jar)
}

func getGitRepos(reposPaths *[]string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		// if path == "/Users/hh/xdev/repositive" {
		// 	return filepath.SkipDir
		// }
		if isGitRepo(path) {
			*reposPaths = append(*reposPaths, path)
			return filepath.SkipDir
		}
		return nil
	}
}

func isGitRepo(path string) bool {
	gitPath := path + "/.git"
	_, err := os.Stat(gitPath)
	return err == nil
}

func modifiedRepos(reposPaths []string) []string {
	return filterReposIfStatus(reposPaths, isClean)
}

func unmodifiedRepos(reposPaths []string) []string {
	return filterReposIfStatus(reposPaths, isDirty)
}

func filterReposIfStatus(reposPaths []string, f func(string)bool) []string {
	//TODO: This is only checking either the current or the master branch. What
	// should be the proper check?
	modifiedRepos := []string{}
	for _, repoPath := range reposPaths {
		out, _ := exec.Command("git", "-C", repoPath, "status", "--porcelain").Output()
		if f(string(out)) {
			modifiedRepos = append(modifiedRepos, repoPath)
		}
	}
	return modifiedRepos

}

func isDirty(s string) bool  {
	return s == ""
}

func isClean(s string) bool  {
	return s != ""
}
