package goruitines

import (
	"fmt"
	"time"
)

func Demo2() {
	//kanallar şu işe yarar : bir iş parçacığından diğer iş parçacığına veri göndermek için kullanılır.
	k := make(chan bool, 2) //2 bool oluşturan kanal

	go func() {

		time.Sleep(time.Second * 5)

		k <- true //kanala değer döndürme işlemi

		time.Sleep(time.Second * 2)

		k <- false
	}()

	//ana iş parçacığı k kanalına 2 değer gelene kadar bekleyecek
	fmt.Println(<-k, <-k) //çıktı: true false  2 sn true içn bekliyor sonra 2 sn false için bekliyor. sonra ekrana yazdıyrpzu.
	time.Sleep(time.Second * 2)
	fmt.Println("main bitti")
}

// k := make(chan bool) ->>>> kanal oluşturma
// k <- true  ->>>>örnke atama
// a := <-k ->>>> kanaldan değer okuma
// <- k ->>>> sadece kanala değer gelmesini beklemesini
