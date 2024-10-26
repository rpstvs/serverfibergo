package main

type Response struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
	Book   string `json:"book"`
}
