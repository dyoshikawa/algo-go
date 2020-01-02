package p127

import "strings"

func Invoke(n int, strs []string) []string {
	var res []string
	m := make(map[string]struct{})
	for _, s := range strs {
		splited := strings.Split(s, " ")
		ord := splited[0]
		w := splited[1]
		if ord == "insert" {
			// insert
			m[w] = struct{}{}
		} else {
			// find
			var r string
			if _, exists := m[w]; exists {
				r = "yes"
			} else {
				r = "no"
			}
			res = append(res, r)
		}
	}
	return res
}
