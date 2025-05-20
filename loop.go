package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Update function is called every tick (by the ebitenengine library) to update the game logic
func (g *Game) Update() error {

	g.movePlayer()
	g.moveAI()

	return nil
}

// Draw function is called every frame to draw the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.drawMaze(screen)

	// Draw the player (red ball)
	playerRadius := float32(playerSize / 2)
	vector.DrawFilledCircle(screen, g.playerX, g.playerY, playerRadius, color.RGBA{255, 0, 0, 255}, false)

	// Draw the AI ball (blue ball)
	aiRadius := float32(aiSize / 2)
	vector.DrawFilledCircle(screen, g.aiX, g.aiY, aiRadius, color.RGBA{0, 0, 255, 255}, false)

	if showPath {
		g.drawPath(screen)
	}
}

// Layout takes the native window size and returns the game's logical screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) drawMaze(screen *ebiten.Image) {
	// Draw the maze tiles
	for y := range g.maze {
		for x := range g.maze[y] {
			tileX := float32(x * tileSize)
			tileY := float32(y * tileSize)

			if g.maze[y][x] == 1 {
				vector.DrawFilledRect(screen, tileX, tileY, tileSize, tileSize, color.Black, false)
			} else {
				vector.DrawFilledRect(screen, tileX, tileY, tileSize, tileSize, color.RGBA{200, 200, 200, 255}, false)
			}
		}
	}
}

func (g *Game) drawPath(screen *ebiten.Image) {
	if len(g.aiPath) > 1 {
		for i := range len(g.aiPath) - 1 {
			fromNode := g.aiPath[i]
			toNode := g.aiPath[i+1]
			fromX := float32(fromNode.X*tileSize + tileSize/2)
			fromY := float32(fromNode.Y*tileSize + tileSize/2)
			toX := float32(toNode.X*tileSize + tileSize/2)
			toY := float32(toNode.Y*tileSize + tileSize/2)
			vector.StrokeLine(screen, fromX, fromY, toX, toY, 2, color.Gray{}, false)
		}
	}
}
