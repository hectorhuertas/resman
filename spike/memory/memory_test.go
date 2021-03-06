package memory

import (
	"fmt"
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
	t.Run("find without filter", func(t *testing.T) {
		//t.Skip("not sure about the need for find")
		memory := New()
		// Given a repo with several gits
		expected := 3
		for i := 1; i <= expected; i++ {
			g, _ := git.New("test" + string(i))
			memory.Create(g)
		}
		// When find without filter
		gits, _ := memory.Find()
		//gits := []git.Git{} // placeholder to avoid errors

		// Then it gets all of them
		actual := len(gits)
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
	t.Run("find with several filters", func(t *testing.T) {
		t.Skip("not sure how deal with filters, probably not worth to test fake adapter")
		memory := New()
		// Given a repo with several gits
		g, _ := git.New("aa")
		memory.Create(g)
		g, _ = git.New("ab")
		memory.Create(g)
		g, _ = git.New("bb")
		memory.Create(g)
		// When find with filters
		fo := git.FindOptions{Visibility: "private", Owner: "bob"}
		gits, _ := memory.Find(fo)
		// or
		//f1, _ := git.Repository.Filter.New("visibility", "private")
		//f2, _ := git.Repository.Filter.New("contains", "esman")
		//f3, _ := git.Repository.Filter.New("language", "go")
		//f4, _ := git.Repository.Filter.New("owner", "hectorhuertas")
		//gits, _ := memory.Find(memory.Visibility("private"), memory.Language("go"))

		// Then it finds proper git
		fmt.Printf("%+v\n", fo)
		fmt.Printf("%+v\n", gits)
		if len(gits) != 1 || gits[0].Name() != "ab" {
			t.Error("filters not applied")
		}
	})
}
