package main

import "fmt"

var (
	aple    = 5.99
	pear    = 7
	myMoney = 23
)

func main() {

	fmt.Println("в нас дано", aple, pear, myMoney)
	sumAple := float64(myMoney) / float64(aple)
	fmt.Println("3. Скільки яблук ми можемо купити?", int(sumAple))

	sumPear := float64(myMoney) / float64(pear)
	fmt.Println("Скільки груш ми можемо купити?", int(sumPear))

	// milkPrice := 200
	// milkBear := 100

	// carPrise := 20000.399
	// fmt.Println("в нас дано", carPrise, milkPrice, milkBear)
	// m := milkPrice > int(carPrise) || milkPrice > milkBear
	// fmt.Println("в нас дано", m)

	// fmt.Println("в нас дано", carPrise, milkPrice, milkBear)
	// carMoreExpensive := carPrise > milkPrice && carPrise > milkBear
	// fmt.Println("машина дорожча в обов випадках", carMoreExpensive)

	// fmt.Println(milkPrice, milkBear)

	// price := milkPrice < milkBear
	// fmt.Println("чи молоко дешевше?", price)
	// fmt.Println(milkPrice, milkBear)
	// isPriceEqual := (milkPrice == milkBear)
	// fmt.Println("чи одинакова ціна?", isPriceEqual)

	// myMOney := 301
	// fmt.Println("дано", milkPrice, milkBear)
	// fmt.Println("чи вистачить мені грошей", (milkPrice+milkBear) < myMOney)
	// ОРИГІНАЛЬНИЙ ТИП І ЗНАЧЕННЯ ЗМІННЛОЇ ЗАЛИШАЮТЬСЯ ТАКИМИ ЯК БУЛИ ПОПРИ КОНВЕРТУВАННЯ
	// x := 15
	// fmt.Println(f1, x)
	// mult := f1 * float32(x)
	// fmt.Println(mult)

	// y := 10
	// fmt.Println(y)
	// sum := y * x
	// fmt.Println(sum)

	// КОНВЕРТУВАННЯ FLOAT32, 64
	// someInt := 10
	// fmt.Println("дано ", f1, someInt)
	// mult := f1 * float32(someInt)
	// fmt.Println(mult)

	// 	b, a := "a", "s"
	// 	println(b, a)

	// 	b, a = "s", "a"
	// 	println(b, a)
	// }
	// var (
	// 	i1 = 2
	// 	u2 = 4
	// )
	// i1++
	// //u2 = u2 - 1
	// fmt.Print(i1, u2)

	// sum := i1 + u2
	// fmt.Print(sum)

}
