package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func notify(hydrateAmount int) {
	message := fmt.Sprintf("Drink %dml of water", hydrateAmount)
	fmt.Println(message)
	// TODO: Make the image smaller
	beeep.Notify("Hydrate!", message, "../../assets/images/water-droplet.png")
}

func tick(hydrateInterval int, hydrateAmount int) {
	ticker := time.NewTicker(time.Duration(hydrateInterval) * time.Second)
	done := make(chan bool)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			go notify(hydrateAmount)
		}
	}
}

func main() {
	hydrateIntervalPtr := flag.Int("interval", 3600, "Hydrate interval in seconds")
	hydrateAmountPtr := flag.Int("amount", 250, "Hydrate amount in millilitres")

	flag.Parse()

	hydrateInterval := *hydrateIntervalPtr
	hydrateAmount := *hydrateAmountPtr

	fmt.Fprintf(os.Stdout, "I'll remind you to hydrate by %dml every %d seconds\n", hydrateAmount, hydrateInterval)

	go tick(hydrateInterval, hydrateAmount)
	select {}
	// TODO: Unit test???

	// TODO: Once passed interval, keep notifying every 30 seconds until
	//   - The user confirms that they've hydrated in the CLI
}
