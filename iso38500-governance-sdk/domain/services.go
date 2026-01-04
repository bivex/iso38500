/**
 * Copyright (c) 2026 Bivex
 *
 * Author: Bivex
 * Available for contact via email: support@b-b.top
 * For up-to-date contact information:
 * https://github.com/bivex
 *
 * Created: 2026-01-04T06:25:49
 * Last Updated: 2026-01-04T06:33:58
 *
 * Licensed under the MIT License.
 * Commercial licensing available upon request.
 */

package domain

import (
	"context"
	"fmt"
	"time"
)

// EvaluationService handles the evaluation principle of ISO 38500
type EvaluationService struct {
	applicationRepo ApplicationRepository
	agreementRepo   GovernanceAgreementRepository
	portfolioRepo   ApplicationPortfolioRepository
	kpiRepo         KPIRepository
	riskRepo        RiskRepository
}

// NewEvaluationService creates a new evaluation service
func NewEvaluationService(appRepo ApplicationRepository, agreementRepo GovernanceAgreementRepository, portfolioRepo ApplicationPortfolioRepository, kpiRepo KPIRepository, riskRepo RiskRepository) *EvaluationService {
	return &EvaluationService{
		applicationRepo: appRepo,
		agreementRepo:   agreementRepo,
		portfolioRepo:   portfolioRepo,
		kpiRepo:         kpiRepo,
		riskRepo:        riskRepo,
	}
}

// EvaluateApplication performs a comprehensive evaluation of an application
func (s *EvaluationService) EvaluateApplication(ctx context.Context, appID ApplicationID, evaluator string) (*ApplicationAssessment, error) {
	// Get application
	app, err := s.applicationRepo.FindByID(ctx, appID)
	if err != nil {
		return nil, fmt.Errorf("failed to find application: %w", err)
	}

	// Get governance agreement (not used in current implementation but may be needed for future enhancements)
	_, err = s.agreementRepo.FindByApplicationID(ctx, appID)
	if err != nil {
		return nil, fmt.Errorf("failed to find governance agreement: %w", err)
	}

	// Assess technical health
	technicalHealth := s.assessTechnicalHealth(app)

	// Assess business value
	businessValue := s.assessBusinessValue(app)

	// Determine risk level
	riskLevel := s.determineRiskLevel(technicalHealth, businessValue)

	// Generate recommendations
	recommendations := s.generateRecommendations(technicalHealth, businessValue, riskLevel)

	assessment := &ApplicationAssessment{
		ApplicationID:   appID,
		TechnicalHealth: technicalHealth,
		BusinessValue:   businessValue,
		RiskLevel:       riskLevel,
		Recommendations: recommendations,
	}

	return assessment, nil
}

// EvaluatePortfolio performs evaluation of the entire portfolio
func (s *EvaluationService) EvaluatePortfolio(ctx context.Context, portfolioID PortfolioID) (*PortfolioHealthAssessment, error) {
	// Get portfolio and its applications
	portfolio, err := s.portfolioRepo.FindByID(ctx, portfolioID)
	if err != nil {
		return nil, fmt.Errorf("failed to find portfolio: %w", err)
	}

	apps := portfolio.Applications

	totalApps := len(apps)
	activeApps := 0
	deprecatedApps := 0
	redundantApps := 0
	totalCost := 0.0
	riskDistribution := make(map[RiskLevel]int)

	assessments := make([]ApplicationAssessment, 0, totalApps)

	for _, app := range apps {
		assessment, err := s.EvaluateApplication(ctx, app.ID, "system")
		if err != nil {
			continue // Skip failed assessments
		}
		assessments = append(assessments, *assessment)

		// Count by status
		switch app.Status {
		case StatusActive:
			activeApps++
		case StatusDeprecated:
			deprecatedApps++
		case StatusRetired:
			// Retired apps don't count toward active metrics
		}

		riskDistribution[assessment.RiskLevel]++
	}

	// Calculate average age (simplified)
	avgAge := s.calculateAverageApplicationAge(apps)

	assessment := &PortfolioHealthAssessment{
		TotalApplications:     totalApps,
		ActiveApplications:    activeApps,
		DeprecatedApplications: deprecatedApps,
		RedundantApplications: redundantApps,
		TotalCost:            totalCost,
		AverageApplicationAge: avgAge,
		RiskDistribution:     riskDistribution,
	}

	return assessment, nil
}

