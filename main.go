package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numStr string
	fmt.Print("Введите вашу оценку в формате от 0 до 100: ")
	fmt.Scanln(&numStr)
	numStr = strings.TrimSpace(numStr)

	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("Ошибка: '%s' — это не целое число. Пожалуйста, введите число от 0 до 100.\n", numStr)
		return
	}

	letter, err := transDigLetter(num)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("Ваша оценка  -%d- или -%s-\n", num, letter)
	}

}

func transDigLetter(num int) (string, error) {

	switch {
	case num < 0 || num > 100:
		return "", fmt.Errorf("недопустимое значение: %d (должно быть от 0 до 100)", num)
	case num >= 90:
		return "A", nil
	case num >= 80:
		return "B", nil
	case num >= 70:
		return "C", nil
	case num >= 60:
		return "D", nil

	default:
		return "F", nil
	}

}

// package main55

// import (
// 	"fmt"
// )

// func main() {
// 	// 10 вызовов с разными значениями isTrap
// 	movePirate(true)
// 	movePirate(true)
// 	movePirate(false)
// 	movePirate(true)
// 	movePirate(false)
// 	movePirate(false)
// 	movePirate(false)
// 	movePirate(false)
// 	movePirate(false)
// 	movePirate(true)
// }

// var stepCounter int
// var crashStep int

// func movePirate(isTrap bool) {
// 	stepCounter++

// 	if crashStep >= 3 {
// 		return
// 	}

// 	fmt.Printf("Пират переместился на плиту %d\n", stepCounter)

// 	if isTrap {
// 		crashStep++
// 		if crashStep == 3 {
// 			fmt.Println("Пират убит")
// 		} else {
// 			fmt.Println("Пират ранен")
// 		}
// 	}

// 	if stepCounter == 10 && crashStep < 3 {
// 		fmt.Println("Пират преодолел все ловушки")
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	traps := []bool{false, true, false, false, true, false, false, true, false, false}

// 	for i := 0; i < 10; i++ {
// 		movePirate(traps[i], i+1)
// 	}
// }

// func movePirate(isTrap bool, step int) {
// 	// Просто выводим перемещение, независимо от isTrap
// 	fmt.Printf("Пират переместился на плиту %d\n", step)

// 	// Если нужно дополнительно выводить информацию о ловушке:
// 	if isTrap {
// 		fmt.Printf("  На плите %d ловушка!\n", step)
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	isTrap :=
// 	totalCount := 0
// 	totalCount = movePirate(totalCount, isTrap)

// 	fmt.Printf("Пират переместился на плиту %d", totalCount)

// }
// func movePirate(currentCount int, isTrap bool) int {

// 	if !isTrap {
// 		return currentCount + 1
// 	}
// 	return currentCount + 1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	fmt.Print(sumOfDigits(1111111111111111111))
// }

// func sumOfDigits(a int) int {
// 	if a < 0 {
// 		a = -a
// 	}
// 	if a == 0 {

// 		return 0
// 	}
// 	return a%10 + sumOfDigits(a/10)

// }

// package main

// import (
// 	"fmt"
// )

// type Day int

// const (
// 	_ Day = iota
// 	Monday
// 	Tuesday
// 	Wednesday
// 	Thursday
// 	Friday
// 	Saturday
// 	Sunday
// )

// func main() {

// 	result := isWeekend(7)

// 	fmt.Println(result)

// }

// func isWeekend(Day Day) bool {
// 	return Day == Saturday || Day == Sunday

// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	id := "Alekc"
// 	resultFinal, err := userProfile(id)
// 	if err != nil {
// 		fmt.Printf("Ошибка %s: \n", err)
// 	} else {
// 		fmt.Println(resultFinal)
// 	}

// }

// func userProfile(id string) (string, error) {

// 	result_kop, err := fetchUserInfo(id)
// 	if err != nil {
// 		return "", fmt.Errorf("fetch error: %w", err)
// 	}
// 	result_rub := float64(result_kop) / 100.0

// 	result_str := fmt.Sprintf("Пользователь с id %s имеет на счету %.2f руб.", id, result_rub)
// 	return result_str, nil

// }

// func fetchUserInfo(id string) (int, error) {
// 	result := 356
// 	if id == "Alekc" {
// 		return result, nil
// 	}
// 	return 0, errors.New("fetch error: mistake")

// }

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	var a, b float64
// 	a = 8.4
// 	b = 0.5
// 	s := "divide"
// 	result, err := calculate(a, b, s)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(result)

// }
// func calculate(a, b float64, s string) (float64, error) {

// 	switch s {
// 	case "add":
// 		return a + b, nil
// 	case "subtract":
// 		return a - b, nil
// 	case "multiply":
// 		return a * b, nil
// 	case "divide":
// 		if b != 0 {
// 			return a / b, nil
// 		} else {
// 			return 0, errors.New("division by zero")
// 		}

