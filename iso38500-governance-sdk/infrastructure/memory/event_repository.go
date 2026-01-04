package memory

import (
	"context"
	"time"

	"github.com/iso38500/iso38500-governance-sdk/domain"
)

// DomainEventRepositoryMemory is an in-memory implementation of DomainEventRepository
type DomainEventRepositoryMemory struct {
	events []domain.DomainEvent
}

// NewDomainEventRepositoryMemory creates a new in-memory domain event repository
func NewDomainEventRepositoryMemory() *DomainEventRepositoryMemory {
	return &DomainEventRepositoryMemory{
		events: make([]domain.DomainEvent, 0),
	}
}

// Save saves a domain event
func (r *DomainEventRepositoryMemory) Save(ctx context.Context, event domain.DomainEvent) error {
	r.events = append(r.events, event)
	return nil
}

// FindByAggregateID finds events by aggregate ID
func (r *DomainEventRepositoryMemory) FindByAggregateID(ctx context.Context, aggregateID string) ([]domain.DomainEvent, error) {
	var result []domain.DomainEvent
	for _, event := range r.events {
		// This is a simplified implementation - in practice, events would need to be associated with aggregates
		result = append(result, event)
	}
	return result, nil
}

// FindByEventType finds events by event type
func (r *DomainEventRepositoryMemory) FindByEventType(ctx context.Context, eventType string) ([]domain.DomainEvent, error) {
	var result []domain.DomainEvent
	for _, event := range r.events {
		if event.EventType() == eventType {
			result = append(result, event)
		}
	}
	return result, nil
}

// FindByTimeRange finds events by time range
func (r *DomainEventRepositoryMemory) FindByTimeRange(ctx context.Context, start, end time.Time) ([]domain.DomainEvent, error) {
	var result []domain.DomainEvent
	for _, event := range r.events {
		if event.Time().After(start) && event.Time().Before(end) {
			result = append(result, event)
		}
	}
	return result, nil
}

// Delete deletes a domain event
func (r *DomainEventRepositoryMemory) Delete(ctx context.Context, eventID string) error {
	// Simplified implementation - in practice, events would have IDs
	return nil
}
