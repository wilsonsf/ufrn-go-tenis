package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	playerScore int8
	skill       float32
	name        string
	opponent    *Player
}

func (p *Player) score() {
	p.playerScore++
	fmt.Printf("%s: ganhou um ponto. (%d)\n", p.name, p.playerScore)
}

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	player1 := Player{name: "Jogador 1"}
	player2 := Player{name: "Jogador 2"}

	makeOpponents(&player1, &player2)

	//player1toPlayer2 := make(chan bool)
	//player2toPlayer1 := make(chan bool)
	//gameIsRunning := make(chan bool)
	hasTheBall := make(chan bool)

	//if player1.hasTheBall {
	//	fmt.Printf("%s has the ball.\n", player1.name)
	//	fmt.Printf("Score: %d\n", player1.playerScore)
	//}
	//
	//if player2.hasTheBall {
	//	fmt.Printf("%s has the ball.\n", player2.name)
	//}

	go gameRoutine(&player1, hasTheBall)
	go gameRoutine(&player2, hasTheBall)

	hasTheBall <- true

	endGame(player1, player2)

	wg.Wait()
}

func makeOpponents(player1 *Player, player2 *Player) {
	player1.opponent = player2
	player2.opponent = player1
}
func endGame(player Player, player2 Player) {
	fmt.Println("End of game!")
	showScore(player, player2)
}

func showScore(player Player, player2 Player) (int, error) {
	return fmt.Printf("Score: %d-%d\n", player.playerScore, player2.playerScore)
}

func gameRoutine(player * Player, hasTheBall chan bool) {
	for ; player.playerScore < 4; {

		//defer wg.Done()

		keepPlaying := <-hasTheBall

		if !keepPlaying {
			break
		}

		player.receiveBall()
		player.play()
		//chance := rand.Float32() * 100
		//fmt.Printf("Chance de sucesso: %f", chance)
		//if chance > 70 {
		//	fmt.Printf("%s acertou a bola!\n", player.name)
		//	//devolve a bola
		//} else {
		//	fmt.Printf("%s errou a bola. :(\n",player.name)
		//	// O outro ganha ponto
		//}

		hasTheBall <- true
	}
	hasTheBall <- false
}

func (p *Player) play() {
	fmt.Printf("%s jogando.\n\n", p.name)

	chance := rand.Float32() * 100
	//fmt.Printf("Chance de sucesso: %.2f\n", chance)
	if chance > 70 {
		fmt.Printf("%s acertou a bola!\n", p.name)
		//devolve a bola
	} else {
		fmt.Printf("%s errou a bola. :(\n", p.name)
		p.opponent.score()
	}
}
func (p *Player) receiveBall() {
	fmt.Printf("%s recebeu a bola!\n", p.name)
}
