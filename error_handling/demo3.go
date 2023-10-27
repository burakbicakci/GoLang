package errorhandling

import (
	"fmt"
)

type borderException struct {
	message   string
	paremeter int
}

func (b *borderException) Error() string { //pointer koyduk ki değer değişsin
	return fmt.Sprintf("%d  --------%s", b.paremeter, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{"1-100 arasında bir sayı giriniz", tahmin}
	}
	return "Bildiniz", nil
}
