package shooting

import (
	"math/rand"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/shape"
	"github.com/masa213f/stg/resource"
)

// 雲の移動速度
const cloudSpeed = 4

// 生成可能な雲の最大数
const cloudMaxNum = 1000

// 雲の画像の数
const cloudImageMax = 3

// 雲のイメージの一覧
var cloudImage = resource.ImageCloud

// 雲のイメージのサイズ
var cloudImageSize = [cloudImageMax]struct {
	w int
	h int
}{
	{128, 48},
	{128, 48},
	{96, 48},
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

type background struct {
	clouds *cloudList
}

func newBackground() *background {
	b := &background{
		clouds: newCloudList(),
	}
	// ゲーム開始時に表示される雲を準備する
	for i := -100; i < constant.ScreenWidth; i += cloudSpeed {
		b.clouds.new(rand.Intn(cloudImageMax), i, constant.ScreenHeight-100+rand.Intn(100))
	}
	return b
}

func (b *background) update() {
	b.clouds.new(rand.Intn(cloudImageMax), constant.ScreenWidth, constant.ScreenHeight-100+rand.Intn(100))
	b.clouds.updateAll()
}

func (b *background) draw() {
	draw.ImageAt(resource.ImageBackground, -100, 0)
	b.clouds.drawAll()
}
