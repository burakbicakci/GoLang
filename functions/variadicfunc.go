package functions

func Demo2(sayilar ...int) int { // kaç tane sayı girilirse girilsin hepsini toplayacak fonksiyon yani belirli sayıda istemiyor.
	toplam := 0

	for i := 0; i < len(sayilar); i++ {
		toplam += sayilar[i]
	}
	return toplam
}
