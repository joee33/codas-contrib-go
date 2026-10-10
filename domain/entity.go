package domain

type Entity[T any] interface {
	MetaCarrier
	ID() T
	SetID(id T)
	IsPending() bool
	IsDeleted() bool
	MarkAsDeleted()
	MarkAsPending()
}

type entity[T any] struct {
	*metaCarrier
	id      T
	pending bool
	deleted bool
}

func NewEntity[T any](id T, meta ...*Meta) Entity[T] {
	return &entity[T]{NewMetaCarrier(meta...), id, true, false}
}

func NewEntityNoID[T any](meta ...*Meta) Entity[T] {
	var zero T
	return &entity[T]{NewMetaCarrier(meta...), zero, true, false}
}

func RestoreEntity[T any](id T, meta ...*Meta) Entity[T] {
	return &entity[T]{NewMetaCarrier(meta...), id, false, false}
}

func (e *entity[T]) ID() T           { return e.id }
func (e *entity[T]) SetID(id T)      { e.id = id }
func (e *entity[T]) IsPending() bool { return e.pending }
func (e *entity[T]) IsDeleted() bool { return e.deleted }
func (e *entity[T]) MarkAsDeleted()  { e.pending = true; e.deleted = true }
func (e *entity[T]) MarkAsPending()  { e.pending = true }
