import (
	"strconv"
)

func evalRPN(tokens []string) int {
    var stack []int

    for _, t := range tokens {
        switch t {
            case "+", "-", "*", "/":
                n1, n2 := stack[len(stack) - 2], stack[len(stack) - 1]
                stack = stack[:len(stack) - 2]

                var num int
                if t == "+" {
                    num = n1 + n2
                }
                if t == "-" {
                    num = n1 - n2
                }
                if t == "*" {
                    num = n1 * n2
                }
                if t == "/" {
                    num = n1 / n2
                }

                stack = append(stack, num)
            default:
                num, _ := strconv.Atoi(t)
                stack = append(stack, num)
        }
    }

    return stack[0]
}
