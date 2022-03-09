package snake

import (
	"errors"
	"math/rand"
	"time"
)

// Board represents the game board.
type Board struct {
	rows     int
	cols     int
	food     *Food
	snake    *Snake
	points   int
	gameOver bool
}

// NewBoard generates a new Board with giving a size.
func NewBoard(rows int, cols int) *Board {
	rand.Seed(time.Now().UnixNano())

	board := &Board{
		rows: rows,
		cols: cols,
	}
	// start in top-left corner
	board.snake = NewSnake([]Coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}})
	board.placeFood()

	return board
}

// Update updates the board state.
func (b *Board) Update(input *Input) error {
	if b.gameOver {
		return nil
	}

	if newDir, ok := input.Dir(); ok {
		b.snake.ChangeDirection(newDir)
	}
	return nil
}

func (b *Board) placeFood() {
	var x, y int

	for {
		x = rand.Intn(b.cols)
		y = rand.Intn(b.rows)

		// make sure we don't put a food on a snake
		if !b.snake.HeadHits(x, y) {
			break
		}
	}

	b.food = NewFood(x, y)
}

func (b *Board) moveSnake() error {
	// remove tail first, add 1 in front
	b.snake.Move()

	if b.snakeLeftBoard() || b.snake.HeadHitsBody() {
		return errors.New("game over")
	}

	if b.snake.HeadHits(b.food.x, b.food.y) {
		// so the next move the snake grows
		b.snake.justAte = true

		b.placeFood()
		b.points++
	}

	return nil
}

func (b *Board) snakeLeftBoard() bool {
	head := b.snake.Head()
	return head.x > b.cols-1 || head.y > b.rows-1 || head.x < 0 || head.y < 0
}
