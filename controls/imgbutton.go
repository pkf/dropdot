package main

import (
	"gform"
)

type ImgButtonState byte

const (
	IBSNormal ImgButtonState = 1 << iota
	IBSMouseHover
	IBSClick
)

type ImgButton struct {
	gform.CustomControl

	resNormal, resMouseOver, resClick ResBmpInfo
	state ImgButtonState
}

func NewImgButton(parent gform.Controller, resNormal, resMouseOver, resClick ResBmpInfo) *ImgButton {
	b := new(ImgButton)
	b.Init(parent)

	b.resNormal = resNormal
	b.resMouseOver = resMouseOver
	b.resClick = resClick

	b.state = IBSNormal

	//Load normal state bitmap to set initial size
	temp := newResBitmap(&b.resNormal)
	defer temp.Dispose()
	b.SetSize(temp.Size())

	gform.RegMsgHandler(b)

	b.OnPaint().Attach(imgButton_OnPaint)
	b.OnMouseHover().Attach(imgButton_OnMouseHover)
	b.OnMouseLeave().Attach(imgButton_OnMouseLeave)
	b.OnLBDown().Attach(imgButton_OnLBDown)
	b.OnLBUp().Attach(imgButton_OnLBUp)

	return b
}

func imgButton_OnPaint(arg *gform.EventArg) {
	if data, ok := arg.Data().(*gform.PaintEventData); ok {
		var b *ImgButton
		if b, ok = arg.Sender().(*ImgButton); ok {
			var bmp *gform.Bitmap
			switch b.state {
				case IBSNormal:
					bmp = newResBitmap(&b.resNormal)
				case IBSMouseHover:
					bmp = newResBitmap(&b.resMouseOver)
				case IBSClick:
					bmp = newResBitmap(&b.resClick)
			}
			if bmp != nil {
				defer bmp.Dispose()

				if b.Height() != bmp.Height() || b.Width() != bmp.Width() {
					b.SetSize(bmp.Size())
				}

				data.Canvas.DrawBitmap(bmp, 0, 0)
			}
		}
	}
}

func imgButton_OnMouseHover(arg *gform.EventArg) {
	if b, ok := arg.Sender().(*ImgButton); ok {
		b.state = IBSMouseHover
		b.Invalidate(false)
	}
}

func imgButton_OnMouseLeave(arg *gform.EventArg) {
	if b, ok := arg.Sender().(*ImgButton); ok {
		b.state = IBSNormal
		b.Invalidate(false)
	}
}

func imgButton_OnLBDown(arg *gform.EventArg) {
	if b, ok := arg.Sender().(*ImgButton); ok {
		b.state = IBSClick
		b.Invalidate(false)
	}
}

func imgButton_OnLBUp(arg *gform.EventArg) {
	if b, ok := arg.Sender().(*ImgButton); ok {
		b.state = IBSMouseHover
		b.Invalidate(false)
	}
}