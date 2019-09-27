package rpg

import (
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

var (
	dieMatcher = regexp.MustCompile(`(\d+)d(\d+)`)
	opa        = map[string]struct {
		// Operator Precedence and supported operators
		prec   int
		rAssoc bool
	}{
		"^": {4, true},
		"*": {3, false},
		"/": {3, false},
		"+": {2, false},
		"-": {2, false},
	}
)

// Roll takes a string encompassing some manner of addition and subtraction of dice and integers and flattens it into a single value if possible
func Roll(diceStr string) int {
	expression := dieMatcher.ReplaceAllStringFunc(diceStr, func(match string) string {
		return strconv.Itoa(rollDice(match))
	})
	infix := expToInfix(expression)
	postfix := parseInfix(infix)
	return eval(postfix)
}

// Converts an expression of integers and non-integers into infix
func expToInfix(expression string) (r string) {
	for _, t := range expression {
		token := string(t)
		if _, err := strconv.Atoi(token); err == nil {
			r += token
		} else if token != " " {
			r += " " + token + " "
		}
	}
	return
}

// Shunting-yard algorith implementation
func parseInfix(infix string) (r string) {
	var stack []string
	for _, t := range strings.Fields(infix) {
		token := string(t)
		switch token {
		case "(":
			stack = append(stack, token)
		case ")":
			var operator string
			for {
				// pop ("(" or operator)
				operator, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if operator == "(" {
					break // discard "("
				}
				r += " " + operator
			}
		default:
			if o, isOp := opa[token]; isOp {
				// token is operator
				for len(stack) > 0 {
					operator := stack[len(stack)-1]
					if p, isOp := opa[operator]; !isOp || o.prec > p.prec ||
						o.prec == p.prec && o.rAssoc {
						break
					}
					stack = stack[:len(stack)-1]
					r += " " + operator
				}
				stack = append(stack, token)
			} else {
				// token is operand
				if r > "" {
					r += " "
				}
				r += token
			}
		}
	}
	for len(stack) > 0 {
		r += " " + stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return
}

// Calculate the final result of a string in reverse Polish notation (RPN).
func eval(rpn string) int {
	var stack []int
	for _, token := range strings.Fields(rpn) {
		switch token {
		case "+":
			if len(stack) == 0 {
				return 0
			}
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			if len(stack) == 0 {
				return 0
			}
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			if len(stack) == 0 {
				return 0
			}
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			if len(stack) == 0 {
				return 0
			}
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "^":
			if len(stack) == 0 {
				return 0
			}
			stack[len(stack)-2] = int(math.Pow(float64(stack[len(stack)-2]), float64(stack[len(stack)-1])))
			stack = stack[:len(stack)-1]
		default:
			i, _ := strconv.Atoi(token)
			stack = append(stack, i)
		}
	}
	if len(stack) > 0 {
		return stack[0]
	}
	return 0
}

func rollDice(xdy string) int {
	var matches []string
	matches = dieMatcher.FindStringSubmatch(xdy)
	var total int
	d, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0
	}
	sides, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0
	}
	for i := 0; i < d; i++ {
		total += rand.Intn(sides) + 1
	}
	return total
}
