package bitboard

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"io"
	"os"

	"github.com/pqviet030188/advance-chess-ai/uint96"
)

type BoardDictionary struct {

	// map from little endian to big endian, square (0-9) -> uint96
	Indices *map[uint32]map[uint32]map[uint32]map[uint8]uint96.Uint96
}

type BoardDictionaryKey struct {
	Occupancy uint96.Uint96
	Square    uint8
}

func NewBoardDictionary() *BoardDictionary {
	var ret = new(BoardDictionary)

	ret.Indices = &map[uint32]map[uint32]map[uint32]map[uint8]uint96.Uint96{}

	return ret
}

func (dict *BoardDictionary) KeyCount() uint32 {
	c := uint32(0)
	for lowKey := range *dict.Indices {
		for midKey := range (*dict.Indices)[lowKey] {
			for hiKey := range (*dict.Indices)[lowKey][midKey] {
				for range (*dict.Indices)[lowKey][midKey][hiKey] {
					c++
				}
			}
		}
	}

	return c
}

func (dict *BoardDictionary) Keys() []BoardDictionaryKey {
	var keys []BoardDictionaryKey

	for lowKey := range *dict.Indices {
		for midKey := range (*dict.Indices)[lowKey] {
			for hiKey := range (*dict.Indices)[lowKey][midKey] {
				for square := range (*dict.Indices)[lowKey][midKey][hiKey] {
					key := BoardDictionaryKey{
						Occupancy: uint96.Uint96{
							Lo:  lowKey,
							Mid: midKey,
							Hi:  hiKey,
						},
						Square: square,
					}
					keys = append(keys, key)
				}
			}
		}
	}

	return keys
}

func (dict *BoardDictionary) Get(key *uint96.Uint96, square uint8) (*uint96.Uint96, bool) {
	search, loFound := (*dict.Indices)[key.Lo]
	if !loFound {
		return nil, false
	}

	midSearch, midFound := search[key.Mid]
	if !midFound {
		return nil, false
	}

	res, hiFound := midSearch[key.Hi]
	if !hiFound {
		return nil, false
	}

	sres, found := res[square]
	if !found {
		return nil, false
	}

	copied := sres.Copy()
	return &copied, true
}

func (dict *BoardDictionary) Put(key *uint96.Uint96, square uint8, value *uint96.Uint96) {
	// dict.indices = map[uint32]map[uint32]map[uint32]map[uint8]uint96.Uint96{}

	indices := dict.Indices

	if _, has := (*indices)[key.Lo]; !has {
		(*indices)[key.Lo] = map[uint32]map[uint32]map[uint8]uint96.Uint96{}
	}

	if _, has := (*indices)[key.Lo][key.Mid]; !has {
		(*indices)[key.Lo][key.Mid] = map[uint32]map[uint8]uint96.Uint96{}
	}

	if _, has := (*indices)[key.Lo][key.Mid][key.Hi]; !has {
		(*indices)[key.Lo][key.Mid][key.Hi] = map[uint8]uint96.Uint96{}
	}

	(*indices)[key.Lo][key.Mid][key.Hi][square] = *value
}

func (dict *BoardDictionary) Serialise() *bytes.Buffer {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(dict)

	if err != nil {
		panic(err)
	}

	return buffer
}

func NewBoardDictionaryFromBytes(buffer *bytes.Buffer) *BoardDictionary {

	ret := NewBoardDictionary()
	decoder := gob.NewDecoder(buffer)

	// Decoding
	err := decoder.Decode(ret)
	if err != nil {
		panic(err)
	}

	return ret
}

func (dict *BoardDictionary) ToFile(oPath string) {

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

func NewBoardDictionaryFromFile(oPath string) *BoardDictionary {
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

	return NewBoardDictionaryFromBytes(data)
}
