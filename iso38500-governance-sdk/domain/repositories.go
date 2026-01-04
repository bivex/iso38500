package domain

import (
	"context"
	"time"
)

// ApplicationRepository defines the interface for application data access
type ApplicationRepository interface {
	Save(ctx context.Context, app Application) error
	FindByID(ctx context.Context, id ApplicationID) (Application, error)
	FindByName(ctx context.Context, name string) (Application, error)
	FindAll(ctx context.Context) ([]Application, error)
	FindByPortfolioID(ctx context.Context, portfolioID PortfolioID) ([]Application, error)
	Update(ctx context.Context, app Application) error
	Delete(ctx context.Context, id ApplicationID) error
	Exists(ctx context.Context, id ApplicationID) (bool, error)
}

// GovernanceAgreementRepository defines the interface for governance agreement data access
type GovernanceAgreementRepository interface {
	Save(ctx context.Context, agreement GovernanceAgreement) error
	FindByID(ctx context.Context, id GovernanceAgreementID) (GovernanceAgreement, error)
	FindByApplicationID(ctx context.Context, appID ApplicationID) (GovernanceAgreement, error)
	FindAll(ctx context.Context) ([]GovernanceAgreement, error)
	FindByStatus(ctx context.Context, status AgreementStatus) ([]GovernanceAgreement, error)
	Update(ctx context.Context, agreement GovernanceAgreement) error
	Delete(ctx context.Context, id GovernanceAgreementID) error
	Exists(ctx context.Context, id GovernanceAgreementID) (bool, error)
}

// ApplicationPortfolioRepository defines the interface for portfolio data access
type ApplicationPortfolioRepository interface {
	Save(ctx context.Context, portfolio ApplicationPortfolio) error
	FindByID(ctx context.Context, id PortfolioID) (ApplicationPortfolio, error)
	FindByOwner(ctx context.Context, owner string) ([]ApplicationPortfolio, error)
	FindAll(ctx context.Context) ([]ApplicationPortfolio, error)
	Update(ctx context.Context, portfolio ApplicationPortfolio) error
	Delete(ctx context.Context, id PortfolioID) error
	Exists(ctx context.Context, id PortfolioID) (bool, error)
	AddApplication(ctx context.Context, portfolioID PortfolioID, appID ApplicationID) error
	RemoveApplication(ctx context.Context, portfolioID PortfolioID, appID ApplicationID) error
}

// ChangeRequestRepository defines the interface for change request data access
type ChangeRequestRepository interface {
	Save(ctx context.Context, cr ChangeRequest) error
	FindByID(ctx context.Context, id string) (ChangeRequest, error)
	FindByApplicationID(ctx context.Context, appID ApplicationID) ([]ChangeRequest, error)
	FindByStatus(ctx context.Context, status ChangeRequestStatus) ([]ChangeRequest, error)
	FindByPriority(ctx context.Context, priority Priority) ([]ChangeRequest, error)
	Update(ctx context.Context, cr ChangeRequest) error
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) (bool, error)
}

// IncidentRepository defines the interface for incident data access
type IncidentRepository interface {
	Save(ctx context.Context, incident Incident) error
	FindByID(ctx context.Context, id string) (Incident, error)
	FindByApplicationID(ctx context.Context, appID ApplicationID) ([]Incident, error)
	FindByStatus(ctx context.Context, status IncidentStatus) ([]Incident, error)
	FindBySeverity(ctx context.Context, severity int) ([]Incident, error)
	Update(ctx context.Context, incident Incident) error
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) (bool, error)
}

// AuditRepository defines the interface for audit data access
type AuditRepository interface {
	Save(ctx context.Context, audit Audit) error
	FindByID(ctx context.Context, id string) (Audit, error)
	FindByApplicationID(ctx context.Context, appID ApplicationID) ([]Audit, error)
	FindByStatus(ctx context.Context, status AuditStatus) ([]Audit, error)
	FindByPeriod(ctx context.Context, start, end time.Time) ([]Audit, error)
	Update(ctx context.Context, audit Audit) error
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) (bool, error)
}

// KPIRepository defines the interface for KPI data access
type KPIRepository interface {
	Save(ctx context.Context, kpi KPI) error
	FindByID(ctx context.Context, id string) (KPI, error)
	FindAll(ctx context.Context) ([]KPI, error)
	FindByCategory(ctx context.Context, category string) ([]KPI, error)
	Update(ctx context.Context, kpi KPI) error
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) (bool, error)
}

// KPIMeasurementRepository defines the interface for KPI measurement data access
type KPIMeasurementRepository interface {
	Save(ctx context.Context, measurement KPIMeasurement) error
	FindByKPIID(ctx context.Context, kpiID string) ([]KPIMeasurement, error)
	FindByPeriod(ctx context.Context, kpiID string, start, end time.Time) ([]KPIMeasurement, error)
	FindLatest(ctx context.Context, kpiID string) (KPIMeasurement, error)
	Delete(ctx context.Context, kpiID string, measuredAt time.Time) error
}

