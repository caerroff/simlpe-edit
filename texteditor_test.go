package main

import (
	"fmt"
	"testing"
)

func TestNewLinkedChar(t *testing.T) {
	test := LinkedChar{
		char: 'A',
		next: &LinkedChar{
			char: 'B',
			next: &LinkedChar{
				char: 'C',
			},
		},
	}

	secondCharacter := test.next
	if secondCharacter.char != 'B' {
		t.Errorf("Second character not as set. Expecting %q, got %q\n", 'B', secondCharacter)
	}

	thirdCharacter := secondCharacter.next
	if thirdCharacter.char != 'C' {
		t.Errorf("Third character not as set. Expecting %q, got %q\n", 'C', thirdCharacter)
	}

	fmt.Println("TestNewLinkedChar passed")
}

func TestAddOnLinkedChar(t *testing.T) {
	rope := LinkedChar{
		char: 'A',
	}

	rope.Push('B')

	if rope.next == nil {
		t.Error("Rope.next is nil after calling Push")
		return
	}

	if rope.next.char != 'B' {
		t.Errorf("Second character is not as expected. Expected %q, got %q\n", 'B', rope.next.char)
	}

	fmt.Println("TestAddOnLinkedChar passed")
}

func TestCreateLinkedRope(t *testing.T) {
	charLinked := LinkedChar{
		char: 'A',
	}

	textRope := LinkedRope{
		rope: &TextRope{
			firstChar: &charLinked,
		},
	}

	if textRope.rope.firstChar.char != 'A' {
		t.Errorf("Error, expecting %q, got %q", 'A', textRope.rope.firstChar.char)
		return
	}

	fmt.Println("TestCreateLinkedRope passed")
}

func TestPushLinkedRope(t *testing.T) {
	charLinked := LinkedChar{
		char: 'A',
	}

	secondLinked := LinkedChar{
		char: 'B',
	}

	textRope := LinkedRope{
		rope: &TextRope{
			firstChar: &charLinked,
		},
	}

	secondRope := TextRope{
		firstChar: &secondLinked,
	}

	textRope.Push(&secondRope)

	if textRope.next.rope.firstChar.char != 'B' {
		t.Errorf("Error, expecting %q, got %q", 'B', textRope.rope.firstChar.char)
		return
	}

	fmt.Println("TestPushLinkedRope passed")
}
