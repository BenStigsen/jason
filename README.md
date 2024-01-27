# Jason
Go module to better process and use dynamic JSON.

## Example
```json
{
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
}
```

```go
type VideoMetadata struct {
	Creator string `json:"creator"`
	Format  string `json:"format"`
}

func main() {
	data, err := os.ReadFile("videodata.json")
	if err != nil {
		panic(err)
	}

	content, err := jason.Unmarshal(data)
	if err != nil {
		panic(err)
	}

	videos := content.GetObjectArray("videos")
	for _, video := range videos {
		fmt.Println("title:", video.GetString("title"))
		fmt.Println("tags:", video.GetStringArray("tags"))
		fmt.Println("seconds:", video.GetNumber("seconds"))

		if video.IsValid("metadata") {
			meta := VideoMetadata{}
			video.GetObject("metadata").Unmarshal(&meta)
			fmt.Println("metadata:", meta)
		}

		fmt.Println("---")
	}
}
```

**Output:**
```
title: My video!
tags: [drama romantic]
seconds: 513
---
title: Another one!
tags: [comedy]
seconds: 123
metadata: {benjamin mp4}
---
```

Another way to access the metadata of a video would be:
```go
content.GetObjectArray("videos")[1].GetString("metadata", "creator")
```
