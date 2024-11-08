package utils

func banner(fileName string) string {
	var file string
	switch fileName {
	case "standard":
		file = "banners/standard.txt"
	case "shadow":
		file = "banners/shadow.txt"
	case "thinkertoy":
		file = "banners/thinkertoy.txt"
	default:
		PrintError("banner")
	}
	return file
}
