package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
)

type Team struct {
	name string;
	mp, w, d, l, p int;
}

func Tally(reader io.Reader, writer io.Writer) error {
	teams := []Team{}
	file := csv.NewReader(reader)
	file.Comma = ';'
	file.Comment = '#'
	for {
		res, err := file.Read()
		if err == io.EOF {
			break
		}
		if len(res) == 0 {
			continue
		}
		if len(res) != 3 {
			return errors.New("error")
		}
		t1, t2 := -1, -1
		for i, val := range teams {
			if val.name == res[0] {
				t1 = i
			}
			if val.name == res[1] {
				t2 = i
			}
		}
		if t1 == -1 {
			teams = append(teams, Team{})
			t1 = len(teams) - 1
			teams[t1].name = res[0]
		}
		if t2 == -1 {
			teams = append(teams, Team{})
			t2 = len(teams) - 1
			teams[t2].name = res[1]
		}
		teams[t1].mp++
		teams[t2].mp++
		switch res[2] {
		case "win":
			teams[t1].w++
			teams[t1].p += 3
			teams[t2].l++
		case "loss":
			teams[t1].l++
			teams[t2].w++
			teams[t2].p += 3
		case "draw":
			teams[t1].d++
			teams[t1].p++
			teams[t2].d++
			teams[t2].p++
		default:
			return errors.New("error")
		}
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].p == teams[j].p {
			return teams[i].name < teams[j].name
		}
		return teams[i].p > teams[j].p
	})

	fmt.Fprintf(writer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, val := range teams {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			val.name, val.mp, val.w, val.d, val.l, val.p)
	}

	return nil
}
