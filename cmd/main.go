package main

import (
	g "github.com/AllenDang/giu"
	"grouper/common/aui"
)

// η¨εΊε₯ε£
func main() {
	// GUI  // g.MasterWindowFlagsNotResizable  // MasterWindowFlagsMaximized
	wnd := g.NewMasterWindow("Grouper π", 730, 600, g.MasterWindowFlagsNotResizable)
	// wnd.SetDropCallback(onDrop)
	wnd.Run(aui.Loop)
}
