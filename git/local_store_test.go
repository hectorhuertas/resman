package git_test

import (
	"testing"

	"github.com/hectorhuertas/resman/git"
	"github.com/hectorhuertas/resman/git/memlocal"
)

func testLocalStorePort(ls git.LocalStore, t *testing.T) {
	t.Run("single add, get and delete", func(t *testing.T) {
		// Given
		loc := ls.GenerateLocation()
		l := git.Local{Location: loc}

		// When
		err := ls.Add(l)

		// Then
		if err != nil {
			t.Error("cannot add local")
		}

		// When
		got, err := ls.Get(loc)

		// Then
		if err != nil {
			t.Error("cannot get local")
		}
		if got != l {
			t.Errorf("got wrong local: %v, expected: %v", got, l)
		}

		// When
		ls.Delete(got.Location)

		// Then
		_, err = ls.Get(loc)
		if err == nil {
			t.Error("cannot delete local")
		}
	})
	t.Run("multiple add, get and delete", func(t *testing.T) {
		// Given
		loc1 := ls.GenerateLocation()
		loc2 := ls.GenerateLocation()
		loc3 := ls.GenerateLocation()
		l1 := git.Local{Location: loc1}
		l2 := git.Local{Location: loc2}
		l3 := git.Local{Location: loc3}
		ls.Add(l1)
		ls.Add(l2)
		ls.Add(l3)
		defer ls.Delete(loc1)
		defer ls.Delete(loc2)
		defer ls.Delete(loc3)

		// When
		got, _ := ls.Get(loc2)

		// Then
		if got != l2 {
			t.Errorf("got wrong local: %v, expected: %v", got, l2)
		}

		// When
		ls.Delete(loc3)

		// Then
		_, err := ls.Get(loc3)
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
		loc := ls.GenerateLocation()
		l1 := git.Local{Location: loc}
		l2 := git.Local{Location: loc}
		defer ls.Delete(loc)
		defer ls.Delete(loc)
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
		times := 3
		for i := 1; i <= times; i++ {
			loc := ls.GenerateLocation()
			l := git.Local{Location: loc}
			ls.Add(l)
			defer ls.Delete(loc)
		}

		// When
		localsNumber := len(ls.GetAll())

		// Then
		if localsNumber != times {
			t.Errorf("got %v, expected %v", localsNumber, times)
		}
	})
}

// fmt.Printf("got: %+v\n", locals)
func TestLocalStorePort(t *testing.T) {
	var localStores []git.LocalStore
	localStores = append(localStores, memlocal.New())
	// To do: add non-memory store only if running integration tests
	//localStores = append(localStores, local.New())

	for _, ls := range localStores {
		before := len(ls.GetAll())

		testLocalStorePort(ls, t)

		after := len(ls.GetAll())

		if before != after {
			t.Fatal("IMPORTANT ERROR: Store left in a different state after testing")
		}
	}
}
