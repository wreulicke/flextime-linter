package a

import (
	x "time"

	fff "github.com/Songmu/flextime"
)

func b() {
	x.Now() // want "Prefer use flextime"
	fff.Now()
}
