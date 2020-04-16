package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var length int = 100000
	var sequence []int
	var sortedSequence []int
	sequence = generateSequence(length)

	//fmt.Println("Sequence before merge sort:\n", sequence)

	sortedSequence = sort(sequence)

	//fmt.Println("Sequence after merge sort:\n", sortedSequence)

	var sorted bool = testAlgo(sortedSequence)
	if sorted == true {
		fmt.Println("The algorithm successfully sorted the array!")
	} else {
		fmt.Println("The algorithm did not successfully sort the array :(")
	}
}

func testAlgo(seq []int) bool {
	var pass bool = false
	for i := 1; i < len(seq)-1; i++ {
		if seq[i] > seq[i-1] {
			pass = true
		} else {
			pass = false
		}
	}
	return pass
}

func generateSequence(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice

}

func merge(l, r []int) (result []int) {
	result = make([]int, len(l)+len(r))

	i := 0

	for len(l) > 0 && len(r) > 0 {
		if l[0] < r[0] {
			result[i] = l[0]
			l = l[1:]
		} else {
			result[i] = r[0]
			r = r[1:]
		}
		i++
	}
	for j := 0; j < len(l); j++ {
		result[i] = l[j]
		i++
	}
	for j := 0; j < len(r); j++ {
		result[i] = r[j]
		i++
	}

	return
}

func sort(sequence []int) []int {
	if len(sequence) == 1 {
		return sequence
	}

	mid := int(len(sequence) / 2)
	var leftSeq = make([]int, mid)
	var rightSeq = make([]int, len(sequence)-mid)

	for i := 0; i < len(sequence); i++ {
		if i < mid {
			leftSeq[i] = sequence[i]
		} else {
			rightSeq[i-mid] = sequence[i]
		}
	}

	return merge(sort(leftSeq), sort(rightSeq))
}
