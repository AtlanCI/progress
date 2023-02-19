package internal

import (
	"strconv"
	"testing"
	"time"
)

func ExampleSingleBar() {
	b := NewSingleBar("Prefix")

	b.UpdateDisplay(&DisplayProps{
		Prefix: "Prefix",
		Suffix: "Suffix",
	})

	n := 3

	go func() {
		time.Sleep(time.Second)
		for i := 0; i < n; i++ {
			b.UpdateDisplay(&DisplayProps{
				Prefix: "Prefix" + strconv.Itoa(i),
				Suffix: "Suffix" + strconv.Itoa(i),
			})
			time.Sleep(time.Second)
		}
	}()

	b.StartRenderLoop()

	time.Sleep(time.Second * time.Duration(n+1))

	b.UpdateDisplay(&DisplayProps{
		Mode: ModeDone,
		// Mode:   progress.ModeError,
		Prefix: "Prefix",
	})

	b.StopRenderLoop()
}

func TestExampleOutput(t *testing.T) {
	ExampleSingleBar()
}
