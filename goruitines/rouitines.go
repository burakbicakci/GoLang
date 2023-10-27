package goruitines

import (
	"fmt"
	"time"
)

func CiftSayilar() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("çift sayı:", i)
			time.Sleep(time.Second * 1) // bu sayede bekleyerek yazma işlemi yapıyor.
		}
	}

}

func TekSayilar() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("tek sayı:", i)
		time.Sleep(time.Second * 1)
	}
}

func DemoMain() {

	fmt.Println("Merhaba Go")
	go DemoYazdir("Eşzamanlı yazdırma") //burdaki amacımız eşzamanlı çalıştırmak
	time.Sleep(time.Second * 2)         // ekranda önce bu yazılsın diye bekletiyoruz.

}

func DemoYazdir(mesaj string) {
	fmt.Println(mesaj)
}

// go programlamada diğer dillerin aksına main de go çağırmadna önce başına go  yazarak routineleri çalıştırabiliriz.
