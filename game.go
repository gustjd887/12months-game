package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Game struct {
	score      int
	totalRound int
	correct    int
}

func NewGame() *Game {
	return &Game{
		score:      0,
		totalRound: 12,
		correct:    0,
	}
}

func (g *Game) Start() {
	fmt.Println("=== 12개월 영단어 맞추기 게임 ===")
	fmt.Println("12개월의 영단어를 맞추세요!")
	fmt.Println()

	for i := 1; i <= g.totalRound; i++ {
		if !g.askQuestion(i) {
			fmt.Println("게임을 종료합니다.")
			break
		}
	}

	g.showResult()
}

func (g *Game) askQuestion(round int) bool {
	// 1부터 12까지의 랜덤 숫자 생성
	monthNum := rand.Intn(12) + 1
	correctAnswer := getMonthByNumber(monthNum)

	fmt.Printf("[%d/%d 라운드] %d월의 영문명을 입력하세요: ", round, g.totalRound, monthNum)

	var userAnswer string
	fmt.Scanln(&userAnswer)

	// 대소문자 무시하고 비교
	if strings.EqualFold(userAnswer, correctAnswer) {
		fmt.Printf("✓ 정답! (%s)\n", correctAnswer)
		g.correct++
		g.score += 10
	} else {
		fmt.Printf("✗ 오답! 정답: %s\n", correctAnswer)
	}
	fmt.Println()

	return true
}

func (g *Game) showResult() {
	fmt.Println("=== 게임 결과 ===")
	fmt.Printf("총 점수: %d점\n", g.score)
	fmt.Printf("정답: %d/%d\n", g.correct, g.totalRound)
	fmt.Printf("정답률: %.1f%%\n", float64(g.correct)*100/float64(g.totalRound))

	if g.correct == g.totalRound {
		fmt.Println("🎉 완벽합니다! 축하합니다!")
	} else if g.correct >= 10 {
		fmt.Println("🌟 거의 다 맞추셨어요!")
	} else if g.correct >= 8 {
		fmt.Println("👍 좋은 점수입니다!")
	} else {
		fmt.Println("💪 더 열심히 공부해보세요!")
	}
}
