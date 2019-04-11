package git_test

import (
	"testing"

	"github.com/hectorhuertas/resman/git"
)

func TestScan(t *testing.T) {
	t.Skip()
	// need to decide how to model the local remote
	t.Run("", func(t *testing.T) {
		//ls := memlocal.New()
		// Given
		//l := git.Local{"location"}
		//ls.Add(l)
		//rs := memremote.New()

		//	gs := NewService(ls, rs)
		// When
		//got := gs.Scan()

		// Then
		got := git.Git{"placeholder"}
		want := git.Git{"location"}
		if got != want {
			t.Error()
		}

	})

}
