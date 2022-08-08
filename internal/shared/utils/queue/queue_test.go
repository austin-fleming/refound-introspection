package queue

import (
	"testing"
)

func TestCreateEmptyStringQueue(t *testing.T) {
	testQ := Create[string](0)

	// Initial
	if capacity := testQ.Capacity(); capacity != 1 {
		t.Errorf("\nFor: %s \nExpected: %d \nGot: %d\n", "capacity", 1, capacity)
	}
	if length := testQ.Length(); length != 0 {
		t.Errorf("\nFor: %s \nExpected: %d \nGot: %d\n", "length", 0, length)
	}
	if isEmpty := testQ.IsEmpty(); isEmpty != true {
		t.Errorf("\nFor: %s \nExpected: %t \nGot: %t\n", "isEmpty", true, isEmpty)
	}
	if isFull := testQ.IsFull(); isFull != false {
		t.Errorf("\nFor: %s \nExpected: %t \nGot: %t\n", "isFull", false, isFull)
	}
	if peekResult := testQ.Peek(); peekResult != "" {
		t.Errorf("\nFor: %s \nExpected: %s \nGot: %s\n", "peek", "", peekResult)
	}
	if hasResult, result := testQ.Dequeue(); hasResult != false || result != "" {
		t.Errorf("\nFor: %s \nExpected: %t & %s \nGot: %t & %s\n", "dequeue", false, "", hasResult, result)
	}
	if hasResult, result := testQ.LazyDequeue(); hasResult != false && result != "" {
		t.Errorf("\nFor: %s \nExpected: %t & %s \nGot: %t & %s\n", "lazyDequeue", false, "", hasResult, result)
	}
	if error := testQ.Enqueue("firstItem"); error != nil {
		t.Errorf("\nFor: %s \nExpected: %s \nGot: %s\n", "enqueue", "no error", error)
	}
	if item := testQ.Peek(); item != "firstItem" {
		t.Errorf("\nFor: %s \nExpected: %s \nGot: %s\n", "peek", "firstItem", item)
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := Create[string](3)

	for _, item := range []string{"first", "second", "third"} {
		if error := q.Enqueue(item); error != nil {
			t.Errorf("\nFor: %s \nExpected: %s \nGot: %s\n", "enqueue", "no error", error)
		}
	}
}
