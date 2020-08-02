package Modules

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {

}

func (calc) operate(input string, operator string) int {
	cleanInput := strings.Split(input, operator)
	operator1 := parser(cleanInput[0])
	operator2 := parser(cleanInput[1])
	switch operator {
	case "+":
		fmt.Println(operator1+operator2)
		return operator1 + operator2
	case "-":
		fmt.Println(operator1-operator2)
		return operator1-operator2
	case "*":
		fmt.Println(operator1*operator2)
		return operator1*operator2
	case "/":
		fmt.Println(operator1/operator2)
		return operator1/operator2
	default:
		fmt.Println("The operator is invalid")
		return 0
	}
}

func parser (input string) int {
	operator, _ := strconv.Atoi(input)
	return operator
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	return operation
}
