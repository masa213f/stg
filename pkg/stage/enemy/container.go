package enemy

import "github.com/masa213f/stg/pkg/shape"

type Enemy interface {
	Update()
	Draw()
	Damage(d int) (score int)
	IsDisabled() bool
	IsInvincible() bool
	GetHitRect() *shape.Rect
}

type Container struct {
	enemies []Enemy
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) UpdateAll() {
	newList := []Enemy{}
	for _, e := range c.enemies {
		if e.IsDisabled() {
			continue
		}
		e.Update()
		if !e.IsDisabled() {
			newList = append(newList, e)
		}
	}
	c.enemies = newList
}

func (c *Container) DrawAll() {
	for _, e := range c.enemies {
		if e.IsDisabled() {
			continue
		}
		e.Draw()
	}
}

func (c *Container) Add(e ...Enemy) {
	c.enemies = append(c.enemies, e...)
}

func (c *Container) Count() int {
	return len(c.enemies)
}

func (c *Container) GetList() []Enemy {
	return c.enemies
}
