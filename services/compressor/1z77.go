package compressor

import (
	"bytes"
	"zip-like/utils"
)

func CompressLZ77(input []byte) ([]byte, error) {
	var result bytes.Buffer
	ht := utils.NewHashTable()
	i := 0
	for i < len(input) {
		matchLen := 0
		matchOffset := 0
		for j := i + 1; j < len(input); j++ {
			len := j - i
			if len > 255 {
				break
			}
			if data, exists := ht.Get(string(input[i : i+len])); exists && bytes.Equal(input[i:i+len], data) {
				matchLen = len
				matchOffset = j - i
			}
		}

		if matchLen > 0 {
			result.WriteByte(0x00) // Indicator for LZ77
			result.WriteByte(byte(matchLen))
			result.WriteByte(byte(matchOffset))
			i += matchLen
		} else {
			result.WriteByte(0x01) // Indicator for literal
			result.WriteByte(input[i])
			ht.Set(string(input[i:i+1]), []byte{input[i]})
			i++
		}
	}
	return result.Bytes(), nil
}

func DecompressLZ77(input []byte) ([]byte, error) {
	var result bytes.Buffer
	ht := utils.NewHashTable()
	i := 0
	for i < len(input) {
		if input[i] == 0x00 {
			matchLen := int(input[i+1])
			matchOffset := int(input[i+2])
			pos := result.Len() - matchOffset
			for j := 0; j < matchLen; j++ {
				result.WriteByte(result.Bytes()[pos+j])
			}
			i += 3
		} else {
			result.WriteByte(input[i+1])
			ht.Set(string([]byte{input[i+1]}), []byte{input[i+1]})
			i += 2
		}
	}
	return result.Bytes(), nil
}
