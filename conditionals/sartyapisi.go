package conditionals

import "fmt"

func SartYapisi() {

	//yapı olarak diğer dillerdeki gibi

	var hesap float32 = 1000
	var cekilecek float32 = 900

	if cekilecek > hesap {
		fmt.Println("bakiye yetersiz")
	} else {
		fmt.Println("paranizi alabilirsiniz")
		hesap = hesap - cekilecek
	}

	fmt.Println(hesap)

	// iç içe şartlar da yapabiliriz
	var hesap2 float32 = 1000
	var cekilecek2 float32 = 900
	var limit float32 = 500

	if cekilecek2 > hesap2 {
		fmt.Println("bakiye yetersiz")
	} else if cekilecek2 > limit {
		fmt.Println("limit yetersiz")
	} else {
		fmt.Println("paranizi alabilirsiniz")
		hesap2 = hesap2 - cekilecek2
	}
	fmt.Println(hesap2)

}
