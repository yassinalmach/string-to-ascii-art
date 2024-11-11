package reverse

import "strings"

// pickBanner checks which banner i will work with
func pickBanner(file string) string {
	if strings.ContainsAny(string(file), "o0") {
		return "banners/thinkertoy.txt"
	} else if containsOnly(file) {
		return "banners/shadow.txt"
	}

	return "banners/standard.txt"
}

// containsOnly it checks whether the file have only specific characters
func containsOnly(file string) bool {
	for i := 0; i < len(file); i++ {
		if file[i] != ' ' && file[i] != '\n' {
			if file[i:i+2] != "_|" && file[i-1:i+1] != "_|" {
				return false
			}
		}
	}
	return true
}
