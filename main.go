package main
import (
  "fmt"
)


func fizzbuzz(num int) string {

	if num % 15 == 0  {
		return "FizzBuzz"
	} else if num % 5 == 0 {
		return "Buzz"
	} else if num % 3 == 0 {
		return "Fizz"
	} else {
		return fmt.Sprintf("%d", num)
	}
}


func main() {
	for i := 1; i < 100; i++ {
		fmt.Print(fizzbuzz(i), " ")
	}
}
