package domain

import (
	"errors"
	"fmt"
	"time"
)

// ApplicationPortfolioAggregate represents the application portfolio aggregate
type ApplicationPortfolioAggregate struct {
	portfolio     ApplicationPortfolio
	domainEvents  []DomainEvent
}

// NewApplicationPortfolioAggregate creates a new portfolio aggregate
func NewApplicationPortfolioAggregate(id PortfolioID, name, description, owner string) (*ApplicationPortfolioAggregate, error) {
	if id == "" {
		return nil, errors.New("portfolio ID cannot be empty")
	}
	if name == "" {
		return nil, errors.New("portfolio name cannot be empty")
	}
	if owner == "" {
		return nil, errors.New("portfolio owner cannot be empty")
	}

	portfolio := ApplicationPortfolio{
		ID:          id,
		Name:        name,
		Description: description,
		Owner:       owner,
		Applications: []Application{},
		KPIs:        []KPI{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	aggregate := &ApplicationPortfolioAggregate{
		portfolio:    portfolio,
		domainEvents: []DomainEvent{},
	}

	// Add domain event
	aggregate.addDomainEvent(PortfolioCreatedEvent{
		PortfolioID: id,
		Name:        name,
		Owner:       owner,
		OccurredAt:  time.Now(),
	})

	return aggregate, nil
}

// AddApplication adds an application to the portfolio with business rules
func (a *ApplicationPortfolioAggregate) AddApplication(app Application) error {
	// Business invariant: Application must be valid
	if err := app.Validate(); err != nil {
		return fmt.Errorf("invalid application: %w", err)
	}

	// Business invariant: No duplicate applications
	for _, existing := range a.portfolio.Applications {
		if existing.ID == app.ID {
			return errors.New("application already exists in portfolio")
		}
		if existing.Name == app.Name {
			return errors.New("application with same name already exists in portfolio")
		}
	}

	// Business invariant: Application must have an active governance agreement
	if app.GovernanceAgreementID == "" {
		return errors.New("application must have a governance agreement")
	}

	a.portfolio.Applications = append(a.portfolio.Applications, app)
	a.portfolio.UpdatedAt = time.Now()

			// Add domain event
			a.addDomainEvent(ApplicationAddedToPortfolioEvent{
				PortfolioID:          a.portfolio.ID,
				ApplicationID:        app.ID,
				ApplicationName:      app.Name,
				GovernanceAgreementID: app.GovernanceAgreementID,
				OccurredAt:           time.Now(),
			})

	return nil
}

// RemoveApplication removes an application from the portfolio
func (a *ApplicationPortfolioAggregate) RemoveApplication(appID ApplicationID) error {
	for i, app := range a.portfolio.Applications {
		if app.ID == appID {
			removedApp := app
			a.portfolio.Applications = append(a.portfolio.Applications[:i], a.portfolio.Applications[i+1:]...)
			a.portfolio.UpdatedAt = time.Now()

			// Add domain event
			a.addDomainEvent(ApplicationRemovedFromPortfolioEvent{
				PortfolioID:     a.portfolio.ID,
				ApplicationID:   removedApp.ID,
				ApplicationName: removedApp.Name,
				OccurredAt:      time.Now(),
			})

			return nil
		}
	}
	return errors.New("application not found in portfolio")
}

// UpdateApplication updates an existing application
func (a *ApplicationPortfolioAggregate) UpdateApplication(app Application) error {
	if err := app.Validate(); err != nil {
		return fmt.Errorf("invalid application: %w", err)
	}

	for i, existing := range a.portfolio.Applications {
		if existing.ID == app.ID {
			a.portfolio.Applications[i] = app
			a.portfolio.UpdatedAt = time.Now()

			// Add domain event
			a.addDomainEvent(ApplicationUpdatedEvent{
				PortfolioID:     a.portfolio.ID,
				ApplicationID:   app.ID,
				ApplicationName: app.Name,
				OccurredAt:      time.Now(),
			})

			return nil
		}
	}
	return errors.New("application not found in portfolio")
}

// GetPortfolio returns the portfolio
func (a *ApplicationPortfolioAggregate) GetPortfolio() ApplicationPortfolio {
	return a.portfolio
}

// GetDomainEvents returns the domain events
func (a *ApplicationPortfolioAggregate) GetDomainEvents() []DomainEvent {
	return a.domainEvents
}

// ClearDomainEvents clears the domain events
func (a *ApplicationPortfolioAggregate) ClearDomainEvents() {
	a.domainEvents = []DomainEvent{}
}

// addDomainEvent adds a domain event to the aggregate
func (a *ApplicationPortfolioAggregate) addDomainEvent(event DomainEvent) {
	a.domainEvents = append(a.domainEvents, event)
}

// GovernanceAgreementAggregate represents the governance agreement aggregate
type GovernanceAgreementAggregate struct {
	agreement     GovernanceAgreement
	domainEvents  []DomainEvent
}

// NewGovernanceAgreementAggregate creates a new governance agreement aggregate
func NewGovernanceAgreementAggregate(id GovernanceAgreementID, applicationID ApplicationID, title string) (*GovernanceAgreementAggregate, error) {
	if id == "" {
		return nil, errors.New("governance agreement ID cannot be empty")
	}
	if applicationID == "" {
		return nil, errors.New("application ID cannot be empty")
	}
	if title == "" {
		return nil, errors.New("governance agreement title cannot be empty")
	}

	agreement := GovernanceAgreement{
		ID:             id,
		ApplicationID:  applicationID,
		Title:          title,
		Version:        "1.0",
		Status:         AgreementDraft,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	aggregate := &GovernanceAgreementAggregate{
		agreement:    agreement,
		domainEvents: []DomainEvent{},
	}

	// Add domain event
	aggregate.addDomainEvent(GovernanceAgreementCreatedEvent{
		AgreementID:   id,
		ApplicationID: applicationID,
		Title:         title,
		OccurredAt:    time.Now(),
	})

	return aggregate, nil
}

// UpdateStrategy updates the strategy component
func (a *GovernanceAgreementAggregate) UpdateStrategy(strategy Strategy) error {
	a.agreement.Strategy = strategy
	a.agreement.UpdatedAt = time.Now()

	a.addDomainEvent(GovernanceAgreementUpdatedEvent{
		AgreementID: a.agreement.ID,
		Component:   "strategy",
		OccurredAt:  time.Now(),
	})

	return nil
}

// UpdateAcquisition updates the acquisition component
func (a *GovernanceAgreementAggregate) UpdateAcquisition(acquisition Acquisition) error {
	a.agreement.Acquisition = acquisition
	a.agreement.UpdatedAt = time.Now()

	a.addDomainEvent(GovernanceAgreementUpdatedEvent{
		AgreementID: a.agreement.ID,
		Component:   "acquisition",
		OccurredAt:  time.Now(),
	})

	return nil
}

// UpdatePerformance updates the performance component
func (a *GovernanceAgreementAggregate) UpdatePerformance(performance Performance) error {
	a.agreement.Performance = performance
	a.agreement.UpdatedAt = time.Now()

	a.addDomainEvent(GovernanceAgreementUpdatedEvent{
		AgreementID: a.agreement.ID,
		Component:   "performance",
		OccurredAt:  time.Now(),
	})

	return nil
}

// UpdateConformance updates the conformance component
func (a *GovernanceAgreementAggregate) UpdateConformance(conformance Conformance) error {
	a.agreement.Conformance = conformance
	a.agreement.UpdatedAt = time.Now()

	a.addDomainEvent(GovernanceAgreementUpdatedEvent{
		AgreementID: a.agreement.ID,
		Component:   "conformance",
		OccurredAt:  time.Now(),
	})

	return nil
}

// UpdateImplementation updates the implementation component
func (a *GovernanceAgreementAggregate) UpdateImplementation(implementation Implementation) error {
	a.agreement.Implementation = implementation
	a.agreement.UpdatedAt = time.Now()

	a.addDomainEvent(GovernanceAgreementUpdatedEvent{
		AgreementID: a.agreement.ID,
		Component:   "implementation",
		OccurredAt:  time.Now(),
	})

	return nil
}

// Approve approves the governance agreement
func (a *GovernanceAgreementAggregate) Approve() error {
	if a.agreement.Status != AgreementDraft {
		return errors.New("only draft agreements can be approved")
	}

	a.agreement.Status = AgreementApproved
	a.agreement.UpdatedAt = time.Now()

	a.addDomainEvent(GovernanceAgreementApprovedEvent{
		AgreementID: a.agreement.ID,
		OccurredAt:  time.Now(),
	})

	return nil
}

// Activate activates the governance agreement
func (a *GovernanceAgreementAggregate) Activate() error {
	if a.agreement.Status != AgreementApproved {
		return errors.New("only approved agreements can be activated")
	}

	a.agreement.Status = AgreementActive
	a.agreement.UpdatedAt = time.Now()

	a.addDomainEvent(GovernanceAgreementActivatedEvent{
		AgreementID: a.agreement.ID,
		OccurredAt:  time.Now(),
	})

	return nil
}

// GetAgreement returns the governance agreement
func (a *GovernanceAgreementAggregate) GetAgreement() GovernanceAgreement {
	return a.agreement
}

// GetDomainEvents returns the domain events
func (a *GovernanceAgreementAggregate) GetDomainEvents() []DomainEvent {
	return a.domainEvents
}

// ClearDomainEvents clears the domain events
func (a *GovernanceAgreementAggregate) ClearDomainEvents() {
	a.domainEvents = []DomainEvent{}
}

// addDomainEvent adds a domain event to the aggregate
func (a *GovernanceAgreementAggregate) addDomainEvent(event DomainEvent) {
	a.domainEvents = append(a.domainEvents, event)
}
