package domain

type Aggregate[T any, E any] interface {
	Entity[T]
	AddEvent(event DomainEvent[any, T])
	PendingEvents() []DomainEvent[any, T]
}

type aggregate[T any] struct {
	Entity[T]
	pendingEvents []DomainEvent[any, T]
}

func NewAggregate[T any](id T, meta ...*Meta) Aggregate[T] {
	return &aggregate[T]{NewEntity(id, meta...), nil}
}
func NewAggregateNoID[T any](meta ...*Meta) Aggregate[T] {
	return &aggregate[T]{NewEntityNoID[T](meta...), nil}
}
func RestoreAggregate[T any](id T, meta ...*Meta) Aggregate[T] {
	return &aggregate[T]{RestoreEntity(id, meta...), nil}
}
func (a *aggregate[T]) PendingEvents() []DomainEvent[any, T] { return a.pendingEvents }
func (a *aggregate[T]) AddEvent(event DomainEvent[any, T]) {
	a.pendingEvents = append(a.pendingEvents, event)
}
