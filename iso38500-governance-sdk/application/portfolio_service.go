package application

import (
	"context"
	"fmt"
	"time"

	"github.com/iso38500/iso38500-governance-sdk/domain"
)

// PortfolioService provides application services for portfolio management
type PortfolioService struct {
	portfolioRepo domain.ApplicationPortfolioRepository
	appRepo       domain.ApplicationRepository
	agreementRepo domain.GovernanceAgreementRepository
	eventRepo     domain.DomainEventRepository
}

// NewPortfolioService creates a new portfolio service
func NewPortfolioService(
	portfolioRepo domain.ApplicationPortfolioRepository,
	appRepo domain.ApplicationRepository,
	agreementRepo domain.GovernanceAgreementRepository,
	eventRepo domain.DomainEventRepository,
) *PortfolioService {
	return &PortfolioService{
		portfolioRepo: portfolioRepo,
		appRepo:       appRepo,
		agreementRepo: agreementRepo,
		eventRepo:     eventRepo,
	}
}

// CreatePortfolio creates a new application portfolio
func (s *PortfolioService) CreatePortfolio(ctx context.Context, cmd CreatePortfolioCommand) (*domain.ApplicationPortfolio, error) {
	// Create aggregate
	aggregate, err := domain.NewApplicationPortfolioAggregate(
		cmd.ID,
		cmd.Name,
		cmd.Description,
		cmd.Owner,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create portfolio aggregate: %w", err)
	}

	// Save to repository
	portfolio := aggregate.GetPortfolio()
	err = s.portfolioRepo.Save(ctx, portfolio)
	if err != nil {
		return nil, fmt.Errorf("failed to save portfolio: %w", err)
	}

	// Save domain events
	for _, event := range aggregate.GetDomainEvents() {
		err = s.eventRepo.Save(ctx, event)
		if err != nil {
			// Log error but don't fail the operation
			fmt.Printf("Failed to save domain event: %v\n", err)
		}
	}

	return &portfolio, nil
}

// AddApplicationToPortfolio adds an application to a portfolio
func (s *PortfolioService) AddApplicationToPortfolio(ctx context.Context, cmd AddApplicationToPortfolioCommand) error {
	// Verify application exists
	app, err := s.appRepo.FindByID(ctx, cmd.ApplicationID)
	if err != nil {
		return fmt.Errorf("application not found: %w", err)
	}

	// Verify governance agreement exists
	_, err = s.agreementRepo.FindByApplicationID(ctx, cmd.ApplicationID)
	if err != nil {
		return fmt.Errorf("governance agreement not found for application: %w", err)
	}

	// Get portfolio
	portfolio, err := s.portfolioRepo.FindByID(ctx, cmd.PortfolioID)
	if err != nil {
		return fmt.Errorf("portfolio not found: %w", err)
	}

	// Check if application is already in portfolio
	for _, existingApp := range portfolio.Applications {
		if existingApp.ID == cmd.ApplicationID {
			return fmt.Errorf("application already exists in portfolio")
		}
	}

	// Add application to portfolio
	portfolio.Applications = append(portfolio.Applications, app)
	portfolio.UpdatedAt = time.Now()

	err = s.portfolioRepo.Save(ctx, portfolio)
	if err != nil {
		return fmt.Errorf("failed to save updated portfolio: %w", err)
	}

	// Publish domain event
	event := domain.ApplicationAddedToPortfolioEvent{
		PortfolioID:          cmd.PortfolioID,
		ApplicationID:        cmd.ApplicationID,
		ApplicationName:      app.Name,
		GovernanceAgreementID: app.GovernanceAgreementID,
		OccurredAt:           time.Now(),
	}

	err = s.eventRepo.Save(ctx, event)
	if err != nil {
		fmt.Printf("Failed to save domain event: %v\n", err)
	}

	return nil
}

