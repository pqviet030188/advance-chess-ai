package board_dictionary_tests

import (
	"crypto/rand"
	"encoding/binary"
	"testing"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/uint96"
)

func TestNilOnNew(t *testing.T) {
	lookup := NewBoardDictionary()
	if lookup == nil {
		t.Errorf("Exepected to not nil")
	}
}

func RandomiseSquare() uint8 {

	squareRandBuf := make([]byte, 2)
	rand.Read(squareRandBuf)
	return uint8(binary.LittleEndian.Uint16(squareRandBuf[:2]))
}

func TestGetPut(t *testing.T) {
	lookup := NewBoardDictionary()
	for range 10 {
		key := uint96.RandUInt96()
		value := uint96.RandUInt96()
		square := RandomiseSquare()

		lookup.Put(&key, square, &value)
		foundValue, ok := lookup.Get(&key, square)

		if !ok {
			t.Errorf("Exepected to find the value given the key")
		}

		if value != *foundValue {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %s, want: %s.", foundValue.Str(), value.Str())
		}
	}
}

func TestKeyCount(t *testing.T) {
	lookup := NewBoardDictionary()
	count := lookup.KeyCount()
	if count != 0 {
		t.Errorf("Exepected to be empty")
	}

	for i := range 2000 {
		for {
			key := uint96.RandUInt96()
			square := RandomiseSquare()
			_, ok := lookup.Get(&key, square)
			if ok {
				continue
			}

			value := uint96.RandUInt96()
			lookup.Put(&key, square, &value)
			break
		}

		count := lookup.KeyCount()

		if uint32(i)+1 != count {
			t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", count, uint32(i)+1)
		}
	}
}

func TestSerialiseDeserialise(t *testing.T) {
	lookup := NewBoardDictionary()

	for i := range 100 {
		for {
			key := uint96.RandUInt96()
			square := RandomiseSquare()

			_, ok := lookup.Get(&key, square)
			if ok {
				continue
			}

			value := uint96.RandUInt96()

			lookup.Put(&key, square, &value)

			buffer1 := lookup.Serialise()
			lookup1 := NewBoardDictionaryFromBytes(buffer1)
			if lookup1.KeyCount() != (uint32(i) + 1) {
				t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", lookup1.KeyCount(), uint32(i)+1)
			}

			lvalue, ok := lookup1.Get(&key, square)
			if !ok || !lvalue.Equals(value) {
				t.Errorf("Expected to be equals, Result was incorrect, got: %s, want: %s.", lvalue.Str(), value.Str())
			}

			break
		}
	}
}

func TestFileInOut(t *testing.T) {
	lookup := NewBoardDictionary()

	for i := range 100 {
		for {
			key := uint96.RandUInt96()
			square := RandomiseSquare()

			_, ok := lookup.Get(&key, square)
			if ok {
				continue
			}

			value := uint96.RandUInt96()

			lookup.Put(&key, square, &value)
			lookup.ToFile("../../artifacts/movelookup")

			lookup1 := NewBoardDictionaryFromFile("../../artifacts/movelookup")

			if lookup1.KeyCount() != (uint32(i) + 1) {
				t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", lookup1.KeyCount(), uint32(i)+1)
			}

			lvalue, ok := lookup1.Get(&key, square)
			if !ok || !lvalue.Equals(value) {
				t.Errorf("Expected to be equals, Result was incorrect, got: %s, want: %s.", lvalue.Str(), value.Str())
			}

			break
		}
	}
}
