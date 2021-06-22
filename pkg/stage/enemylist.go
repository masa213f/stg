package stage

type EnemyList struct {
	list []Enemy
}

func newEnemyList() *EnemyList {
	return &EnemyList{}
}

func (l *EnemyList) Add(e []Enemy) {
	l.list = append(l.list, e...)
}

func (l *EnemyList) Update() {
	newList := []Enemy{}
	for _, e := range l.list {
		if e.IsDisabled() {
			continue
		}
		e.Update()
		if !e.IsDisabled() {
			newList = append(newList, e)
		}
	}
}

func (l *EnemyList) Draw() {
	for _, e := range l.list {
		if e.IsDisabled() {
			continue
		}
		e.Draw()
	}
}

func (l *EnemyList) GetList() []Enemy {
	return l.list
}
