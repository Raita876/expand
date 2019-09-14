package expand

import "strings"

// Cartesian method
func Cartesian(list [][]string) []string {
	strList := cartesianHelper(list, [][]string{nil})
	return toQueryParameterStrList(strList)
}

func cartesianHelper(list [][]string, result [][]string) [][]string {
	if 0 == len(list) {
		return result
	}
	var s [][]string
	for _, r := range result {
		for _, e := range list[0] {
			s = append(s, append(r, e))
		}
	}
	return cartesianHelper(list[1:], s)
}

func toQueryParameterStrList(str2DArray [][]string) []string {
	var qpsList []string
	for _, sl := range str2DArray {
		qps := strings.Join(sl, "&")
		qpsList = append(qpsList, qps)
	}

	return qpsList
}
