package loops

import (
	"fmt"
)

func Demo2() {
	sayimiz := 20

	tahminEdilen := 0

	fmt.Println("Tahmin edilen sayiyi giriniz: ")
	fmt.Scanln(&tahminEdilen)

	for sayimiz != tahminEdilen {
		if tahminEdilen == sayimiz {
			fmt.Println("Tebrikler sayiyi buldunuz.")
		} else if tahminEdilen < sayimiz {
			fmt.Println("Daha büyük bir sayi giriniz.")
			fmt.Scanln(&tahminEdilen)
		} else if tahminEdilen > sayimiz {
			fmt.Println("Daha küçük bir sayi giriniz.")
			fmt.Scanln(&tahminEdilen)
		}
	}
}
