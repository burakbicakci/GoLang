package loops

import (
	"fmt"
)

func Demo4() {
	//kullanıcıdan sayı alıcak ve asal ya da asal olmama bulma
	sayi := 0
	fmt.Println("Bir sayi giriniz: ")
	fmt.Scanln(&sayi)

	deger := true

	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			deger = false
			break
		}
	}

	if deger == true {
		fmt.Println("Asal sayidir.")
	} else {
		fmt.Println("Asal sayi degildir.")
	}

}
