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
	"strings"
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
	businessValue := s.assessBusinessValue(ctx, app)

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
	score := 3 // Base score

	// Analyze version maturity (semantic versioning indicates better practices)
	versionScore := s.analyzeVersionMaturity(app.Version)
	score += versionScore

	// Security provisions analysis
	securityScore := s.analyzeSecurityProvisions(app.SecurityProvisions)
	score += securityScore

	// Documentation and catalogue completeness
	documentationScore := s.analyzeDocumentationCompleteness(app.Catalogue)
	score += documentationScore

	// Age-based depreciation (older apps may have accumulated technical debt)
	ageScore := s.analyzeApplicationAge(app.CreatedAt, app.UpdatedAt)
	score += ageScore

	// Application status impact
	statusScore := s.analyzeApplicationStatus(app.Status)
	score += statusScore

	// Ensure score is within bounds
	if score < 1 {
		score = 1
	}
	if score > 5 {
		score = 5
	}

	// Calculate individual metrics based on overall score with some variance
	basePercentage := float64(score) * 20.0 // Base percentage

	return TechnicalHealth{
		CodeQuality:      s.adjustScoreWithVariance(score, 0.8, 1.2),
		Documentation:    s.adjustScoreWithVariance(score, 0.9, 1.1),
		TestCoverage:     basePercentage + float64(securityScore)*5.0, // Security affects testing
		SecurityScore:    s.adjustScoreWithVariance(score+securityScore, 0.7, 1.3),
		PerformanceScore: s.adjustScoreWithVariance(score+ageScore, 0.8, 1.2),
	}
}

// analyzeVersionMaturity evaluates version string for maturity indicators
func (s *EvaluationService) analyzeVersionMaturity(version string) int {
	if version == "" {
		return -1 // Penalty for no version
	}

	// Check for semantic versioning (major.minor.patch)
	parts := strings.Split(version, ".")
	if len(parts) >= 3 {
		// Semantic versioning indicates better development practices
		return 1
	}

	// Check for development/pre-release indicators
	lowerVersion := strings.ToLower(version)
	if strings.Contains(lowerVersion, "dev") ||
	   strings.Contains(lowerVersion, "alpha") ||
	   strings.Contains(lowerVersion, "beta") ||
	   strings.Contains(lowerVersion, "rc") {
		return 0 // Neutral for development versions
	}

	return 0 // Neutral for other version formats
}

// analyzeSecurityProvisions evaluates security measures in place
func (s *EvaluationService) analyzeSecurityProvisions(provisions SecurityProvisions) int {
	score := 0

	// Data confidentiality measures
	if len(provisions.DataConfidentiality) > 0 {
		score++
		if len(provisions.DataConfidentiality) > 2 {
			score++ // Bonus for comprehensive confidentiality
		}
	}

	// Data integrity measures
	if len(provisions.DataIntegrity) > 0 {
		score++
		if len(provisions.DataIntegrity) > 2 {
			score++ // Bonus for comprehensive integrity
		}
	}

	// Application authenticity measures
	if len(provisions.ApplicationAuthenticity) > 0 {
		score++
	}

	// Roles and permissions (access control)
	if len(provisions.RolesAndPermissions) > 0 {
		score++
		if len(provisions.RolesAndPermissions) > 3 {
			score++ // Bonus for comprehensive role management
		}
	}

	// SLA-based availability (indirect security measure)
	if provisions.ApplicationAvailability.ResponseTime > 0 {
		score++
	}

	return score - 2 // Normalize (subtract base expectation)
}

// analyzeDocumentationCompleteness evaluates documentation quality
func (s *EvaluationService) analyzeDocumentationCompleteness(catalogue ApplicationCatalogue) int {
	score := 0

	// Recent updates indicate active maintenance
	if !catalogue.LastUpdated.IsZero() {
		daysSinceUpdate := time.Since(catalogue.LastUpdated).Hours() / 24
		if daysSinceUpdate < 90 { // Updated within 3 months
			score += 2
		} else if daysSinceUpdate < 365 { // Updated within a year
			score++
		}
	} else {
		score-- // Penalty for no update date
	}

	// Comprehensive functionality documentation
	if len(catalogue.Functionality) > 0 {
		score++
		if len(catalogue.Functionality) > 5 {
			score++ // Bonus for detailed functionality
		}
	}

	return score
}

// analyzeApplicationAge evaluates age-related technical debt
func (s *EvaluationService) analyzeApplicationAge(createdAt, updatedAt time.Time) int {
	if createdAt.IsZero() {
		return 0 // No age data available
	}

	ageInDays := time.Since(createdAt).Hours() / 24

	// Very old applications may have accumulated technical debt
	if ageInDays > 365*5 { // Over 5 years old
		return -2
	} else if ageInDays > 365*2 { // Over 2 years old
		return -1
	}

	// Recently updated applications are better maintained
	if !updatedAt.IsZero() {
		daysSinceUpdate := time.Since(updatedAt).Hours() / 24
		if daysSinceUpdate < 90 { // Updated within 3 months
			return 1
		} else if daysSinceUpdate < 180 { // Updated within 6 months
			return 0
		}
	}

	return 0
}

