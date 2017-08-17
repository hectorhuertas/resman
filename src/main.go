package main

import "fmt"

// import "log"
import "os"

// import "os/exec"

func main() {
	p1 := "/usr"
	p2 := "/Users/hh/xdev/go/resource-manager"
	// // binary, lookErr := exec.LookPath("git")
	// // if lookErr != nil {
	// //     log.Fatal(lookErr)
	// // }
	// // cmdX := binary
	// // args := []string{"status",""}
	// // output, err := exec.Command(cmdX, args...).Output()
	// cmd := exec.Command("/usr/bin/git", "status")
	// output, err := cmd.Output()
	// // fmt.Println(cmdX, output, "\n", args)
	// fmt.Println(cmd, output, err)
	// out, err := exec.Command("git", "status", "--porcelsain").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("The date is %s\n", out)
	// out, err := exec.Command("git", "status", "--porcelsain").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s",bob.Name(),"\nJANDER\n", err)
	// fmt.Printf("%s", bob.Name())
	fmt.Printf("%s: %t\n", p1, isGitRepo(p1))
	fmt.Printf("%s: %t\n", p2, isGitRepo(p2))
	// fmt.Printf("%t\n", isGitRepo("/usr"))
}

func isGitRepo(path string) bool {
	gitPath := path + "/.git"
	_, err := os.Stat(gitPath)
	return err == nil
}
