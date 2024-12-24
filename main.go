package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const GRAVITY float32 = 50

type Ball struct {
	pos        rl.Vector2
	prevPos    rl.Vector2
	vel        rl.Vector2
	friction   float32
	elasticity float32
	radius     float32
	color      rl.Color
	dragging   bool
}

var ball Ball
var pressOffset rl.Vector2

func main() {

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(800, 450, "Go Test - Bouncing Ball")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	initGame()

	for !rl.WindowShouldClose() {
		update()
		draw()
	}

}

func initGame() {

	ball = Ball{
		pos: rl.Vector2{
			X: float32(rl.GetScreenWidth() / 2),
			Y: float32(rl.GetScreenHeight() / 2),
		},
		vel: rl.Vector2{
			X: 200,
			Y: 200,
		},
		friction:   0.99,
		elasticity: 0.9,
		radius:     50,
		color:      rl.Blue,
		dragging:   false,
	}

}

func update() {

	var delta float32 = rl.GetFrameTime()
	var mousePos rl.Vector2 = rl.GetMousePosition()

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if rl.CheckCollisionPointCircle(mousePos, ball.pos, ball.radius) {
			ball.dragging = true
			pressOffset.X = mousePos.X - ball.pos.X
			pressOffset.Y = mousePos.Y - ball.pos.Y
		}
	} else if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		ball.dragging = false
	}

	if ball.dragging {

		ball.pos.X = mousePos.X - pressOffset.X
		ball.pos.Y = mousePos.Y - pressOffset.Y

		ball.vel.X = (ball.pos.X - ball.prevPos.X) / delta
		ball.vel.Y = (ball.pos.Y - ball.prevPos.Y) / delta

		ball.prevPos = ball.pos

	} else {

		ball.pos.X += ball.vel.X * delta
		ball.pos.Y += ball.vel.Y * delta

		if ball.pos.X+ball.radius > float32(rl.GetScreenWidth()) {
			ball.pos.X = float32(rl.GetScreenWidth()) - ball.radius
			ball.vel.X = -ball.vel.X * ball.elasticity
		} else if ball.pos.X-ball.radius < 0 {
			ball.pos.X = ball.radius
			ball.vel.X = -ball.vel.X * ball.elasticity
		}

		if ball.pos.Y+ball.radius > float32(rl.GetScreenHeight()) {
			ball.pos.Y = float32(rl.GetScreenHeight()) - ball.radius
			ball.vel.Y = -ball.vel.Y * ball.elasticity
		} else if ball.pos.Y-ball.radius < 0 {
			ball.pos.Y = ball.radius
			ball.vel.Y = -ball.vel.Y * ball.elasticity
		}

		ball.vel.X *= ball.friction
		ball.vel.Y = ball.vel.Y*ball.friction + GRAVITY

	}

}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	drawBall(&ball)
	rl.EndDrawing()
}

func drawBall(ball *Ball) {
	rl.DrawCircleV(ball.pos, ball.radius, ball.color)
}
