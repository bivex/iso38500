package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/iso38500/iso38500-governance-sdk/domain"
)

// ApplicationPortfolioRepositoryMemory is an in-memory implementation of ApplicationPortfolioRepository
type ApplicationPortfolioRepositoryMemory struct {
	mu        sync.RWMutex
	portfolios map[domain.PortfolioID]domain.ApplicationPortfolio
	byOwner   map[string][]domain.PortfolioID
}

// NewApplicationPortfolioRepositoryMemory creates a new in-memory portfolio repository
func NewApplicationPortfolioRepositoryMemory() *ApplicationPortfolioRepositoryMemory {
	return &ApplicationPortfolioRepositoryMemory{
		portfolios: make(map[domain.PortfolioID]domain.ApplicationPortfolio),
		byOwner:   make(map[string][]domain.PortfolioID),
	}
}

// Save saves an application portfolio
func (r *ApplicationPortfolioRepositoryMemory) Save(ctx context.Context, portfolio domain.ApplicationPortfolio) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.portfolios[portfolio.ID] = portfolio

	// Update owner index
	r.byOwner[portfolio.Owner] = append(r.byOwner[portfolio.Owner], portfolio.ID)
	return nil
}

// FindByID finds a portfolio by ID
func (r *ApplicationPortfolioRepositoryMemory) FindByID(ctx context.Context, id domain.PortfolioID) (domain.ApplicationPortfolio, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	portfolio, exists := r.portfolios[id]
	if !exists {
		return domain.ApplicationPortfolio{}, errors.New("portfolio not found")
	}
	return portfolio, nil
}

// FindByOwner finds portfolios by owner
func (r *ApplicationPortfolioRepositoryMemory) FindByOwner(ctx context.Context, owner string) ([]domain.ApplicationPortfolio, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	portfolioIDs, exists := r.byOwner[owner]
	if !exists {
		return []domain.ApplicationPortfolio{}, nil
	}

	portfolios := make([]domain.ApplicationPortfolio, 0, len(portfolioIDs))
	for _, id := range portfolioIDs {
		if portfolio, exists := r.portfolios[id]; exists {
			portfolios = append(portfolios, portfolio)
		}
	}
	return portfolios, nil
}

// FindAll finds all portfolios
func (r *ApplicationPortfolioRepositoryMemory) FindAll(ctx context.Context) ([]domain.ApplicationPortfolio, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	portfolios := make([]domain.ApplicationPortfolio, 0, len(r.portfolios))
	for _, portfolio := range r.portfolios {
		portfolios = append(portfolios, portfolio)
	}
	return portfolios, nil
}

// Update updates a portfolio
func (r *ApplicationPortfolioRepositoryMemory) Update(ctx context.Context, portfolio domain.ApplicationPortfolio) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.portfolios[portfolio.ID]; !exists {
		return errors.New("portfolio not found")
	}

	r.portfolios[portfolio.ID] = portfolio
	return nil
}

// Delete deletes a portfolio
func (r *ApplicationPortfolioRepositoryMemory) Delete(ctx context.Context, id domain.PortfolioID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	portfolio, exists := r.portfolios[id]
	if !exists {
		return errors.New("portfolio not found")
	}

	delete(r.portfolios, id)

	// Remove from owner index
	ownerPortfolios := r.byOwner[portfolio.Owner]
	for i, pid := range ownerPortfolios {
		if pid == id {
			r.byOwner[portfolio.Owner] = append(ownerPortfolios[:i], ownerPortfolios[i+1:]...)
			break
		}
	}

	return nil
}

// Exists checks if a portfolio exists
func (r *ApplicationPortfolioRepositoryMemory) Exists(ctx context.Context, id domain.PortfolioID) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.portfolios[id]
	return exists, nil
}

// AddApplication adds an application to a portfolio
func (r *ApplicationPortfolioRepositoryMemory) AddApplication(ctx context.Context, portfolioID domain.PortfolioID, appID domain.ApplicationID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	portfolio, exists := r.portfolios[portfolioID]
	if !exists {
		return errors.New("portfolio not found")
	}

	// Check if application is already in portfolio
	for _, app := range portfolio.Applications {
		if app.ID == appID {
			return errors.New("application already in portfolio")
		}
	}

	// Note: In a real implementation, we'd fetch the application from the application repository
	// For this memory implementation, we'll create a placeholder
	placeholderApp := domain.Application{ID: appID}
	portfolio.Applications = append(portfolio.Applications, placeholderApp)
	r.portfolios[portfolioID] = portfolio

	return nil
}

// RemoveApplication removes an application from a portfolio
func (r *ApplicationPortfolioRepositoryMemory) RemoveApplication(ctx context.Context, portfolioID domain.PortfolioID, appID domain.ApplicationID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	portfolio, exists := r.portfolios[portfolioID]
	if !exists {
		return errors.New("portfolio not found")
	}

	// Find and remove application
	for i, app := range portfolio.Applications {
		if app.ID == appID {
			portfolio.Applications = append(portfolio.Applications[:i], portfolio.Applications[i+1:]...)
			r.portfolios[portfolioID] = portfolio
			return nil
		}
	}

	return errors.New("application not found in portfolio")
}
