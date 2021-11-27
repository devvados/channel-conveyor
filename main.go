package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var input string
	var c = make(chan int, 2)
	var wg sync.WaitGroup

	fmt.Print("Пожалуйста, введите число: ")
	_, _ = fmt.Scan(&input)

	switch input {
	case "стоп":
		{
			fmt.Print("Конец работы программы. Выход...\n")
			return
		}
	default:
		{
			number, err := strconv.Atoi(input)
			if err != nil {
				fmt.Print("Ошибка обработки введенных данных. Выход...\n")
				return
			} else {
				wg.Add(2)
				go func() {
					defer wg.Done()
					square(number, c)
				}()
				go func() {
					defer wg.Done()
					defer close(c)
					multiplyByTwo(c)
				}()
			}
		}
	}

	wg.Wait()
	fmt.Println("----------")
	for val := range c {
		fmt.Println(val)
	}
}

func square(number int, channel chan int) {
	res := number * number
	channel <- res
}

func multiplyByTwo(channel chan int) {
	num1 := <-channel
	num2 := num1 * 2
	channel <- num1
	channel <- num2
}