// analyzeApplicationStatus evaluates status impact on technical health
func (s *EvaluationService) analyzeApplicationStatus(status ApplicationStatus) int {
	switch status {
	case StatusActive:
		return 1 // Active apps are well-maintained
	case StatusDeprecated:
		return -1 // Deprecated apps may have issues
	case StatusRetired:
		return -2 // Retired apps have significant issues
	case StatusPlanned:
		return 0 // Planned apps are new, no technical debt yet
	default:
		return 0
	}
}

// adjustScoreWithVariance adds realistic variance to scores
func (s *EvaluationService) adjustScoreWithVariance(baseScore int, minFactor, maxFactor float64) int {
	// Simple deterministic variance based on base score
	// In a real system, this could use random factors
	variance := (float64(baseScore) * 0.1) // 10% variance
	if variance > 0.5 {
		variance = 0.5
	}
	if variance < -0.5 {
		variance = -0.5
	}

	adjusted := float64(baseScore) + variance
	if adjusted < 1 {
		adjusted = 1
	}
	if adjusted > 5 {
		adjusted = 5
	}

	return int(adjusted + 0.5) // Round to nearest integer
}

// assessBusinessValue evaluates the business value of an application
func (s *EvaluationService) assessBusinessValue(ctx context.Context, app Application) BusinessValueAssessment {
	// Get governance agreement for business context
	var agreement *GovernanceAgreement
	if s.agreementRepo != nil {
		if govAgreement, err := s.agreementRepo.FindByApplicationID(ctx, app.ID); err == nil {
			agreement = &govAgreement
		}
	}

	// Calculate usage metrics based on application attributes
	usageMetrics := s.calculateUsageMetrics(app, agreement)

	// Calculate business alignment based on governance agreement
	businessAlignment := s.calculateBusinessAlignment(app, agreement)

	// Calculate cost efficiency based on application status and maintenance
	costEfficiency := s.calculateCostEfficiency(app, agreement)

	// Calculate user satisfaction based on application health and governance
	userSatisfaction := s.calculateUserSatisfaction(app, agreement)

	return BusinessValueAssessment{
		UsageMetrics:     usageMetrics,
		BusinessAlignment: businessAlignment,
		CostEfficiency:   costEfficiency,
		UserSatisfaction: userSatisfaction,
	}
}

// calculateUsageMetrics derives usage metrics from application attributes
func (s *EvaluationService) calculateUsageMetrics(app Application, agreement *GovernanceAgreement) UsageMetrics {
	// Base metrics derived from application characteristics
	activeUsers := 50   // Base active users
	transactionVolume := 1000 // Base transactions

	// Scale based on application status and governance
	switch app.Status {
	case StatusActive:
		activeUsers *= 2
		transactionVolume *= 3
	case StatusDeprecated:
		activeUsers /= 2
		transactionVolume /= 2
	case StatusRetired:
		activeUsers /= 4
		transactionVolume /= 4
	}

	// Governance agreement indicates higher usage
	if agreement != nil {
		activeUsers = int(float64(activeUsers) * 1.5)
		transactionVolume = int(float64(transactionVolume) * 1.8)
	}

	// Age affects usage patterns
	if !app.CreatedAt.IsZero() {
		ageInYears := time.Since(app.CreatedAt).Hours() / (24 * 365)
		if ageInYears > 3 {
			// Mature applications typically have higher usage
			activeUsers = int(float64(activeUsers) * 1.3)
			transactionVolume = int(float64(transactionVolume) * 1.4)
		}
	}

	// Calculate uptime based on technical health proxy
	uptimePercentage := 99.0 // Base uptime
	if len(app.SecurityProvisions.RolesAndPermissions) > 0 {
		uptimePercentage += 0.5 // Better security = better uptime
	}
	if !app.UpdatedAt.IsZero() && time.Since(app.UpdatedAt).Hours() < 24*30 {
		uptimePercentage += 0.4 // Recently updated = better maintenance
	}

	// Response time based on application complexity
	responseTime := time.Millisecond * 300 // Base response time
	if strings.Contains(strings.ToLower(app.Name), "legacy") {
		responseTime += time.Millisecond * 200 // Legacy systems slower
	}
	if len(app.SecurityProvisions.DataIntegrity) > 0 {
		responseTime += time.Millisecond * 50 // Security measures overhead
	}

	return UsageMetrics{
		ActiveUsers:       activeUsers,
		TransactionVolume: transactionVolume,
		UptimePercentage:  uptimePercentage,
		ResponseTime:      responseTime,
	}
}

