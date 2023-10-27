package errorhandling

import (
	"fmt"
	"os"
)

//type assertion

func Demo1() {

	f, err := os.Open("demo.txt") //burda dosya açıyoruz os kütüphanesinden
	//dosya bulunursa nill bulunmazsa hata döner.
	if err != nil { // eğer hata varsa
		if pErr, ok := err.(*os.PathError); ok { //ok,if true ise demek
			fmt.Println("Dosya Bulunamadı", pErr.Path) //pErr hatayı buna ata demek cevap oalrak Dosya Bulunamadı demo.txt verir
			return
		} else {
			fmt.Println("Dosya Yok")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
