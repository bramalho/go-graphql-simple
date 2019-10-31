package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

type Artist struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Album struct {
	ID     string `json:"id,omitempty"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
	Type   string `json:"type"`
}

type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

var artists []Artist = []Artist{
	Artist{
		ID:   "1",
		Name: "Led Zeppelin",
		Type: "artist",
	},
}

var albums []Album = []Album{
	Album{
		ID:     "lz-led-zeppelin",
		Artist: "1",
		Title:  "Led Zeppelin",
		Year:   "1969",
		Type:   "album",
	},
}

var songs []Song = []Song{
	Song{
		ID:       "1",
		Album:    "lz-led-zeppelin",
		Title:    "Good Times Bad Times",
		Duration: "2:46",
		Type:     "song",
	},
	Song{
		ID:       "2",
		Album:    "lz-led-zeppelin",
		Title:    "Babe I'm Gonna Leave You",
		Duration: "6:42",
		Type:     "song",
	},
}

func main() {
	artistType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Artist",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	albumType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Album",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"artist": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.String,
			},
			"genre": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	songType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Song",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"album": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"duration": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"artists": &graphql.Field{
				Type: graphql.NewList(artistType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return artists, nil
				},
			},
			"albums": &graphql.Field{
				Type: albumType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["id"].(string)
					for _, album := range albums {
						if album.ID == id {
							return album, nil
						}
					}
					return nil, nil
				},
			},
			"songs": &graphql.Field{
				Type: graphql.NewList(songType),
				Args: graphql.FieldConfigArgument{
					"album": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					album := params.Args["album"].(string)
					filtered := Filter(songs, func(v Song) bool {
						return strings.Contains(v.Album, album)
					})
					return filtered, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})

	err = http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Filter(songs []Song, f func(Song) bool) []Song {
	vsf := make([]Song, 0)
	for _, v := range songs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
