package main

import "fmt"

func MediaAritmetica(nums []int) float64 {
	soma:= 0
	for _,num := range nums {
	soma += num
	}
	return float64(soma) / float64 (len(nums))

	}
