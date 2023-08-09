package main

import "fmt"

type Leeters struct {
	letters []byte
	len     int
	index   int
}

func NewLeeters() *Leeters {
	leeters := []byte{}
	for i := 'a'; i <= 'z'; i++ {
		leeters = append(leeters, byte(i))
	}
	len := 26
	index := 0

	return &Leeters{
		letters: leeters,
		len:     len,
		index:   index,
	}
}

func (L *Leeters) GetLen() int {
	return L.len
}

func (L *Leeters) GetLerrter() []byte {
	return L.letters
}

func (L *Leeters) GetIndex() int {
	return L.index
}

func (L *Leeters) Move(x byte) int {
	answer := 0
	if L.letters[L.index] == x {
		return answer
	}

	if x > L.letters[L.index] {
		if x-L.letters[L.index] <= 13 {
			answer = int(x - L.letters[L.index])
		} else {
			answer = -(L.len - int(x-L.letters[L.index]))
		}
	} else {
		if L.letters[L.index]-x <= 13 {
			answer = -int(L.letters[L.index] - x)
		} else {
			answer = L.len - int(L.letters[L.index]-x)
		}
	}

	if answer < 0 {
		if L.index+answer < 0 {
			L.index = (L.len + (L.index + answer)) % 26
		} else {
			L.index = (L.index + answer) % 26
		}
	} else {
		L.index = (L.index + answer) % 26
	}

	return answer
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func minTimeToType(word string) int {
	answer := 0
	leeters := NewLeeters()
	for i := range word {
		if move := leeters.Move(word[i]); move == 0 {
			answer++
			continue
		} else {
			answer = answer + abs(move)
			answer++
			continue
		}
	}
	return answer
}

func main() {
	word := "zjpc"
	fmt.Println(minTimeToType(word))
}
