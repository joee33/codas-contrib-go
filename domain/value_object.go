package domain

type ValueObject interface {
	MetaCarrier
	IsPending() bool
	IsDeleted() bool
	MarkAsDeleted()
}

type valueObject struct {
	*metaCarrier
	pending bool
	deleted bool
}

func NewValueObject(meta ...*Meta) ValueObject {
	return &valueObject{NewMetaCarrier(meta...), true, false}
}

func RestoreValueObject(meta ...*Meta) ValueObject {
	return &valueObject{NewMetaCarrier(meta...), false, false}
}

func (e *valueObject) IsPending() bool { return e.pending }
func (e *valueObject) IsDeleted() bool { return e.deleted }
func (e *valueObject) MarkAsDeleted()  { e.pending = true; e.deleted = true }
