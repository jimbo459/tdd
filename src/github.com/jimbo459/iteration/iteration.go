package iteration

func Repeat(chars string) string {
	var repeatedString string
	repeatCount := 5

	for x := 0; x < repeatCount; x++ {
		repeatedString += chars
	}
	return repeatedString
}
