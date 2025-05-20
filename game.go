package main

import "github.com/matteo00gm/go-astar"

type Game struct {
	maze [][]int // maze grid: 0 for path, 1 for wall
	a    *astar.Astar

	playerX, playerY float32 // Player position in pixels (center of the ball)

	aiX, aiY   float32        // AI ball position in pixels (center of the ball)
	aiPath     []astar.Coords // The path AI is currently following
	aiPathStep int            // Index of the current step in the aiPath the AI is moving towards

	// Store player's grid position when the AI's path was last calculated
	lastPathCalcPlayerGridX, lastPathCalcPlayerGridY int
}

// NewGame creates and initializes a new Game struct.
func NewGame() *Game {

	// Find the starting position for the player (the first '0' tile)
	playerStartX, playerStartY := 0, 0
	for y := range gameMap {
		for x := range gameMap[y] {
			if gameMap[y][x] == 0 {
				playerStartX, playerStartY = x, y
				goto foundPlayerStart
			}
		}
	}
foundPlayerStart:

	// Find a starting position for the AI (the last '0' tile)
	aiStartX, aiStartY := 0, 0
	for y := len(gameMap) - 1; y >= 0; y-- {
		for x := len(gameMap[y]) - 1; x >= 0; x-- {
			if gameMap[y][x] == 0 {
				aiStartX, aiStartY = x, y
				goto foundAIStart
			}
		}
	}
foundAIStart:

	// Initialize A*
	a := astar.New(gameMap, &astar.EuclideanHeuristic{})

	// "playerStartX*tileSize" calculates the pixels on the X.
	// "tileSize/2" is adding half a tileSize. doing so, we get the position of the tile's center
	// doing this for both axys of both player and AI
	game := &Game{
		maze:    gameMap,
		a:       a,
		playerX: float32(playerStartX*tileSize + tileSize/2),
		playerY: float32(playerStartY*tileSize + tileSize/2),
		aiX:     float32(aiStartX*tileSize + tileSize/2),
		aiY:     float32(aiStartY*tileSize + tileSize/2),

		lastPathCalcPlayerGridX: playerStartX,
		lastPathCalcPlayerGridY: playerStartY,
	}

	game.calculateAIPath()

	return game
}
