package loops

func Demo1() {
	i := 1

	for i <= 10 { //for yapısı kullanımı ya da i:=1; i<=10; i++ şeklinde de kullanılabilir.
		println(i)
		i++
	}
}

func DemoSwitch() {
	switch 1 {
	case 1:
		println("1")
	case 2:
		println("2")
	default:
		println("default")
	}
}
