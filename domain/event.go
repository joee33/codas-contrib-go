package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type DomainEventType interface {
	String() string
	Kind() any
}

type DomainEvent[T any] interface {
	Ctx() context.Context
	EventID() string
	Type() DomainEventType
	String() string
	OccurredAt() time.Time
	AggregateRootID() T
	Payload() any
}

// BaseDomainEvent 领域事件
type domainEvent[T any] struct {
	ctx             context.Context
	id              string
	aggregateRootID T
	eventType       DomainEventType
	payload         any
	occurredAt      time.Time
}

// NewBaseDomainEvent 创建新的基础领域事件
func NewDomainEvent[T any](ctx context.Context, eventType DomainEventType, aggregateRootID T, payload any) DomainEvent[T] {
	return &domainEvent[T]{
		ctx:             ctx,
		id:              uuid.New().String(),
		aggregateRootID: aggregateRootID,
		eventType:       eventType,
		payload:         payload,
		occurredAt:      time.Now(),
	}
}

func (e *domainEvent[T]) String() string {
	return fmt.Sprintf("DomainEvent[%s][%s]:AggregateRootID=%v:OccurredAt=%s", e.eventType, e.id, e.aggregateRootID, e.occurredAt.Format(time.RFC3339))
}

func (e *domainEvent[T]) Ctx() context.Context {
	return e.ctx
}

// EventID 返回事件ID
func (e *domainEvent[T]) EventID() string {
	return e.id
}

func (e *domainEvent[T]) Type() DomainEventType {
	return e.eventType
}

// OccurredAt 返回事件发生时间
func (e *domainEvent[T]) OccurredAt() time.Time {
	return e.occurredAt
}

// AggregateID 返回聚合根ID
func (e *domainEvent[T]) AggregateRootID() T {
	return e.aggregateRootID
}

func (e *domainEvent[T]) Payload() any {
	return e.payload
}
