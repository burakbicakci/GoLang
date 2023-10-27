package errorhandling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz")
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayı giriniz", nil // bir string bir error döndürüyoruz nil demek hata yok demek
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayı giriniz", nil
	}

	return "Bildiniz", nil // bir string le bildniz yolluyoruz error olmadığında nil yani null geliyor
}

func Demo3() {
	mesaj, _ := TahminEt(20) //_ demek hata yok demek
	fmt.Println(mesaj)

}
