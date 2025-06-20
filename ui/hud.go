package ui

import (
	"gamejam/log"
	"gamejam/util"
	"image"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type HUD struct {
	bg        *ebiten.Image
	rect      image.Rectangle
	attackBtn *Button
	//attackLabel *ebiten.Image
	moveBtn *Button
	stopBtn *Button
	log     *slog.Logger
}

func NewHUD(font text.Face) *HUD {
	c := &HUD{
		rect: image.Rectangle{Min: image.Pt(0, 450), Max: image.Pt(300, 600)},
		//attackLabel: util.LoadImage("ui/keys/z.png"),
		log: log.NewLogger().With("for", "HUD"),
	}
	c.bg = util.ScaleImage(util.LoadImage("ui/btn/controls-bg.png"), float32(c.rect.Dx()), float32(c.rect.Dy()))

	c.attackBtn = NewButton(font,
		WithRect(image.Rectangle{Min: image.Pt(c.rect.Min.X+20, c.rect.Min.Y+20), Max: image.Pt(c.rect.Min.X+70, c.rect.Min.Y+70)}),
		WithClickFunc(func() {
			c.log.Info("atkbtnclicked")
		}),
		WithImage(util.LoadImage("ui/btn/atk-btn.png"), util.LoadImage("ui/btn/atk-btn-pressed.png")),
		WithKeyActivation(ebiten.KeyZ),
	)
	c.moveBtn = NewButton(font,
		WithRect(image.Rectangle{Min: image.Pt(c.rect.Min.X+80, c.rect.Min.Y+20), Max: image.Pt(c.rect.Min.X+130, c.rect.Min.Y+70)}),
		WithImage(util.LoadImage("ui/btn/move-btn.png"), util.LoadImage("ui/btn/move-btn-pressed.png")),
		WithClickFunc(func() {
			c.log.Info("movebtnclicked")
		}),
		WithKeyActivation(ebiten.KeyX),
	)
	c.stopBtn = NewButton(font,
		WithRect(image.Rectangle{Min: image.Pt(c.rect.Min.X+140, c.rect.Min.Y+20), Max: image.Pt(c.rect.Min.X+190, c.rect.Min.Y+70)}),
		WithImage(util.LoadImage("ui/btn/stop-btn.png"), util.LoadImage("ui/btn/stop-btn-pressed.png")),
		WithClickFunc(func() {
			c.log.Info("stopbtnclicked")
		}),
		WithKeyActivation(ebiten.KeyC),
	)
	return c
}

func (c *HUD) Update() {
	c.attackBtn.Update()
	c.stopBtn.Update()
	c.moveBtn.Update()
}

func (c *HUD) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(c.rect.Min.X), float64(c.rect.Min.Y))
	screen.DrawImage(c.bg, opts)
	c.attackBtn.Draw(screen)
	c.stopBtn.Draw(screen)
	c.moveBtn.Draw(screen)
}
