package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

//Write a program to sort an array of integers.
//The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
//Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted
//subarrays into one large sorted array.
//
//The program should prompt the user to input a series of integers.
//Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete,
//the main goroutine should print the entire sorted list.

func main() {

	var integers []int
	fmt.Println("Keep adding integers to sort, press q when you decide you wanna sort")
	for {
		var integerStr string
		fmt.Print("Integer to add:\t")
		_, _ = fmt.Scan(&integerStr)
		if integerStr == "q" {
			break
		}
		integer, _ := strconv.Atoi(integerStr)
		integers = append(integers, integer)
	}

	// Big amounts of integers, worth it, otherwise not
	//rand.Seed(time.Now().Unix())
	//integers = rand.Perm(100000)
	fmt.Printf("Sorting %v length array\n", len(integers))
	tmp1 := make([]int, len(integers))
	tmp2 := make([]int, len(integers))
	copy(tmp1, integers)
	copy(tmp2, integers)
	ConcurrentSort(tmp1)
	InlineSort(tmp2)
}

func InlineSort(integers []int) {
	start := time.Now()
	Sort(integers)
	fmt.Printf("Inlinle sort took %v --> %v\n", time.Since(start), integers)
}

func ConcurrentSort(integers []int) {
	start := time.Now()
	dividedArr := SubArrayEvenly(integers)
	channel := make(chan []int)
	for _, part := range dividedArr {
		go SortChannel(part, channel)
	}
	// don't really need to receive the results, as slices are pointers itself,
	// so the dividedArr slider is being modified
	<-channel
	<-channel
	<-channel
	<-channel

	var joined []int
	for _, part := range dividedArr {
		joined = append(joined, part...)
	}
	Sort(joined)
	fmt.Printf("Concurrent sort took %v --> %v\n", time.Since(start), joined)
}

func SubArrayEvenly(integers []int) [][]int {
	var divided [][]int

	chunkSize := int(math.Ceil(float64(len(integers)) / float64(4)))
	for i := 0; i <= len(integers)/chunkSize; i++ {
		start := i * chunkSize
		end := int(math.Min(float64((i+1)*chunkSize), float64(len(integers))))

		part := integers[start:end]
		if len(part) > 0 {
			divided = append(divided, part)
		}
	}
	return divided

}

func SortChannel(integers []int, c chan []int) {
	Sort(integers)
	c <- integers
}

func Sort(integers []int) {
	/* Using manual sorting cause I dunno which algorithm sort.Ints uses, and with this Bubble Sort, goroutines will
	improve the performance */
	for {
		swaps := 0
		for i, _ := range integers {
			if len(integers)-1 == i {
				break
			}
			if integers[i] > integers[i+1] {
				SwapCopy(integers, i)
				swaps += 1
			}
		}
		if swaps == 0 {
			break
		}
	}
}

func SwapCopy(integers []int, position int) {
	integers[position], integers[position+1] = integers[position+1], integers[position]
}