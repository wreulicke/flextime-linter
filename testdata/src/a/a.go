package a

import (
	"time" // want "Prefer use flextime"
)

func main() {
	time.Now() // want "Prefer use flextime"
}
