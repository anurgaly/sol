
// #printDigits #7

// func main() {
// 	for i := '0'; i <= '9'; i++ {
// 		z01.PrintRune(i)
// 	}
// 	z01.PrintRune('\n')
// }

// displayalpham #6

// func main() {
// 	for _, ch := range "aBcDeFgHiJkLmNoPqRsTuVwXyZ" {
// 		z01.PrintRune(ch)
// 	}
// 	z01.PrintRune('\n')
// }

// #displayalrevm #5

// func main() {
// 	for _, ch := range "zYxWvUtSrQpOnMlKjIhGfEdCbA" {
// 		z01.PrintRune(ch)
// 	}
// 	z01.PrintRune('\n')
// }

// #paramCount #4

// func main() {
// 	args := os.Args[1:]
// 	lenth := len(args)
// 	if len(args) != 0 {
// 		printNbr(lenth)
// 		z01.PrintRune('\n')
// 	} else {
// 		return
// 	}
// }
// func printNbr(n int) {
// 	t := 1
// 	if n < 0 {
// 		t = -1
// 		z01.PrintRune('-')
// 	}
// 	if n != 0 {
// 		f := (n / 10) * t
// 		if f != 0 {
// 			printNbr(f)
// 		}
// 		k := (n % 10 * t) + '0'
// 		z01.PrintRune(rune(k))
// 		// z01.PrintRune('\n')
// 	} else {
// 		z01.PrintRune('0')
// 	}
// }

// #hello #3

// func main() {
// 	for _, ch := range "Hello World!" {
// 		z01.PrintRune(ch)
// 	}
// 	z01.PrintRune('\n')
// }

// #DisplayLastParam #2

// func main() {
// 	arg := os.Args[:]
// 	if len(arg) == 1 {
// 		z01.PrintRune('\n')
// 		return
// 	} else {
// 		last := arg[len(arg)-1]
// 		for _, ch := range last {
// 			z01.PrintRune(ch)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// #DisplayFirstParam #1

// func main() {
// args := os.Args[:]
// if len(args) == 1 {
// z01.PrintRune('\n')
// return
// }
// first := args[1]
// if first != "" {
// for _, ch := range first {
// z01.PrintRune(ch)
// }
// z01.PrintRune('\n')
// } else {
// z01.PrintRune('\n')
// }
//}