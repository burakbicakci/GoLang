package goruitines

import (
	"fmt"
	"sync" // mutex'i kullanmak için
)

var mt sync.Mutex // bir global mutex nesnesi oluşturduk

func paraÇek(bakiye *float64, çekilecekMiktar float64, wg *sync.WaitGroup) {
	/*
	* mt isimli mutex'i bu işlem yapılırken kilitliyoruz.
	* bu sayede mt mutex'ini başka işlemler kullanamıyor.
	 */
	mt.Lock()

	//devamı asnekron işlemlerden sonra
	*bakiye -= 15
	fmt.Printf("Yeni Bakiye: %.2f\n", *bakiye)

	mt.Unlock() //diğer işlemlerin mutex'i kullanabilmesi için mutex'i açıyoruz.
	fmt.Println("Çekme işlemi tamamlandı.")

	wg.Done() // işlem tamamlandığında wg.Add(2) ile 2 arttırdığımız değeri 1 azaltıyoruz.
}

// bu fonksiyonda yukarıdaki ile aynı mantıkta
func paraYatır(bakiye *float64, yatırılacakMiktar float64, wg *sync.WaitGroup) {
	mt.Lock()
	*bakiye += 65
	fmt.Printf("Yeni Bakiye: %.2f\n", *bakiye)
	mt.Unlock()
	fmt.Println("Yatırma işlemi tamamlandı.")
	wg.Done()
}

func DemoMutex() {

	//bu fonksiyonunu amacı bir bekleyenler grubu oluşturmak

	var wg sync.WaitGroup

	//2 fonksiyonu da bekleyeceğimiz için Add'e 2 yazalım
	wg.Add(2)

	//fonksiyonlarımızın kullancağı bakiye değişkenimiz
	var bakiye float64 = 100
	fmt.Printf("İlk Bakiye: %.2f\n", bakiye)

	// iki fonksiyonu da aynı anda çalıştırıyoruz.
	go paraÇek(&bakiye, 25, &wg)
	go paraYatır(&bakiye, 65, &wg)

	/*
	* ana iş parçacığı tamamlandığında asenkron çalışan fonksiyonları beklemez.
	* beklemediğinde de asenkron fonksiyonlar çalışmadan program sonlanır.
	* ana iş parçacığının asenkron işlemleri beklemesi için waitgroup sonucunun 0 olmasını bekleriz.
	* wg.Add(2) yazarak 2 adet wg.Done() fonksiyonu çalıştığında wg.Add(0) olur ve
	* wg.Wait() tamamlanır ve program başka işlemler yapılmıyor ise sonlanır.
	 */
	wg.Wait()
}
