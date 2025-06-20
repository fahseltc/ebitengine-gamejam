package scene

import (
	"gamejam/fonts"
	"gamejam/ui"
	"gamejam/util"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type MenuScene struct {
	BaseScene
	startBtn *ui.Button
	bg       *ebiten.Image
	txt      string
	fonts    *fonts.All
}

func NewMenuScene(fonts *fonts.All) *MenuScene {
	scene := &MenuScene{
		bg:    util.LoadImage("ui/menu-bg.png"),
		txt:   "ANTony & CleopatROACH",
		fonts: fonts,
	}
	scene.startBtn = ui.NewButton(fonts.Med, ui.WithText("START"), ui.WithRect(image.Rectangle{
		Min: image.Point{X: 250, Y: 520},
		Max: image.Point{X: 550, Y: 570},
	}), ui.WithClickFunc(func() {
		scene.BaseScene.sm.SwitchTo(NewPlayScene(scene.fonts))
	}))
	return scene
}

func (s *MenuScene) Update() error {
	s.startBtn.Update()
	return nil
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.bg, nil)
	util.DrawCenteredText(screen, s.fonts.Med, s.txt, 400, 50, nil)

	s.startBtn.Draw(screen)
}
