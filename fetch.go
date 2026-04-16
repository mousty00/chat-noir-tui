package main

import (
	"encoding/json"
	"io"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

const apiBase = "https://api.chatnoir.fun/api"

func fetchCatsCmd() tea.Cmd {
	return func() tea.Msg {
		body, err := get(apiBase + "/cats")
		if err != nil {
			return catsMsg{err: err}
		}
		var r catsResponse
		if err := json.Unmarshal(body, &r); err != nil {
			return catsMsg{err: err}
		}
		return catsMsg{
			cats:       r.Data.Result,
			totalItems: r.Data.TotalItems,
			totalPages: r.Data.TotalPages,
		}
	}
}

func fetchCategoriesCmd() tea.Cmd {
	return func() tea.Msg {
		body, err := get(apiBase + "/cats/categories/")
		if err != nil {
			return categoriesMsg{err: err}
		}
		var r categoriesResponse
		if err := json.Unmarshal(body, &r); err != nil {
			return categoriesMsg{err: err}
		}
		return categoriesMsg{categories: r.Data}
	}
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
