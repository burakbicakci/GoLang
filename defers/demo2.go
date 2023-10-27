package errors

import "fmt"

func Demo1(sayi int32) string {
	defer DeferFucn()
	if sayi%2 == 0 {
		return "Sayı çifttir"
	}
	if sayi%2 != 0 {
		return "Sayı tektir"
	}
	return "Belli DEğil"
}

func DeferFucn() {
	fmt.Println("Bitti")

}
