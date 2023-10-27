package errors

import (
	"fmt"
)

func C() {
	fmt.Println("C fonks çalıştı")
}

func A() {
	fmt.Println("A fonks çalıştı")
}

func B() {

	defer A() // defer ile fonksiyonun sonunda çalıştırılmasını sağlıyoruz.
	defer C()
	fmt.Println("B fonks çalıştı") //ilk b çalışır
}
