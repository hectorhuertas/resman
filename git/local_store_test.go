package git_test

import (
	"testing"

	"github.com/hectorhuertas/resman/git"
	"github.com/hectorhuertas/resman/git/memlocal"
)

func testLocalStorePort(ls git.LocalStore, t *testing.T) {
	t.Run("single add, get and delete", func(t *testing.T) {
		// Given
		newLocal := git.Local{Location: "somewhere"}

		// When
		err := ls.Add(newLocal)

		// Then
		if err != nil {
			t.Error("cannot add local")
		}

		// When
		got, err := ls.Get(newLocal.Location)

		// Then
		if err != nil {
			t.Error("cannot get local")
		}
		if got != newLocal {
			t.Errorf("got wrong local: %v, expected: %v", got, newLocal)
		}

		// When
		ls.Delete(got.Location)

		// Then
		_, err = ls.Get(newLocal.Location)
		if err == nil {
			t.Error("cannot delete local")
		}

		//fmt.Printf("got: %+v\n", got)
		//fmt.Printf("want: %+v\n", want)

	})
	t.Run("multiple add, get and delete", func(t *testing.T) {
		// Given
		newLocal1 := git.Local{Location: "somewhere"}
		ls.Add(newLocal1)
		newLocal2 := git.Local{Location: "here"}
		ls.Add(newLocal2)
		newLocal3 := git.Local{Location: "there"}
		ls.Add(newLocal3)

		// When
		got, _ := ls.Get(newLocal2.Location)

		// Then
		if got != newLocal2 {
			t.Errorf("got wrong local: %v, expected: %v", got, newLocal2)
		}

		// When
		ls.Delete(newLocal3.Location)

		// Then
		_, err := ls.Get(newLocal3.Location)
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
