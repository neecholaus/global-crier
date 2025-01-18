package main

import (
	"nick/global-crier/headlines"
)

func main() {

	headlines.ReprocessExistingHeadlines()

	// for _, s := range headlines.Sources {
	// 	start := time.Now()

	// 	hlines, err := headlines.GetHeadlinesFromSource(s)
	// 	if err != nil {
	// 		fmt.Printf("ERR (%s - %s) pull failed (%s)\n", s.Publication, s.Name, err)
	// 		continue
	// 	}

	// 	duration := time.Since(start)
	// 	fmt.Printf("(%s - %s) pull took (%.1f) seconds to pull (%d) headlines\n", s.Publication, s.Name, duration.Seconds(), len(hlines))

	// 	headlines.ProcessHeadlines(hlines)
	// }
}
