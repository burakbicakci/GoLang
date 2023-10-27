package degiskenlerDosya

import "fmt"

func Degiskenler() {

	//var degiskenAdi veriTipi = deger
	//string, int, float, bool, const

	var metin string = "Merhaba Go" // string degisken tanimlama
	fmt.Println("degiskenler")
	fmt.Println(metin)
	fmt.Print("\n") // alt satira gecmeama print yaparsam alt satıra geçmiyor println alt geçme

	var kdv int = 10 // int degisken tanimlama
	fmt.Println(kdv)
	fmt.Println(100 + kdv) // degisken ile toplama islemi

	var kdv2 float32 = 0.18 // float degisken tanimlama
	fmt.Println(kdv2)
	fmt.Printf("veri tipi : %T", kdv2) // veri tipini yazdirma
	fmt.Print("\n")

	var durum bool // bool degisken tanimlama

	var metin2 string = "Merhaba"
	var metin3 string = "Go"
	durum = metin2 == metin3 // string karsilastirma ve ! koymak tersini alir

	const pi float32 = 3.14 // const sabit tanimlama

	fmt.Println(durum)

}
