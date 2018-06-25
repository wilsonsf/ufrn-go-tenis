# Partida de Tênis

## Regras
- 2 jogadores  
- Cada jogador joga, quando a bola está na sua quadra
- Pode acertar ou não (aleatoreidade)

>  Definir uma taxa de sucesso (70%?)  
  Random.float entre 0 e 1 (já está em percentual)  
  Adiciona a taxa de sucesso calcula o max entre 0 e 1.  
    Ou: reduz a taxa de sucesso do valor calculado, se for menor que 0, errou.


## Pontuação
#### GAME
Cada GAME tem contagem de pontos, 0-15-30-40
Se não houver vantagem de 2 pontos no 4º ponto, entra na segunda forma de contagem
 0-0 1-0 2-0 ou 0-1 0-2
  Empatou, volta pra 0-0

#### SET
Um SET tem vários GAME.
  A contagem de GAME vai de 1 a 6.
  Ganha quem tiver ao menos 6 pontos, com 2 de vantagem.


#### MATCH
Um MATCH é composto de vários SET
  Por ora, são 5 SET por MATCH.
  Ganha quem ganhar ao menos 3 com 2 de vantagem.


## Objetivos
#### Primeira etapa
Inicialmente:
1 Match - 1 set - 1 Game
Ganha o game que atingir o número de pontos P primeiro.

#### Segunda etapa
1. Personalizar o número de pontos P.
2. Personalizar o número de GAMEs.
3. Personalizar o número de SETs.

## Algoritmo
```
Começa a partida

Enquanto condição de match não for verdadeira
  Enquanto condição de set não for verdadeira
    Enquanto condição de game não for verdadeira
      Contabiliza os pontos
      Quem tá com a bola reage
        Se reagir correto, passa a bola.
        Se reagir errado, o outro ganha ponto.
      Se a pontuação atingiu P, termina game em funçào do vencedor.
    Termina o game.
  Termina o set.
Termina o match.
```






