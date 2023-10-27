package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	userId    int    `json:"userId"`
	id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") // burada get ile istek attık

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close() // defer ile response body kapatıldı

	bodyBtytes, _ := ioutil.ReadAll(response.Body) // response body okundu

	bodyString := string(bodyBtytes) // body stringe çevrildi
	fmt.Println(bodyString)          // body string yazdırıldı json verisi geldi

	var todo Todo                             // todo nesnesi oluşturuldu
	json.Unmarshal([]byte(bodyString), &todo) // json verisi todo nesnesine çevrildi
	fmt.Println(todo)

}

func Demo2() {

	todo := Todo{1, 2, "Alışveriş", false} // todo nesnesi oluşturduk
	jsonTodo, err := json.Marshal(todo)    // todo nesnesini json verisine çevirdik marshal ile

	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", //burada post isteği attık
		"application/json;charset=utf-8", // burada content type belirttik bunun sebebi json verisi gönderdiğimizi belirtmek
		bytes.NewBuffer(jsonTodo))        //burada ise json verisini byte a çevirdik

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close() // defer ile response body kapatıldı

	bodyBtytes, _ := ioutil.ReadAll(response.Body) // response body okundu

	bodyString := string(bodyBtytes) // body stringe çevrildi
	fmt.Println(bodyString)          // body string yazdırıldı json verisi geldi

	var todoResponse Todo
	json.Unmarshal([]byte(bodyString), &todoResponse) // json verisi todo nesnesine çevrildi unmasrshalın içine yazdık
	fmt.Println(todoResponse)
}

//marshal = json verisine çevirme
// unmarshal = json verisinden nesneye çevirme
// yani ilk olarak bir nesne oluşturuyoruz ve bu nesneyi json verisine çeviriyoruz eğer geri çevirme yapmak istersek unmarshal ile json verisinden nesneye çeviriyoruz
