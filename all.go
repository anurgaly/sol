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
// #reduceInt #15

// func ReduceInt(a []int, f func(int, int) int) {
// 	result := a[0]
// 	for i := 1; i <= len(a)-1; i++ {
// 		result = f(result, a[i])
// 	}
// 	r := strconv.Itoa(result)
// 	for _, i := range r {
// 		z01.PrintRune(i)
// 	}
// 	z01.PrintRune('\n')
// }
// func main() {
// 	mul := func(acc int, cur int) int {
// 		return acc * cur
// 	}
// 	sum := func(acc int, cur int) int {
// 		return acc + cur
// 	}
// 	div := func(acc int, cur int) int {
// 		return acc / cur
// 	}
// 	as := []int{500, 2}
// 	ReduceInt(as, mul)
// 	ReduceInt(as, sum)
// 	ReduceInt(as, div)
// }

// #lastword #14

// func main() {
// 	args := os.Args[1:]
// 	arg := args[0]
// 	result := ""
// 	if len(args) == 1 {
// 		for i := len(arg) - 1; i >= 0; {
// 			if arg[i] != ' ' {
// 				result = string(arg[i]) + result
// 				i--
// 			} else {
// 				if result == "" {
// 					i--
// 				} else {
// 					break
// 				}
// 			}
// 		}
// 		if result != "" {
// 			for _, ch := range result {
// 				z01.PrintRune(ch)
// 			}
// 			z01.PrintRune('\n')
// 		} else {
// 			return
// 		}
// 	}
// }

// #rot13 #13

// func main() {
// 	args := os.Args[:]
// 	if len(args) == 1 {
// 		return
// 	}
// 	arg := args[1]
// 	if len(arg) != 0 {
// 		for _, ch := range arg {
// 			if ch >= 'a' && ch <= 'n' || ch >= 'A' && ch <= 'N' {
// 				z01.PrintRune(ch + 13)
// 			} else if ch >= 'm' && ch <= 'z' || ch >= 'M' && ch <= 'Z' {
// 				z01.PrintRune(ch - 13)
// 			} else {
// 				z01.PrintRune(ch)
// 			}
// 		}
// 		z01.PrintRune('\n')
// 	} else {
// 		return
// 	}
// }

// #lastRune #12

// func lastRune(s string) rune {
// 	return rune(s[(len(s) - 1)])
// }
// func main() {
// 	z01.PrintRune(lastRune("Hello!"))
// 	z01.PrintRune(lastRune("Salut!"))
// 	z01.PrintRune(lastRune("Ola!"))
// 	z01.PrintRune('\n')
// }

// #firstRune #12

// func firstRune(s string) rune {
// 	return rune(s[0])
// }
// func main() {
// 	z01.PrintRune(firstRune("Hello!"))
// 	z01.PrintRune(firstRune("Salut!"))
// 	z01.PrintRune(firstRune("Ola!"))
// 	z01.PrintRune('\n')
// }
// #compare #19

// func compare(a, b string) int {
// 	if a == b {
// 		return 0
// 	} else if a > b {
// 		return 1
// 	}
// 	return -1
// }
// func main() {
// 	fmt.Println(compare("Hello!", "Hello!"))
// 	fmt.Println(compare("Salut!", "lut!"))
// 	fmt.Println(compare("Ola!", "Ol"))
// }

// alphamirror #18

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	arg := args[0]
// 	for _, ch := range arg {
// 		if ch >= 'A' && ch <= 'Z' {
// 			z01.PrintRune('Z' - ch + 'A')
// 		} else if ch >= 'a' && ch <= 'z' {
// 			z01.PrintRune('z' - ch + 'a')
// 		} else {
// 			z01.PrintRune(ch)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// #tabmult #17

// func printSTR(s string) {
// 	for _, ch := range s {
// 		z01.PrintRune(ch)
// 	}
// }
// func printInt(n int) {
// 	if n != 0 {
// 		i := n / 10
// 		if i != 0 {
// 			printInt(i)
// 		}
// 		k := (n % 10) + '0'
// 		z01.PrintRune(rune(k))
// 	}
// }
// func mul(i int, n int) int {
// 	return i * n
// }
// func strToint(s string) int {
// 	i, _ := strconv.Atoi(s)
// 	return i
// }
// func main() {
// 	arg := os.Args[1:]
// 	if len(arg) != 1 {
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	n := arg[0]
// 	for i := 1; i <= 9; i++ {
// 		printInt(i)
// 		printSTR(" x ")
// 		printSTR(n)
// 		printSTR(" = ")
// 		printInt(mul(i, strToint(n)))
// 		z01.PrintRune('\n')
// 	}
// }

// #searchReplace #17

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 3 {
// 		return
// 	}
// 	str := args[0]
// 	first := args[1]
// 	second := []rune(args[2])
// 	for _, ch := range str {
// 		if string(ch) == first {
// 			z01.PrintRune(second[0])
// 		} else {
// 			z01.PrintRune(ch)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// #chunk #16

// func Chunk(slice []int, size int) {
// 	var str []int
// 	result := [][]int{}
// 	if len(slice) == 0 {
// 		fmt.Println(result)
// 		return
// 	}
// 	if size == 0 {
// 		fmt.Println('\n')
// 		return
// 	}
// 	for len(slice) > size {
// 		str = slice[:size]
// 		result = append(result, str)
// 		slice = slice[size:]
// 	}
// 	result = append(result, slice)
// 	fmt.Println(result)
// }
// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 10)
// }


// func FindPrevPrime(nb int) int {
// 	count := 0
// 	var res int
// 	for x := nb; x > 1; x-- {
// 	 for i := 1; i <= x; i++ {
// 	  if x%i == 0 {
// 	   count++
// 	   res = i
// 	  }
// 	 }
// 	 if count == 2 {
// 	  return res
// 	 }
// 	 count = 0
// 	}
// 	return 0
//    }