// assessTechnicalHealth evaluates the technical health of an application
func (s *EvaluationService) assessTechnicalHealth(app Application) TechnicalHealth {
	// Simplified scoring logic - in real implementation, this would analyze
	// code quality metrics, test coverage, security scans, etc.
	score := 3 // Default neutral score

	// Factor in security provisions
	if len(app.SecurityProvisions.DataConfidentiality) > 0 {
		score++
	}
	if len(app.SecurityProvisions.DataIntegrity) > 0 {
		score++
	}

	// Factor in documentation
	if app.Catalogue.LastUpdated.IsZero() {
		score--
	}

	// Ensure score is within bounds
	if score < 1 {
		score = 1
	}
	if score > 5 {
		score = 5
	}

	return TechnicalHealth{
		CodeQuality:      score,
		Documentation:    score,
		TestCoverage:     float64(score) * 20, // Convert to percentage
		SecurityScore:    score,
		PerformanceScore: score,
	}
}

// assessBusinessValue evaluates the business value of an application
func (s *EvaluationService) assessBusinessValue(app Application) BusinessValueAssessment {
	// Simplified assessment - in real implementation, this would analyze
	// usage metrics, business alignment, cost efficiency, etc.
	return BusinessValueAssessment{
		UsageMetrics: UsageMetrics{
			ActiveUsers:       100, // Mock data
			TransactionVolume: 1000,
			UptimePercentage:  99.9,
			ResponseTime:      time.Millisecond * 200,
		},
		BusinessAlignment: 85.0,
		CostEfficiency:    75.0,
		UserSatisfaction:  80.0,
	}
}

// determineRiskLevel calculates the overall risk level
func (s *EvaluationService) determineRiskLevel(techHealth TechnicalHealth, businessValue BusinessValueAssessment) RiskLevel {
	avgScore := (techHealth.CodeQuality + techHealth.SecurityScore + techHealth.PerformanceScore) / 3

	if avgScore <= 2 || businessValue.CostEfficiency < 50 {
		return RiskCritical
	}
	if avgScore <= 3 || businessValue.CostEfficiency < 70 {
		return RiskHigh
	}
	if avgScore <= 4 {
		return RiskMedium
	}
	return RiskLow
}

// generateRecommendations creates recommendations based on assessment
func (s *EvaluationService) generateRecommendations(techHealth TechnicalHealth, businessValue BusinessValueAssessment, riskLevel RiskLevel) []Recommendation {
	recommendations := []Recommendation{}

	if techHealth.SecurityScore < 3 {
		recommendations = append(recommendations, Recommendation{
			ID:             "sec-001",
			Type:           RecModernize,
			Description:    "Improve security measures and implement additional security controls",
			Priority:       PriorityHigh,
			EstimatedEffort: time.Hour * 80,
			BusinessImpact:  "Reduce security risks and ensure compliance",
		})
	}

	if techHealth.CodeQuality < 3 {
		recommendations = append(recommendations, Recommendation{
			ID:             "tech-001",
			Type:           RecEnhance,
			Description:    "Refactor code to improve quality and maintainability",
			Priority:       PriorityMedium,
			EstimatedEffort: time.Hour * 120,
			BusinessImpact:  "Reduce technical debt and improve development velocity",
		})
	}

	if businessValue.CostEfficiency < 70 {
		recommendations = append(recommendations, Recommendation{
			ID:             "cost-001",
			Type:           RecReplace,
			Description:    "Evaluate more cost-effective alternatives",
			Priority:       PriorityMedium,
			EstimatedEffort: time.Hour * 40,
			BusinessImpact:  "Reduce operational costs",
		})
	}

	if riskLevel == RiskCritical {
		recommendations = append(recommendations, Recommendation{
			ID:             "risk-001",
			Type:           RecRetire,
			Description:    "Consider retiring or replacing this high-risk application",
			Priority:       PriorityCritical,
			EstimatedEffort: time.Hour * 160,
			BusinessImpact:  "Eliminate critical business and technical risks",
		})
	}

	return recommendations
}

// calculateAverageApplicationAge calculates the average age of applications
func (s *EvaluationService) calculateAverageApplicationAge(apps []Application) time.Duration {
	if len(apps) == 0 {
		return 0
	}

	totalAge := time.Duration(0)
	for _, app := range apps {
		age := time.Since(app.CreatedAt)
		totalAge += age
	}

	return totalAge / time.Duration(len(apps))
}

// DirectionService handles the direction principle of ISO 38500
type DirectionService struct {
	agreementRepo GovernanceAgreementRepository
}

// NewDirectionService creates a new direction service
func NewDirectionService(agreementRepo GovernanceAgreementRepository) *DirectionService {
	return &DirectionService{
		agreementRepo: agreementRepo,
	}
}

