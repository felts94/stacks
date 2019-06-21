package stacks

import (
	"testing"
)

func TestStack(t *testing.T) {
	stringStack := &Stack{}
	stringStack.Push("First")
	stringStack.Push("Second")
	exV := stringStack.Pop()
	if exV != "Second" {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", exV, "Second")
	}
	stringStack.Push("Third")
	exV = stringStack.Pop()
	if exV != "Third" {
		t.Errorf("Item was incorrect, got: %v, want: %v.", exV, "Third")
	}
	exV = stringStack.Pop()
	if exV != "First" {
		t.Errorf("Item was incorrect, got: %v, want: %v.", exV, "First")
	}

	shouldBeNil := stringStack.Pop()
	if shouldBeNil != nil {
		t.Errorf("Item was incorrect, got: %v, want: %v.", shouldBeNil, "nil")
	}
}
