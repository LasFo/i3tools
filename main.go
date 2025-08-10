package main

import (
	"fmt"
	"log"

	"go.i3wm.org/i3"
)

func main() {
	// Create a new workspace on the second contain and move the currently
	// focused container there.

	wss, err := i3.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
	}

	var nextWS int64
	for _, ws := range wss {
		nextWS = max(nextWS, ws.Num)
	}
	nextWS++

	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	focusedCon := tree.Root.FindFocused(func(n *i3.Node) bool {
		return n.Focused
	})
	// Hard coding the output for now because it always only want to move to
	// my secondary monitor. Never to my primary monitor.
	i3.RunCommand(fmt.Sprintf(`workspace %[1]d; move workspace to output HDMI-0; [con_id="%[2]d"] focus; move container to workspace number %[1]d`, nextWS, focusedCon.ID))
}
