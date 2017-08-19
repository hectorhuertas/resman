package gitLocal

// import "fmt"

// import "log"
import "os"
import "path/filepath"
import "os/exec"

func GetRepos(path string) []string  {
	repos := []string{}
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if isGitRepo(path) {
			repos = append(repos, path)
			return filepath.SkipDir
		}
		return nil
	})
	return repos
}
func GetGitRepos(reposPaths *[]string) filepath.WalkFunc {
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

func GetModifiedRepos(path string) []string  {
	repos := GetRepos(path)
	return ModifiedRepos(repos)
}
func GetUnmodifiedRepos(path string) []string  {
	repos := GetRepos(path)
	return UnmodifiedRepos(repos)
}

func isGitRepo(path string) bool {
	gitPath := path + "/.git"
	_, err := os.Stat(gitPath)
	return err == nil
}

func ModifiedRepos(reposPaths []string) []string {
	return filterReposIfStatus(reposPaths, isClean)
}

func UnmodifiedRepos(reposPaths []string) []string {
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
