package ui

import (
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
)

//-----------------------------------------------------------
type ControlArea struct {
    flex                *tview.Flex
    addButton           *tview.Button
    searchButton        *tview.Button
    conAllButton        *tview.Button
    disconAllButton     *tview.Button
    closeButton         *tview.Button
    focusElements       []tview.Primitive
    focusIndex          int
    parent              *UI
}
//-----------------------------------------------------------

func New_Control_Area() *ControlArea {
    flex := tview.NewFlex().
        SetDirection(tview.FlexColumn)

    addButton := tview.NewButton("Add")
    searchButton := tview.NewButton("Search")
    conAllButton := tview.NewButton("Connect All")
    disconAllButton := tview.NewButton("Disconnect All")
    closeButton := tview.NewButton("Close")

    gap := tview.NewTextView()

    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(addButton, 0, 1, false)
    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(searchButton, 0, 1, false)
    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(conAllButton, 0, 1, false)
    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(disconAllButton, 0, 1, false)
    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(closeButton, 0, 1, false)
    flex.AddItem(gap, 1, 1, false)

    focusElements := []tview.Primitive{
        addButton,
        searchButton,
        conAllButton,
        disconAllButton,
        closeButton,
    }

    controlArea := &ControlArea{
        flex:               flex,
        addButton:          addButton,
        searchButton:       searchButton,
        conAllButton:       conAllButton,
        disconAllButton:    disconAllButton,
        closeButton:        closeButton,
        focusElements:      focusElements,
        focusIndex:         0,
        parent:             nil,
    }

    controlArea.Handle_Control_Buttons()

    return controlArea
}

func (cA *ControlArea) Get_Primitive() tview.Primitive {
    return cA.flex
}

func (cA *ControlArea) Handle_Arrow_Keys(ui *UI, event *tcell.EventKey) {
    switch event.Key() {
        case tcell.KeyRight:
            cA.focusIndex = (cA.focusIndex + 1) % len(cA.focusElements)
        case tcell.KeyLeft:
            cA.focusIndex = (cA.focusIndex - 1 + len(cA.focusElements)) % len(cA.focusElements)
    }
    ui.app.SetFocus(cA.focusElements[cA.focusIndex])
}

