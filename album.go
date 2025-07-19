package main

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "bluemoon", Artist: "Jack", Price: 25.60},
	{ID: "2", Title: "redmoon", Artist: "jaff", Price: 35.60},
	{ID: "1", Title: "whitemoon", Artist: "chris", Price: 45.60},
}
