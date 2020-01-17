package iteration

func Repeat(chars string) string {
	var repeatedString string
	for x := 0; x < 5; x++ {
		repeatedString += chars
	}
	return repeatedString
}
