package arrays

import "fmt"

func Demo1() {

	// var sayilar [5]int   //burda 5 elemanlı bir dizi oluşturduk.
	// fmt.Println(sayilar) // değer vermesek bile hepsine 0 atar
	// sayilar[2] = 5       //burda 2. indexe 5 değerini atadık.
	// fmt.Println(sayilar)

	// dizileri dolaşm

	var sehirler = [5]string{"Ankara", "İstanbul", "İzmir", "Adana", "Bursa"}

	fmt.Println(sehirler) // bu şekilde de yazzabiliriz

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	var sayilar2 = [5]int{4, 2, 5, 8, 9}
	enBuyuk := sayilar2[0]

	for i := 0; i < len(sayilar2); i++ {
		if sayilar2[i] > enBuyuk {
			enBuyuk = sayilar2[i]
		}
	}
	fmt.Println("En büyük sayı:", enBuyuk)
}
