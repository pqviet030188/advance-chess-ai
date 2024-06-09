package bitboard

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"io"
	"os"

	"github.com/pqviet030188/advance-chess-ai/uint96"
)

type FactBoardDictionary struct {

	// square->mask types->board
	Indices *map[uint8]map[uint8]uint96.Uint96
}

type FactBoardDictionaryKey struct {
	Square   uint8
	MaskType uint8
}

func NewFactBoardDictionary() *FactBoardDictionary {
	var ret = new(FactBoardDictionary)
	ret.Indices = &map[uint8]map[uint8]uint96.Uint96{}
	return ret
}

func (dict *FactBoardDictionary) KeyCount() uint32 {
	c := uint32(0)
	for square := range *dict.Indices {
		for range (*dict.Indices)[square] {
			c++
		}
	}

	return c
}

func (dict *FactBoardDictionary) Keys() []FactBoardDictionaryKey {
	var keys []FactBoardDictionaryKey

	for square := range *dict.Indices {
		for maskType := range (*dict.Indices)[square] {

			key := FactBoardDictionaryKey{
				MaskType: maskType,
				Square:   square,
			}

			keys = append(keys, key)
		}
	}

	return keys
}

func (dict *FactBoardDictionary) Get(square uint8, maskType uint8) (*uint96.Uint96, bool) {
	search, found := (*dict.Indices)[square]
	if !found {
		return nil, false
	}

	maskSearch, found := search[maskType]
	if !found {
		return nil, false
	}

	return &maskSearch, true
}

func (dict *FactBoardDictionary) Put(square uint8, maskType uint8, value *uint96.Uint96) {
	indices := dict.Indices

	if _, has := (*indices)[square]; !has {
		(*indices)[square] = map[uint8]uint96.Uint96{}
	}

	(*indices)[square][maskType] = *value
}

func (dict *FactBoardDictionary) Serialise() *bytes.Buffer {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(dict)

	if err != nil {
		panic(err)
	}

	return buffer
}

func NewFactBoardDictionaryFromBytes(buffer *bytes.Buffer) *FactBoardDictionary {

	ret := NewFactBoardDictionary()
	decoder := gob.NewDecoder(buffer)

	// Decoding
	err := decoder.Decode(ret)
	if err != nil {
		panic(err)
	}

	return ret
}

func (dict *FactBoardDictionary) ToFile(oPath string) {

	data := dict.Serialise()
	reader := bufio.NewReader(data)

	fo, err := os.Create(oPath)

	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a writer
	writer := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)

	for {
		// read a chunk
		n, err := reader.Read(buf)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}

		// write a chunk
		if _, err := writer.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err = writer.Flush(); err != nil {
		panic(err)
	}
}

func NewFactBoardDictionaryFromFile(oPath string) *FactBoardDictionary {
	fi, err := os.Open(oPath)
	if err != nil {
		panic(err)
	}

	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// make a read buffer
	data := new(bytes.Buffer)
	reader := bufio.NewReader(fi)
	writer := bufio.NewWriter(data)

	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := writer.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err = writer.Flush(); err != nil {
		panic(err)
	}

	return NewFactBoardDictionaryFromBytes(data)
}
