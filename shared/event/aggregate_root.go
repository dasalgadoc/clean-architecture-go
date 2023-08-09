package event

type AggregateRoot struct {
	events []Event
}

func (a *AggregateRoot) Record(evt Event) {
	a.events = append(a.events, evt)
}

func (a *AggregateRoot) PullEvents() []Event {
	events := a.events
	a.events = []Event{}

	return events
}
