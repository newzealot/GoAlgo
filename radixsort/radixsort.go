package radixsort

import (
	"fmt"
	"math"
)

func GetDigit(num, p int) int {
	if num < 0 {
		num = 0 - num
	}
	return (num / int(math.Pow(10, float64(p)))) % 10
}

func DigitCount(num int) int {
	if num < 0 {
		num = 0 - num
	}
	s := fmt.Sprint(num)
	return len(s)
}

func MostDigits(arr []int) (mostDigits int) {
	for _, v := range arr {
		if d := DigitCount(v); d > mostDigits {
			mostDigits = d
		}
	}
	return
}

func RadixSort(arr []int) []int {
	maxLoops := MostDigits(arr)
	buckets := [10][]int{}
	inputArr := arr
	for i := 0; i < maxLoops; i++ {
		outputArr := make([]int, 0)
		// sort to bucket
		for _, v := range inputArr {
			digit := GetDigit(v, i)
			buckets[digit] = append(buckets[digit], v)
		}
		// 2D bucket to 1D slice
		for j := 0; j < 10; j++ {
			outputArr = append(outputArr, buckets[j]...)
		}
		// reset output and buckets for the next pass
		inputArr = outputArr
		buckets = [10][]int{}
	}
	return inputArr
}
