// #strlen #9

// func main() {
// 	l := strLen("Hello World!")
// 	fmt.Println(l)
// }
// func strLen(s string) int {
// 	return len(s)
// }

// #countDown #8

// func main() {
// 	for _, ch := range "9876543210" {
// 		z01.PrintRune(ch)
// 	}
// 	z01.PrintRune('\n')
// }


// #wdmatch #11

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 2 {
// 		return
// 	}
// 	first := args[0]
// 	second := args[1]
// 	result := ""
// 	for _, f := range first {
// 		for i, s := range second {
// 			if f == s {
// 				result += string(s)
// 				second = second[i:]
// 				break
// 			}
// 		}
// 	}
// 	for _, ch := range result {
// 		if result == first {
// 			z01.PrintRune(ch)
// 		} else {
// 			return
// 		}
// 	}
// 	z01.PrintRune('\n')
// }