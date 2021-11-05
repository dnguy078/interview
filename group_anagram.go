package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[string][]string)
	for _, word := range strs {
		sortedWord := sortWord(word)
		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}
	var anagramList [][]string

	for _, list := range anagramGroups {
		anagramList = append(anagramList, list)
	}
	return anagramList
}

func sortWord(word string) string {
	sortedBytes := []byte(word)
	sort.Slice(sortedBytes, func(i, j int) bool {
		return sortedBytes[i] < sortedBytes[j]
	})

	return string(sortedBytes)
}
