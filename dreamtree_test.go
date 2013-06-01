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

// Insertion
func Test2(test *testing.T) {
	t := Create()
	t.Insert("a")
	t.Insert("f")
	t.Insert("foo")
	t.Insert("bar")

	node := Get(t.Root, "f")
	if node.Value != "f" {
		test.Errorf("Get failure")
	}

	t.Remove("f")
}

// Deletion
func Test3(test *testing.T) {
	t := Create()
	t.Insert("a")
	t.Insert("f")
	t.Insert("foo")
	t.Insert("bar")

	node := Get(t.Root, "f")
	if node.Value != "f" {
		test.Errorf("Get failure")
	}

	t.Remove("f")

	node = Get(t.Root, "f")
	if node != nil {
		test.Errorf("Deletion failure")
	}

	if t.Size != 3 {
		test.Errorf("Incorrect size")
	}

	t.Remove("f")

	if t.Size != 3 {
		test.Errorf("Incorrect size")
	}
}

// Balance test 1
func Test4(test *testing.T) {
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
func Test5(test *testing.T) {
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

/*
func Test6(test *testing.T) {
	t := Create()
	i := 0
	var max int = 1e4
	for i = 0; i < max; i++ {
		t.Insert(string(i))
	}

	t.Balance()

	if t.Size != max {
		test.Errorf("Mass insert fail")
	}

	t.Balance()
}
*/
