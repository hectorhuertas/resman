package main

import "fmt"

// import "log"
import "os"
import  "resource-manager/gitLocal"

func main() {
	CODEPATH := os.Getenv("HOME") + "/xdev"
	fmt.Printf("CODEPATH: %s\n", CODEPATH)
	// // cmdX := binary
	// // args := []string{"status",""}
	// // output, err := exec.Command(cmdX, args...).Output()
	// cmd := exec.Command("/usr/bin/git", "status")
	// output, err := cmd.Output()
	// // fmt.Println(cmdX, output, "\n", args)
	fmt.Printf("Clean: %v\n", gitLocal.GetUnmodifiedRepos(CODEPATH))
	fmt.Printf("Dirty: %v\n", gitLocal.GetModifiedRepos(CODEPATH))

	fmt.Printf("All: %v\n", gitLocal.GetRepos(CODEPATH))
}
