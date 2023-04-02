package twelve

import "fmt"

var gave []string = []string{
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"a Partridge in a Pear Tree",
}

var nums []string = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

func Verse(i int) string {
	var gifts string
	for j := i - 1; j >= 0; j-- {
		if j == 0 && i != 1 {
			gifts += "and "
		}
		gifts += gave[len(gave) - 1 - j]
		if j == 0 {
			gifts += "."
		} else {
			gifts += ", "
		}
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s",
		nums[i-1], gifts)
}

func Song() string {
	res := ""
	for i := 1; i <= 12; i++ {
		res += Verse(i)
		if i != 12 {
			res += "\n"
		}
	}
	return res
}
