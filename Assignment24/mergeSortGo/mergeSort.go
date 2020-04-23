package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	runTest(50)
}

func runTest(numTests int) {
	var length = 1000
	var sequence []int
	var sortedSequence []int
	var times = make([]float64, numTests)
	var sorted bool

	for i := 0; i < numTests; i++ {
		sequence = generateSequence(length)
		startTime := time.Now()
		sortedSequence = sort(sequence)
		totalTime := time.Now().Sub(startTime).Milliseconds()
		times[i] = float64(totalTime) / 1000.00
		sorted = testAlgo(sortedSequence)
		if !sorted {
			fmt.Printf("Sequence of size: %d failed to sort.", length)
		}
		length += 2000
	}
	// for i := 0; i < len(times); i++ {
	fmt.Print(times)
	// }

}

func testAlgo(seq []int) bool {
	var pass bool = false
	for i := 1; i < len(seq)-1; i++ {
		if seq[i] >= seq[i-1] {
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
