package main

import "fmt"

func main()  {
	videos := getVideos()
	for i := 0; i < len(videos); i++ {
		videos[i].Description = videos[i].Description + "\n Wagwan?"
	}
	
	fmt.Println(videos)

	saveVideos(videos)
}