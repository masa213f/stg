package manager

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/masa213f/stg/pkg/constant"
	"github.com/masa213f/stg/pkg/scene"
	"github.com/masa213f/stg/pkg/util"
)

// Manager controls scene handlers. This implements the ebiten.Game interface.
type Manager struct {
	tick            uint64
	currentScene    int
	event           scene.Event
	screen          util.Screen
	handlers        []scene.Handler
	transitionTable map[int]map[scene.Event]transition

	// for debug
	debug           bool
	memStat         runtime.MemStats
	privTotalAlloc  uint64
	privNumGC       uint32
	debugStatusLine string
}

type transition struct {
	scene     int
	keepState bool
}

var errExit = errors.New("exit")

const (
	sceneDummy = iota
	sceneTitle
	sceneMenu
	sceneConfig
	scenePlay
	sceneGameOver
	sceneStageClear
	scenePause
	numOfScene
)

// New creates a new Manager instance.
func New(debugMode bool) *Manager {
	ebiten.SetWindowTitle(constant.WindowTitle)
	ebiten.SetWindowSize(constant.WindowWidth, constant.WindowHeight)
	screen := util.NewScreen(constant.ScreenWidth, constant.ScreenHeight, debugMode)

	mgr := &Manager{
		event:           scene.EventNone,
		screen:          screen,
		handlers:        make([]scene.Handler, numOfScene),
		transitionTable: map[int]map[scene.Event]transition{},
		debug:           debugMode,
	}
	input := util.NewCombinedInput(util.NewKeyboardInput(), util.NewGamepadInput())
	ctrl := util.NewControl(input)

	scenes := []*struct {
		id int
		h  scene.Handler
	}{
		{sceneDummy, scene.NewDummy()},
		{sceneTitle, scene.NewTitle(screen, ctrl)},
		{sceneMenu, scene.NewMenu(screen, ctrl, audio)},
		{sceneConfig, scene.NewConfig(screen, ctrl)},
		{scenePlay, scene.NewPlay(screen, ctrl, audio)},
		{sceneGameOver, scene.NewGameOver(screen, ctrl)},
		{sceneStageClear, scene.NewStageClear(screen, ctrl)},
		{scenePause, scene.NewPause(screen, ctrl)},
	}
	for _, s := range scenes {
		mgr.handlers[s.id] = s.h
		mgr.transitionTable[s.id] = map[scene.Event]transition{}
	}

	transitions := []*struct {
		from      int
		to        int
		event     scene.Event
		keepState bool
	}{
		{from: sceneTitle, to: sceneMenu, event: scene.EventNext},
		{from: sceneMenu, to: scenePlay, event: scene.MenuEventPlay},
		{from: sceneMenu, to: sceneConfig, event: scene.MenuEventConfig},
		{from: sceneMenu, to: sceneDummy, event: scene.EventExit},
		{from: sceneConfig, to: sceneMenu, event: scene.EventBack},
		{from: scenePlay, to: scenePause, event: scene.GameEventPause},
		{from: scenePlay, to: sceneGameOver, event: scene.GameEventGameOver},
		{from: scenePlay, to: sceneStageClear, event: scene.GameEventStageClear},
		{from: scenePause, to: scenePlay, event: scene.EventBack, keepState: true},
		{from: scenePause, to: sceneGameOver, event: scene.GameEventRetire},
		{from: sceneGameOver, to: sceneMenu, event: scene.EventNext},
		{from: sceneStageClear, to: sceneMenu, event: scene.EventNext},
	}
	for _, tr := range transitions {
		mgr.transitionTable[tr.from][tr.event] = transition{scene: tr.to, keepState: tr.keepState}
	}
	mgr.currentScene = sceneTitle
	return mgr
}

func (m *Manager) RunGame() error {
	err := ebiten.RunGame(m)
	if err == errExit {
		return nil
	}
	return err
}

// Layout returns screen size.
func (m *Manager) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constant.ScreenWidth, constant.ScreenHeight
}

// Update executes a Update function of the current scene.
func (m *Manager) Update() error {
	m.tick++
	if m.debug && m.tick%60 == 0 {
		runtime.ReadMemStats(&m.memStat)
		m.debugStatusLine = fmt.Sprintf("FPS: %0.1f\nTPS: %0.1f\n", ebiten.CurrentFPS(), ebiten.CurrentTPS())
		m.debugStatusLine += fmt.Sprintf("Heap: %2d (inuse: %2d, idle: %2d)\n", bToMib(m.memStat.HeapSys), bToMib(m.memStat.HeapInuse), bToMib(m.memStat.HeapIdle))
		m.debugStatusLine += fmt.Sprintf("Sys: %2d (stack: %2d, heap: %2d)\n", bToMib(m.memStat.Sys), bToMib(m.memStat.StackSys), bToMib(m.memStat.HeapIdle))
		m.debugStatusLine += fmt.Sprintf("Alloc/Sec: %d, NumGC/Sec: %d", bToMib(m.memStat.TotalAlloc-m.privTotalAlloc), m.memStat.NumGC-m.privNumGC)
		m.privTotalAlloc = m.memStat.TotalAlloc
		m.privNumGC = m.memStat.NumGC
	}

	if m.event != scene.EventNone {
		next, ok := m.transitionTable[m.currentScene][m.event]
		if !ok {
			return fmt.Errorf("invalid transition; [scene %d, event %d] is not registered", m.currentScene, m.event)
		}
		if !next.keepState {
			m.handlers[next.scene].Reset()
		}
		m.currentScene = next.scene
	}
	m.event = m.handlers[m.currentScene].Update()
	if m.event == scene.EventExit {
		return errExit
	}
	return nil
}

func bToMib(b uint64) uint64 {
	return b / 1024 / 1024
}

// Draw executes a Draw function of the current scene.
func (m *Manager) Draw(screen *ebiten.Image) {
	m.screen.SetImage(screen)
	m.handlers[m.currentScene].Draw()
	ebitenutil.DebugPrint(screen, m.debugStatusLine)
}
