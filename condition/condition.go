package main

import (
	"fmt"
	"log"
	"math/rand"
)

// contoh penggunanan if else untuk angka
func oddOrEven(number int) {
	if number%2 == 0 {
		println("Even number") //jika sisa maka masuk sini
	} else {
		println("Odd number") // jika tidak sisa maka masuk ke sini
	}
}

//contoh fizz
/*
	number = bisa dibagi 3 & 5 yaitu fizzBuzz
	number = bisa dibagi 3 yaitu fizz
	number = bisa dibagi 5 yaitu buzz
	number = ngga bisa diagi 3 & 5 unknown/ tdk ditemukan
*/
func fizzBuzz(number int) {
	if number%3 == 0 && number%5 == 0 {
		println("fizzBuzz") //jika kondisi memenuhi maka output yg keluar ini
	} else if number%3 == 0 {
		println("fizz") //jika angka 3 tdk sama dg 0 maka output ini
	} else if number%5 == 0 {
		println("Buzz") // jika angka 5 tdk sama dg 0 maka output ini
	} else {
		println("unknown") //jika kondisi tdk terpenuhi maka output ini
	}
}

// contoh fungsi untuk random number
func randomNumber() {
	//range output : 0-9
	random := rand.Intn(10) //memunculkan random number yang dimulai dari 0 yang ditaruh acak
	println(random)

	//range output :1-10, supaya angka 10 bisa diambil
	min := 100
	max := 200
	//randomRange := rand.Intn(200-100)+100
	randomRange := rand.Intn(max-min) + min
	println("randomMin:", randomRange)

}

// contoh untuk inisial variabel
func initialIf() {
	//random := rand.Intn(10)
	if random := rand.Intn(10); random < 5 {
		println("low")
	} else if random >= 5 {
		println("high")
	}

	/*
		if random < 5 {
			println("low")
		}else if random >= 5 {
			println("high")
		}
	*/
}

// contoh menginput 3 variabel
func guessNumber() {
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)
	angka3 := angka1 - angka2

	fmt.Printf("Berapa hasil %d - %d = ?", angka1, angka2)
	//println("berapa hasil dari 7-3 =?")

	//decclare var jawab utk menampunng inputan dari console
	var jawab int              //untuk menampung ketikan dari console
	_, err := fmt.Scan(&jawab) //hasil inputan akan ditampung dg tipe integer
	if err != nil {            //
		log.Fatal(err)
	}
	if jawab == angka3 {
		println("Jawaban anda benar")
	} else {
		println("Jawaban anda salah")
	}
	println("Jawaban :", jawab)
}

func multipleArgs() {
	fmt.Println("Input 3 angka :")
	var angka1, angka2, angka3 int
	n, err := fmt.Scan(&angka1, &angka2, &angka3) // n utk menampung inputan
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("jawaban :", angka1, angka2, angka3, n)
}
func tebakHari(day int) {
	switch n := day; {
	case n == 0:
		fmt.Println("Sunday")
	case n == 1:
		fmt.Println("Monday")
	case n == 2:
		fmt.Println("Tuesday")
	case n == 3:
		fmt.Println("Wednesday")
	case n == 4:
		fmt.Println("Thursday")
	case n == 5:
		fmt.Println("Friday")
	case n == 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Hari Tidak Ditemukan")
	}

}

func main() {
	oddOrEven(8)
	fizzBuzz(15)
	randomNumber()
	initialIf()
	guessNumber()
	tebakHari(5)
	multipleArgs()
}
