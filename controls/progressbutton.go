package main

import (
    "gform"
)

type ProgressButton struct {
    gform.CustomControl

    bgColor, foreColor, highlightColor, tempColor gform.Color
}

func NewProgressButton(parent gform.Controller, bgColor, foreColor gform.Color) *ProgressButton {
    pb := new(ProgressButton)
    pb.Init(parent)
    pb.bgColor, pb.foreColor, pb.highlightColor = bgColor, foreColor, gform.RGB(150, 150, 150)
    pb.tempColor = pb.bgColor

    gform.RegMsgHandler(pb)

    pb.OnPaint().Attach(progressButton_OnPaint)
    pb.OnMouseHover().Attach(progressButton_OnMouseHover)
    pb.OnMouseLeave().Attach(progressButton_OnMouseLeave)

    return pb
}

func (this *ProgressButton) BackgroundColor() *gform.Color {
    return &this.bgColor
}

func (this *ProgressButton) ForeColor() *gform.Color {
    return &this.foreColor
}

func progressButton_OnMouseHover(arg *gform.EventArg) {
    if pb, ok := arg.Sender().(*ProgressButton); ok {
        pb.bgColor = pb.highlightColor
        pb.Invalidate(false)
    }
}

func progressButton_OnMouseLeave(arg *gform.EventArg) {
    if pb, ok := arg.Sender().(*ProgressButton); ok {
        pb.bgColor = pb.tempColor
        pb.Invalidate(false)
    }
}

func progressButton_OnPaint(arg *gform.EventArg) {
    rc := arg.Sender().ClientRect()

    if pb, ok := arg.Sender().(*ProgressButton); ok {
        pen := gform.NewNullPen()
        defer pen.Dispose()
        brush := gform.NewSolidColorBrush(pb.bgColor)
        defer brush.Dispose()
        
        var data *gform.PaintEventData
        if data, ok = arg.Data().(*gform.PaintEventData); ok {
            data.Canvas.DrawRect(rc, pen, brush)
        }
    }
}