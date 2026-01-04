package application

import (
	"context"
	"fmt"
	"time"

	"github.com/iso38500/iso38500-governance-sdk/domain"
)

// GovernanceService provides application services for governance management
type GovernanceService struct {
	agreementRepo  domain.GovernanceAgreementRepository
	appRepo        domain.ApplicationRepository
	eventRepo      domain.DomainEventRepository
	evalService    *domain.EvaluationService
	directService  *domain.DirectionService
	monitorService *domain.MonitoringService
}

// NewGovernanceService creates a new governance service
func NewGovernanceService(
	agreementRepo domain.GovernanceAgreementRepository,
	appRepo domain.ApplicationRepository,
	eventRepo domain.DomainEventRepository,
	evalService *domain.EvaluationService,
	directService *domain.DirectionService,
	monitorService *domain.MonitoringService,
) *GovernanceService {
	return &GovernanceService{
		agreementRepo:  agreementRepo,
		appRepo:        appRepo,
		eventRepo:      eventRepo,
		evalService:    evalService,
		directService:  directService,
		monitorService: monitorService,
	}
}

// CreateGovernanceAgreement creates a new governance agreement
func (s *GovernanceService) CreateGovernanceAgreement(ctx context.Context, cmd CreateGovernanceAgreementCommand) (*domain.GovernanceAgreement, error) {
	// Verify application exists
	_, err := s.appRepo.FindByID(ctx, cmd.ApplicationID)
	if err != nil {
		return nil, fmt.Errorf("application not found: %w", err)
	}

	// Create aggregate
	aggregate, err := domain.NewGovernanceAgreementAggregate(cmd.ID, cmd.ApplicationID, cmd.Title)
	if err != nil {
		return nil, fmt.Errorf("failed to create governance agreement aggregate: %w", err)
	}

	// Save to repository
	agreement := aggregate.GetAgreement()
	err = s.agreementRepo.Save(ctx, agreement)
	if err != nil {
		return nil, fmt.Errorf("failed to save governance agreement: %w", err)
	}

	// Save domain events
	for _, event := range aggregate.GetDomainEvents() {
		err = s.eventRepo.Save(ctx, event)
		if err != nil {
			fmt.Printf("Failed to save domain event: %v\n", err)
		}
	}

	return &agreement, nil
}

// UpdateStrategy updates the strategy component of a governance agreement
func (s *GovernanceService) UpdateStrategy(ctx context.Context, cmd UpdateStrategyCommand) error {
	agreement, err := s.agreementRepo.FindByID(ctx, cmd.AgreementID)
	if err != nil {
		return fmt.Errorf("governance agreement not found: %w", err)
	}

	agreement.Strategy = cmd.Strategy

	err = s.agreementRepo.Update(ctx, agreement)
	if err != nil {
		return fmt.Errorf("failed to update strategy: %w", err)
	}

	return nil
}

// UpdateAcquisition updates the acquisition component of a governance agreement
func (s *GovernanceService) UpdateAcquisition(ctx context.Context, cmd UpdateAcquisitionCommand) error {
	agreement, err := s.agreementRepo.FindByID(ctx, cmd.AgreementID)
	if err != nil {
		return fmt.Errorf("governance agreement not found: %w", err)
	}

	agreement.Acquisition = cmd.Acquisition

	err = s.agreementRepo.Update(ctx, agreement)
	if err != nil {
		return fmt.Errorf("failed to update acquisition: %w", err)
	}

	return nil
}

// UpdatePerformance updates the performance component of a governance agreement
func (s *GovernanceService) UpdatePerformance(ctx context.Context, cmd UpdatePerformanceCommand) error {
	agreement, err := s.agreementRepo.FindByID(ctx, cmd.AgreementID)
	if err != nil {
		return fmt.Errorf("governance agreement not found: %w", err)
	}

	agreement.Performance = cmd.Performance

	err = s.agreementRepo.Update(ctx, agreement)
	if err != nil {
		return fmt.Errorf("failed to update performance: %w", err)
	}

	return nil
}

// UpdateConformance updates the conformance component of a governance agreement
func (s *GovernanceService) UpdateConformance(ctx context.Context, cmd UpdateConformanceCommand) error {
	agreement, err := s.agreementRepo.FindByID(ctx, cmd.AgreementID)
	if err != nil {
		return fmt.Errorf("governance agreement not found: %w", err)
	}

	agreement.Conformance = cmd.Conformance

	err = s.agreementRepo.Update(ctx, agreement)
	if err != nil {
		return fmt.Errorf("failed to update conformance: %w", err)
	}

	return nil
}

