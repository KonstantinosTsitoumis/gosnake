package entities

import (
	"fmt"
	"time"
)

type buff struct {
	expiresAt time.Time
	name      string
}

type body struct {
	position Position
	previous *body
	next     *body
}

type Snake struct {
	direction string
	head      *body
	buff      *buff
}

func NewSnake(x int, y int) Snake {
	head := body{
		position: Position{
			X: x,
			Y: y,
		},
		next:     nil,
		previous: nil,
	}

	head.previous = &body{
		position: Position{
			X: head.position.X - 1,
			Y: head.position.Y,
		},
		previous: nil,
		next:     &head,
	}

	return Snake{
		head:      &head,
		direction: Direction[Right],
		buff:      nil,
	}
}

func (snake *Snake) GetSnakeBodyPositions() []Position {
	body := snake.head

	result := []Position{}

	for body != nil {
		result = append(result, body.position)

		body = body.previous
	}

	return result
}

func (snake *Snake) Move(nextCellPosition Position) {
	newHead := &body{
		position: nextCellPosition,
		previous: snake.head,
		next:     nil,
	}

	snake.head.next = newHead
	snake.head = newHead

	PreLastBody := snake.GetLastBody().next

	PreLastBody.previous = nil
}

func (snake *Snake) GetDirection() string {
	return snake.direction
}

func (snake *Snake) SetDirection(direction string) {
	head := snake.head
	previous := head.previous

	moveX, moveY := DirectionToXY(direction)

	if head.position.X+moveX == previous.position.X && head.position.Y+moveY == previous.position.Y {
		return
	}

	snake.direction = direction
}

func (snake *Snake) GetLastBody() *body {
	body := snake.head
	previous := body.previous

	for previous != nil {
		body = previous
		previous = body.previous
	}

	return body
}

func (snake *Snake) eat(nextCellPosition Position) {
	newHead := &body{
		position: nextCellPosition,
		previous: snake.head,
		next:     nil,
	}

	snake.head.next = newHead
	snake.head = newHead
}

// Update
func (snake *Snake) Update(direction string, apple *Apple, energyDrink *EnergyDrink, gameover *bool) {
	if snake.buff != nil && snake.buff.expiresAt.Before(time.Now()) {
		fmt.Println("Buff expired!")
		snake.buff = nil
	}

	if snake.GetDirection() != direction {
		snake.SetDirection(direction)
	}

	moveX, moveY := DirectionToXY(snake.direction)

	var speed int
	if snake.buff != nil && snake.buff.name == "speed" {
		speed = 2
	} else {
		speed = 1
	}

	for i := 0; i < speed; i++ {
		nextPosition := Position{
			X: snake.head.position.X + moveX,
			Y: snake.head.position.Y + moveY,
		}

		// is gameover
		for _, position := range snake.GetSnakeBodyPositions() {
			if nextPosition == position {
				*gameover = true
				return
			}
		}

		// is action
		if nextPosition == apple.Position {
			snake.eat(nextPosition)
			apple.Eaten = true
		} else if nextPosition == energyDrink.Position {
			snake.Move(nextPosition)
			snake.buff = &buff{
				name:      "speed",
				expiresAt: time.Now().Add(3 * time.Second),
			}
			energyDrink.Eaten = true
		} else {
			snake.Move(nextPosition)
		}
	}
}

// Helper
func DirectionToXY(direction string) (int, int) {
	moveX := 0
	moveY := 0

	switch direction {
	case Direction[Up]:
		moveY = -1
		break
	case Direction[Down]:
		moveY = 1
		break
	case Direction[Left]:
		moveX = -1
		break
	case Direction[Right]:
		moveX = 1
		break
	}

	return moveX, moveY
}
