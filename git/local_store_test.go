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
	})
	t.Run("multiple add, get and delete", func(t *testing.T) {
		// Given
		newLocal1 := git.Local{Location: "somewhere"}
		defer ls.Delete(newLocal1.Location)
		ls.Add(newLocal1)
		newLocal2 := git.Local{Location: "here"}
		defer ls.Delete(newLocal2.Location)
		ls.Add(newLocal2)
		newLocal3 := git.Local{Location: "there"}
		defer ls.Delete(newLocal3.Location)
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
	})
	t.Run("cannot add empty location", func(t *testing.T) {
		// Given
		l := git.Local{Location: ""}

		// When
		err := ls.Add(l)

		// Then
		if err == nil {
			t.Error()
		}
	})
	t.Run("cannot add in existing location", func(t *testing.T) {
		// Given
		l1 := git.Local{Location: "somewhere"}
		defer ls.Delete(l1.Location)
		l2 := git.Local{Location: "somewhere"}
		defer ls.Delete(l2.Location)
		ls.Add(l1)

		// When
		err := ls.Add(l2)

		// Then
		if err == nil {
			t.Error()
		}
	})
	t.Run("get all", func(t *testing.T) {
		// Given
		l1 := git.Local{Location: "somewhere"}
		defer ls.Delete(l1.Location)
		l2 := git.Local{Location: "there"}
		defer ls.Delete(l2.Location)
		ls.Add(l1)
		ls.Add(l2)

		// When
		locals := ls.GetAll()

		// Then
		if 2 != len(locals) {
			t.Errorf("got %v, expected 2", len(locals))
		}
	})
}

// fmt.Printf("got: %+v\n", locals)
func TestLocalStorePort(t *testing.T) {
	memls := memlocal.New()
	// ls := local.New()
	before := len(memls.GetAll())

	testLocalStorePort(memls, t)
	after := len(memls.GetAll())
	if before != after {
		t.Fatal("IMPORTANT ERROR: Store left in a different state after testing")
	}
	//testLocalStorePort(ls,t)
}
