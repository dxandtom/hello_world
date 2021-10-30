package main

import "fmt"

func add(inputNum1, inputNum2 int) int {
	return inputNum1 + inputNum2
}


func minus(inputNum1, inputNum2 int) int {
	return inputNum1 - inputNum2
}
func mcl(inputNum1, inputNum2 int) int {
	return inputNum1 * inputNum2
}

func main() {
	var inputNum1,inputNum2 int
	var ans int
	var cmd string
    var history []int
	history = make([]int, 0)

	for true {
		fmt.Scanln(&inputNum1, &cmd, &inputNum2)
		switch cmd {
		case "+":
			ans = add(inputNum1, inputNum2)
		case "-":
			ans = minus(inputNum1, inputNum2)
		case "*":
			ans = mcl(inputNum1, inputNum2)
		case "_":
			ans = minus(inputNum1, inputNum2)
		}
		history = append(history, ans)
		fmt.Println(ans)
        for _, historyNum := range history{
			fmt.Print(historyNum, " ")
		}
		fmt.Println( )
	}
}

