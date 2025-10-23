package main

import "iter"

type targetList []int

func (list targetList) FilterBefore(match func(int) bool) targetList {
	filtered := make(targetList, 0, len(list)) // The bottleneck
	for _, e := range list {
		if match(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func (list targetList) FilterAfter(match func(int) bool) targetList {
	n := 0
	for _, e := range list {
		if match(e) {
			list[n] = e
			n++
		}
	}
	list = list[:n]
	return list
}

func (list targetList) FilerIter(match func(int) bool) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, element := range list {
			if match(element) {
				if !yield(element) {
					return
				}
			}
		}
	}
}

func main() {
	var list targetList = make([]int, 100)
	for i := 0; i < 100; i++ {
		list[i] = i
	}

	_ = list.FilterAfter(func(i int) bool {
		return i%2 == 0
	})
}
