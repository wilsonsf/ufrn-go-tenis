package main

/*
	Author: Wilson Farias
	Date: 2018.06.24
*/
import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

/* Player define uma estrutura que encapsula os dados do jogador e sua pontuação. */
type Player struct {
	playerScore int8    // Pontuação do jogador por GAME
	gameScore   int8    // Quantidade de GAMEs ganhos
	setScore    int8    // Quantidade de SETs ganhos
	skill       float32 // Chance de acertar a bola
	name        string
	opponent    *Player // Referência para o oponente
}

/* Aumenta a pontuação do jogador e notifica. */
func (p *Player) score() {
	p.playerScore++
	fmt.Printf("%s ganhou um ponto. (%d)\n", p.name, p.playerScore)
}

/* Utilizado para aguardar final do Game. */
var wg sync.WaitGroup

/* Utilizado para definir o limite de pontos por Game. */
var gameScoreLimit = int8(4)
var gamesPerSetLimit = int8(1)
var setScoreLimit = int8(1)

func main() {
	/* Garantindo aleatoriedade das execuções */
	rand.Seed(time.Now().UnixNano())

	fmt.Print("Limite de pontos por GAME: ")
	fmt.Scanln(&gameScoreLimit)

	player1 := Player{name: "Gabriela", skill: float32(0.65)}
	player2 := Player{name: "Wilson", skill: float32(0.40)}

	makeOpponents(&player1, &player2)

	/* GAMES LOOP */
	//for ; keepRunningGames(player1, player2) ; {
	hasTheBall := make(chan bool)

	wg.Add(2)

	go gameRoutine(&player1, hasTheBall)
	go gameRoutine(&player2, hasTheBall)

	hasTheBall <- true

	wg.Wait()

	endGame(player1, player2)
	computeGameScores(&player1, &player2)
	//}
	//endSet(player1, player2)
}

func computeGameScores(p1 *Player, p2 *Player) {
	if p1.playerScore > p2.playerScore {
		p1.gameScore++
	} else {
		p2.gameScore++
	}
	p1.playerScore = 0
	p2.playerScore = 0
}

func keepRunningGames(p1 Player, p2 Player) bool {
	return p1.gameScore < gamesPerSetLimit &&
		p2.gameScore < gamesPerSetLimit
}

/* Adiciona referência de um jogador para o outro, e vice-versa. */
func makeOpponents(player1 *Player, player2 *Player) {
	player1.opponent = player2
	player2.opponent = player1
}

/* Encerra a o GAME anunciando o placar. */
func endGame(p1 Player, p2 Player) {
	fmt.Println("\n\nEnd of game!")
	showGameScore(p1, p2)
}

/* Exibe a pontuação do jogadores no GAME. */
func showGameScore(player Player, player2 Player) (int, error) {
	return fmt.Printf("Score: %d-%d\n", player.playerScore, player2.playerScore)
}

/* Encapsula a rotina de um GAME em goroutine. */
func gameRoutine(player *Player, hasTheBall chan bool) {
	defer wg.Done()
	for ; true; {

		keepPlaying, ok := <-hasTheBall

		if !ok {
			return
		}

		player.receiveBall()
		//gameOver :=
		player.play()

		if ok && !isGameOver(*player) {
			hasTheBall <- keepPlaying
		}
	}
	close(hasTheBall)
}

/* Anuncia quem está com a bola. */
func (p Player) receiveBall() {
	fmt.Printf("\n%s está com a bola!\n", p.name)
}

/* Encapsula a rotina de jogo, usando valores aleatórios e trata pontuação do GAME. */
func (p *Player) play() bool {
	fmt.Printf("%s jogando.\n", p.name)

	chance := rand.Float32()
	if chance-p.skill > 0 {
		fmt.Printf("%s acertou a bola!\n", p.name)
		return true
	}

	fmt.Printf("%s errou a bola. :(\n", p.name)
	p.opponent.score()
	return isGameOver(*p)
}

/* Calcula se deve encerrar o GAME. */
func isGameOver(p Player) bool {
	scoreDifference := p.playerScore - p.opponent.playerScore

	if scoreDifference < 0 {
		scoreDifference = -scoreDifference
	}

	return p.playerScore >= gameScoreLimit &&
		scoreDifference >= 2
}
