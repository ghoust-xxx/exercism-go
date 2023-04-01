package bottlesong

import (
	"fmt"
	"strings"
)

var nums []string = []string{
	"Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two",
	"One", "No",
}

func Recite(startBottles, takeDown int) []string {
	var res []string
	for i := len(nums) - startBottles - 1; i < len(nums) - startBottles - 1 + takeDown ; i++ {
		bot1 := "bottles"
		bot2 := "bottles"
		if i == len(nums)-2 {
			bot1 = "bottle"
		} else if i == len(nums)-3 {
			bot2 = "bottle"
		}
		res = append(res, fmt.Sprintf("%s green %s hanging on the wall,", nums[i], bot1))
		res = append(res, fmt.Sprintf("%s green %s hanging on the wall,", nums[i], bot1))
		res = append(res, "And if one green bottle should accidentally fall,")
		res = append(res, fmt.Sprintf("There'll be %s green %s hanging on the wall.",
			strings.ToLower(nums[i+1]), bot2))

		if i != len(nums) - startBottles - 2 + takeDown {
			res = append(res, "")
		}
	}

	return res
}
