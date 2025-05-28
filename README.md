# Chase Game
This project is a simple 2D maze chase game developed in Go, designed to showcase the practical implementation of a custom A* pathfinding algorithm. The player navigates a maze while being pursued by an AI opponent that intelligently tracks the player using the A* algorithm.

## Demonstration
![Chase Game](/assets/chase.gif)

## Game Overview
In this game, you control a red ball (the player) within a grid-based maze. Your objective is to evade a blue ball (the AI opponent) that will actively try to find the optimal path to your current location. The AI's calculated path is visually rendered on the screen, allowing you to observe the A* algorithm in action.

## Features
* **2D Maze Environment:** A static, predefined maze built with walls (1s) and walkable spaces (0s).
* **Player Control:** Move the red player ball using the arrow keys (Up, Down, Left, Right). Movement is continuous while a key is held.
* **A\* Powered AI:** The blue AI ball uses a custom A\* pathfinding implementation (`github.com/matteo00gm/go-astar`) to determine the shortest path to the player.
* **Visible AI Path:** The path calculated by the A\* algorithm for the AI is drawn directly on the maze, providing a real-time visualization of the pathfinding process.
* **Optimized AI Recalculation:** The AI only recalculates its path when the player moves a significant distance (more than 3 tiles) or when the AI has completed its current path, ensuring smooth gameplay and efficient resource usage.
* **Collision Detection:** Basic collision detection prevents both the player and the AI from moving through maze walls.

## How the AI Uses A\*
The AI utilizes the `github.com/matteo00gm/go-astar` library, which implements the A\* search algorithm. Here's how it's integrated:

* **Maze as a Grid:** The game's 2D maze (represented as a `[][]int` grid) is provided to the A\* solver.
* **Start and End Points:** For each pathfinding request, the AI's current grid position serves as the `start` node, and the player's current grid position serves as the `end` (goal) node.
* **Path Calculation:** The `astar.FindPath` function is called, which computes the most efficient route through the walkable maze cells from the AI's position to the player's position.
* **Movement Along Path:** The AI then moves step-by-step along the calculated path. It prioritizes reaching the center of each target tile in the path before advancing to the next. This ensures smooth movement and avoids "stepping back" glitches.
* **Strategic Recalculation:** To prevent constant, unnecessary computations, the A\* path is only recalculated under specific conditions:
    * When the player moves more than 8 tiles away from the position where the last path was calculated.
    * When the AI has no current path (e.g., at the start of the game).
    * When the AI has successfully reached the end of its current path.

## Libraries
* **Ebitengine:** (`github.com/hajimehoshi/ebiten/v2`) used for rendering graphics, handling input, and managing the game loop.
* **go-astar:** Custom A\* pathfinding library (`github.com/matteo00gm/go-astar`) used for the AI's navigation.

## How to Run
To run this game locally, follow these steps:

* **Clone the Repository:** `git clone https://github.com/matteo00gm/chase-game.git`
* **Navigate into the project directory:** `cd chase-game`
* **Install Dependencies:**
    * `go get github.com/hajimehoshi/ebiten/v2`
    * `go get github.com/matteo00gm/go-astar`
* **Run the Game:** `go run main.go`

A new game window should appear, allowing you to play the maze chase game.

## License
This project is open-sourced under the MIT License.
