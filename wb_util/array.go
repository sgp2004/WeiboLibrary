package wb_util

func FindIntPos(s int, array []int) int {
	for p, v := range array {
		if v == s {
			return p
		}
	}
	return -1
}

func FindStrPos(s string, array []string) int {
	for p, v := range array {
		if v == s {
			return p
		}
	}
	return -1
}
