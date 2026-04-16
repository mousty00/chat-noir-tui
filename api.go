package main

type Cat struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Color      string      `json:"color"`
	Category   CatCategory `json:"category"`
	Image      string      `json:"image"`
	SourceName string      `json:"sourceName"`
}

type CatCategory struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	MediaTypeHint string `json:"mediaTypeHint"`
}

type catsMsg struct {
	cats       []Cat
	totalItems int
	totalPages int
	err        error
}

type categoriesMsg struct {
	categories []CatCategory
	err        error
}

type catsResponse struct {
	Data struct {
		Result     []Cat `json:"result"`
		TotalItems int   `json:"totalItems"`
		TotalPages int   `json:"totalPages"`
	} `json:"data"`
}

type categoriesResponse struct {
	Data []CatCategory `json:"data"`
}
