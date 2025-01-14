package ussd

import (
	"fmt"
)

// Menu for USSD
type Menu struct {
	Items    []*menuItem
	ZeroItem *menuItem
}

type menuItem struct {
	Name  string
	Route route
}

// NewMenu creates a new Menu
func NewMenu() *Menu {
	return &Menu{
		Items: make([]*menuItem, 0),
	}
}

// Add to USSD menu.
func (m *Menu) Add(name, ctrl, action string) *Menu {
	item := &menuItem{name, route{ctrl, action}}
	m.Items = append(m.Items, item)
	return m
}

// AddZero adds an item at the bottom of USSD menu.
// This item always routes to a choice of "0".
func (m *Menu) AddZero(name, ctrl, action string) *Menu {
	m.ZeroItem = &menuItem{name, route{ctrl, action}}
	return m
}

// render USSD menu.
func (m Menu) render() string {
	msg := StrEmpty

	for _, item := range m.Items {
		msg += fmt.Sprintf("%v"+StrNewLine, item.Name)
	}
	if m.ZeroItem != nil {
		msg += "0. " + m.ZeroItem.Name + StrNewLine
	}
	return msg
}
