package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	var buffer bytes.Buffer
	for _, s := range strs {
		buffer.WriteString(fmt.Sprintf("%d#%s", len(s), s))
	}

	return buffer.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	var results []string
	i := 0
	for i < len(strs)-1 {
		j := i
		for string(strs[j]) != "#" {
			j++
		}
		length, err := strconv.ParseInt(string(strs[i:j]), 10, 32)
		if err != nil {
			return results
		}
		results = append(results, string(strs[j+1:j+1+int(length)]))
		i = j + 1 + int(length)
	}
	return results
}
