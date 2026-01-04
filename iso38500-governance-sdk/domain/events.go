package domain

import "time"

// DomainEvent represents a domain event
type DomainEvent interface {
	EventType() string
	Time() time.Time
}

// PortfolioCreatedEvent represents a portfolio creation event
type PortfolioCreatedEvent struct {
	PortfolioID PortfolioID
	Name        string
	Owner       string
	OccurredAt  time.Time
}

func (e PortfolioCreatedEvent) EventType() string {
	return "PortfolioCreated"
}

func (e PortfolioCreatedEvent) Time() time.Time {
	return e.OccurredAt
}

// ApplicationAddedToPortfolioEvent represents an application addition event
type ApplicationAddedToPortfolioEvent struct {
	PortfolioID          PortfolioID
	ApplicationID        ApplicationID
	ApplicationName      string
	GovernanceAgreementID GovernanceAgreementID
	OccurredAt           time.Time
}

func (e ApplicationAddedToPortfolioEvent) EventType() string {
	return "ApplicationAddedToPortfolio"
}

func (e ApplicationAddedToPortfolioEvent) Time() time.Time {
	return e.OccurredAt
}

// ApplicationRemovedFromPortfolioEvent represents an application removal event
type ApplicationRemovedFromPortfolioEvent struct {
	PortfolioID     PortfolioID
	ApplicationID   ApplicationID
	ApplicationName string
	OccurredAt      time.Time
}

func (e ApplicationRemovedFromPortfolioEvent) EventType() string {
	return "ApplicationRemovedFromPortfolio"
}

func (e ApplicationRemovedFromPortfolioEvent) Time() time.Time {
	return e.OccurredAt
}

// ApplicationUpdatedEvent represents an application update event
type ApplicationUpdatedEvent struct {
	PortfolioID     PortfolioID
	ApplicationID   ApplicationID
	ApplicationName string
	OccurredAt      time.Time
}

func (e ApplicationUpdatedEvent) EventType() string {
	return "ApplicationUpdated"
}

func (e ApplicationUpdatedEvent) Time() time.Time {
	return e.OccurredAt
}

// GovernanceAgreementCreatedEvent represents a governance agreement creation event
type GovernanceAgreementCreatedEvent struct {
	AgreementID   GovernanceAgreementID
	ApplicationID ApplicationID
	Title         string
	OccurredAt    time.Time
}

func (e GovernanceAgreementCreatedEvent) EventType() string {
	return "GovernanceAgreementCreated"
}

func (e GovernanceAgreementCreatedEvent) Time() time.Time {
	return e.OccurredAt
}

// GovernanceAgreementUpdatedEvent represents a governance agreement update event
type GovernanceAgreementUpdatedEvent struct {
	AgreementID GovernanceAgreementID
	Component   string
	OccurredAt  time.Time
}

func (e GovernanceAgreementUpdatedEvent) EventType() string {
	return "GovernanceAgreementUpdated"
}

func (e GovernanceAgreementUpdatedEvent) Time() time.Time {
	return e.OccurredAt
}

// GovernanceAgreementApprovedEvent represents a governance agreement approval event
type GovernanceAgreementApprovedEvent struct {
	AgreementID GovernanceAgreementID
	OccurredAt  time.Time
}

func (e GovernanceAgreementApprovedEvent) EventType() string {
	return "GovernanceAgreementApproved"
}

func (e GovernanceAgreementApprovedEvent) Time() time.Time {
	return e.OccurredAt
}

// GovernanceAgreementActivatedEvent represents a governance agreement activation event
type GovernanceAgreementActivatedEvent struct {
	AgreementID GovernanceAgreementID
	OccurredAt  time.Time
}

func (e GovernanceAgreementActivatedEvent) EventType() string {
	return "GovernanceAgreementActivated"
}

func (e GovernanceAgreementActivatedEvent) Time() time.Time {
	return e.OccurredAt
}

// GovernanceEvaluationCompletedEvent represents a governance evaluation completion event
type GovernanceEvaluationCompletedEvent struct {
	AgreementID     GovernanceAgreementID
	Evaluator       string
	Findings        []string
	Recommendations []string
	OccurredAt      time.Time
}

