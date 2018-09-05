package main

import (
	"fmt"
	"sync"
)

//自定义集合实现，存byte类型
type Set struct {
	Mu      sync.Mutex
	mapByte map[byte]bool
}

func NewSet(items ...byte) *Set {
	s := &Set{
		mapByte: make(map[byte]bool, len(items)),
	}
	for _, item := range items {
		s.insert(item)
	}
	return s
}

func (hs *Set) insert(c byte) error {
	hs.Mu.Lock()
	defer hs.Mu.Unlock()
	hs.mapByte[c] = true
	return nil
}

func (hs *Set) del(c byte) error {
	hs.Mu.Lock()
	defer hs.Mu.Unlock()
	hs.mapByte[c] = false
	return nil
}
func (hs *Set) delDeep(c byte) error {
	hs.Mu.Lock()
	defer hs.Mu.Unlock()
	delete(hs.mapByte, c)
	return nil
}

func (hs *Set) contains(c byte) (exists bool) {
	exists = false
	hs.Mu.Lock()
	defer hs.Mu.Unlock()
	if _, ok := hs.mapByte[c]; ok {
		exists = true
	}
	return
}

var vowels = []byte("aeiouAEIUO")

//双指针
//使用双指针，指向待反转的两个元音字符，一个指针从头向尾遍历，一个指针从尾到头遍历。
func reverseVowels(s string) (result []byte) {
	result = make([]byte, len(s))
	set := NewSet(vowels...)
	strs := []byte(s)
	i := 0
	j := len(strs) - 1
	for i < j {
		if !set.contains(strs[i]) {
			result[i] = strs[i]
			i++
		} else if !set.contains(strs[j]) {
			result[j] = strs[j]
			j--
		} else {
			result[i] = strs[j]
			result[j] = strs[i]
			i++
			j--

		}
	}

	return
}

func main() {
	s := "ae"

	fmt.Println(string(reverseVowels(s)))
}
