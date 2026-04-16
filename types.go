package main

type viewState int

const (
	viewMenu       viewState = iota
	viewCats
	viewCategories
	viewChat
)

func (v viewState) String() string {
	switch v {
	case viewMenu:
		return "Menu"
	case viewCats:
		return "Cats"
	case viewCategories:
		return "Categories"
	case viewChat:
		return "Chat"
	default:
		return ""
	}
}

type menuEntry struct {
	label string
	desc  string
	view  viewState
}

var menuEntries = []menuEntry{
	{"Cats", "Browse cats from the API", viewCats},
	{"Categories", "Browse cat categories", viewCategories},
	{"Chat", "Open chat interface", viewChat},
}

type chatMessage struct {
	role    string
	content string
}
