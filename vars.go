package main

import (
	"gform"
	"w32"
	"syscall"
)

const (
    IDI_LINK = 100
    IDR_CLOSE = 103
    IDR_CLOSE_MOUSEHOVER = 113
    IDR_CLOSE_CLICK = 106
    IDR_DROPARROW = 108
    IDR_PIN = 110
    IDR_PIN_CHECKED = 112
    IDR_LOAD = 118
	IDR_LOAD_CLICK = 120
	IDR_LOAD_MOUSEHOVER = 121
	IDR_SAVE = 123
	IDR_SAVE_CLICK = 125
	IDR_SAVE_MOUSEHOVER = 127
)

var (
	gMainformTitleBarColor = gform.RGB(139, 190, 37)
	gMainformBkColor = gform.RGB(255, 255, 255)

	gResPNG = syscall.StringToUTF16Ptr("PNG")
	
	gResCloseNormal = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_CLOSE), gMainformTitleBarColor}
	gResCloseMouseOver = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_CLOSE_MOUSEHOVER), gMainformTitleBarColor}
	gResCloseClick = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_CLOSE_CLICK), gMainformTitleBarColor}

	gResPinNormal = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_PIN), gMainformTitleBarColor}
	gResPinChecked = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_PIN_CHECKED), gMainformTitleBarColor}

	gResLoadNormal = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_LOAD), gMainformBkColor}
	gResLoadMouseOver = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_LOAD_MOUSEHOVER), gMainformBkColor}
	gResLoadClick = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_LOAD_CLICK), gMainformBkColor}

	gResSaveNormal = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_SAVE), gMainformBkColor}
	gResSaveMouseOver = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_SAVE_MOUSEHOVER), gMainformBkColor}
	gResSaveClick = ResBmpInfo{gResPNG, w32.MakeIntResource(IDR_SAVE_CLICK), gMainformBkColor}
)