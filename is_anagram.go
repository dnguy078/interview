package main

func isAnagram(s string, t string) bool {
	var occurence [26]int
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		occurence[s[i]-'a']++
		occurence[t[i]-'a']--
	}
	for _, c := range occurence {
		if c != 0 {
			return false
		}
	}
	return true
}
