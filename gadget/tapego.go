package tapego

import (
	"fmt"
)

type TapePlayer struct {
	Charger string
}

func (t TapePlayer) TapePlayer() TapePlayer {
	return t.TapePlayer()
}

func (t TapePlayer) Play(song string) {
	fmt.Println("TapePlayer", "playes", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("TapeRecorder", "stop")
}

type TapeRecorder struct {
	Microphones []int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("TapeRecorder", "playes", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("TapeRecorder", "stop")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording...")
}
