/**
 * Copyright (c) 2026 Bivex
 *
 * Author: Bivex
 * Available for contact via email: support@b-b.top
 * For up-to-date contact information:
 * https://github.com/bivex
 *
 * Created: 2026-01-04T06:24:01
 * Last Updated: 2026-01-04T06:25:15
 *
 * Licensed under the MIT License.
 * Commercial licensing available upon request.
 */

package domain

import (
	"errors"
	"time"
)

// ApplicationID represents a unique identifier for an application
type ApplicationID string

// GovernanceAgreementID represents a unique identifier for a governance agreement
type GovernanceAgreementID string

// PortfolioID represents a unique identifier for an application portfolio
type PortfolioID string

// Application represents a software application within the portfolio
type Application struct {
	ID          ApplicationID
	Name        string
	Description string
	Version     string
	Status      ApplicationStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Governance related
	GovernanceAgreementID GovernanceAgreementID
	Catalogue             ApplicationCatalogue
	Interfaces            []ApplicationInterface
	ConfigurationStandard ConfigurationStandard
	SecurityProvisions    SecurityProvisions
	BusinessContinuity    BusinessContinuity
}

// ApplicationStatus represents the lifecycle status of an application
type ApplicationStatus string

const (
	StatusActive     ApplicationStatus = "active"
	StatusDeprecated ApplicationStatus = "deprecated"
	StatusRetired    ApplicationStatus = "retired"
	StatusPlanned    ApplicationStatus = "planned"
)

// Validate ensures the application has valid data
func (a *Application) Validate() error {
	if a.ID == "" {
		return errors.New("application ID cannot be empty")
	}
	if a.Name == "" {
		return errors.New("application name cannot be empty")
	}
	return nil
}

// GovernanceAgreement represents the governance framework for an application
type GovernanceAgreement struct {
	ID          GovernanceAgreementID
	ApplicationID ApplicationID
	Title       string
	Version     string
	Status      AgreementStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Core governance components
	ResponsibilityMatrix    ResponsibilityMatrix
	Strategy               Strategy
	Acquisition            Acquisition
	Performance            Performance
	Conformance            Conformance
	Implementation         Implementation

	// ISO 38500 principles
	Evaluate EvaluatePrinciple
	Direct   DirectPrinciple
	Monitor  MonitorPrinciple
}

// AgreementStatus represents the status of a governance agreement
type AgreementStatus string

const (
	AgreementDraft     AgreementStatus = "draft"
	AgreementApproved  AgreementStatus = "approved"
	AgreementActive    AgreementStatus = "active"
	AgreementSuspended AgreementStatus = "suspended"
	AgreementRetired   AgreementStatus = "retired"
)

// Validate ensures the governance agreement has valid data
func (ga *GovernanceAgreement) Validate() error {
	if ga.ID == "" {
		return errors.New("governance agreement ID cannot be empty")
	}
	if ga.ApplicationID == "" {
		return errors.New("application ID cannot be empty")
	}
	if ga.Title == "" {
		return errors.New("governance agreement title cannot be empty")
	}
	return nil
}

// ApplicationPortfolio represents a collection of applications
type ApplicationPortfolio struct {
	ID          PortfolioID
	Name        string
	Description string
	Owner       string
	Applications []Application
	KPIs        []KPI
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Validate ensures the portfolio has valid data
func (ap *ApplicationPortfolio) Validate() error {
	if ap.ID == "" {
		return errors.New("portfolio ID cannot be empty")
	}
	if ap.Name == "" {
		return errors.New("portfolio name cannot be empty")
	}
	return nil
}

// AddApplication adds an application to the portfolio
func (ap *ApplicationPortfolio) AddApplication(app Application) error {
	if err := app.Validate(); err != nil {
		return err
	}

	// Check for duplicate applications
	for _, existing := range ap.Applications {
		if existing.ID == app.ID {
			return errors.New("application already exists in portfolio")
		}
	}

	ap.Applications = append(ap.Applications, app)
	ap.UpdatedAt = time.Now()
	return nil
}

// RemoveApplication removes an application from the portfolio
func (ap *ApplicationPortfolio) RemoveApplication(appID ApplicationID) error {
	for i, app := range ap.Applications {
		if app.ID == appID {
			ap.Applications = append(ap.Applications[:i], ap.Applications[i+1:]...)
			ap.UpdatedAt = time.Now()
			return nil
		}
	}
	return errors.New("application not found in portfolio")
}
