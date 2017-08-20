package main

import "fmt"

// import "log"
import "os"
import  "resource-manager/gitLocal"

func main() {
	scan()
}

func scan()  {
	CODEPATH := os.Getenv("HOME") + "/xdev" // Might file for / dir
	fmt.Printf("CODEPATH: %s\n", CODEPATH)

	// fmt.Printf("All: %v\n", gitLocal.GetRepos(CODEPATH))
	// fmt.Printf("Clean: %v\n", gitLocal.GetUnmodifiedRepos(CODEPATH))
	fmt.Printf("Dirty: %v\n", gitLocal.GetModifiedRepos(CODEPATH))
}
