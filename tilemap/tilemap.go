package tilemap

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

// see https://pkg.go.dev/github.com/lafriks/go-tiled
type Tilemap struct {
	renderer       *render.Renderer
	tileMap        *tiled.Map
	StaticBg       *ebiten.Image
	CollisionRects []*image.Rectangle
	Tiles          map[int]*tiled.TilesetTile
}

func NewTilemap() *Tilemap {
	t, err := tiled.LoadFile("assets/tilemap/untitled.tmx") // this wont work in wasm! need to embed files but it breaks
	if err != nil {
		log.Fatalf("unable to load tmx: %v", err.Error())
	}

	r, err := render.NewRenderer(t)
	if err != nil {
		log.Fatal("unable to load tmx renderer")
	}

	r.RenderLayer(0)
	if err != nil {
		log.Fatalf("layer unsupported for rendering: %v", err.Error())
	}
	staticBg := ebiten.NewImageFromImage(r.Result)
	r.Clear()

	tilesIdMap := make(map[int]*tiled.TilesetTile)

	for _, tile := range t.Tilesets[0].Tiles {
		tilesIdMap[int(tile.ID)] = tile
	}

	var mapRects []*image.Rectangle
	for _, object := range t.ObjectGroups[0].Objects {
		rect := image.Rect(int(object.X), int(object.Y), int(object.X+object.Width), int(object.Y+object.Height))
		mapRects = append(mapRects, &rect)
	}

	tm := &Tilemap{
		renderer:       r,
		tileMap:        t,
		StaticBg:       staticBg,
		Tiles:          tilesIdMap,
		CollisionRects: mapRects,
	}

	tm.ToWorld()
	return tm
}

func (tm *Tilemap) GetMap() *tiled.Map {
	return tm.tileMap
}

func (tm *Tilemap) ToWorld() {
	//layer := tm.tileMap.Layers[0].Tiles
}

func (tm *Tilemap) Render(screen *ebiten.Image) {
	// for _, tile := range tm.tileMap.Layers[0].GetTilePosition() {
	// 	if tile != nil {
	// 		//prop := tile.Tileset.Properties.Get("anything") // nil
	// 		// tileRect := tile.Tileset.GetTileRect(tile.ID) // nil ref error ' ts.Image' is nil
	// 		// tileImage := tm.StaticBgT.SubImage(tileRect).(*ebiten.Image)
	// 		// tile.Tileset.Properties.Get("passable")
	// 		// opts := &ebiten.DrawImageOptions{}
	// 		// opts.GeoM.Translate(float64(tileRect.Min.X), float64(tileRect.Min.Y))
	// 		// screen.DrawImage(tileImage, opts)
	// 	}

	// }
}

type Tile struct {
	Type     string
	Passable bool
	//Resource *Resource
}

// // Import the library
// import (
//     etiled "github.com/bird-mtn-dev/ebitengine-tiled"
// )

// // Load the xml output from tiled during the initilization of the Scene.
// // Note that OpenTileMap will attempt to load the associated tilesets and tile images
// Tilemap = etiled.OpenTileMap("assets/tilemap/base.tmx")
// // Defines the draw parameters of the tilemap tiles
// Tilemap.Zoom = 1

// // Call Update on the Tilemap during the ebitengine Update loop
// Tilemap.Update()

// // Call Draw on the Tilemap during the ebitegine Draw loop to draw all the layers in the tilemap
// Tilemap.Draw(worldScreen)

// // This loop will draw all the Object Groups in the Tilemap.
// for idx := range Tilemap.ObjectGroups {
//     Tilemap.ObjectGroups[idx].Draw(worldScreen)
// }

// // You can draw a specific Layer by calling
// Tilemap.GetLayerByName("layer1").Draw(worldScreen)

// // You can draw a specific Object Group by calling
// Tilemap.GetObjectGroupByName("ojbect group 1").Draw(worldScreen)
