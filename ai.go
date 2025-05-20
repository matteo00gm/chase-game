package main

import (
	"math"

	"github.com/matteo00gm/go-astar"
)

// calculateAIPath finds a path from the AI's current position to the player's position.
func (g *Game) calculateAIPath() {

	// Get the AI's current grid position
	aiGridX := int(g.aiX / tileSize)
	aiGridY := int(g.aiY / tileSize)
	aiStartCoords := astar.Coords{X: aiGridX, Y: aiGridY} // Use astar.Coords

	// Get the player's current grid position
	playerGridX := int(g.playerX / tileSize)
	playerGridY := int(g.playerY / tileSize)
	playerTargetCoords := astar.Coords{X: playerGridX, Y: playerGridY} // Use astar.Coords

	// Use the astar instance to find the path
	found, path := g.a.FindPath(aiStartCoords, playerTargetCoords)

	if found {
		//set the new path for the AI
		g.aiPath = path
		g.aiPathStep = 0 // Reset the next step in the path for AI
	} else {
		g.aiPath = nil // No path found
	}
}

func (g *Game) moveAI() {
	//checks if it need to update the current path the ai is taking
	g.checkAndUpdatePath()
	//checks if it is arrived to the destination. else, keep moving
	g.checkAndMoveAI()
}

func (g *Game) calcPlayerMovement() (playerGridX, playerGridY int, movement float64) {
	// Get the player's current grid position
	playerGridX = int(g.playerX / tileSize)
	playerGridY = int(g.playerY / tileSize)

	// Calculate the distance from the current player position to the last calculated one
	delta := float64((playerGridX-g.lastPathCalcPlayerGridX)*(playerGridX-g.lastPathCalcPlayerGridX) + (playerGridY-g.lastPathCalcPlayerGridY)*(playerGridY-g.lastPathCalcPlayerGridY))
	return playerGridX, playerGridY, math.Sqrt(delta)
}

func (g *Game) checkAndUpdatePath() {
	//calculate and return the starting x,y player position and the distance he made
	currentPlayerX, currentPlayerY, distance := g.calcPlayerMovement()

	// Recalculate AI path if:
	// 1. The player has moved more than 'recalculationDistance' tiles away from the last path calculation position.
	// 2. The AI currently has no path.
	// 3. The AI has reached the end of its current path.
	if distance > float64(recalculationDistance) || len(g.aiPath) == 0 || g.aiPathStep >= len(g.aiPath) {
		g.calculateAIPath()
		// Update the last path calculation position with the player's current grid position
		g.lastPathCalcPlayerGridX = currentPlayerX
		g.lastPathCalcPlayerGridY = currentPlayerY
	}
}

func (g *Game) checkAndMoveAI() {
	// Move the AI along the calculated path
	if len(g.aiPath) > 0 && g.aiPathStep < len(g.aiPath) {
		targetNode := g.aiPath[g.aiPathStep]
		targetX := float32(targetNode.X*tileSize + tileSize/2)
		targetY := float32(targetNode.Y*tileSize + tileSize/2)

		// Calculate direction vector towards the target node center
		dx := targetX - g.aiX
		dy := targetY - g.aiY

		// Calculate distance to the target node center
		distanceToTarget := float32(math.Sqrt(float64(dx*dx + dy*dy)))

		// Check if the AI is close enough to the center of the current target tile
		if distanceToTarget < tileCenterTolerance {
			// Snap to the center of the target node and move to the next step
			g.aiX = targetX
			g.aiY = targetY
			g.aiPathStep++
		} else {
			// Move the AI ball towards the target node center
			// Normalize direction vector and scale by AI move speed
			moveX := (dx / distanceToTarget) * aiMoveSpeed
			moveY := (dy / distanceToTarget) * aiMoveSpeed
			g.aiX += moveX
			g.aiY += moveY
		}
	}
}