// RiskRepository defines the interface for risk data access
type RiskRepository interface {
	Save(ctx context.Context, risk Risk) error
	FindByID(ctx context.Context, id string) (Risk, error)
	FindAll(ctx context.Context) ([]Risk, error)
	FindByLevel(ctx context.Context, level RiskLevel) ([]Risk, error)
	FindByCategory(ctx context.Context, category string) ([]Risk, error)
	Update(ctx context.Context, risk Risk) error
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) (bool, error)
}

// MitigationPlanRepository defines the interface for mitigation plan data access
type MitigationPlanRepository interface {
	Save(ctx context.Context, plan MitigationPlan) error
	FindByRiskID(ctx context.Context, riskID string) (MitigationPlan, error)
	FindAll(ctx context.Context) ([]MitigationPlan, error)
	Update(ctx context.Context, plan MitigationPlan) error
	Delete(ctx context.Context, riskID string) error
	Exists(ctx context.Context, riskID string) (bool, error)
}

// ComplianceRepository defines the interface for compliance data access
type ComplianceRepository interface {
	SaveRequirement(ctx context.Context, req interface{}) error
	FindLegalRequirements(ctx context.Context, appID ApplicationID) ([]LegalRequirement, error)
	FindContractualRequirements(ctx context.Context, appID ApplicationID) ([]ContractualRequirement, error)
	FindIndustryStandards(ctx context.Context, appID ApplicationID) ([]IndustryStandard, error)
	UpdateComplianceStatus(ctx context.Context, reqType, reqID string, status ComplianceStatus) error
}

// DomainEventRepository defines the interface for domain event data access
type DomainEventRepository interface {
	Save(ctx context.Context, event DomainEvent) error
	FindByAggregateID(ctx context.Context, aggregateID string) ([]DomainEvent, error)
	FindByEventType(ctx context.Context, eventType string) ([]DomainEvent, error)
	FindByTimeRange(ctx context.Context, start, end time.Time) ([]DomainEvent, error)
	Delete(ctx context.Context, eventID string) error
}

// ChangeRequest represents a change request entity
type ChangeRequest struct {
	ID            string
	ApplicationID ApplicationID
	Requester     string
	Type          ChangeType
	Priority      Priority
	Status        ChangeRequestStatus
	Title         string
	Description   string
	BusinessCase  string
	Impact        string
	Risk          string
	Approvals     []Approval
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// ChangeRequestStatus represents the status of a change request
type ChangeRequestStatus string

const (
	ChangeStatusDraft     ChangeRequestStatus = "draft"
	ChangeStatusSubmitted ChangeRequestStatus = "submitted"
	ChangeStatusApproved  ChangeRequestStatus = "approved"
	ChangeStatusRejected  ChangeRequestStatus = "rejected"
	ChangeStatusImplemented ChangeRequestStatus = "implemented"
	ChangeStatusClosed    ChangeRequestStatus = "closed"
)

// Approval represents an approval for a change request
type Approval struct {
	Approver    string
	Role        string
	Status      ApprovalStatus
	Comments    string
	ApprovedAt  time.Time
}

// ApprovalStatus represents the status of an approval
type ApprovalStatus string

const (
	ApprovalPending   ApprovalStatus = "pending"
	ApprovalApproved  ApprovalStatus = "approved"
	ApprovalRejected  ApprovalStatus = "rejected"
)

// Incident represents an incident entity
type Incident struct {
	ID            string
	ApplicationID ApplicationID
	Reporter      string
	Severity      int
	Status        IncidentStatus
	Title         string
	Description   string
	Impact        string
	RootCause     string
	Resolution    string
	TimeToResolve time.Duration
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ResolvedAt    time.Time
}

// IncidentStatus represents the status of an incident
type IncidentStatus string

const (
	IncidentStatusOpen      IncidentStatus = "open"
	IncidentStatusInvestigating IncidentStatus = "investigating"
	IncidentStatusResolved   IncidentStatus = "resolved"
	IncidentStatusClosed     IncidentStatus = "closed"
)

// Audit represents an audit entity
type Audit struct {
	ID            string
	ApplicationID ApplicationID
	Auditor       string
	Type          AuditType
	Status        AuditStatus
	Scope         string
	Findings      []AuditFinding
	Recommendations []string
	StartedAt     time.Time
	CompletedAt   time.Time
}

// AuditType represents the type of audit
type AuditType string

const (
	AuditTypeSecurity    AuditType = "security"
	AuditTypeCompliance  AuditType = "compliance"
	AuditTypePerformance AuditType = "performance"
	AuditTypeOperational AuditType = "operational"
)

// AuditStatus represents the status of an audit
type AuditStatus string

const (
	AuditStatusPlanned    AuditStatus = "planned"
	AuditStatusInProgress AuditStatus = "in_progress"
	AuditStatusCompleted  AuditStatus = "completed"
	AuditStatusOverdue    AuditStatus = "overdue"
)

// AuditFinding represents an audit finding
type AuditFinding struct {
	ID          string
	Severity    string
	Category    string
	Description string
	Evidence    string
	Remediation string
}
