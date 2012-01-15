package main

import (
    "gform"
)

var (
    mainform *Mainform
)

func main() {
    gform.Init()

    mainform = NewMainform(nil)
    mainform.SetCaption("Dropdot")

    mainform.Center()
    mainform.Show()

    gform.RunMainLoop()
}