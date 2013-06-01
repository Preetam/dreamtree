package dreamtree

import (
	"testing"
)

// Size test
func Test1(test *testing.T) {
	t := Create()
	if t.Size != 0 {
		test.Errorf("Nonzero size")
	}

	t.Insert("a")
	if t.Size != 1 {
		test.Errorf("Size didn't increase")
	}
}

// Balance test 1
func Test2(test *testing.T) {
	t := Create()
	t.Insert("b")
	t.Insert("d")
	t.Insert("a")
	t.Insert("c")
	ch := make(chan int)
	go func() {
		t.Height(ch)
	}()
	h1 := <-ch
	t.Balance()
	go func() {
		t.Height(ch)
	}()
	h2 := <-ch

	if h1 < h2 {
		test.Errorf("Balance didn't decrease height")
	}
}

// Balance test 2
func Test3(test *testing.T) {
	t := Create()
	t.Insert("a")
	t.Insert("b")
	t.Insert("c")
	t.Insert("d")
	ch := make(chan int)
	go func() {
		t.Height(ch)
	}()
	h1 := <-ch
	t.Balance()
	go func() {
		t.Height(ch)
	}()
	h2 := <-ch
	if h1 < h2 {
		test.Errorf("Balance didn't decrease height")
	}
}
