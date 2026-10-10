package domain

type Aggregate[T any] interface {
	Entity[T]
	AddEvent(event any)
	PendingEvents() []any
}

type aggregate[T any] struct {
	Entity[T]
	pendingEvents []any
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
func (a *aggregate[T]) PendingEvents() []any { return a.pendingEvents }
func (a *aggregate[T]) AddEvent(event any)   { a.pendingEvents = append(a.pendingEvents, event) }
