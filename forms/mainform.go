package main

import (
    "w32"
    "gform"
)

type Mainform struct {
    gform.Form
}

func NewMainform(parent gform.Controller) *Mainform {
    mf := Mainform{*gform.NewForm(parent)}

    mf.EnableDragMove(true)
    mf.SetSize(200, 200)
    mf.EnableSizable(false)
    mf.EnableDragAcceptFiles(true)
    gform.ToggleStyle(mf.Handle(), false, w32.WS_CAPTION)
    gform.ToggleExStyle(mf.Handle(), false, w32.WS_EX_CLIENTEDGE)

    gform.RegMsgHandler(&mf)

    mf.OnPaint().Attach(mainform_OnPaint)

    btnClose := NewImgButton(&mf, gResCloseNormal, gResCloseMouseOver, gResCloseClick)
    btnClose.SetPos(mf.Width()-btnClose.Width()-3, 3)
    btnClose.OnLBUp().Attach(btnClose_OnLBUp)

    x, y := btnClose.Pos()
    btnPin := NewStateButton(&mf, gResPinNormal, gResPinChecked)
    btnPin.SetPos(x-btnPin.Width()-3, y)
    btnPin.OnStateChange().Attach(btnPin_OnStateChange)

    btnLoad := NewImgButton(&mf, gResLoadNormal, gResLoadMouseOver, gResLoadClick)
    btnLoad.SetPos(mf.Width()-btnLoad.Width()-3, 130)

    x, y = btnLoad.Pos()
    btnSave := NewImgButton(&mf, gResSaveNormal, gResSaveMouseOver, gResSaveClick)
    btnSave.SetPos(x-btnSave.Width()-3, y)

    return &mf
}

func mainform_OnPaint(arg *gform.EventArg) {
    var mf *Mainform
    var ok bool
    var data *gform.PaintEventData

    if mf, ok = arg.Sender().(*Mainform); ok {
        if data, ok = arg.Data().(*gform.PaintEventData); ok {
            gResMainformBkColor := gform.RGB(255, 255, 255)
            bkBrush := gform.NewSolidColorBrush(gResMainformBkColor)
            defer bkBrush.Dispose()
            borderBrush := gform.NewSolidColorBrush(gform.RGB(154, 154, 154))
            defer borderBrush.Dispose()
            borderPen := gform.NewPen(w32.PS_COSMETIC|w32.PS_SOLID, 1, borderBrush)
            defer borderPen.Dispose()
            data.Canvas.DrawRect(mf.ClientRect(), borderPen, bkBrush)

            // Draw title bar
            w := mf.Width()
            titleRect := gform.NewRect(1, 1, w-1, 22)
            titleBrush := gform.NewSolidColorBrush(gMainformTitleBarColor)
            defer titleBrush.Dispose()
            data.Canvas.FillRect(titleRect, titleBrush)
            
            // Draw title text
            f := gform.NewFont("Bauhaus 93", 9, 0)
            defer f.Dispose()
            titleRect.Inflate(-5, 0)
            data.Canvas.DrawText(mf.Caption(), titleRect, w32.DT_LEFT|w32.DT_VCENTER|w32.DT_SINGLELINE, f, gResMainformBkColor)

            // Draw "Drop Here"
            f = gform.NewFont("Bauhaus 93", 25, gform.FontBold)
            defer f.Dispose()
            bodyRect := gform.NewRect(1, 50, w-1, 120)
            data.Canvas.DrawText("Drop Here", bodyRect, w32.DT_CENTER|w32.DT_VCENTER|w32.DT_SINGLELINE, f, gform.RGB(187, 187, 187))

            // Draw drop arrow
            if bmp, err := gform.NewBitmapFromResource(gform.GetAppInstance(), w32.MakeIntResource(IDR_DROPARROW), gResPNG, gResMainformBkColor); err == nil {
                data.Canvas.DrawBitmap(bmp, (w-bmp.Width())/2, 110)
                bmp.Dispose()
            }

            // Draw bottom panel
            h := mf.Height()
            bottomRect := gform.NewRect(1, 150, w-1, h-1)
            data.Canvas.FillRect(bottomRect, titleBrush)
        }
    }
}

func btnClose_OnLBUp(arg *gform.EventArg) {
    gform.Exit()
}

func btnPin_OnStateChange(arg *gform.EventArg) {
    if b, ok := arg.Sender().(*StateButton); ok {
        mainform.EnableTopMost(b.Checked())
    }
}