package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)
 
func main(){

	top := 100
	rand.Seed(time.Now().UnixNano())
	ans := rand.Intn(top)
	//fmt.Println("ans is ",ans)

	reader := bufio.NewReader(os.Stdin)

	for{
		input,err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("error and try again")
			continue
		}
		input = strings.TrimSuffix(input,"\r\n")

		guess,err := strconv.Atoi(input)
		
		if err != nil{
			fmt.Println("no num and try again ",err)
			continue
		}

		if guess > ans {
			fmt.Println(guess," is bigger ")
		} else if guess < ans {
			fmt.Println(guess," is smaller")
		}else {
			break
		}
	}

	fmt.Println("you wim !")
}