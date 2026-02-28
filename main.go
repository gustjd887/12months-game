package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 난수 시드 설정
	rand.Seed(time.Now().UnixNano())

	for {
		game := NewGame()
		game.Start()

		fmt.Print("\n다시 플레이하시겠습니까? (y/n): ")
		var answer string
		fmt.Scanln(&answer)

		if answer != "y" && answer != "Y" {
			fmt.Println("게임을 종료합니다. 안녕히 가세요!")
			break
		}
		fmt.Println()
	}
}
