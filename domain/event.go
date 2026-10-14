package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type contextKey struct{}

var _eventPayloadKey = contextKey{}

type DomainEvent[T any, S any] interface {
	Ctx() context.Context
	EventID() string
	Type() T
	String() string
	OccurredAt() time.Time
	AggregateRootID() S
	Payload() any
}

// BaseDomainEvent 领域事件
type domainEvent[T any, S any] struct {
	context         context.Context
	id              string
	aggregateRootID S
	eventType       T
	occurredAt      time.Time
}

// NewBaseDomainEvent 创建新的基础领域事件
func NewDomainEvent[T any, S any](ctx context.Context, eventType T, aggregateRootID S, payload any) DomainEvent[T, S] {
	if payload != nil {
		ctx = context.WithValue(ctx, _eventPayloadKey, payload)
	}

	return &domainEvent[T, S]{
		context:         ctx,
		id:              uuid.New().String(),
		aggregateRootID: aggregateRootID,
		eventType:       eventType,
		occurredAt:      time.Now(),
	}
}

func (e *domainEvent[T, S]) String() string {
	return fmt.Sprintf("DomainEvent[%v][%s]:AggregateRootID=%v:OccurredAt=%s", e.eventType, e.id, e.aggregateRootID, e.occurredAt.Format(time.RFC3339))
}

func (e *domainEvent[T, S]) Ctx() context.Context {
	return e.context
}

// EventID 返回事件ID
func (e *domainEvent[T, S]) EventID() string {
	return e.id
}

func (e *domainEvent[T, S]) Type() T {
	return e.eventType
}

// OccurredAt 返回事件发生时间
func (e *domainEvent[T, S]) OccurredAt() time.Time {
	return e.occurredAt
}

// AggregateID 返回聚合根ID
func (e *domainEvent[T, S]) AggregateRootID() S {
	return e.aggregateRootID
}

func (e *domainEvent[T, S]) Payload() any {
	return e.context.Value(_eventPayloadKey)
}
