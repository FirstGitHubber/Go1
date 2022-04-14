package main

import "fmt"

func main() {
  var n, d1, d2 int
n=2466
  fmt.Print("Введите трехзначное число: ")
  fmt.Scanln(&n)
if n>999 {fmt.Print("Ошибка")
}else{
	d1 = n % 10
	n = n / 10
	d2 = n % 10
	n = n / 10

  fmt.Print("Первое число: ", n)
  fmt.Print("Третее число: ", d2)
  fmt.Print("Второе число: ", d1)
	}
}