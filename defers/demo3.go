package errors

import "fmt"

type product struct { //struct tanımlama diğer diller ile aynı
	productName string
	price       int
}

func (p product) Save() { //sadece struck function tanımlama böyle farklı olarak
	fmt.Println("Ürün kaydedildi", p.productName)
	defer Log()
}

func Demo3() {
	p := product{productName: "Laptop", price: 5000}
	defer p.Save()                                   //önce işlem başarılı çalıştı ardından p.Save çalıştırdı ve ürün kaydedildi yazdı.
	p = product{productName: "Telefon", price: 2000} // telefon görmez defer gördüğü andan öncekini kaydeder ve işlemi yapar.
	fmt.Println("İşlem Başarılı")
}

func Log() {
	fmt.Println("Loglandı")
}
