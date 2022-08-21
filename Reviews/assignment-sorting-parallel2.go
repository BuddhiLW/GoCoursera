package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"sync"
)

func swap(seq []int, i int) {
	temp := seq[i]
	seq[i] = seq[i+1]
	seq[i+1] = temp
}

func sortInt(seq []int, wg *sync.WaitGroup) {
	fmt.Printf("sort the seq: %v and it's len: %d\n", seq, len(seq))

	for i := 0; i < len(seq)-1; i++ {
		for j := 0; j < len(seq)-1-i; j++ {
			if seq[j] > seq[j+1] {
				swap(seq, j)
			}
		}
	}

	//fmt.Printf("end sort the seq: %v and it's len: %d\n", seq, len(seq))
	
	if wg != nil {
		wg.Done()
	}
}

func main() {
	fmt.Printf("Please input a series of integers >> ")

	var indexSpace int
	var strInput, str1, str2 string
	listInt := []int {}

	reader := bufio.NewReader(os.Stdin)
	strInput,_ = reader.ReadString('\n')
	strInput = strings.TrimSpace(strInput)
	var strInputLowcase = strings.ToLower(strInput)
	//var strLen = len(strInput)

	//indexSpace = strings.Index(strInputLowcase, " ")
	//str1 = strInputLowcase[0:indexSpace]
	//str2 = strInputLowcase[indexSpace+1:len(strInputLowcase)]

	str2 = strInputLowcase

	for {
		indexSpace = strings.Index(str2, " ")
		if indexSpace == -1 {
			value, err:= strconv.Atoi(str2)
			if err != nil {
				fmt.Printf("The series of integar is incorrect !!!\n")
				break
			} else {
				listInt = append(listInt, value)
			}

			break
		} else {
			str1 = str2[0:indexSpace]
			str2 = str2[indexSpace+1:len(str2)]
			//fmt.Printf("str1: %s str2: %s\n", str1, str2)
			value, err:= strconv.Atoi(str1)
			if err != nil {
				fmt.Printf("The series of integar is incorrect !!!\n")
				break
			} else {
				listInt = append(listInt, value)
			}
		}
	}

	var wg sync.WaitGroup

	var lenInt = len(listInt)
	//fmt.Printf("1 listInt: %v len: %d\n", listInt, lenInt)

	var remLen = lenInt
	var newLen int
	indexStart := 0
	for i:=4; i>0; i-- {
		newLen = remLen/i
		remLen = remLen - newLen
		//fmt.Printf("i: %d len: %d 1/4len: %d curlen: %d indexStart: %d\n", i, lenInt, newLen, remLen, indexStart)
		go sortInt(listInt[indexStart:indexStart+newLen], &wg)
		wg.Add(1)
		indexStart = indexStart + newLen
	}

	wg.Wait()

	fmt.Printf("The series of integer after sorting in goroutine is: %v\n", listInt)

	sortInt(listInt, nil)
	fmt.Printf("The series of sorted integer is: %v\n", listInt)
}

/*
11 10 9 8 7 6 5 4 3 2 1 0 12
*/