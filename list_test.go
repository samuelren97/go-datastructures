package datastructures_test

import (
	"datastructures"
	"testing"
)

func listenForPanic(t *testing.T) {
	if r := recover(); r != nil {
		t.Log("Expected panic")
	}
}

func TestList_Add_NewElement_OneMore(t *testing.T) {
	// Arrange
	list := datastructures.NewList[int]()

	// Act
	list.Add(12)

	// Assert
	if list.Count != 1 {
		t.Errorf("Wanted: 1\nGot: %d", list.Count)
	}
}

func TestList_Remove_RemoveElement0_OneLess(t *testing.T) {
	list := datastructures.NewList[int]()
	list.Add(23)
	list.Add(42)

	list.Remove(0)

	if list.Count != 1 {
		t.Errorf("Wanted: 1\n Got: %d", list.Count)
	}
}

func TestList_Remove_IndexUnder0_Panic(t *testing.T) {
	list := datastructures.NewList[int]()
	list.Add(42)

	defer listenForPanic(t)
	list.Remove(-1)

	t.Errorf("Expected panic, but no panic occured")
}
func TestList_Remove_IndexOnCount_Panic(t *testing.T) {
	list := datastructures.NewList[int]()
	list.Add(42)

	defer listenForPanic(t)
	list.Remove(1)

	t.Errorf("Expected panic, but no panic occured")
}

func TestList_Remove_IndexOverCount_Panic(t *testing.T) {
	list := datastructures.NewList[int]()
	list.Add(42)

	defer listenForPanic(t)
	list.Remove(2)

	t.Errorf("Expected panic, but no panic occured")
}

func TestList_EmptyList_Panic(t *testing.T) {
	list := datastructures.NewList[int]()

	defer listenForPanic(t)
	list.Remove(42)

	t.Errorf("Expected panic, but no panic occured")
}

func TestList_Get_Index0_Got42(t *testing.T) {
	list := datastructures.NewList[int]()
	wanted := 42
	list.Add(wanted)

	got := list.Get(0)

	if got != wanted {
		t.Errorf("Wanted: %d\nGot: %d", wanted, got)
	}
}

func TestList_Get_IndexUnder0_Panic(t *testing.T) {
	list := datastructures.NewList[int]()
	list.Add(42)

	defer listenForPanic(t)
	list.Get(-1)

	t.Errorf("Expected panic never occured")
}

func TestList_Get_IndexAtCount_Panic(t *testing.T) {
	list := datastructures.NewList[int]()
	list.Add(42)

	defer listenForPanic(t)
	list.Get(1)

	t.Errorf("Expected panic never occured")
}

func TestList_Get_IndexOverCount_Panic(t *testing.T) {
	list := datastructures.NewList[int]()
	list.Add(42)

	defer listenForPanic(t)
	list.Get(2)

	t.Errorf("Expected panic never occured")
}

func TestList_Get_EmptyList_Panic(t *testing.T) {
	list := datastructures.NewList[int]()

	defer listenForPanic(t)
	list.Remove(42)

	t.Errorf("Expected panic never occured")
}

func TestList_Find_Value_In_List(t *testing.T) {
	//Arrange
	list := datastructures.NewList[int]()
	for i := 0; i < 3; i++ {
		list.Add(i)
	}
	valueToFind := 1
	expectedIndex := 1

	//Act
	indexFound := list.FindIndex(func(i int) bool {
		return i == valueToFind
	})

	//Assert
	if indexFound != expectedIndex {
		t.Errorf("wanted index: %d got index: %d", expectedIndex, indexFound)
	}
}

func TestList_Find_Value_Not_In_List(t *testing.T) {
	//Arrange
	list := datastructures.NewList[int]()
	for i := 0; i < 3; i++ {
		list.Add(i)
	}
	valueToFind := 42
	expectedIndex := -1

	//Act
	indexFound := list.FindIndex(func(i int) bool {
		return i == valueToFind
	})

	if indexFound != expectedIndex {
		t.Errorf("wanted index: %d got index: %d", expectedIndex, indexFound)
	}
}

func TestList_ForEach_ListOfThree(t *testing.T) {
	list := datastructures.NewList[int]()
	for i := 0; i < 3; i++ {
		list.Add(6)
	}

	count := 0
	list.ForEach(func(i int) {
		if i != 6 {
			t.Errorf("Wanted: 6\nGot: %d", i)
		}
		count++
	})

	if count != 3 {
		t.Errorf("Wanted 3 interations\nGot: %d", count)
	}
}

func TestList_BubbleSort_SortedNumbers(t *testing.T) {
	list := datastructures.NewList[int]()
	list.Add(3)
	list.Add(2)
	list.Add(1)

	list.BubbleSort(func(a, b int) bool {
		return a > b
	})

	for i := 0; i < 3; i++ {
		if list.Get(i) != i+1 {
			t.Errorf("i=%d || The list is not sorted: %d", i, list.Get(i))
		}
	}
}
