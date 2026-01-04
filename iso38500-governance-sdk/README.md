# ISO 38500 Governance Framework SDK

A comprehensive Go SDK implementing the ISO 38500 corporate governance of information technology framework. This SDK provides a structured approach to managing application portfolios and governance agreements based on the three core principles: **Evaluate**, **Direct**, and **Monitor**.

## Overview

The ISO 38500 standard provides guidance for governing IT within organizations. This SDK implements the framework through:

- **Domain-Driven Design** principles for clean architecture
- **Hexagonal Architecture** (Ports & Adapters) for infrastructure independence
- **CQRS** and **Event Sourcing** patterns for scalability
- **SOLID** principles for maintainable code

## Architecture

The SDK is organized into four main layers:

```
┌─────────────────┐
│   Application   │  Use Cases & Application Services
├─────────────────┤
│     Domain      │  Business Logic & Rules
├─────────────────┤
│ Infrastructure  │  External Concerns (DB, APIs, etc.)
├─────────────────┤
│   Presentation  │  REST APIs, CLI, etc. (future)
└─────────────────┘
```

### Core Components

- **Domain Layer**: Contains entities, value objects, aggregates, domain services, and domain events
- **Application Layer**: Contains use case implementations and application services
- **Infrastructure Layer**: Contains repository implementations and external integrations
- **Events**: Domain events for governance actions and audit trails

## Key Features

### 📋 Application Portfolio Management
- Create and manage application portfolios
- Track application lifecycle and dependencies
- Monitor portfolio health and optimization opportunities

### 📄 Governance Agreements
- Structured governance agreements based on ISO 38500
- RACI matrices for responsibility assignment
- Comprehensive governance components (Strategy, Acquisition, Performance, etc.)

### 🔍 Evaluation Framework
- Application health assessments
- Technical and business value evaluation
- Risk assessment and recommendations

### 🎯 Strategic Direction
- Set organizational objectives and initiatives
- Resource allocation and budgeting
- Policy and standards establishment

### 📊 Monitoring & Compliance
- KPI tracking and measurement
- Compliance monitoring
- Risk management and mitigation

### 🔄 Change Management
- Change request lifecycle
- Incident management
- Audit trails and reporting

## Installation

```bash
go get github.com/iso38500/iso38500-governance-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "github.com/iso38500/iso38500-governance-sdk/application"
    "github.com/iso38500/iso38500-governance-sdk/infrastructure/memory"
)

func main() {
    // Initialize repositories
    appRepo := memory.NewApplicationRepositoryMemory()
    govRepo := memory.NewGovernanceAgreementRepositoryMemory()
    portfolioRepo := memory.NewApplicationPortfolioRepositoryMemory()
    eventRepo := memory.NewDomainEventRepositoryMemory()

    // Initialize services
    portfolioService := application.NewPortfolioService(portfolioRepo, appRepo, govRepo, eventRepo)

    // Create a portfolio
    portfolio, err := portfolioService.CreatePortfolio(context.Background(), application.CreatePortfolioCommand{
        ID:          "portfolio-001",
        Name:        "Enterprise Applications",
        Description: "Main enterprise application portfolio",
        Owner:       "IT Governance Team",
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("Created portfolio: %s\n", portfolio.Name)
}
```

## ISO 38500 Principles Implementation

### 1. Evaluate Principle
Assess the current and future use of IT to ensure alignment with organizational objectives.

```go
evaluationService := domain.NewEvaluationService(appRepo, govRepo, kpiRepo, riskRepo)

// Evaluate an application
assessment, err := evaluationService.EvaluateApplication(ctx, appID, "evaluator")
if err != nil {
    // handle error
}

fmt.Printf("Risk Level: %s\n", assessment.RiskLevel)
fmt.Printf("Recommendations: %d\n", len(assessment.Recommendations))
```

### 2. Direct Principle
Ensure that IT activities are aligned with organizational objectives and resources are used responsibly.

```go
directionService := domain.NewDirectionService(govRepo)

// Set strategic direction
err := governanceService.SetStrategicDirection(ctx, application.SetStrategicDirectionCommand{
    AgreementID: agreementID,
    Director:    "CIO",
    Objectives: []domain.StrategicObjective{...},
    Initiatives: []domain.StrategicInitiative{...},
})
```

### 3. Monitor Principle
Monitor IT activities to ensure compliance with organizational objectives and policies.

```go
monitoringService := domain.NewMonitoringService(kpiRepo, measurementRepo, riskRepo, govRepo)

// Monitor governance
result, err := governanceService.MonitorGovernance(ctx, application.MonitorGovernanceCommand{
    AgreementID: agreementID,
})
```

## Domain Concepts

### Core Entities
- **Application**: Software applications within the portfolio
- **ApplicationPortfolio**: Collections of applications
- **GovernanceAgreement**: Governance framework for applications
- **ChangeRequest**: Change management requests
- **Incident**: System incidents and resolutions
- **Audit**: Compliance and operational audits

### Value Objects
- **ResponsibilityMatrix**: RACI matrices for stakeholder roles
- **KPIs**: Key Performance Indicators
- **Risk**: Risk assessments and mitigations
- **SLA**: Service Level Agreements
- **SecurityProvisions**: Security requirements and measures

### Aggregates
- **ApplicationPortfolioAggregate**: Manages portfolio consistency
- **GovernanceAgreementAggregate**: Ensures governance agreement integrity

## Repository Pattern

The SDK uses repository interfaces to abstract data access:

```go
type ApplicationRepository interface {
    Save(ctx context.Context, app Application) error
    FindByID(ctx context.Context, id ApplicationID) (Application, error)
    FindAll(ctx context.Context) ([]Application, error)
    // ... other methods
}
```

### Available Implementations
- **Memory**: In-memory storage for testing and development
- **Database**: Planned implementations for PostgreSQL, MySQL, MongoDB
- **File**: JSON file-based persistence

## Domain Events

The SDK implements domain events for audit trails and event sourcing:

```go
// Example domain events
PortfolioCreatedEvent
ApplicationAddedToPortfolioEvent
GovernanceAgreementApprovedEvent
ChangeRequestCreatedEvent
IncidentResolvedEvent
```

## Examples

Run the comprehensive example:

```bash
cd examples
go run main.go
```

This demonstrates:
- Creating applications and portfolios
- Establishing governance agreements
- Evaluating applications
- Setting strategic direction
- Monitoring governance activities

## Testing

```bash
go test ./...
```

## Contributing

1. Follow Domain-Driven Design principles
2. Maintain hexagonal architecture
3. Write comprehensive tests
4. Update documentation
5. Follow Go coding standards

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## References

- [ISO/IEC 38500:2015 - Information technology — Governance of IT for the organization](https://www.iso.org/standard/62816.html)
- Domain-Driven Design by Eric Evans
- Clean Architecture by Robert C. Martin
- Implementing Domain-Driven Design by Vaughn Vernon

## Roadmap

- [ ] REST API implementation
- [ ] Database adapters (PostgreSQL, MySQL)
- [ ] Message queue integration
- [ ] Advanced reporting and dashboards
- [ ] Integration with popular ITSM tools
- [ ] Compliance automation features
- [ ] Multi-tenant support
- [ ] GraphQL API
- [ ] Kubernetes operator
