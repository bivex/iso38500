package application

import (
	"context"
	"fmt"
	"time"

	"github.com/iso38500/iso38500-governance-sdk/domain"
)

// ChangeManagementService provides application services for change management
type ChangeManagementService struct {
	changeRequestRepo domain.ChangeRequestRepository
	incidentRepo      domain.IncidentRepository
	auditRepo         domain.AuditRepository
	appRepo           domain.ApplicationRepository
	eventRepo         domain.DomainEventRepository
}

// NewChangeManagementService creates a new change management service
func NewChangeManagementService(
	changeRequestRepo domain.ChangeRequestRepository,
	incidentRepo domain.IncidentRepository,
	auditRepo domain.AuditRepository,
	appRepo domain.ApplicationRepository,
	eventRepo domain.DomainEventRepository,
) *ChangeManagementService {
	return &ChangeManagementService{
		changeRequestRepo: changeRequestRepo,
		incidentRepo:      incidentRepo,
		auditRepo:         auditRepo,
		appRepo:           appRepo,
		eventRepo:         eventRepo,
	}
}

// CreateChangeRequest creates a new change request
func (s *ChangeManagementService) CreateChangeRequest(ctx context.Context, cmd CreateChangeRequestCommand) (*domain.ChangeRequest, error) {
	// Verify application exists
	_, err := s.appRepo.FindByID(ctx, cmd.ApplicationID)
	if err != nil {
		return nil, fmt.Errorf("application not found: %w", err)
	}

	changeRequest := domain.ChangeRequest{
		ID:            cmd.ID,
		ApplicationID: cmd.ApplicationID,
		Requester:     cmd.Requester,
		Type:          cmd.Type,
		Priority:      cmd.Priority,
		Status:        domain.ChangeStatusDraft,
		Title:         cmd.Title,
		Description:   cmd.Description,
		BusinessCase:  cmd.BusinessCase,
		Impact:        cmd.Impact,
		Risk:          cmd.Risk,
		Approvals:     []domain.Approval{},
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err = s.changeRequestRepo.Save(ctx, changeRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to save change request: %w", err)
	}

	// Publish domain event
	event := domain.ChangeRequestCreatedEvent{
		ChangeRequestID: changeRequest.ID,
		ApplicationID:   changeRequest.ApplicationID,
		Requester:       changeRequest.Requester,
		Type:           changeRequest.Type,
		Priority:       changeRequest.Priority,
		Description:    changeRequest.Description,
		OccurredAt:     time.Now(),
	}

	err = s.eventRepo.Save(ctx, event)
	if err != nil {
		fmt.Printf("Failed to save domain event: %v\n", err)
	}

	return &changeRequest, nil
}

// ApproveChangeRequest approves a change request
func (s *ChangeManagementService) ApproveChangeRequest(ctx context.Context, cmd ApproveChangeRequestCommand) error {
	changeRequest, err := s.changeRequestRepo.FindByID(ctx, cmd.ChangeRequestID)
	if err != nil {
		return fmt.Errorf("change request not found: %w", err)
	}

	if changeRequest.Status != domain.ChangeStatusSubmitted {
		return fmt.Errorf("change request is not in submitted status")
	}

	// Add approval
	approval := domain.Approval{
		Approver:   cmd.Approver,
		Role:       cmd.Role,
		Status:     domain.ApprovalApproved,
		Comments:   cmd.Comments,
		ApprovedAt: time.Now(),
	}

	changeRequest.Approvals = append(changeRequest.Approvals, approval)
	changeRequest.Status = domain.ChangeStatusApproved
	changeRequest.UpdatedAt = time.Now()

	err = s.changeRequestRepo.Update(ctx, changeRequest)
	if err != nil {
		return fmt.Errorf("failed to update change request: %w", err)
	}

	// Publish domain event
	event := domain.ChangeRequestApprovedEvent{
		ChangeRequestID: cmd.ChangeRequestID,
		Approver:        cmd.Approver,
		OccurredAt:      time.Now(),
	}

	err = s.eventRepo.Save(ctx, event)
	if err != nil {
		fmt.Printf("Failed to save domain event: %v\n", err)
	}

	return nil
}

// RejectChangeRequest rejects a change request
func (s *ChangeManagementService) RejectChangeRequest(ctx context.Context, cmd RejectChangeRequestCommand) error {
	changeRequest, err := s.changeRequestRepo.FindByID(ctx, cmd.ChangeRequestID)
	if err != nil {
		return fmt.Errorf("change request not found: %w", err)
	}

	if changeRequest.Status != domain.ChangeStatusSubmitted {
		return fmt.Errorf("change request is not in submitted status")
	}

	// Add rejection
	approval := domain.Approval{
		Approver:   cmd.Approver,
		Role:       cmd.Role,
		Status:     domain.ApprovalRejected,
		Comments:   cmd.Comments,
		ApprovedAt: time.Now(),
	}

	changeRequest.Approvals = append(changeRequest.Approvals, approval)
	changeRequest.Status = domain.ChangeStatusRejected
	changeRequest.UpdatedAt = time.Now()

	err = s.changeRequestRepo.Update(ctx, changeRequest)
	if err != nil {
		return fmt.Errorf("failed to update change request: %w", err)
	}

	return nil
}

// SubmitChangeRequest submits a change request for approval
func (s *ChangeManagementService) SubmitChangeRequest(ctx context.Context, changeRequestID string) error {
	changeRequest, err := s.changeRequestRepo.FindByID(ctx, changeRequestID)
	if err != nil {
		return fmt.Errorf("change request not found: %w", err)
	}

	if changeRequest.Status != domain.ChangeStatusDraft {
		return fmt.Errorf("change request is not in draft status")
	}

	changeRequest.Status = domain.ChangeStatusSubmitted
	changeRequest.UpdatedAt = time.Now()

	err = s.changeRequestRepo.Update(ctx, changeRequest)
	if err != nil {
		return fmt.Errorf("failed to submit change request: %w", err)
	}

	return nil
}

// ReportIncident reports a new incident
func (s *ChangeManagementService) ReportIncident(ctx context.Context, cmd ReportIncidentCommand) (*domain.Incident, error) {
	// Verify application exists
	_, err := s.appRepo.FindByID(ctx, cmd.ApplicationID)
	if err != nil {
		return nil, fmt.Errorf("application not found: %w", err)
	}

	incident := domain.Incident{
		ID:            cmd.ID,
		ApplicationID: cmd.ApplicationID,
		Reporter:      cmd.Reporter,
		Severity:      cmd.Severity,
		Status:        domain.IncidentStatusOpen,
		Title:         cmd.Title,
		Description:   cmd.Description,
		Impact:        cmd.Impact,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err = s.incidentRepo.Save(ctx, incident)
	if err != nil {
		return nil, fmt.Errorf("failed to save incident: %w", err)
	}

	// Publish domain event
	event := domain.IncidentReportedEvent{
		IncidentID:    incident.ID,
		ApplicationID: incident.ApplicationID,
		Reporter:      incident.Reporter,
		Severity:      incident.Severity,
		Description:   incident.Description,
		OccurredAt:    time.Now(),
	}

	err = s.eventRepo.Save(ctx, event)
	if err != nil {
		fmt.Printf("Failed to save domain event: %v\n", err)
	}

	return &incident, nil
}

// ResolveIncident resolves an incident
func (s *ChangeManagementService) ResolveIncident(ctx context.Context, cmd ResolveIncidentCommand) error {
	incident, err := s.incidentRepo.FindByID(ctx, cmd.IncidentID)
	if err != nil {
		return fmt.Errorf("incident not found: %w", err)
	}

	if incident.Status == domain.IncidentStatusResolved || incident.Status == domain.IncidentStatusClosed {
		return fmt.Errorf("incident is already resolved or closed")
	}

	incident.Status = domain.IncidentStatusResolved
	incident.Resolution = cmd.Resolution
	incident.RootCause = cmd.RootCause
	incident.TimeToResolve = time.Since(incident.CreatedAt)
	incident.ResolvedAt = time.Now()
	incident.UpdatedAt = time.Now()

	err = s.incidentRepo.Update(ctx, incident)
	if err != nil {
		return fmt.Errorf("failed to resolve incident: %w", err)
	}

	// Publish domain event
	event := domain.IncidentResolvedEvent{
		IncidentID:    incident.ID,
		Resolver:      cmd.Resolver,
		Resolution:    cmd.Resolution,
		TimeToResolve: incident.TimeToResolve,
		OccurredAt:    time.Now(),
	}

	err = s.eventRepo.Save(ctx, event)
	if err != nil {
		fmt.Printf("Failed to save domain event: %v\n", err)
	}

	return nil
}

// CreateAudit creates a new audit
func (s *ChangeManagementService) CreateAudit(ctx context.Context, cmd CreateAuditCommand) (*domain.Audit, error) {
	// Verify application exists
	_, err := s.appRepo.FindByID(ctx, cmd.ApplicationID)
	if err != nil {
		return nil, fmt.Errorf("application not found: %w", err)
	}

	audit := domain.Audit{
		ID:            cmd.ID,
		ApplicationID: cmd.ApplicationID,
		Auditor:       cmd.Auditor,
		Type:          cmd.Type,
		Status:        domain.AuditStatusPlanned,
		Scope:         cmd.Scope,
		Findings:      []domain.AuditFinding{},
		StartedAt:     cmd.StartDate,
	}

	err = s.auditRepo.Save(ctx, audit)
	if err != nil {
		return nil, fmt.Errorf("failed to save audit: %w", err)
	}

	return &audit, nil
}

// CompleteAudit completes an audit
func (s *ChangeManagementService) CompleteAudit(ctx context.Context, cmd CompleteAuditCommand) error {
	audit, err := s.auditRepo.FindByID(ctx, cmd.AuditID)
	if err != nil {
		return fmt.Errorf("audit not found: %w", err)
	}

	if audit.Status != domain.AuditStatusInProgress {
		return fmt.Errorf("audit is not in progress")
	}

	audit.Status = domain.AuditStatusCompleted
	audit.CompletedAt = time.Now()
	audit.Findings = cmd.Findings
	audit.Recommendations = cmd.Recommendations

	err = s.auditRepo.Update(ctx, audit)
	if err != nil {
		return fmt.Errorf("failed to complete audit: %w", err)
	}

	// Convert findings to string slice for event
	findings := make([]string, len(cmd.Findings))
	for i, finding := range cmd.Findings {
		findings[i] = finding.Description
	}

	// Publish domain event
	event := domain.AuditCompletedEvent{
		AuditID:       audit.ID,
		ApplicationID: audit.ApplicationID,
		Auditor:       audit.Auditor,
		Scope:         audit.Scope,
		Findings:      findings,
		Status:        string(audit.Status),
		OccurredAt:    time.Now(),
	}

	err = s.eventRepo.Save(ctx, event)
	if err != nil {
		fmt.Printf("Failed to save domain event: %v\n", err)
	}

	return nil
}

// GetChangeRequestsByApplication retrieves change requests for an application
func (s *ChangeManagementService) GetChangeRequestsByApplication(ctx context.Context, appID domain.ApplicationID) ([]domain.ChangeRequest, error) {
	changeRequests, err := s.changeRequestRepo.FindByApplicationID(ctx, appID)
	if err != nil {
		return nil, fmt.Errorf("failed to get change requests: %w", err)
	}
	return changeRequests, nil
}

// GetIncidentsByApplication retrieves incidents for an application
func (s *ChangeManagementService) GetIncidentsByApplication(ctx context.Context, appID domain.ApplicationID) ([]domain.Incident, error) {
	incidents, err := s.incidentRepo.FindByApplicationID(ctx, appID)
	if err != nil {
		return nil, fmt.Errorf("failed to get incidents: %w", err)
	}
	return incidents, nil
}

// GetAuditsByApplication retrieves audits for an application
func (s *ChangeManagementService) GetAuditsByApplication(ctx context.Context, appID domain.ApplicationID) ([]domain.Audit, error) {
	audits, err := s.auditRepo.FindByApplicationID(ctx, appID)
	if err != nil {
		return nil, fmt.Errorf("failed to get audits: %w", err)
	}
	return audits, nil
}

// Commands for Change Management Service

type CreateChangeRequestCommand struct {
	ID            string
	ApplicationID domain.ApplicationID
	Requester     string
	Type          domain.ChangeType
	Priority      domain.Priority
	Title         string
	Description   string
	BusinessCase  string
	Impact        string
	Risk          string
}

type ApproveChangeRequestCommand struct {
	ChangeRequestID string
	Approver        string
	Role            string
	Comments        string
}

type RejectChangeRequestCommand struct {
	ChangeRequestID string
	Approver        string
	Role            string
	Comments        string
}

type ReportIncidentCommand struct {
	ID            string
	ApplicationID domain.ApplicationID
	Reporter      string
	Severity      int
	Title         string
	Description   string
	Impact        string
}

type ResolveIncidentCommand struct {
	IncidentID string
	Resolver   string
	Resolution string
	RootCause  string
}

type CreateAuditCommand struct {
	ID            string
	ApplicationID domain.ApplicationID
	Auditor       string
	Type          domain.AuditType
	Scope         string
	StartDate     time.Time
}

type CompleteAuditCommand struct {
	AuditID        string
	Findings       []domain.AuditFinding
	Recommendations []string
}
