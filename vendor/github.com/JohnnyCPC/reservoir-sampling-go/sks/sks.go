package sks

import (
	"math/rand"
	"sort"
	"time"
)

func SelectKItems(streams []int, n int, k int) []int {

	var i int
	reservoir := make([]int, k)
	rand.Seed(time.Now().UnixNano())

	for i = 0; i < k; i++ {
		reservoir[i] = streams[i]
	}
	for ; i < n; i++ {
		j := rand.Intn(i + 1)
		if j < k {
			reservoir[j] = streams[i]
		}
	}
	sort.Ints(reservoir)

	return reservoir
}
