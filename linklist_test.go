package datastructures_test

import (
	"testing"

	"github.com/samuelren97/go-datastructures"
)

func TestLinkList_Add_OneMore(t *testing.T) {
	list := datastructures.NewLinkList[int]()

	list.Add(42)

	if list.Count != 1 {
		t.Errorf("Wanted: 1\nGot: %d", list.Count)
	}
}

func TestLinkList_Get_OneElement_Ok(t *testing.T) {
	list := datastructures.NewLinkList[int]()
	wanted := 3
	list.Add(wanted)

	got := list.Get(0)

	if got != wanted {
		t.Errorf("Wanted: %d\nGot: %d", wanted, got)
	}
}

func TestLinkList_Get_TwoElements_Ok(t *testing.T) {
	list := datastructures.NewLinkList[int]()
	wanted := 2
	list.Add(1)
	list.Add(wanted)

	got := list.Get(1)

	if got != wanted {
		t.Errorf("Wanted: %d\nGot: %d", wanted, got)
	}
}

func TestLinkList_Get_ThreeElements_Ok(t *testing.T) {
	list := datastructures.NewLinkList[int]()
	wanted := 2
	list.Add(1)
	list.Add(wanted)
	list.Add(3)

	got := list.Get(1)

	if got != wanted {
		t.Errorf("Wanted: %d\nGot: %d", wanted, got)
	}
}

func TestLinkList_Remove_OneElementInList_Ok(t *testing.T) {
	list := datastructures.NewLinkList[int]()
	list.Add(42)

	list.Remove(0)

	if list.Count != 0 {
		t.Errorf("Wanted: 0\nGot: %d", list.Count)
	}
}

func TestLinkList_Remove_TwoElementsInList_Ok(t *testing.T) {
	list := datastructures.NewLinkList[int]()
	list.Add(42)
	list.Add(23)

	list.Remove(0)

	if list.Count != 1 {
		t.Errorf("Wanted: 0\nGot: %d", list.Count)
	}
	// TODO: Test case of getting value to compare with expected
}

func TestLinkList_Remove_ThreeElementsDeleteSecond_Ok(t *testing.T) {
	list := datastructures.NewLinkList[int]()
	wanted1 := 42
	list.Add(wanted1)
	list.Add(666)
	wanted2 := 23
	list.Add(wanted2)

	list.Remove(1)

	if list.Count != 2 {
		t.Errorf("Wanted: 0\nGot: %d", list.Count)
	}

	got1 := list.Get(0)
	got2 := list.Get(1)

	if got1 != wanted1 {
		t.Errorf("Wanted: %d\nGot: %d", wanted1, got1)
	}

	if got2 != wanted2 {
		t.Errorf("Wanted: %d\nGot: %d", wanted2, got2)
	}
}

func TestLinkList_ForEach_Ok(t *testing.T) {
	linkList := datastructures.NewLinkList[string]()
	linkList.Add("s")
	linkList.Add("a")
	linkList.Add("m")

	arr := make([]string, 0)
	linkList.ForEach(func(s string) {
		arr = append(arr, s)
	})

	if len(arr) != linkList.Count {
		t.Errorf("Lengths are not the same")
	}

	for i := 0; i < linkList.Count; i++ {
		if arr[i] != linkList.Get(i) {
			t.Errorf("Wanted: %s\nGot: %s", linkList.Get(i), arr[i])
		}
	}
}

func TestLinkList_Enqueue_Ok(t *testing.T) {
	linkList := datastructures.NewLinkList[string]()
	linkList.Add("a")
	linkList.Add("m")

	linkList.Enqueue("s")
	expected := "sam"

	var result string = ""
	linkList.ForEach(func(s string) { result += s })

	if result != expected {
		t.Errorf("Wanted: %s Got: %s", expected, result)
	}
}

func TestLinkList_Dequeue_Ok(t *testing.T) {
	linklist := datastructures.NewLinkList[string]()
	linklist.Add("s")
	linklist.Add("a")
	linklist.Add("m")
	linklist.Add("y")
	expected := "sam"
	expectedOutput := "y"

	output := linklist.Dequeue()
	var result string = ""
	linklist.ForEach(func(s string) { result += s })

	if result != expected {
		t.Errorf("Wanted: %s Got: %s", expected, result)
	}

	if output != expectedOutput {
		t.Errorf("Wanted output: %s, Got: %s", expectedOutput, output)
	}
}

func TestLinkList_Clear_Ok(t *testing.T) {
	linkList := datastructures.NewLinkList[string]()
	linkList.Add("s")
	linkList.Add("a")
	linkList.Add("m")
	expectedLen := 0

	linkList.Clear()

	if linkList.Count != expectedLen {
		t.Errorf("Wanted: %d Got: %d", expectedLen, linkList.Count)
	}
}
