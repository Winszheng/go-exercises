package iteration

func Repeat(character string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated = repeated + character
	}
	//time.Sleep(10*time.Second)
	return repeated
}