// 	}

// 	return 0, errors.New("unknown operation")
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	name, age := "   dxfg    ", 19
// 	result, err := UserProfileToString(name, age)
// 	if err != nil {
// 		fmt.Println("Получена ошибка от функции:", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(result)

// }

// func UserProfileToString(name string, age int) (string, error) {

// 	if name == "" {
// 		return name, errors.New("empty name")
// 	}
// 	if age < 0 {
// 		return name, errors.New("negative age")
// 	}

// 	if strings.TrimSpace(name) == "" {
// 		return name, errors.New("name cannot contain only spaces")
// 	}

// 	result := fmt.Sprintf("Имя человека: %s, возраст: %d.", strings.TrimSpace(name), age)
// 	return result, nil

// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	fmt.Print("Как ваше имя?\n")
// 	var name string

// 	fmt.Scan(&name)

// 	generateCompliment(name)
// 	// fmt.Print(result)
// }

// func generateCompliment(name string) string {

// 	strings := []string{"Ты великолепен, ", "У тебя потрясающая улыбка, ", "Ты вдохновляешь, "}
// 	randomString := strings[rand.Intn(len(strings))]

// 	result := randomString + name + "!"
// 	return result

// }

/*
"Ты великолепен, [имя]!"
"У тебя потрясающая улыбка, [имя]!"
"Ты вдохновляешь, [имя]!"
*/

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	num := 5
// 	printNumberInfo(num)
// }

// func printNumberInfo(num int) {
// 	if num < 0 {
// 		fmt.Printf("Число %.d отрицательное.\n", num)
// 		if num%2 == 0 {
// 			fmt.Printf("Число %d четное.\n", num)
// 		} else {
// 			fmt.Printf("Число %d нечетное.\n", num)
// 		}
// 	} else if num > 0 {
// 		fmt.Printf("Число %.d положительное.\n", num)
// 		if num%2 == 0 {
// 			fmt.Printf("Число %d четное.\n", num)
// 		} else {
// 			fmt.Printf("Число %d нечетное.\n", num)
// 		}
// 		sq(num)

// 	} else {
// 		fmt.Println("Число равно 0.")
// 		fmt.Println("Число 0 четное.")
// 	}

// }
// func sq(num int) {

// 	sqnum := math.Sqrt(float64(num))

// 	switch {
// 	case sqnum == math.Trunc(sqnum):
// 		fmt.Printf("Квадратный корень числа %.d является целым числом и равен %.0f.\n", num, sqnum)
// 	default:
// 		fmt.Printf("Квадратный корень числа %.d не является целым числом и равен %.5f.\n", num, sqnum)

// 	}

// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var cats = 4
// 	var dogs = 4

// 	PetBattle(cats, dogs)
// }
// func PetBattle(cats, dogs int) {
// 	switch {
// 	case cats > dogs:
// 		fmt.Printf("Котики победили со счетом %d:%d!", cats, dogs)

// 	case cats < dogs:
// 		fmt.Printf("Собачки победили со счетом %d:%d!", dogs, cats)

// 	case cats == dogs:
// 		fmt.Printf("Ничья! Все дружат!")
// 	}

// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print("Введите вес (кг): ")
// 	var weight float64

// 	_, err := fmt.Scan(&weight)
// 	if err != nil {
// 		fmt.Println("Ошибка ввода")
// 		return
// 	}
// 	// fmt.Printf("Ваш вес: %.2f\n", weight)

// 	fmt.Print("Введите рост (см): ")
// 	var height float64
// 	_, errr :=fmt.Scan(&height)
// 	if errr != nil {
// 		fmt.Println("Ошибка ввода")
// 		return
// 	}
// 	// fmt.Printf("Ваш рост: %.2f\n", height/100)
// 	height_m := height/100

// 	imt := weight / (height_m*height_m)
// 	fmt.Printf("Ваш ИМТ: %.2f\n", imt)

// // Недостаточный вес: ИМТ < 18.5
// // Нормальный вес: 18.5 ≤ ИМТ < 25
// // Избыточный вес: 25 ≤ ИМТ < 30
// // Ожирение: ИМТ ≥ 30

// // Введите ваш вес (кг): 70
// // Введите ваш рост (см): 175
// // Ваш ИМТ: 22.86
// // Категория: Нормальный вес
// switch {
// case imt < 18.5:
// 	fmt.Println("Категория: Недостаточный вес")
// case 18.5 <= imt && imt < 25:
// 	fmt.Println("Категория: Нормальный вес")
// case 25 <= imt && imt < 30:
// 	fmt.Println("Категория: Избыточный вес")
// case imt >= 30:
// 	fmt.Println("Категория: Ожирение")
// }