// calculateBusinessAlignment evaluates how well the application aligns with business objectives
func (s *EvaluationService) calculateBusinessAlignment(app Application, agreement *GovernanceAgreement) float64 {
	baseAlignment := 70.0 // Base alignment score

	// Governance agreement significantly improves alignment
	if agreement != nil {
		baseAlignment += 20.0

		// Strategic objectives indicate strong alignment
		if len(agreement.Direct.StrategicDirection.Objectives) > 0 {
			baseAlignment += 5.0
		}

		// Active monitoring improves alignment
		if agreement.Conformance.ComplianceMonitoring.MonitoringFrequency != "" {
			baseAlignment += 5.0
		}
	}

	// Application status affects alignment
	switch app.Status {
	case StatusActive:
		baseAlignment += 5.0
	case StatusPlanned:
		baseAlignment += 2.0 // Future alignment
	case StatusDeprecated:
		baseAlignment -= 10.0 // Misalignment with current strategy
	case StatusRetired:
		baseAlignment -= 20.0 // No alignment
	}

	// Recent updates indicate current alignment
	if !app.UpdatedAt.IsZero() && time.Since(app.UpdatedAt).Hours() < 24*90 {
		baseAlignment += 3.0
	}

	// Ensure bounds
	if baseAlignment > 100.0 {
		baseAlignment = 100.0
	}
	if baseAlignment < 0.0 {
		baseAlignment = 0.0
	}

	return baseAlignment
}

// calculateCostEfficiency evaluates the cost effectiveness of the application
func (s *EvaluationService) calculateCostEfficiency(app Application, agreement *GovernanceAgreement) float64 {
	baseEfficiency := 60.0 // Base efficiency

	// Governance agreements improve cost efficiency through oversight
	if agreement != nil {
		baseEfficiency += 15.0

		// Resource allocation indicates cost management
		if len(agreement.Direct.ResourceAllocation.BudgetAllocations) > 0 {
			baseEfficiency += 10.0
		}
	}

	// Application status affects cost efficiency
	switch app.Status {
	case StatusActive:
		baseEfficiency += 10.0 // Active maintenance
	case StatusDeprecated:
		baseEfficiency -= 15.0 // High maintenance costs
	case StatusRetired:
		baseEfficiency -= 25.0 // Very high maintenance costs
	case StatusPlanned:
		baseEfficiency += 5.0 // Planned efficiency
	}

	// Age affects efficiency (older systems may be more expensive to maintain)
	if !app.CreatedAt.IsZero() {
		ageInYears := time.Since(app.CreatedAt).Hours() / (24 * 365)
		if ageInYears > 5 {
			baseEfficiency -= 10.0
		} else if ageInYears < 1 {
			baseEfficiency += 5.0 // New systems are more efficient
		}
	}

	// Security provisions may indicate higher quality (better efficiency)
	securityMeasures := len(app.SecurityProvisions.DataConfidentiality) +
					   len(app.SecurityProvisions.DataIntegrity) +
					   len(app.SecurityProvisions.RolesAndPermissions)
	if securityMeasures > 3 {
		baseEfficiency += 5.0
	}

	// Ensure bounds
	if baseEfficiency > 100.0 {
		baseEfficiency = 100.0
	}
	if baseEfficiency < 0.0 {
		baseEfficiency = 0.0
	}

	return baseEfficiency
}

// calculateUserSatisfaction estimates user satisfaction based on application factors
func (s *EvaluationService) calculateUserSatisfaction(app Application, agreement *GovernanceAgreement) float64 {
	baseSatisfaction := 65.0 // Base satisfaction

	// Governance agreement indicates better user experience management
	if agreement != nil {
		baseSatisfaction += 15.0

		// Performance metrics indicate user focus
		if len(agreement.Evaluate.PerformanceMetrics) > 0 {
			baseSatisfaction += 5.0
		}
	}

	// Application status affects user satisfaction
	switch app.Status {
	case StatusActive:
		baseSatisfaction += 10.0
	case StatusDeprecated:
		baseSatisfaction -= 15.0 // Users frustrated with deprecated systems
	case StatusRetired:
		baseSatisfaction -= 30.0 // Users very dissatisfied
	case StatusPlanned:
		baseSatisfaction += 5.0 // Anticipation of new system
	}

	// Recent updates indicate better user experience
	if !app.UpdatedAt.IsZero() && time.Since(app.UpdatedAt).Hours() < 24*60 {
		baseSatisfaction += 8.0
	}

	// Security features may affect perceived reliability
	if len(app.SecurityProvisions.RolesAndPermissions) > 0 {
		baseSatisfaction += 3.0
	}

	// Ensure bounds
	if baseSatisfaction > 100.0 {
		baseSatisfaction = 100.0
	}
	if baseSatisfaction < 0.0 {
		baseSatisfaction = 0.0
	}

	return baseSatisfaction
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
