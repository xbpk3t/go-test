package timewheel

type timeWheelSlot struct {
	id    int
	tasks *timeWheelList
}

func newSlot(id int) *timeWheelSlot {
	return &timeWheelSlot{
		id:    id,
		tasks: newList(),
	}
}
