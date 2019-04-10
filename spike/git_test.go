package git_test

import (
	"fmt"
	"testing"

	"github.com/hectorhuertas/resman/git"
)

func TestLocal(t *testing.T) {
	g := git.Git{}
	t.Run("default options", func(t *testing.T) {
		fmt.Println(g)
	})
}
