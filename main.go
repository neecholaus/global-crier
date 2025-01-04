package main

import "nick/global-crier/data"

func main() {
	sources := []data.Source{
		{
			Publication: "Reuters",
			Name:        "Reuters Video List",
			URL:         "https://www.reuters.com/pf/api/v3/content/fetch/video-playlist-by-slug-v1?query=%7B%22slug%22%3A%22%2Fvideo%2Fhp-us-2024-01-04%2F%22%2C%22website%22%3A%22reuters%22%7D&d=240&_website=reuter",
		},
	}

	data.Puller{}.Pull(sources[0])
}
