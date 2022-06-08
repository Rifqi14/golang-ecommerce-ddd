package str

import "strconv"

func StringToInt(t string) int {
	v, err := strconv.Atoi(t)
	if err != nil {
		return 0
	}
	return v
}
