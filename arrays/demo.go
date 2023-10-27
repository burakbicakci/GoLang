package arrays

import "fmt"

func Demo2() {

	// çok boyutlu dizeler

	var sayilar [3][2]int = [3][2]int{{1, 2}, {3, 4}, {5, 6}}

	fmt.Println(sayilar)

	// for ile gezme.

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(sayilar[i][j])
		}
	}

}
