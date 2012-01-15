package main

import (
	"gform"
)

type ResBmpInfo struct {
	ResType *uint16
	ResName *uint16
	BgColor gform.Color
}

func newResBitmap(ri *ResBmpInfo) *gform.Bitmap {
	bmp, err := gform.NewBitmapFromResource(gform.GetAppInstance(), ri.ResName, ri.ResType, ri.BgColor)
	if err != nil {
		panic("Load bitmap from resource failed")
	}

	return bmp
}