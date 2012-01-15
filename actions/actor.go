package main

import (
	"gform"
)

type ExecuteResult struct {
	IsSuccess bool
	FilePath string
	Description string
}

type Actor interface {
	NormalColor() gform.Color
	MouseHoverColor() gform.Color
	ClickColor() gform.Color
	ForeColor() gform.Color

	Caption() string
	Description() string

	//Do action.
	Do(files []string) []ExecuteResult

	//Show configuration form.
	DoConfig() *gform.Form

	//Get JSON string of configuration entity
	Config() string

	//Load JSON string configuration entity
	SetConfig(jsonString string)
}