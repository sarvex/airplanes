package system

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/airplanes/component"
)

type Render struct {
	query     *query.Query
	offscreen *ebiten.Image
}

func NewRenderer() *Render {
	return &Render{
		query: query.NewQuery(
			filter.Contains(component.Position, component.Sprite),
		),
		// TODO figure out the proper size
		offscreen: ebiten.NewImage(1000, 1000),
	}
}

func (r *Render) Draw(w donburi.World, screen *ebiten.Image) {
	camera, ok := query.NewQuery(filter.Contains(component.CameraTag)).FirstEntity(w)
	if !ok {
		panic("no camera found")
	}
	cameraPos := component.GetPosition(camera)

	r.offscreen.Clear()

	var entries []*donburi.Entry
	r.query.EachEntity(w, func(entry *donburi.Entry) {
		entries = append(entries, entry)
	})

	byLayer := lo.GroupBy(entries, func(entry *donburi.Entry) int {
		return component.GetSprite(entry).Layer
	})
	layers := lo.Keys(byLayer)
	sort.Ints(layers)

	for _, layer := range layers {
		for _, entry := range byLayer[layer] {
			position := component.GetPosition(entry)
			sprite := component.GetSprite(entry)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(position.X, position.Y)
			r.offscreen.DrawImage(sprite.Image, op)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-cameraPos.X, -cameraPos.Y)
	screen.DrawImage(r.offscreen, op)
}
