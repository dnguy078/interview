package main

// [1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// [0 0 0 0 1 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// next iteration
// [1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
// [0 0 0 1 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var s1Freq [26]int
	var s2Freq [26]int
	i := 0
	for ; i < len(s1); i++ {
		s1Freq[s1[i]-'a']++
		s2Freq[s2[i]-'a']++
	}
	if compare(s1Freq, s2Freq) {
		return true
	}
	for ; i < len(s2); i++ {
		s2Freq[s2[i-len(s1)]-'a']--
		s2Freq[s2[i]-'a']++
		if compare(s1Freq, s2Freq) {
			return true
		}
	}
	return false
}

func compare(a1, a2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
