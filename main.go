package main

import (
	"fmt"
	"time"

	"github.com/Renjie-Woo/Doraemon/progressBar"
)

func main() {
	fmt.Println("this is a progress bar")
	var title = "demo"
	var current = 12
	var total = 100
	var unit = "Mib"
	var newBar = progressBar.NewProgressBar(title, current, total)
	newBar.SetUnit(unit)
	newBar.SetGraph(">")
	for i := current; i <= total; i++ {
		newBar.Run(i)
		time.Sleep(time.Second / 100)
	}
}
