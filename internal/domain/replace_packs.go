package domain

import (
	"fmt"
)

// ReplacePacks is used to replace available pack sizes, which the app logic uses to calculate order packs.
func (a *Application) ReplacePacks(packs []int) {
	// the following operation should be protected by a Mutex for thread safety
	a.packSizes = packs

	fmt.Printf("Pack sizes replaced: %v", a.packSizes)
}