func (e GovernanceEvaluationCompletedEvent) EventType() string {
	return "GovernanceEvaluationCompleted"
}

func (e GovernanceEvaluationCompletedEvent) Time() time.Time {
	return e.OccurredAt
}

// GovernanceDirectionSetEvent represents a governance direction setting event
type GovernanceDirectionSetEvent struct {
	AgreementID GovernanceAgreementID
	Director    string
	Objectives  []string
	ActionPlans []string
	OccurredAt  time.Time
}

func (e GovernanceDirectionSetEvent) EventType() string {
	return "GovernanceDirectionSet"
}

func (e GovernanceDirectionSetEvent) Time() time.Time {
	return e.OccurredAt
}

// GovernanceMonitoringCompletedEvent represents a governance monitoring completion event
type GovernanceMonitoringCompletedEvent struct {
	AgreementID      GovernanceAgreementID
	Monitor          string
	KPIMeasurements  []string
	ComplianceStatus string
	RiskStatus       string
	OccurredAt       time.Time
}

func (e GovernanceMonitoringCompletedEvent) EventType() string {
	return "GovernanceMonitoringCompleted"
}

func (e GovernanceMonitoringCompletedEvent) Time() time.Time {
	return e.OccurredAt
}

// ChangeRequestCreatedEvent represents a change request creation event
type ChangeRequestCreatedEvent struct {
	ChangeRequestID string
	ApplicationID   ApplicationID
	Requester       string
	Type            ChangeType
	Priority        Priority
	Description     string
	OccurredAt      time.Time
}

func (e ChangeRequestCreatedEvent) EventType() string {
	return "ChangeRequestCreated"
}

func (e ChangeRequestCreatedEvent) Time() time.Time {
	return e.OccurredAt
}

// ChangeRequestApprovedEvent represents a change request approval event
type ChangeRequestApprovedEvent struct {
	ChangeRequestID string
	Approver        string
	OccurredAt      time.Time
}

func (e ChangeRequestApprovedEvent) EventType() string {
	return "ChangeRequestApproved"
}

func (e ChangeRequestApprovedEvent) Time() time.Time {
	return e.OccurredAt
}

// IncidentReportedEvent represents an incident reporting event
type IncidentReportedEvent struct {
	IncidentID     string
	ApplicationID  ApplicationID
	Reporter       string
	Severity       int
	Description    string
	OccurredAt     time.Time
}

func (e IncidentReportedEvent) EventType() string {
	return "IncidentReported"
}

func (e IncidentReportedEvent) Time() time.Time {
	return e.OccurredAt
}

// IncidentResolvedEvent represents an incident resolution event
type IncidentResolvedEvent struct {
	IncidentID     string
	Resolver       string
	Resolution     string
	TimeToResolve  time.Duration
	OccurredAt     time.Time
}

func (e IncidentResolvedEvent) EventType() string {
	return "IncidentResolved"
}

func (e IncidentResolvedEvent) Time() time.Time {
	return e.OccurredAt
}

// ComplianceViolationDetectedEvent represents a compliance violation detection event
type ComplianceViolationDetectedEvent struct {
	ViolationID     string
	ApplicationID   ApplicationID
	RequirementType string
	Description     string
	Severity        string
	OccurredAt      time.Time
}

func (e ComplianceViolationDetectedEvent) EventType() string {
	return "ComplianceViolationDetected"
}

func (e ComplianceViolationDetectedEvent) Time() time.Time {
	return e.OccurredAt
}

// AuditCompletedEvent represents an audit completion event
type AuditCompletedEvent struct {
	AuditID        string
	ApplicationID  ApplicationID
	Auditor        string
	Scope          string
	Findings       []string
	Status         string
	OccurredAt     time.Time
}

func (e AuditCompletedEvent) EventType() string {
	return "AuditCompleted"
}

func (e AuditCompletedEvent) Time() time.Time {
	return e.OccurredAt
}
