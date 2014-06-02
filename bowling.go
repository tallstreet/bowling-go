package bowling

import (
	"fmt"
)

const (
	totalPins = 10

	initialTries = 2

	totalFrames = 10
)

type Game struct {
	triesLeft int

	pastFrames [][]int

	currentFrame []int

	currentFrameHits int

	hasOldStrike bool

	hasLastStrike bool

	hasSpare bool

	totalScore int
}

func NewGame() *Game {
	game := &Game{}
	game.triesLeft = initialTries
	game.pastFrames = make([][]int, 0, totalFrames)
	game.currentFrame = make([]int, 0, initialTries)
	return game
}

func (game *Game) Roll(pins int) error {
	if !game.hasLastStrike && !game.hasOldStrike && !game.hasSpare && len(game.pastFrames) >= totalFrames {
		return fmt.Errorf("Game over")
	}
	if pins < 0 || pins > totalPins {
		return fmt.Errorf("Must roll between %d and %d pins", 0, totalPins)
	}
	if game.currentFrameHits+pins > totalPins {
		return fmt.Errorf("Can't hit more than %d in a frame", totalPins)
	}
	game.updateFrame(pins)
	game.updateScore(pins)

	if game.triesLeft == 0 || pins == totalPins {
		if len(game.pastFrames) < totalFrames {
			if pins == totalPins {
				game.hasLastStrike = true
			} else if game.currentFrameHits == totalPins {
				game.hasSpare = true
			}
		}
		game.nextFrame()
	}

	return nil
}

func (game *Game) Score() int {
	return game.totalScore
}

func (game *Game) nextFrame() {
	// Move to next frame
	game.pastFrames = append(game.pastFrames, game.currentFrame)
	game.currentFrame = make([]int, 0, initialTries)
	game.currentFrameHits = 0
	game.triesLeft = initialTries
}

func (game *Game) updateFrame(pins int) {
	game.currentFrame = append(game.currentFrame, pins)
	game.currentFrameHits += pins
	game.triesLeft--
}

func (game *Game) updateScore(pins int) {
	if game.hasSpare {
		game.totalScore += pins
		game.hasSpare = false
	}

	if game.hasOldStrike {
		game.totalScore += pins
		game.hasOldStrike = false
	}

	if game.hasLastStrike {
		game.totalScore += pins
		game.hasLastStrike = false
		game.hasOldStrike = true
	}

	if len(game.pastFrames) < totalFrames {
		game.totalScore += pins
	}
}
