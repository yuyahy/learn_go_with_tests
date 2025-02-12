package iteration

func Repeat(character string, repeatCnt int) string {
	var repeated string
	for i := 0; i < repeatCnt; i++ {
		repeated += character
	}
	return repeated
}
