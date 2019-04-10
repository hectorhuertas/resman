package memlocal

import (
	"fmt"
	"reflect"
	"testing"
	//	"github.com/hectorhuertas/resman/git"
)

// Need to test(names should change):
// New(location,...excludes)
// Scaj
// FindAllLocations() || Scan() return []locations || []Git(location+name)????
// GetInfo(location) Git,error
// GetInfoByName?() Git,error
// Create(Git) error

func TestLocal(t *testing.T) {
	t.Run("what?", func(t *testing.T) {
		l := New()
		fmt.Printf("%+v\n", l)
		fmt.Printf("%T\n", l)
		fmt.Println(reflect.TypeOf(l))
		fmt.Println(reflect.TypeOf(l))
	})
}