// SetStrategicDirection establishes strategic direction for governance
func (s *DirectionService) SetStrategicDirection(ctx context.Context, agreementID GovernanceAgreementID, director string, objectives []StrategicObjective, initiatives []StrategicInitiative) error {
	agreement, err := s.agreementRepo.FindByID(ctx, agreementID)
	if err != nil {
		return fmt.Errorf("failed to find governance agreement: %w", err)
	}

	// Update the direct principle
	agreement.Direct.StrategicDirection.Objectives = objectives
	agreement.Direct.StrategicDirection.Initiatives = initiatives
	agreement.Direct.LastDirected = time.Now()

	// Create action plans from objectives
	actionPlans := s.createActionPlansFromObjectives(objectives)
	agreement.Direct.ActionPlans = actionPlans

	err = s.agreementRepo.Update(ctx, agreement)
	if err != nil {
		return fmt.Errorf("failed to update governance agreement: %w", err)
	}

	return nil
}

// AllocateResources allocates resources for governance activities
func (s *DirectionService) AllocateResources(ctx context.Context, agreementID GovernanceAgreementID, budgetAllocations []BudgetAllocation, personnelAllocations []PersonnelAllocation) error {
	agreement, err := s.agreementRepo.FindByID(ctx, agreementID)
	if err != nil {
		return fmt.Errorf("failed to find governance agreement: %w", err)
	}

	agreement.Direct.ResourceAllocation.BudgetAllocations = budgetAllocations
	agreement.Direct.ResourceAllocation.PersonnelAllocations = personnelAllocations
	agreement.Direct.LastDirected = time.Now()

	err = s.agreementRepo.Update(ctx, agreement)
	if err != nil {
		return fmt.Errorf("failed to update governance agreement: %w", err)
	}

	return nil
}

// EstablishPolicies establishes governance policies and standards
func (s *DirectionService) EstablishPolicies(ctx context.Context, agreementID GovernanceAgreementID, policies []Policy, standards []Standard, procedures []Procedure) error {
	agreement, err := s.agreementRepo.FindByID(ctx, agreementID)
	if err != nil {
		return fmt.Errorf("failed to find governance agreement: %w", err)
	}

	agreement.Direct.PolicyFramework.Policies = policies
	agreement.Direct.PolicyFramework.Standards = standards
	agreement.Direct.PolicyFramework.Procedures = procedures

	err = s.agreementRepo.Update(ctx, agreement)
	if err != nil {
		return fmt.Errorf("failed to update governance agreement: %w", err)
	}

	return nil
}

// createActionPlansFromObjectives creates action plans from strategic objectives
func (s *DirectionService) createActionPlansFromObjectives(objectives []StrategicObjective) []ActionPlan {
	actionPlans := make([]ActionPlan, len(objectives))

	for i, objective := range objectives {
		actionPlans[i] = ActionPlan{
			ID:          fmt.Sprintf("ap-%d", i+1),
			Name:        fmt.Sprintf("Action Plan for %s", objective.Name),
			Description: fmt.Sprintf("Implementation plan for strategic objective: %s", objective.Description),
			Owner:       "TBD", // To be determined
			Deadline:    objective.Deadline,
			Status:      ActionPending,
			Actions: []Action{
				{
					ID:          fmt.Sprintf("action-%d-1", i+1),
					Description: "Define detailed implementation steps",
					Responsible: "TBD",
					Deadline:    objective.Deadline.AddDate(0, 0, -30), // 30 days before objective deadline
					Status:      ActionPending,
				},
			},
		}
	}

	return actionPlans
}

// MonitoringService handles the monitoring principle of ISO 38500
type MonitoringService struct {
	kpiRepo         KPIRepository
	measurementRepo KPIMeasurementRepository
	riskRepo        RiskRepository
	agreementRepo   GovernanceAgreementRepository
}

// NewMonitoringService creates a new monitoring service
func NewMonitoringService(kpiRepo KPIRepository, measurementRepo KPIMeasurementRepository, riskRepo RiskRepository, agreementRepo GovernanceAgreementRepository) *MonitoringService {
	return &MonitoringService{
		kpiRepo:         kpiRepo,
		measurementRepo: measurementRepo,
		riskRepo:        riskRepo,
		agreementRepo:   agreementRepo,
	}
}

