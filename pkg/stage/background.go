package stage

import (
	"math/rand"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/shape"
	"github.com/masa213f/stg/resource"
)

type Background interface {
	Update()
	Draw()
}

const cloudSpeed = 4

// Maximum number of clouds that can be generated
const cloudMaxNum = 1000

const cloudImageMax = 3

var cloudImage = resource.ImageCloud

var cloudImageSize = [cloudImageMax]struct {
	w int
	h int
}{
	{128, 48},
	{128, 48},
	{96, 48},
}

type backgroundImpl struct {
	clouds *cloudList
}

func newBackground() Background {
	b := &backgroundImpl{
		clouds: newCloudList(),
	}
	// Prepare the clouds that will be displayed at the beginning of the game.
	for i := -100; i < constant.ScreenWidth; i += cloudSpeed {
		b.clouds.new(rand.Intn(cloudImageMax), i, constant.ScreenHeight-100+rand.Intn(100))
	}
	return b
}

func (b *backgroundImpl) Update() {
	b.clouds.new(rand.Intn(cloudImageMax), constant.ScreenWidth, constant.ScreenHeight-100+rand.Intn(100))
	b.clouds.updateAll()
}

func (b *backgroundImpl) Draw() {
	draw.ImageAt(resource.ImageBackground, -100, 0)
	b.clouds.drawAll()
}

type cloud struct {
	id       objectID
	image    *ebiten.Image
	drawRect *shape.Rect
}

type cloudList struct {
	nextID    objectID
	activeNum int
	buffer    []*cloud
}

func newCloudList() *cloudList {
	list := &cloudList{
		buffer: make([]*cloud, cloudMaxNum),
	}
	for i := 0; i < cloudMaxNum; i++ {
		list.buffer[i] = &cloud{
			id:       inactiveObjectID,
			drawRect: &shape.Rect{},
		}
	}
	return list
}

func (list *cloudList) new(imageNum, x, y int) {
	if list.activeNum == cloudMaxNum {
		return
	}
	ent := list.buffer[list.activeNum]
	ent.id = list.nextID
	ent.image = cloudImage[imageNum]
	ent.drawRect.Reset(x, y, cloudImageSize[imageNum].w, cloudImageSize[imageNum].h)
	list.nextID++
	list.activeNum++
}

func (list *cloudList) updateAll() {
	for _, ent := range list.buffer {
		if ent.id == inactiveObjectID {
			break
		}
		if ent.drawRect.MoveX(-cloudSpeed).X1() <= 0 {
			ent.id = inactiveObjectID
			list.activeNum--
		}
	}
	sort.Slice(list.buffer, func(i, j int) bool { return list.buffer[i].id < list.buffer[j].id })
}

func (list *cloudList) drawAll() {
	for _, ent := range list.buffer {
		if ent.id == inactiveObjectID {
			break
		}
		draw.ImageAt(ent.image, ent.drawRect.X0(), ent.drawRect.Y0())
	}
}
