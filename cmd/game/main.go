package main

import (
	"flag"
	"fmt"
	_ "image/png"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/debug"
	"github.com/masa213f/stg/pkg/draw"
	"github.com/masa213f/stg/pkg/scene"
)

var version string

var (
	debugOpt   bool
	versionOpt bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.BoolVar(&debugOpt, "debug", false, "show debug print")
	flag.BoolVar(&versionOpt, "version", false, "show version")
}

func main() {
	flag.Parse()
	if versionOpt {
		fmt.Println(version)
		os.Exit(0)
	}
	debug.SetMode(debugOpt)

	ebiten.SetWindowTitle(constant.WindowTitle)
	ebiten.SetWindowSize(constant.WindowWidth, constant.WindowHeight)
	draw.SetScreenSize(constant.ScreenWidth, constant.ScreenHeight)
	g := &Game{
		debug:   debugOpt,
		manager: scene.NewManager(),
	}
	if err := ebiten.RunGame(g); err != nil && err != scene.Exit {
		panic(err)
	}
}

// Game implements the ebiten.Game interface.
type Game struct {
	tick           uint64
	memStat        runtime.MemStats
	privTotalAlloc uint64
	privNumGC      uint32
	statusLine     string
	debug          bool
	manager        *scene.Manager
}

// Layout returns screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constant.ScreenWidth, constant.ScreenHeight
}

// Update updates a game by one tick.
func (g *Game) Update() error {
	g.tick++
	if debug.GetMode() && g.tick%60 == 0 {
		runtime.ReadMemStats(&g.memStat)
		g.statusLine = fmt.Sprintf("FPS: %0.1f\nTPS: %0.1f\n", ebiten.CurrentFPS(), ebiten.CurrentTPS())
		g.statusLine += fmt.Sprintf("Heap: %2d (inuse: %2d, idle: %2d)\n", bToMib(g.memStat.HeapSys), bToMib(g.memStat.HeapInuse), bToMib(g.memStat.HeapIdle))
		g.statusLine += fmt.Sprintf("Sys: %2d (stack: %2d, heap: %2d)\n", bToMib(g.memStat.Sys), bToMib(g.memStat.StackSys), bToMib(g.memStat.HeapIdle))
		g.statusLine += fmt.Sprintf("Alloc/Sec: %d, NumGC/Sec: %d", bToMib(g.memStat.TotalAlloc-g.privTotalAlloc), g.memStat.NumGC-g.privNumGC)
		g.privTotalAlloc = g.memStat.TotalAlloc
		g.privNumGC = g.memStat.NumGC
	}
	return g.manager.Update()
}

func bToMib(b uint64) uint64 {
	return b / 1024 / 1024
}

// Draw draws the game screen by one frame.
func (g *Game) Draw(screen *ebiten.Image) {
	draw.SetScreen(screen)
	g.manager.Draw()
	ebitenutil.DebugPrint(screen, g.statusLine)
}
