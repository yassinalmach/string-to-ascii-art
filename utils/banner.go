package utils

func banner(fileName string) string {
	var file string
	switch fileName {
	case "standard":
		file = "Banners/standard.txt"
	case "shadow":
		file = "Banners/shadow.txt"
	case "thinkertoy":
		file = "Banners/thinkertoy.txt"
	default:
		PrintError("banner")
	}
	return file
}
