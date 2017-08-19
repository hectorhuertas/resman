package main

import "fmt"

// import "log"
import "os"
import "path/filepath"
// import "os/exec"
import  "resource-manager/local"

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
	filepath.Walk(CODEPATH, local.GetGitRepos(&reposPaths))

	bob := local.UnmodifiedRepos(reposPaths)
	fmt.Printf("Clean: %v\n", bob)

	jar := local.ModifiedRepos(reposPaths)
	fmt.Printf("Dirty: %v\n", jar)
}
