package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/iso38500/iso38500-governance-sdk/domain"
)

// ApplicationRepositoryMemory is an in-memory implementation of ApplicationRepository
type ApplicationRepositoryMemory struct {
	mu           sync.RWMutex
	applications map[domain.ApplicationID]domain.Application
	portfolios   map[domain.PortfolioID][]domain.ApplicationID
}

// NewApplicationRepositoryMemory creates a new in-memory application repository
func NewApplicationRepositoryMemory() *ApplicationRepositoryMemory {
	return &ApplicationRepositoryMemory{
		applications: make(map[domain.ApplicationID]domain.Application),
		portfolios:   make(map[domain.PortfolioID][]domain.ApplicationID),
	}
}

// Save saves an application
func (r *ApplicationRepositoryMemory) Save(ctx context.Context, app domain.Application) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.applications[app.ID] = app
	return nil
}

// FindByID finds an application by ID
func (r *ApplicationRepositoryMemory) FindByID(ctx context.Context, id domain.ApplicationID) (domain.Application, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	app, exists := r.applications[id]
	if !exists {
		return domain.Application{}, errors.New("application not found")
	}
	return app, nil
}

// FindByName finds an application by name
func (r *ApplicationRepositoryMemory) FindByName(ctx context.Context, name string) (domain.Application, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, app := range r.applications {
		if app.Name == name {
			return app, nil
		}
	}
	return domain.Application{}, errors.New("application not found")
}

// FindAll finds all applications
func (r *ApplicationRepositoryMemory) FindAll(ctx context.Context) ([]domain.Application, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	apps := make([]domain.Application, 0, len(r.applications))
	for _, app := range r.applications {
		apps = append(apps, app)
	}
	return apps, nil
}

// FindByPortfolioID finds applications by portfolio ID
func (r *ApplicationRepositoryMemory) FindByPortfolioID(ctx context.Context, portfolioID domain.PortfolioID) ([]domain.Application, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	appIDs, exists := r.portfolios[portfolioID]
	if !exists {
		return []domain.Application{}, nil
	}

	apps := make([]domain.Application, 0, len(appIDs))
	for _, appID := range appIDs {
		if app, exists := r.applications[appID]; exists {
			apps = append(apps, app)
		}
	}
	return apps, nil
}

// Update updates an application
func (r *ApplicationRepositoryMemory) Update(ctx context.Context, app domain.Application) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.applications[app.ID]; !exists {
		return errors.New("application not found")
	}

	r.applications[app.ID] = app
	return nil
}

// Delete deletes an application
func (r *ApplicationRepositoryMemory) Delete(ctx context.Context, id domain.ApplicationID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.applications[id]; !exists {
		return errors.New("application not found")
	}

	delete(r.applications, id)
	return nil
}

// Exists checks if an application exists
func (r *ApplicationRepositoryMemory) Exists(ctx context.Context, id domain.ApplicationID) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.applications[id]
	return exists, nil
}
