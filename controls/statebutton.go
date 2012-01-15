package main

import (
	"gform"
)

type StateButton struct {
	gform.CustomControl

	resNormal, resChecked ResBmpInfo
	isChecked bool

	onStateChange gform.EventManager
}

func NewStateButton(parent gform.Controller, resNormal, resChecked ResBmpInfo) *StateButton {
	b := new(StateButton)
	b.Init(parent)

	b.resNormal = resNormal
	b.resChecked = resChecked

	b.isChecked = false

	//Load normal state bitmap to set initial size
	temp := newResBitmap(&b.resNormal)
	defer temp.Dispose()
	b.SetSize(temp.Size())

	gform.RegMsgHandler(b)

	b.OnPaint().Attach(stateButton_OnPaint)
	b.OnLBUp().Attach(stateButton_OnLBUp)
	
	return b
}

func (this *StateButton) OnStateChange() *gform.EventManager {
	return &this.onStateChange
}

func (this *StateButton) Checked() bool {
	return this.isChecked
}

func (this *StateButton) SetChecked(b bool) {
	if this.isChecked != b {
		this.isChecked = b
		this.Invalidate(false)

		this.onStateChange.Fire(gform.NewEventArg(this, nil))
	}
}

func stateButton_OnPaint(arg *gform.EventArg) {
	if b, ok := arg.Sender().(*StateButton); ok {
		var data *gform.PaintEventData
		if data, ok = arg.Data().(*gform.PaintEventData); ok {
			var bmp *gform.Bitmap
			if b.isChecked {
				bmp = newResBitmap(&b.resChecked)
			} else {
				bmp = newResBitmap(&b.resNormal)
			}
			defer bmp.Dispose()

			data.Canvas.DrawBitmap(bmp, 0, 0)
		}
	}
}

func stateButton_OnLBUp(arg *gform.EventArg) {
	if b, ok := arg.Sender().(*StateButton); ok {
		b.SetChecked(!b.Checked())
	}
}