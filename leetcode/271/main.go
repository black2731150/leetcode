package main

import "fmt"

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	key := "#A!$"
	encode := ""
	for _, v := range strs {
		encode = encode + v + key
	}
	return encode
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	key := "#A!$"
	decode := []string{}
	last := 0
	for i := 0; i+3 < len(strs); i++ {
		if strs[i:i+4] == key {
			decode = append(decode, strs[last:i])
			last = i + 4
		}
	}

	return decode
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));

func main() {
	var codec Codec
	strs := []string{"Hello", "World"}
	encode := codec.Encode(strs)
	fmt.Println(encode)
	decode := codec.Decode(encode)
	fmt.Println(decode)
}
