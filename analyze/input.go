package analyze

import (
	"fmt"

	"github.com/JohnnyCPC/tw-lottery-linebot-go/combinations"
)

func BuildInputData(in []int) (res []string) {

	var reary []int
	lslice := make([]int, 40)
	combinations := combinations.All(in)

	for i := 0; i < len(combinations); i++ {
		lslice[0] = 1
		reary = make([]int, len(combinations[i]))
		for j := 0; j < len(combinations[i]); j++ {
			reary[j] = combinations[i][j]
			lslice[combinations[i][j]] = 1
		}

		final := uint64(lslice[0])

		for i := 1; i < len(lslice); i++ {
			final <<= 1
			final += uint64(lslice[i])
		}
		final -= 0x8000000000

		reshex := fmt.Sprintf("%010X", final)

		res = append(res, reshex)

		lslice = make([]int, 40)
	}
	return
}
