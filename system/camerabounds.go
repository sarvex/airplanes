package system

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"

	"github.com/m110/airplanes/component"
)

type CameraBounds struct {
	query *query.Query
}

func NewCameraBounds() *CameraBounds {
	return &CameraBounds{
		query: query.NewQuery(filter.Contains(
			component.CameraTag,
			component.Position,
		)),
	}
}

func (b *CameraBounds) Update(w donburi.World) {
	b.query.EachEntity(w, func(entry *donburi.Entry) {
		position := component.GetPosition(entry)
		if position.X < 0 {
			position.X = 0
		}

		if position.Y < 0 {
			position.Y = 0
		}
	})
}
