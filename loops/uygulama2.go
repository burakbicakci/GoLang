package loops

import (
	"fmt"
)

func Demo3() {
	//tahmin sayısı ekleme

	sayimiz := 20

	tahminEdilen := 0

	tahminSayisi := 0

	fmt.Println("Tahmin edilen sayiyi giriniz: ")
	fmt.Scanln(&tahminEdilen)

	for sayimiz != tahminEdilen {

		if tahminEdilen < sayimiz {
			fmt.Println("Daha büyük bir sayi giriniz.")
			fmt.Scanln(&tahminEdilen)
			tahminSayisi++
		} else if tahminEdilen > sayimiz {
			fmt.Println("Daha küçük bir sayi giriniz.")
			fmt.Scanln(&tahminEdilen)
			tahminSayisi++
		}
	}
	fmt.Println("Tebrikler sayiyi buldunuz.")
	fmt.Println("Tahmin sayisi: ", tahminSayisi)
}
