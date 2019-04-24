package algorithms

import (
	"fmt"
	"sort"
)
var ToBeSorted [10] int = [10] int {100,51,90,77,5,7,8,12,10,7}
// Les https://en.wikipedia.org/wiki/Bubble_sort
func Bubble_sort_modified(list [10]int) {
	// Deres kode her
	fmt.Println(list)
	n := len(list)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if list[i-1] > list[i] {
				fmt.Println("swapping", list)
				list[i-1], list[i] = list[i], list[i-1]
				swapped = true

			}

		}
	}
	fmt.Println(list)
}

// Implementering av Bubble_sort algoritmen
func Bubble_sort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

func Test_Bubble_sort_map() {

	m := map[string]int {
		"a":	5,
		"b":	6,
		"c":	7,
		" ":	17,
		"e":	10,
	}


	fmt.Println(Bubble_sort_map(m))
}

func Bubble_sort_map(toSort map[string]int) PairList {
	pl := make(PairList, len(toSort))
	i := 0
	for k, v := range toSort {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key		string
	Value	int
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less() int {
	return len(p)
}

func (p PairList) Swap() int {
	return len(p)
}

// Implementering av Quicksort algoritmen
func QSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
