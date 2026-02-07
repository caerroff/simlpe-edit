package main

/* TextRope, representing one line in the edtior. It knows the characters of the line, and the next line */
type TextRope struct {
	firstChar *LinkedChar
	length    uint16
}

/*Stores all lines for the editor */
type LinkedRope struct {
	rope *TextRope
	next *LinkedRope
}

func (lr *LinkedRope) Push(rope *TextRope) {
	currentRope := lr
	for currentRope.next != nil {
		currentRope = lr.next
	}
	var newRope *LinkedRope = &LinkedRope{
		rope: rope,
	}
	currentRope.next = newRope
}

func (lr *LinkedRope) AddAtIndex(rope *TextRope, index int16) *LinkedRope {
	currentRope := lr
	var currentIndex int16 = 0
	for currentRope.next != nil && currentIndex != index {
		currentRope = lr.next
		currentIndex++
	}

	var newRope *LinkedRope = &LinkedRope{
		rope: rope,
	}
	if currentIndex == 0 {
		newRope.next = lr
		return newRope
	} else {
		oldNext := currentRope.next
		currentRope.next = newRope
		newRope.next = oldNext
		return lr
	}
}

type LinkedChar struct {
	char rune
	next *LinkedChar
}

func (lc *LinkedChar) Push(char rune) {
	currentRope := lc
	for currentRope.next != nil {
		currentRope = lc.next
	}
	var newChar LinkedChar = LinkedChar{
		char: char,
	}
	currentRope.next = &newChar
}
