package main

import "fmt"

// LFSR represents a Linear Feedback Shift Register.
type LFSR struct {
	state        uint16
	feedbackTaps []int
	registerLen  int
}

// NewLFSR creates a new LFSR with the given initial state and feedback taps.
func NewLFSR(initialState uint16, feedbackTaps []int, registerLen int) *LFSR {
	return &LFSR{
		state:        initialState,
		feedbackTaps: feedbackTaps,
		registerLen:  registerLen,
	}
}

// Shift performs one shift operation on the LFSR.
func (l *LFSR) Shift() {
	feedbackBit := uint16(0)
	for _, tap := range l.feedbackTaps {
		feedbackBit ^= (l.state >> (tap - 1)) & 1 // XOR the tapped bits
	}

	// Shift the state to the right and set the leftmost bit to the feedback result
	l.state = (l.state >> 1) | (feedbackBit << (l.registerLen - 1))
}

// GenerateSequence generates a pseudo-random sequence of the specified length.
func (l *LFSR) GenerateSequence(length int) []uint16 {
	sequence := make([]uint16, length)
	for i := 0; i < length; i++ {
		sequence[i] = l.state & 1 // Append the least significant bit
		l.Shift()                 // Shift the LFSR state
	}
	return sequence
}

func main() {
	x1 := encryptMsg("b2", generate16BitsStreamX())

	x2 := encryptMsg("b9", generate16BitsStreamX())

	fmt.Println(x1, x2)
}

// 1 0 0 1 0 0 1 1

// 0011 0010

// 0011 1001

// 1 0 0 1 1 0 0 0

//

func generate16BitsStreamX() []uint16 {
	lfsr1 := NewLFSR(uint16(0b001), []int{2, 1}, 3)
	R1 := lfsr1.GenerateSequence(16)

	lfsr2 := NewLFSR(uint16(0b01001), []int{5, 1, 2, 3}, 5)
	R3 := lfsr2.GenerateSequence(16)

	lfsr := NewLFSR(uint16(0b1011), []int{2, 1}, 4)
	R2 := lfsr.GenerateSequence(16) // Generate 15 bits

	//fmt.Println("Generated R1 Sequence:", R1)
	//fmt.Println("Generated R2 Sequence:", R2)
	//fmt.Println("Generated R3 Sequence:", R3)

	streamX := []uint16{}
	R2Point := -1
	R3Point := -1
	for i := 0; i < 16; i++ {
		if R1[i] == 1 {
			R2Point++
			if R3Point == -1 {
				streamX = append(streamX, R2[R2Point]^0)
			} else {
				streamX = append(streamX, R2[R2Point]^R3[R3Point])
			}
		} else {
			R3Point++
			if R2Point == -1 {
				streamX = append(streamX, R3[R3Point]^0)
			} else {
				streamX = append(streamX, R3[R3Point]^R2[R2Point])
			}
		}
	}

	//fmt.Println("The first 16 bit of stream is", streamX)
	return streamX
}