// }

// func main() {
// 	fmt.Print("Введите время в формате от 0 до 24: ")
// 	var clock int
// 	_, err := fmt.Scan(&clock)
// 	if err != nil {
// 		fmt.Println("Ошибка ввода")
// 		return
// 	}
// 	switch {
// 	case (clock > 24 || clock < 0):
// 		fmt.Printf("Неверно задано время - %d.", clock)

// 	case (clock >= 6 && clock < 12):
// 		fmt.Printf("Сейчас %dч. - утро.", clock)

// 	case (clock >= 12 && clock < 18):
// 		fmt.Printf("Сейчас %dч. - день.", clock)

// 	case (clock >= 18 && clock < 23):
// 		fmt.Printf("Сейчас %dч. - вечер.", clock)

// 	case (clock >= 23 || clock < 6):
// 		fmt.Printf("Сейчас %dч. - ночь.", clock)

// 	default:
// 		fmt.Println("Неверно задано время.")
// 	}

// }

// func main() {
// 	var null any
// 	var val any = null

// 	switch x := val.(type) {
// 	case int, float64, string, bool:
// 		fmt.Printf("В переменной val находится тип %T.\n", x)
// 	default:
// 		fmt.Println("В переменной val находится неизвестный тип данных.")
// 	}

// }

// func main() {
// 	temp := 45
// 	if temp < 0 {
// 		fmt.Println("Город замерзает! Верните лето.")
// 	}
// 	if temp >= 0 && temp <= 35 {
// 		fmt.Println("Температура в норме. Продолжаем писать код.")

// 	}
// 	if temp > 35 {
// 		fmt.Println("Город в огне! Яичницу можно жарить на асфальте.")
// 	}

// }

// func main() {

// 	// В программе у нас уже есть полученные данные, которые лежат в переменных:

// 	// age - целое число, переменная хранит возраст пользователя.
// 	// role - строка, роль пользователя на сайте, значения могут быть: "admin", "moderator", "user".
// 	// status - строка, статус подписки, значения могут быть: "active", "inactive", "paused".
// 	// Требования
// 	// Доступ может быть предоставлен только в том случае, если у пользователя статус активный, за некоторым исключением:

// 	// Если пользователь несовершеннолетний (меньше 18 лет), в таком случае, контент ему не предоставляется.
// 	// Если роль пользователя admin или moderator, то он имеет доступ к контенту в любом случае, не важно сколько ему лет и какой у него статус.
// 	// Если роль пользователя не соответствует ни одному из трех возможных значений, доступ должен быть запрещен.
// 	// Необходимо вывести булев тип данных (true/false), как результат работы нашего приложения.
// 	// var role string
// 	// var status string
// 	// var age int

// 	age := 18
// 	role := "user"
// 	status := "active"

// 	result := ((role == "admin") || (role == "moderator")) || (age >= 18 && status == "active" && role == "user")
// 	fmt.Println(result)

// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// // Исходное число: 53.2
// // Исходное число, увеличенное на 10%: 58.52000
// // Исходное число является четным: false
// // Предпоследняя цифра целой части исходного числа: 5

// func main() {
// 	random := 54.3

// 	fmt.Printf("Исходное число: %.1f\n", random)

// 	fmt.Printf("Исходное число, увеличенное на 10%%: %.5f\n", random*1.10)

// 	fmt.Println("Исходное число является четным:", math.Mod(random, 2) == 0)

// 	fmt.Printf("Предпоследняя цифра целой части исходного числа: %d\n", int(random/10))
// }

// func main() {
// 	var firstname, lastname string
// 	var age int
// 	fmt.Print("Введите ваше имя и фамилию: ")
// 	fmt.Scanln(&firstname, &lastname)
// 	fmt.Print("Введите ваш возраст: ")
// 	fmt.Scanln(&age)

// 	fmt.Printf(
// 		"Приятно познакомиться, %s. Я 5 лет назад познакомился с человеком, у которого тоже фамилия %s, вам тогда было %d. Как молоды мы были!",
// 		firstname, lastname, age-5)
// }

// func main() {
// var str string
// str = "𝓗𝓮𝓵𝓵𝓸, мой друг."
// fmt.Println(str, len(str), utf8.RuneCountInString(str))

// }

// func main() {
// 	discountPercent := 11.111
// 	productPrice := 100.0
// //	var c float64 = 0.0
// 	c := 0.0
// 	c_t := 0.0
// 	c_t = (productPrice - (productPrice * discountPercent / 100))
// 	c = math.Floor((productPrice - (productPrice * discountPercent / 100)) * 100) / 100
// 	fmt.Println(c_t)
// 	fmt.Println(c)
// }
