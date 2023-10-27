package errorhandling

import "fmt"

func dogrula(i interface{}) { // böyle kullanmamaızın sebebi her türlü veri tipini alabiliyor olması
	sayi := i.(int)
	fmt.Println(sayi)
}

func Demo2() {
	// 	// var sayi interface{} = "selam" //  string kullanamazsın diye hata verir
	// 	// dogrula(sayi)

	var sayi2 interface{} = 5
	dogrula(sayi2)

}