// UpdateImplementation updates the implementation component of a governance agreement
func (s *GovernanceService) UpdateImplementation(ctx context.Context, cmd UpdateImplementationCommand) error {
	agreement, err := s.agreementRepo.FindByID(ctx, cmd.AgreementID)
	if err != nil {
		return fmt.Errorf("governance agreement not found: %w", err)
	}

	agreement.Implementation = cmd.Implementation

	err = s.agreementRepo.Update(ctx, agreement)
	if err != nil {
		return fmt.Errorf("failed to update implementation: %w", err)
	}

	return nil
}

// ApproveGovernanceAgreement approves a governance agreement
func (s *GovernanceService) ApproveGovernanceAgreement(ctx context.Context, cmd ApproveGovernanceAgreementCommand) error {
	// Get agreement
	agreement, err := s.agreementRepo.FindByID(ctx, cmd.AgreementID)
	if err != nil {
		return fmt.Errorf("governance agreement not found: %w", err)
	}

	if agreement.Status != domain.AgreementDraft {
		return fmt.Errorf("only draft agreements can be approved")
	}

	// Update agreement status
	agreement.Status = domain.AgreementApproved
	agreement.UpdatedAt = time.Now()

	err = s.agreementRepo.Update(ctx, agreement)
	if err != nil {
		return fmt.Errorf("failed to save approved agreement: %w", err)
	}

	// Publish domain event
	event := domain.GovernanceAgreementApprovedEvent{
		AgreementID: cmd.AgreementID,
		OccurredAt:  time.Now(),
	}

	err = s.eventRepo.Save(ctx, event)
	if err != nil {
		fmt.Printf("Failed to save domain event: %v\n", err)
	}

	return nil
}

// ActivateGovernanceAgreement activates a governance agreement
func (s *GovernanceService) ActivateGovernanceAgreement(ctx context.Context, cmd ActivateGovernanceAgreementCommand) error {
	// Get agreement
	agreement, err := s.agreementRepo.FindByID(ctx, cmd.AgreementID)
	if err != nil {
		return fmt.Errorf("governance agreement not found: %w", err)
	}

	if agreement.Status != domain.AgreementApproved {
		return fmt.Errorf("only approved agreements can be activated")
	}

	// Update agreement status
	agreement.Status = domain.AgreementActive
	agreement.UpdatedAt = time.Now()

	err = s.agreementRepo.Update(ctx, agreement)
	if err != nil {
		return fmt.Errorf("failed to save activated agreement: %w", err)
	}

	// Publish domain event
	event := domain.GovernanceAgreementActivatedEvent{
		AgreementID: cmd.AgreementID,
		OccurredAt:  time.Now(),
	}

	err = s.eventRepo.Save(ctx, event)
	if err != nil {
		fmt.Printf("Failed to save domain event: %v\n", err)
	}

	return nil
}

// EvaluateApplication performs evaluation of an application
func (s *GovernanceService) EvaluateApplication(ctx context.Context, cmd EvaluateApplicationCommand) (*domain.ApplicationAssessment, error) {
	assessment, err := s.evalService.EvaluateApplication(ctx, cmd.ApplicationID, cmd.Evaluator)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate application: %w", err)
	}

	return assessment, nil
}

// EvaluatePortfolio performs evaluation of a portfolio
func (s *GovernanceService) EvaluatePortfolio(ctx context.Context, cmd EvaluatePortfolioCommand) (*domain.PortfolioHealthAssessment, error) {
	assessment, err := s.evalService.EvaluatePortfolio(ctx, cmd.PortfolioID)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate portfolio: %w", err)
	}

	return assessment, nil
}

// SetStrategicDirection sets strategic direction for governance
func (s *GovernanceService) SetStrategicDirection(ctx context.Context, cmd SetStrategicDirectionCommand) error {
	err := s.directService.SetStrategicDirection(ctx, cmd.AgreementID, cmd.Director, cmd.Objectives, cmd.Initiatives)
	if err != nil {
		return fmt.Errorf("failed to set strategic direction: %w", err)
	}

	return nil
}

// AllocateResources allocates resources for governance activities
func (s *GovernanceService) AllocateResources(ctx context.Context, cmd AllocateResourcesCommand) error {
	err := s.directService.AllocateResources(ctx, cmd.AgreementID, cmd.BudgetAllocations, cmd.PersonnelAllocations)
	if err != nil {
		return fmt.Errorf("failed to allocate resources: %w", err)
	}

	return nil
}

