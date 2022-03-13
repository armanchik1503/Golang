package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("ParsePhone(%q) = %q\n", "123-456-7890", ParsePhone("123-456-7890"))
	fmt.Println(Anagram("cat", "arman"))
	arr := []int32{1, 2, 2, 4, 4, 29, 17, 12, 11}
	fmt.Println(FindEvens(arr))
	fmt.Println(SliceProduct([]int32{1, 2, 3, 4}))
	fmt.Println(Unique(arr))
	fmt.Printf("InvertMap(%v) = %v\n", map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}, InvertMap(map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}))
	fmt.Printf("TopCharacters(%v, %v) = %v", "123123122", 2, TopCharacters("123123122", 2))
}

func ParsePhone(phone string) string {
	digits := make([]byte, 10)
	index := 0
	for _, ch := range phone {
		if unicode.IsDigit(ch) {
			digits[index] = byte(ch)
			index++
		}
	}

	fst := string(digits[:3])
	mid := string(digits[3:6])
	end := string(digits[6:10])
	return fmt.Sprintf("(%v) %v-%v", fst, mid, end)
}

func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	hash := make(map[string]int)

	for _, r := range s1 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}

	for _, r := range s2 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}

	var isAnagram bool = true
	for _, value := range hash {
		if value%2 != 0 {
			isAnagram = false
			break
		}
	}

	return isAnagram
}

func FindEvens(e []int32) []int32 {
	var result []int32
	for i := 0; i < len(e); i++ {
		if e[i]%2 == 0 {
			result = append(result, e[i])
		}
	}
	return result
}

func SliceProduct(e []int32) int32 {
	var total int32
	for i := 0; i < len(e); i++ {
		total += e[i]
	}
	return total
}

func Unique(e []int32) []int32 {
	keys := make(map[int]bool)
	var list []int32
	for _, entry := range e {
		if _, value := keys[int(entry)]; !value {
			keys[int(entry)] = true
			list = append(list, entry)
		}
	}
	return list
}

func InvertMap(kv map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range kv {
		n[v] = k
	}
	return n
}

func TopCharacters(s string, k int) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	for r, cnt := range counts {
		if cnt <= k {
			delete(counts, r)
		}
	}
	return counts
}
