package domain

type Meta struct {
	ID any // 持久化 id
	PO any // 持久化对象，如ORM对象
}

type MetaCarrier interface {
	Meta() *Meta
}

type metaCarrier struct {
	meta *Meta
}

func NewMetaCarrier(meta ...*Meta) *metaCarrier {
	if len(meta) == 0 || meta[0] == nil {
		return &metaCarrier{&Meta{}}
	}
	return &metaCarrier{meta[0]}
}

func (m *metaCarrier) Meta() *Meta {
	return m.meta
}
