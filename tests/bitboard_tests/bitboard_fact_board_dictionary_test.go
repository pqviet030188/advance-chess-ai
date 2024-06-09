package bitboard_tests

import (
	"testing"

	. "github.com/pqviet030188/advance-chess-ai/bitboard"
	"github.com/pqviet030188/advance-chess-ai/uint96"
)

func TestNilOnNewFactBoardDictionary(t *testing.T) {
	lookup := NewFactBoardDictionary()
	if lookup == nil {
		t.Errorf("Exepected to not nil")
	}
}

func TestGetPutFactBoardDictionary(t *testing.T) {
	lookup := NewFactBoardDictionary()
	for range 10 {
		value := uint96.RandUInt96()
		square := RandomiseSquare()
		maskType := RandomiseSquare()

		lookup.Put(square, maskType, &value)
		foundValue, ok := lookup.Get(square, maskType)

		if !ok {
			t.Errorf("Exepected to find the value given the key")
		}

		if value != *foundValue {
			t.Errorf("Expected values to be the same, Result was incorrect, got: %s, want: %s.", foundValue.Str(), value.Str())
		}
	}
}

func TestKeyCountFactBoardDictionary(t *testing.T) {
	lookup := NewFactBoardDictionary()
	count := lookup.KeyCount()
	if count != 0 {
		t.Errorf("Exepected to be empty")
	}

	for i := range 2000 {
		for {
			value := uint96.RandUInt96()
			square := RandomiseSquare()
			maskType := RandomiseSquare()

			_, ok := lookup.Get(square, maskType)
			if ok {
				continue
			}

			lookup.Put(square, maskType, &value)
			break
		}

		count := lookup.KeyCount()

		if uint32(i)+1 != count {
			t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", count, uint32(i)+1)
		}
	}
}

func TestKeysFactBoardDictionary(t *testing.T) {
	lookup := NewFactBoardDictionary()
	count := lookup.KeyCount()
	if count != 0 {
		t.Errorf("Exepected to be empty")
	}

	keyDict := map[FactBoardDictionaryKey]bool{}
	retKeyDict := map[FactBoardDictionaryKey]bool{}

	for range 2 {
		for {

			value := uint96.RandUInt96()
			square := RandomiseSquare()
			maskType := RandomiseSquare()

			_, ok := lookup.Get(square, maskType)
			if ok {
				continue
			}

			lookup.Put(square, maskType, &value)
			bkey := FactBoardDictionaryKey{
				Square:   square,
				MaskType: maskType,
			}

			keyDict[bkey] = true
			break
		}
	}

	keys := lookup.Keys()
	for _, key := range keys {
		bkey := FactBoardDictionaryKey{
			Square:   key.Square,
			MaskType: key.MaskType,
		}

		retKeyDict[bkey] = true
	}

	for key := range keyDict {
		if has := retKeyDict[key]; !has {
			t.Errorf("Expected to contain the key, Result was incorrect, got: %d, %d, %t, want: %d, %d, %t.", key.Square, key.MaskType, has,
				key.Square, key.MaskType, true)
		}
	}
}

func TestSerialiseDeserialiseFactBoardDictionary(t *testing.T) {
	lookup := NewFactBoardDictionary()

	for i := range 100 {
		for {
			value := uint96.RandUInt96()
			square := RandomiseSquare()
			maskType := RandomiseSquare()

			_, ok := lookup.Get(square, maskType)
			if ok {
				continue
			}

			lookup.Put(square, maskType, &value)

			buffer1 := lookup.Serialise()
			lookup1 := NewFactBoardDictionaryFromBytes(buffer1)
			if lookup1.KeyCount() != (uint32(i) + 1) {
				t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", lookup1.KeyCount(), uint32(i)+1)
			}

			lvalue, ok := lookup1.Get(square, maskType)
			if !ok || !lvalue.Equals(value) {
				t.Errorf("Expected to be equals, Result was incorrect, got: %s, want: %s.", lvalue.Str(), value.Str())
			}

			break
		}
	}
}

func TestFileInOutFactBoardDictionary(t *testing.T) {
	lookup := NewFactBoardDictionary()

	for i := range 100 {
		for {
			value := uint96.RandUInt96()
			square := RandomiseSquare()
			maskType := RandomiseSquare()

			_, ok := lookup.Get(square, maskType)
			if ok {
				continue
			}

			lookup.Put(square, maskType, &value)
			lookup.ToFile("../../artifacts/factlookup")

			lookup1 := NewFactBoardDictionaryFromFile("../../artifacts/factlookup")

			if lookup1.KeyCount() != (uint32(i) + 1) {
				t.Errorf("Expected to be equals, Result was incorrect, got: %d, want: %d.", lookup1.KeyCount(), uint32(i)+1)
			}

			lvalue, ok := lookup1.Get(square, maskType)
			if !ok || !lvalue.Equals(value) {
				t.Errorf("Expected to be equals, Result was incorrect, got: %s, want: %s.", lvalue.Str(), value.Str())
			}

			break
		}
	}
}
