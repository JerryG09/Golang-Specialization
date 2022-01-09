package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
    "sync"
)

/*
    Write a program to sort an array of integers. The program should partition the array into 4 parts,
    each of which is sorted by a different goroutine.

    Each partition should be of approximately equal size.
    Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

    The program should prompt the user to input a series of integers.
    Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
    When sorting is complete, the main goroutine should print the entire sorted list.
*/

func sortArr(arr []int, wait *sync.WaitGroup) {
    fmt.Println("Subarray:", arr)
    sort.Ints(arr)
    wait.Done()
}

func main() {
    var wait sync.WaitGroup
    numSlices := 4

    fmt.Print("Enter a sequence of integers (separated by spaces): ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    sequence := scanner.Text()

    inputArr := strings.Fields(sequence)
    intArr := make([]int, 0, len(inputArr))

    for _, n := range inputArr {
        v, _ := strconv.Atoi(n)
        intArr = append(intArr, v)
    }

    wait.Add(numSlices)
    intervalSize := len(intArr) / numSlices

    fmt.Println("Array to be sorted:", intArr)

    go sortArr(intArr[ : intervalSize], &wait)
    go sortArr(intArr[ intervalSize : intervalSize * 2], &wait)
    go sortArr(intArr[intervalSize * 2 : intervalSize * 3], &wait)
    go sortArr(intArr[intervalSize * 3: ], &wait)

    wait.Wait()

    fmt.Println("Array after all subarray is sorted:", intArr)
    sort.Ints(intArr)
    fmt.Println("The sorted list is", intArr)
}