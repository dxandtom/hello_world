package main

import "fmt"

//func f2(a, b int, c string) (int, string) {
     //sum := a + b
	//return sum, c
//}
func f2(a, b int, c string) (sum int, cc string) {
	sum = a + b
	cc = c
	return
}



func main() {
	//var inputName1, inputName2,inputName3,inputName4,inputName5 string
	//fmt.Scanln(&inputName1,&inputName2,&inputName3,&inputName4,&inputName5)
	//fmt.Println(inputName5,inputName4,inputName3,inputName2,inputName1)
	fmt.Println(f2(5, 6, "dx"))
}





