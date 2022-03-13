package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("========== Problem 1: Sorting Names ==========")
	people := PersonSlice{NewPerson("Ziyi", "Yan"), NewPerson("Xuefei", "Chen"), NewPerson("Zida", "Liu")}
	fmt.Printf("Before: %v\n", people)
	sort.Sort(people)
	fmt.Printf("After: %v\n", people)

	fmt.Println("========== Problem 2: IsPalindrome Redux ==========")
	first := NewPerson("Mr.", "First")
	second := NewPerson("Mr.", "Second")
	them := PersonSlice{first, second, first}
	fmt.Printf("IsPalindrome(%v) = %v\n", them, IsPalindrome(them))

	fmt.Println("========== Problem 3: Functional Programming ==========")
	add := func(x, y int) int { return x + y }
	fmt.Printf("Fold([]int{1, 2, 3, 4, 5}, 0, add) = %v", Fold([]int{1, 2, 3, 4, 5}, 0, add))
	mult := func(x, y int) int { return x * y }
	fmt.Printf("Fold([]int{1, 2, 3, 4, 5}, 1, mult) = %v", Fold([]int{1, 2, 3, 4, 5}, 1, mult))
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

func NewPerson(first, last string) *Person {
	person := new(Person)
	person.ID = nextPersonID
	person.FirstName = first
	person.LastName = last
	nextPersonID++
	return person
}

var nextPersonID int = 1

func (p *Person) String() string {
	return fmt.Sprintf("%s %s(%v)", p.FirstName, p.LastName, p.ID)
}

type PersonSlice []*Person

func (ps PersonSlice) Len() int {
	return len([]*Person(ps))
}

func (ps PersonSlice) Less(i, j int) bool {
	if c := strings.Compare(ps[i].LastName, ps[j].LastName); c != 0 {
		return c < 0
	}
	if c := strings.Compare(ps[i].FirstName, ps[j].FirstName); c != 0 {
		return c < 0
	}
	return ps[i].ID < ps[j].ID
}

func (ps PersonSlice) Swap(i, j int) {
	p := ps[j]
	ps[j] = ps[i]
	ps[i] = p
}

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if s.Less(i, s.Len()-1-i) || s.Less(s.Len()-1-i, i) {
			return false
		}
	}
	return true
}

func Fold(s []int, v int, f func(int, int) int) int {
	if len(s) == 0 {
		return v
	}

	return Fold(s[1:], f(v, s[0]), f)
}
