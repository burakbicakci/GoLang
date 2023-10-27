package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	vt, err := sql.Open("mysql", "şifre", "kullanıcı adı")

	{ // Yeni tablo oluşturma işlemleri aynı mantık
		sorgu := `
            CREATE TABLE kullanicilar (
                id INT AUTO_INCREMENT,
                kullanici TEXT NOT NULL,
                sifre TEXT NOT NULL,
                tarih DATETIME,
                PRIMARY KEY (id)
            );`
		_, err := db.Exec(sorgu)

	}
	{ // Yeni kayıt ekle
		kullanici := "deneme"
		sifre := "1234"
		tarih := time.Now()                                                                                                         // şimdiki zamanı al
		sonuc, err := db.Exec(`INSERT INTO kullanicilar (kullanici, sifre, tarih) VALUES (?, ?, ?)`, username, password, createdAt) //burdaki exec amacı sorgu çalıştırmak

		eklenenid, err := sonuc.LastInsertId() // son eklenen id yi almak için fonksiyon
		fmt.Println(eklenenid)
	}

	{ //  İstenilen kaydı sorgulama
		var (
			id        int
			kullanici string
			sifre     string
			tarih     time.Time
			sorguid   int = 1
		)
		sorgu := "SELECT id, kullanici, sifre, tarih FROM kullanici WHERE id = ?" //normal sql sorgusu
		err := vt.QueryRow(sorgu, sorguid).Scan(&id, &kullanici, &sifre, &tarih)  //queryrow tek bir kayıt döndürür scan ile değişkenlere atarız

		fmt.Println(id, kullanici, sifre, tarih)
	}

	{ // Tün kayıtları sorgula
		type kullanici struct {
			id        int
			kullanici string
			sifre     string
			tarih     time.Time
		}
		tablo, err := vt.Query(`SELECT id, kullanici, sifre, tarih FROM kullanicilar`)

		defer tablo.Close() // işlem bitince kapat
		var kullanicilar []kullanici
		for tablo.Next() { // tablo next ile döngüye alınır
			var k kullanici
			err := tablo.Scan(&k.id, &k.kullanici, &k.sifre, &k.tarih)

			kullanici = append(kullanicilar, k) // append ile kullanicilar dizisine k nesnesini ekledik
		}
		err := rows.Err() //hata varsa

		fmt.Printf("%#v", kullanicilar)
	}

	//Kayıt Sil
	{
		silinecekid := 1
		_, err := vt.Exec(`DELETE FROM kullanicilar WHERE id = ?`, silinecekid) // exec ile sorgu çalıştırılır

	}
}

//yani sql işlemlerinde exec ile sorgu çalıştırılır ve sonucu alınır ,
// query ile sorgu çalıştırılır ve sonucu alınır ,
// queryrow ile sorgu çalıştırılır ve tek bir kayıt döndürülür
// scan ile değişkenlere atılır
// next ile döngüye alınır
// append ile diziye eklenir
