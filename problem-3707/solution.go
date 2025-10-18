package main

func score(s string) int {

	value := 0

	for i := range s {
		value += int(s[i]-'a') + 1
	}

	return value
}

func scoreBalance(s string) bool {

	j_limit := len(s)

	for j := 1; j < j_limit; j++ {
		if score(s[:j]) == score(s[j:]) {
			return true
		}
	}

	return false
}
