package utils

import "strconv"

func StringToUint(s string) (uint, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return uint(i), nil
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
