package jason

import (
	"fmt"
	"testing"
)

type VideoMetadata struct {
	Creator string `json:"creator"`
	Format  string `json:"format"`
}

func TestJason(t *testing.T) {
	data := `{
		"videos": [
			{
				"title": "My video!",
				"tags": ["drama", "romantic"],
				"seconds": 513,
				"metadata": null
			},
			{
				"title": "Another one!",
				"tags": ["comedy"],
				"seconds": 123,
				"metadata": {
					"creator": "benjamin",
					"format": "mp4"
				}
			}
		]
	}`

	content, err := Unmarshal([]byte(data))
	if err != nil {
		panic(err)
	}

	videos := content.GetObjectArray("videos")
	for _, video := range videos {
		fmt.Println("title:", video.GetString("title"))
		fmt.Println("tags:", video.GetStringArray("tags"))
		fmt.Println("seconds:", video.GetNumber("seconds"))

		if video.IsValid("metadata") {
			metadata := VideoMetadata{}
			video.GetObject("metadata").Unmarshal(&metadata)
			fmt.Println("info:", metadata)
		}

		fmt.Println("---")
	}
}
