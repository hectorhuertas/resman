package main

import "fmt"

// import "log"
import "os"
import "path/filepath"
// import "os/exec"

func main() {
  CODEPATH:=os.Getenv("HOME")+"/xdev"
  fmt.Printf("CODEPATH: %s\n", CODEPATH)
  // fmt.Println("jsalkdfjaslf")
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
  //  log.Fatal(err)
  // }
  // fmt.Printf("The date is %s\n", out)
  // out, err := exec.Command("git", "status", "--porcelsain").Output()
  // if err != nil {
  // log.Fatal(err)
  // }
  // fmt.Printf("%s",bob.Name(),"\nJANDER\n", err)
  // fmt.Printf("%s", bob.Name())
  fmt.Printf("TEST: %s: %t\n", p1, isGitRepo(p1))
  fmt.Printf("TEST: %s: %t\n", p2, isGitRepo(p2))
  // fmt.Printf("%t\n", isGitRepo("/usr"))
  reposPaths := []string{}
  filepath.Walk(CODEPATH, getGitRepos(&reposPaths))
  // kak := make([]string, 0)
  // fmt.Printf("One: %v\n", kak)
  // fmt.Printf("One: %v\n", reposPaths)
  for _, repoPath := range reposPaths {
        // fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
        fmt.Printf("Path: %s\n", repoPath)

        // fmt.Println("")
    }
  // bob = append(bob, "jander")
  // fmt.Printf("Two: %v\n", bob)
}


func getGitRepos(reposPaths *[]string) filepath.WalkFunc{
  return func (path string, f os.FileInfo, err error) error {
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
