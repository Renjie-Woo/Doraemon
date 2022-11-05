package main

import (
	"fmt"
	"github.com/Renjie-Woo/Doraemon/pkg/progressBar"
	"time"
)

func main() {
	fmt.Println("this is a progress bar")
	var title = "demo"
	var current float64 = -1
	var total float64 = 100
	var unit = "Mib"
	var newBar = progressBar.NewProgressBar(title, current, total)
	newBar.SetUnit(unit).
		SetGraph(">")
	for i := current; i <= total; i++ {
		newBar.Run(i)
		time.Sleep(time.Second / 500)
	}
}