// EstablishPolicies establishes governance policies and standards
func (s *GovernanceService) EstablishPolicies(ctx context.Context, cmd EstablishPoliciesCommand) error {
	err := s.directService.EstablishPolicies(ctx, cmd.AgreementID, cmd.Policies, cmd.Standards, cmd.Procedures)
	if err != nil {
		return fmt.Errorf("failed to establish policies: %w", err)
	}

	return nil
}

// MonitorGovernance monitors governance activities
func (s *GovernanceService) MonitorGovernance(ctx context.Context, cmd MonitorGovernanceCommand) (*GovernanceMonitoringResult, error) {
	// Monitor KPIs
	kpiMeasurements, err := s.monitorService.MonitorKPIs(ctx, cmd.AgreementID)
	if err != nil {
		return nil, fmt.Errorf("failed to monitor KPIs: %w", err)
	}

	// Monitor compliance
	compliance, err := s.monitorService.MonitorCompliance(ctx, cmd.AgreementID)
	if err != nil {
		return nil, fmt.Errorf("failed to monitor compliance: %w", err)
	}

	// Monitor risks
	risks, err := s.monitorService.MonitorRisks(ctx, cmd.AgreementID)
	if err != nil {
		return nil, fmt.Errorf("failed to monitor risks: %w", err)
	}

	result := &GovernanceMonitoringResult{
		KPIMeasurements:   kpiMeasurements,
		ComplianceStatus:  compliance,
		RiskStatus:        risks,
	}

	return result, nil
}

// GetGovernanceAgreement retrieves a governance agreement by ID
func (s *GovernanceService) GetGovernanceAgreement(ctx context.Context, agreementID domain.GovernanceAgreementID) (*domain.GovernanceAgreement, error) {
	agreement, err := s.agreementRepo.FindByID(ctx, agreementID)
	if err != nil {
		return nil, fmt.Errorf("failed to get governance agreement: %w", err)
	}
	return &agreement, nil
}

// ListGovernanceAgreements retrieves all governance agreements
func (s *GovernanceService) ListGovernanceAgreements(ctx context.Context) ([]domain.GovernanceAgreement, error) {
	agreements, err := s.agreementRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list governance agreements: %w", err)
	}
	return agreements, nil
}

// Commands for Governance Service

type CreateGovernanceAgreementCommand struct {
	ID            domain.GovernanceAgreementID
	ApplicationID domain.ApplicationID
	Title         string
}

type UpdateStrategyCommand struct {
	AgreementID domain.GovernanceAgreementID
	Strategy    domain.Strategy
}

type UpdateAcquisitionCommand struct {
	AgreementID domain.GovernanceAgreementID
	Acquisition domain.Acquisition
}

type UpdatePerformanceCommand struct {
	AgreementID    domain.GovernanceAgreementID
	Performance    domain.Performance
}

type UpdateConformanceCommand struct {
	AgreementID domain.GovernanceAgreementID
	Conformance domain.Conformance
}

type UpdateImplementationCommand struct {
	AgreementID    domain.GovernanceAgreementID
	Implementation domain.Implementation
}

type ApproveGovernanceAgreementCommand struct {
	AgreementID domain.GovernanceAgreementID
}

type ActivateGovernanceAgreementCommand struct {
	AgreementID domain.GovernanceAgreementID
}

type EvaluateApplicationCommand struct {
	ApplicationID domain.ApplicationID
	Evaluator     string
}

type EvaluatePortfolioCommand struct {
	PortfolioID domain.PortfolioID
}

type SetStrategicDirectionCommand struct {
	AgreementID domain.GovernanceAgreementID
	Director    string
	Objectives  []domain.StrategicObjective
	Initiatives []domain.StrategicInitiative
}

type AllocateResourcesCommand struct {
	AgreementID          domain.GovernanceAgreementID
	BudgetAllocations    []domain.BudgetAllocation
	PersonnelAllocations []domain.PersonnelAllocation
}

type EstablishPoliciesCommand struct {
	AgreementID domain.GovernanceAgreementID
	Policies    []domain.Policy
	Standards   []domain.Standard
	Procedures  []domain.Procedure
}

type MonitorGovernanceCommand struct {
	AgreementID domain.GovernanceAgreementID
}

type GovernanceMonitoringResult struct {
	KPIMeasurements  []domain.KPIMeasurement
	ComplianceStatus *domain.ComplianceMonitoring
	RiskStatus       *domain.RiskMonitoring
}
