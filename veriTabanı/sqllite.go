package veritabanı

import (
	"database/sql" //database ekleme
	"fmt"

	_ "github.com/mattn/go-sqlite3" //sqlite3 ekleme
)

func main() {
	vt, _ := sql.Open("sqlite3", "./veritabanı.db") //veri tabanı dosyamız
	vt.Close()                                      //İşimiz bittikten sonra veri tabanımızı kapatıyoruz
}

// veri ekleme için fonksiyon
func main() {
	vt, _ := sql.Open("sqlite3", "./veritabanı.db")             // veri tabanı dosyamız _ ile hata kontrolü yapmadık
	işlem, _ := vt.Prepare("INSERT INTO kisiler(ad) values(?)") // sorgu hazırlıyoruz insert ekleme için
	veri, _ := işlem.Exec("Mustafa Kemal ATATÜRK")              //exec ile sorguyu çalıştırıyoruz
	id, _ := veri.LastInsertId()                                //Son girişin id numarısını aldık
	vt.Close()                                                  //İşimiz bittikten sonra veri tabanımızı kapatıyoruz
}

// veri güncelleme için fonksiyon
func main() {
	vt, _ := sql.Open("sqlite3", "./veritabanı.db")              // veri tabanı dosyamız _ ile hata kontrolü yapmadık
	id := 1                                                      // Güncellenecek kişinin id'si
	işlem, _ := vt.Prepare("UPDATE kisiler set ad=? where id=?") // prepare ile sorguyu hazırlıyoruz update güncelleme için
	veri, _ := işlem.Exec("Gazi M. K. ATATÜRK", id)              // exec ile sorguyu çalıştırıyoruz
	değişiklik, _ := veri.RowsAffected()                         // Değişen kişinin id'sini aldık
	vt.Close()                                                   //İşimiz bittikten sonra veri tabanımızı kapatıyoruz
}

// veri silme için fonksiyon
func main() {
	vt, _ := sql.Open("sqlite3", "./veritabanı.db")          // veri tabanı dosyamız _ ile hata kontrolü yapmadık
	işlem, _ := vt.Prepare("delete from kisiler where id=?") // prepare ile sorguyu hazırlıyoruz delete sile için
	veri, _ := işlem.Exec(id)                                //Silinecek kişinin id'si
	değişiklik, _ := veri.RowsAffected()                     //Silinen kişinin id'sini aldık
	fmt.Println("Silinen kişinin id'si: ", değişiklik)
	vt.Close()
}

func main() {
	vt, _ := sql.Open("sqlite3", "./veritabanı.db")
	tablo, _ := vt.Query("SELECT * FROM kisiler") // tablo sorgulama
	tablo.Close()                                 //İşimiz bittiği için tablo sorgulamayı kapattık
	vt.Close()                                    //İşimiz bittikten sonra veri tabanımızı kapatıyoruz
}
