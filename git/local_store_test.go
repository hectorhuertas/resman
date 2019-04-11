package git_test

import (
	"testing"

	"github.com/hectorhuertas/resman/git"
	"github.com/hectorhuertas/resman/git/memlocal"
)

func testLocalStorePort(ls git.LocalStore, t *testing.T) {
	t.Run("add, get and delete", func(t *testing.T) {
		// Given
		l := git.Local{Location: "somewhere"}

		// When
		err := ls.Add(l)

		// Then
		if err != nil {
			t.Error("cannot add local")
		}

		// When
		got, err := ls.Get(l.ID())

		// Then
		if err != nil {
			t.Error("cannot get local")
		}
		if got != l {
			t.Errorf("got wrong local: %v, expected: %v", got, l)
		}

		// When
		ls.Delete(got.ID())

		// Then
		_, err = ls.Get(l.ID())
		if err == nil {
			t.Error("cannot delete local")
		}

		//fmt.Printf("got: %+v\n", got)
		//fmt.Printf("want: %+v\n", want)

	})
	t.Run("when adding, empty location throws error", func(t *testing.T) {
		// Given
		l := git.Local{Location: ""}

		// When
		err := ls.Add(l)

		// Then
		if err == nil {
			t.Error()
		}
	})

	// test that you cannot add on the same location
}

func TestLocalStorePort(t *testing.T) {
	memls := memlocal.New()
	// ls := local.New()

	testLocalStorePort(memls, t)
	//testLocalStorePort(ls,t)
}
