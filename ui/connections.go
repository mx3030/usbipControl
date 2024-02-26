package ui

import (
	"github.com/rivo/tview"
    "time"
)

//-----------------------------------------------------------
type ConnectionsArea struct {
    flex            *tview.Flex
    connections     map[int64]*Connection
}
//-----------------------------------------------------------

func New_Connections_Area() *ConnectionsArea {
    flex := tview.NewFlex()
    flex.SetDirection(tview.FlexRow).SetBorder(true)
    connections := make(map[int64]*Connection)
    return &ConnectionsArea{
        flex:           flex,
        connections:    connections,
    }
}

func (conA *ConnectionsArea) Add_Connection() {
    // init connection element
    c := New_Connection()
    c.id = time.Now().Unix()
    c.parent = conA
    // add connection to connectionsArea
    conA.connections[c.id] = c
    gap := tview.NewTextView()
    conA.flex.AddItem(gap, 1, 1, false)
    conA.flex.AddItem(c.flex, 1, 1, false)
}

func (conA *ConnectionsArea) Remove_Connection(c *Connection) {
    delete(conA.connections, c.id)
    conA.flex.RemoveItem(c.flex)
}
