package game

import (
	"gamejam/config"
	"gamejam/fonts"
	"gamejam/log"
	"gamejam/scene"
	"log/slog"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joelschutz/stagehand"
)

var fontPath = "fonts/PressStart2P-Regular.ttf"

type Game struct {
	LastUpdateTime time.Time
	sceneManager   *stagehand.SceneManager[scene.GameState]
	fonts          *fonts.All

	cfg *config.T
	log *slog.Logger
}

func New(cfg *config.T) *Game {
	state := scene.GameState{}
	fonts := fonts.Load(fontPath)
	var manager *stagehand.SceneManager[scene.GameState]
	if cfg.SkipMenu {
		sceneInstance := scene.NewPlayScene(fonts)
		manager = stagehand.NewSceneManager(sceneInstance, state)
	} else {
		sceneInstance := scene.NewMenuScene(fonts)
		manager = stagehand.NewSceneManager(sceneInstance, state)
	}

	return &Game{
		sceneManager: manager,
		fonts:        fonts,
		cfg:          cfg,
		log:          log.NewLogger().With("for", "game"),
	}
}

func (g *Game) Update() error {
	// Pt1: Calculate DT
	if g.LastUpdateTime.IsZero() {
		g.LastUpdateTime = time.Now()
	}
	//dt := time.Since(g.LastUpdateTime).Seconds()

	//
	// call game object updates here
	//
	g.sceneManager.Update()

	// Pt2: Calculate DT for next loop
	g.LastUpdateTime = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//g.tileMap.Draw(screen)
	g.sceneManager.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return g.cfg.Resolutions.Internal.Width, g.cfg.Resolutions.Internal.Height
}
