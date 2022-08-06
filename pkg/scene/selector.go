package scene

type item struct {
	text  string
	value Event
}

type itemSelector struct {
	cursor int
	texts  []string
	values []Event
}

func newItemSelector(items []item) *itemSelector {
	num := len(items)
	selector := &itemSelector{
		texts:  make([]string, num),
		values: make([]Event, num),
	}
	for i, it := range items {
		selector.texts[i] = it.text
		selector.values[i] = it.value
	}
	return selector
}

func (i *itemSelector) getIndex() int {
	return i.cursor
}

func (i *itemSelector) getValue() Event {
	return i.values[i.cursor]
}

func (i *itemSelector) getTexts() []string {
	return i.texts
}

func (i *itemSelector) next() {
	i.cursor = (i.cursor + 1) % len(i.values)
}

func (i *itemSelector) priv() {
	i.cursor = (i.cursor - 1 + len(i.values)) % len(i.values)
}

func (i *itemSelector) first() {
	i.cursor = 0
}

func (i *itemSelector) last() {
	i.cursor = len(i.values) - 1
}

type menuSceneHandler struct {
	items *itemSelector
}
