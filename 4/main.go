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