// MonitorKPIs monitors KPI performance
func (s *MonitoringService) MonitorKPIs(ctx context.Context, agreementID GovernanceAgreementID) ([]KPIMeasurement, error) {
	// Get agreement to find associated KPIs (not used in current implementation but may be needed for future enhancements)
	_, err := s.agreementRepo.FindByID(ctx, agreementID)
	if err != nil {
		return nil, fmt.Errorf("failed to find governance agreement: %w", err)
	}

	// Handle case where repositories are not available (e.g., in demo mode)
	if s.kpiRepo == nil || s.measurementRepo == nil {
		// Return mock data for demonstration
		return []KPIMeasurement{
			{
				KPIID:     "kpi-001",
				Value:     95.5,
				Target:    100.0,
				Achieved:  false,
				MeasuredAt: time.Now(),
				Notes:     "Demo KPI measurement",
			},
			{
				KPIID:     "kpi-002",
				Value:     99.2,
				Target:    98.0,
				Achieved:  true,
				MeasuredAt: time.Now(),
				Notes:     "Demo KPI measurement",
			},
		}, nil
	}

	// For portfolio-level agreements, get all KPIs
	kpis, err := s.kpiRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find KPIs: %w", err)
	}

	measurements := []KPIMeasurement{}

	for _, kpi := range kpis {
		// Get latest measurement
		measurement, err := s.measurementRepo.FindLatest(ctx, kpi.ID)
		if err != nil {
			// Create default measurement if none exists
			measurement = KPIMeasurement{
				KPIID:     kpi.ID,
				Value:     0,
				Target:    kpi.Target,
				Achieved:  false,
				MeasuredAt: time.Now(),
				Notes:     "No measurement available",
			}
		}

		// Update achievement status
		measurement.Achieved = s.isKPITargetAchieved(kpi, measurement)
		measurements = append(measurements, measurement)
	}

	return measurements, nil
}

// MonitorCompliance monitors compliance status
func (s *MonitoringService) MonitorCompliance(ctx context.Context, agreementID GovernanceAgreementID) (*ComplianceMonitoring, error) {
	agreement, err := s.agreementRepo.FindByID(ctx, agreementID)
	if err != nil {
		return nil, fmt.Errorf("failed to find governance agreement: %w", err)
	}

	// Return the compliance monitoring configuration from the agreement
	return &agreement.Conformance.ComplianceMonitoring, nil
}

// MonitorRisks monitors risk status
func (s *MonitoringService) MonitorRisks(ctx context.Context, agreementID GovernanceAgreementID) (*RiskMonitoring, error) {
	// Handle case where risk repository is not available (e.g., in demo mode)
	if s.riskRepo == nil {
		// Return mock risk monitoring data for demonstration
		return &RiskMonitoring{
			RiskIndicators: []RiskIndicator{
				{
					Name:     "Technical Debt",
					Value:    75.0,
					Threshold: 80.0,
					Status:   RiskStatusWarning,
				},
				{
					Name:     "Security Vulnerabilities",
					Value:    25.0,
					Threshold: 50.0,
					Status:   RiskStatusNormal,
				},
			},
			RiskHeatMaps:   []RiskHeatMap{},
			MitigationTracking: []MitigationTracking{},
		}, nil
	}

	risks, err := s.riskRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find risks: %w", err)
	}

	riskIndicators := make([]RiskIndicator, len(risks))
	for i, risk := range risks {
		riskIndicators[i] = RiskIndicator{
			Name:     risk.Name,
			Value:    float64(risk.Probability) * s.convertImpactToNumeric(risk.Impact),
			Threshold: s.getRiskThreshold(risk.Level),
			Status:   s.determineRiskStatus(risk),
		}
	}

	riskMonitoring := &RiskMonitoring{
		RiskIndicators: riskIndicators,
		RiskHeatMaps:   []RiskHeatMap{}, // Would be populated with actual heat map data
		MitigationTracking: []MitigationTracking{}, // Would be populated with actual tracking data
	}

	return riskMonitoring, nil
}

// isKPITargetAchieved determines if a KPI target is achieved
func (s *MonitoringService) isKPITargetAchieved(kpi KPI, measurement KPIMeasurement) bool {
	// Simplified logic - in real implementation, this would consider KPI type and thresholds
	switch kpi.Category {
	case "performance":
		return measurement.Value >= kpi.Target
	case "efficiency":
		return measurement.Value <= kpi.Target // Lower is better for efficiency
	default:
		return measurement.Value >= kpi.Target
	}
}

// convertImpactToNumeric converts risk impact to numeric value
func (s *MonitoringService) convertImpactToNumeric(impact RiskImpact) float64 {
	switch impact {
	case ImpactLow:
		return 1.0
	case ImpactMedium:
		return 2.0
	case ImpactHigh:
		return 3.0
	case ImpactCritical:
		return 4.0
	default:
		return 1.0
	}
}

// getRiskThreshold returns the threshold for a risk level
func (s *MonitoringService) getRiskThreshold(level RiskLevel) float64 {
	switch level {
	case RiskLow:
		return 2.0
	case RiskMedium:
		return 4.0
	case RiskHigh:
		return 8.0
	case RiskCritical:
		return 12.0
	default:
		return 2.0
	}
}

// determineRiskStatus determines the current risk status
func (s *MonitoringService) determineRiskStatus(risk Risk) RiskStatus {
	currentValue := float64(risk.Probability) * s.convertImpactToNumeric(risk.Impact)
	threshold := s.getRiskThreshold(risk.Level)

	if currentValue >= threshold*1.5 {
		return RiskStatusCritical
	}
	if currentValue >= threshold {
		return RiskStatusWarning
	}
	return RiskStatusNormal
}
