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