// RemoveApplicationFromPortfolio removes an application from a portfolio
func (s *PortfolioService) RemoveApplicationFromPortfolio(ctx context.Context, cmd RemoveApplicationFromPortfolioCommand) error {
	// Get portfolio
	portfolio, err := s.portfolioRepo.FindByID(ctx, cmd.PortfolioID)
	if err != nil {
		return fmt.Errorf("portfolio not found: %w", err)
	}

	// Find and remove application
	var removedApp domain.Application
	found := false
	for i, app := range portfolio.Applications {
		if app.ID == cmd.ApplicationID {
			removedApp = app
			portfolio.Applications = append(portfolio.Applications[:i], portfolio.Applications[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("application not found in portfolio")
	}

	portfolio.UpdatedAt = time.Now()

	err = s.portfolioRepo.Save(ctx, portfolio)
	if err != nil {
		return fmt.Errorf("failed to save updated portfolio: %w", err)
	}

	// Publish domain event
	event := domain.ApplicationRemovedFromPortfolioEvent{
		PortfolioID:     cmd.PortfolioID,
		ApplicationID:   cmd.ApplicationID,
		ApplicationName: removedApp.Name,
		OccurredAt:      time.Now(),
	}

	err = s.eventRepo.Save(ctx, event)
	if err != nil {
		fmt.Printf("Failed to save domain event: %v\n", err)
	}

	return nil
}

// GetPortfolio retrieves a portfolio by ID
func (s *PortfolioService) GetPortfolio(ctx context.Context, portfolioID domain.PortfolioID) (*domain.ApplicationPortfolio, error) {
	portfolio, err := s.portfolioRepo.FindByID(ctx, portfolioID)
	if err != nil {
		return nil, fmt.Errorf("failed to get portfolio: %w", err)
	}
	return &portfolio, nil
}

// ListPortfolios retrieves all portfolios
func (s *PortfolioService) ListPortfolios(ctx context.Context) ([]domain.ApplicationPortfolio, error) {
	portfolios, err := s.portfolioRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list portfolios: %w", err)
	}
	return portfolios, nil
}

// ListPortfoliosByOwner retrieves portfolios by owner
func (s *PortfolioService) ListPortfoliosByOwner(ctx context.Context, owner string) ([]domain.ApplicationPortfolio, error) {
	portfolios, err := s.portfolioRepo.FindByOwner(ctx, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to list portfolios by owner: %w", err)
	}
	return portfolios, nil
}

// UpdatePortfolio updates portfolio information
func (s *PortfolioService) UpdatePortfolio(ctx context.Context, cmd UpdatePortfolioCommand) error {
	portfolio, err := s.portfolioRepo.FindByID(ctx, cmd.ID)
	if err != nil {
		return fmt.Errorf("portfolio not found: %w", err)
	}

	// Update fields
	portfolio.Name = cmd.Name
	portfolio.Description = cmd.Description
	portfolio.UpdatedAt = time.Now()

	err = s.portfolioRepo.Save(ctx, portfolio)
	if err != nil {
		return fmt.Errorf("failed to update portfolio: %w", err)
	}

	return nil
}

// DeletePortfolio deletes a portfolio
func (s *PortfolioService) DeletePortfolio(ctx context.Context, portfolioID domain.PortfolioID) error {
	// Check if portfolio has applications
	portfolio, err := s.portfolioRepo.FindByID(ctx, portfolioID)
	if err != nil {
		return fmt.Errorf("portfolio not found: %w", err)
	}

	if len(portfolio.Applications) > 0 {
		return fmt.Errorf("cannot delete portfolio with applications")
	}

	err = s.portfolioRepo.Delete(ctx, portfolioID)
	if err != nil {
		return fmt.Errorf("failed to delete portfolio: %w", err)
	}

	return nil
}

// Commands for Portfolio Service

type CreatePortfolioCommand struct {
	ID          domain.PortfolioID
	Name        string
	Description string
	Owner       string
}

type AddApplicationToPortfolioCommand struct {
	PortfolioID   domain.PortfolioID
	ApplicationID domain.ApplicationID
}

type RemoveApplicationFromPortfolioCommand struct {
	PortfolioID   domain.PortfolioID
	ApplicationID domain.ApplicationID
}

type UpdatePortfolioCommand struct {
	ID          domain.PortfolioID
	Name        string
	Description string
}
