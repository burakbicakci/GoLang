package loops

import (
	"fmt"
)

func Demo5() {

	//arkadaş sayı bulma 220 ve 284 arkadaş sayı : 220'nin bölenlerinin toplamı 284, 284'ün bölenlerinin toplamı 220'yi verir.
	sayi1 := 220
	sayi2 := 284

	toplam1 := 0
	toplam2 := 0

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			toplam1 += i
		}
	}

	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			toplam2 += i
		}
	}

	if toplam1 == sayi2 && toplam2 == sayi1 {
		fmt.Printf("%v ve %v arkadaş sayılardır.", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v arkadaş sayılar değildir.", sayi1, sayi2)
	}
}
