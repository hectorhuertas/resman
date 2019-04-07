package memory

import (
	"testing"

	"github.com/hectorhuertas/resman/git"
)

func TestMemory(t *testing.T) {
	t.Run("create and find", func(t *testing.T) {
		memory := New()
		// Given a git
		name := "testCreateGet"
		expected, _ := git.New(name)

		// When it is created
		memory.Create(expected)

		// Then it can be recovered
		actual, _ := memory.Get(name)
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
	t.Run("get all", func(t *testing.T) {
		memory := New()
		// Given a repo with several gits
		expected := 3
		for i := 1; i <= expected; i++ {
			g, _ := git.New("test" + string(i))
			memory.Create(g)
		}

		// When get all
		gits, _ := memory.GetAll()

		// Then it get all of them
		actual := len(gits)
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
	t.Run("get all when empty", func(t *testing.T) {
		memory := New()
		// Given a repo with no gits
		expected := 0
		// When get all
		gits, _ := memory.GetAll()

		// Then it get all of them
		actual := len(gits)
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
}

// test the difference between network error and user not found
