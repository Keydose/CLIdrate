package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

type hydrateTimer struct {
	startedAt    time.Time
	interval     int
	amount       int
	hydrationDue bool
}

func (t hydrateTimer) getElapsedTimeInSeconds() int {
	return int(time.Since(t.startedAt).Seconds())
}

func (t hydrateTimer) notify() {
	message := fmt.Sprintf("Drink %dml of water", t.amount)
	// TODO: Make the image smaller
	beeep.Notify("Hydrate!", message, "../../assets/images/water-droplet.png")
}

func main() {
	hydrateTimer := hydrateTimer{
		startedAt:    time.Now(),
		interval:     10,
		amount:       250,
		hydrationDue: false,
	}

	// TODO: Unit test???

	// TODO: Once passed interval, keep notifying every 30 seconds until
	//   - The user confirms that they've hydrated in the CLI
	prevElapsed := 0
	for {
		elapsed := hydrateTimer.getElapsedTimeInSeconds()
		if elapsed != prevElapsed {
			prevElapsed = elapsed
			if elapsed%hydrateTimer.interval == 0 {
				hydrateTimer.hydrationDue = true // TODO: Link in with CLI confirmation
				hydrateTimer.notify()
			}
		}
	}
}
