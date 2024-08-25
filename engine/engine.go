package engine

type Engine interface {
	Init()
	Reset()
	SetupPosition()
	Play(moveNotation string)
	Stop()
	Quit()
	IsRunning() bool
}
