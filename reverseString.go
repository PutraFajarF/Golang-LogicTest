package main

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// func main() {
// 	input := "I am Superman"
// 	reverse := Reverse(input)
// 	reverseAgain := Reverse(reverse)

// 	fmt.Println("Origin: %q\n", input)
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Reverse: %q\n", reverse)
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("ReverseAgain: %q\n", reverseAgain)
// }
