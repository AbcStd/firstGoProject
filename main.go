package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	for range 100 {
		random = rand.N(100) + 1
		guesses = 0
		result := play()
		fmt.Println(random, result)
		if result != random {
			fmt.Printf("Неверный ответ. Было загадано число %d, а в ответе получили число %d", random, result)
			os.Exit(-1)
		}
	}
}

var guesses int
var random int

func guess(num int) (int, error) {
	if guesses >= 6 {
		return 0, errors.New("too many attempts")
	}
	guesses++
	if num > random {
		return -1, nil
	}
	if num < random {
		return 1, nil
	}
	return 0, nil
}

func play() int {
	left, right := 1, 100

	for left <= right {
		mid := left + (right-left)/2
		result, err := guess(mid) // ← теперь проверяем ошибку
		if left == right {
			return left
		}

		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			// Если превышено количество попыток, но мы ещё не угадали — плохо
			// В данной задаче мы контролируем вызовы, так что можно вернуть -1

		}

		switch result {
		case 0:
			return mid
		case -1:
			right = mid - 1
		case 1:
			left = mid + 1
		}
	}

	return -1
}
