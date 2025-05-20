package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) movePlayer() {
	currentX, currentY := g.playerX, g.playerY
	newPlayerX, newPlayerY := g.playerX, g.playerY

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		newPlayerY -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		newPlayerY += moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		newPlayerX -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		newPlayerX += moveSpeed
	}

	// Check player collision and update position
	if newPlayerX != currentX {
		if !g.checkCollision(newPlayerX, currentY) {
			g.playerX = newPlayerX
		}
	}
	if newPlayerY != currentY {
		if !g.checkCollision(currentX, newPlayerY) {
			g.playerY = newPlayerY
		}
	}
}

// checkCollision checks if the player at a potential new position collides with any walls.
func (g *Game) checkCollision(newX, newY float32) bool {
	playerHalfSize := float32(playerSize / 2)

	corners := []struct{ x, y float32 }{
		{newX - playerHalfSize, newY - playerHalfSize}, // Top-left
		{newX + playerHalfSize, newY - playerHalfSize}, // Top-right
		{newX - playerHalfSize, newY + playerHalfSize}, // Bottom-left
		{newX + playerHalfSize, newY + playerHalfSize}, // Bottom-right
	}

	for _, corner := range corners {
		if g.isWall(corner.x, corner.y) {
			return true // Collision detected
		}
	}

	return false
}

// isWall checks if a given pixel coordinate is inside a wall tile.
func (g *Game) isWall(x, y float32) bool {
	gridX := int(x / tileSize)
	gridY := int(y / tileSize)

	if gridY < 0 || gridY >= len(g.maze) || gridX < 0 || gridX >= len(g.maze[0]) {
		return true // Treat out of bounds as a wall
	}

	return g.maze[gridY][gridX] == 1
}
