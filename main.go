package main

import "fmt"

func main() {
	fmt.Println(sockMerchant(int32(9), []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}))
}

func sockMerchant(n int32, ar []int32) int32 {
	dic := make(map[int32]int32)
	pairsCount := int32(0)
	for i := int32(0); i < n; i++ {
		dic[ar[i]]++
		if dic[ar[i]]%2 == 0 {
			pairsCount++
		}
	}
	return pairsCount
}

// Sample Input
// STDIN                       Function
// -----                       --------
// 9                           n = 9
// 10 20 20 10 10 30 50 10 20  ar = [10, 20, 20, 10, 10, 30, 50, 10, 20]

// Sample Output
// 3
