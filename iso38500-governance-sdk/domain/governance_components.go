/**
 * Copyright (c) 2026 Bivex
 *
 * Author: Bivex
 * Available for contact via email: support@b-b.top
 * For up-to-date contact information:
 * https://github.com/bivex
 *
 * Created: 2026-01-04T06:25:16
 * Last Updated: 2026-01-04T06:25:16
 *
 * Licensed under the MIT License.
 * Commercial licensing available upon request.
 */

package domain

import "time"

// Strategy represents the strategic governance components
type Strategy struct {
	ICTOperationsManual    ICTOperationsManual
	ApplicationCatalogue   ApplicationCatalogue
	ApplicationInterfaces  []ApplicationInterface
	ConfigurationStandard  ConfigurationStandard
}

// ICTOperationsManual represents the technical operations manual
type ICTOperationsManual struct {
	ApplicationArchitecture string
	InfrastructureConfig    string
	OperatingSystem        string
	ProgrammingLanguage    string
	RightsAndRoles         []RolePermission
	SecurityProvisions     SecurityProvisions
	LastUpdated           time.Time
}

// Acquisition represents acquisition and requirements management
type Acquisition struct {
	RequirementsManagement RequirementsManagement
	CommunicationManagement CommunicationManagement
	BusinessCaseTemplate   string
	PrioritizationMatrix   []PrioritizationRule
	ChangeRequestProcess  ChangeRequestProcess
}

// RequirementsManagement represents the requirements management process
type RequirementsManagement struct {
	GatheringProcess   []RequirementStep
	ValidationProcess  []RequirementStep
	ApprovalWorkflow   []ApprovalStep
	BusinessRules      []BusinessRule
}

// RequirementStep represents a step in the requirements process
type RequirementStep struct {
	StepNumber  int
	Name        string
	Description string
	Responsible string
}

// ApprovalStep represents a step in the approval workflow
type ApprovalStep struct {
	StepNumber   int
	Name         string
	ApproverRole string
	Conditions   string
}

// BusinessRule represents a business rule for requirements
type BusinessRule struct {
	ID          string
	Name        string
	Description string
	Category    string
}

// CommunicationManagement represents communication processes
type CommunicationManagement struct {
	Stakeholders         []Stakeholder
	CommunicationMatrix  ResponsibilityMatrix
	CommunicationTypes   []CommunicationType
	CommunicationSchedule string
}

// Stakeholder represents a stakeholder in the communication process
type Stakeholder struct {
	Name     string
	Role     string
	Contact  string
	RACIRole string // R, A, C, or I
}

// CommunicationType represents a type of communication
type CommunicationType struct {
	Type        string
	Description string
	Frequency    string
	Audience     string
}

// PrioritizationRule represents a rule for prioritizing change requests
type PrioritizationRule struct {
	Criteria    string
	Weight      int
	Description string
}

// ChangeRequestProcess represents the change request process
type ChangeRequestProcess struct {
	Types         []ChangeType
	ApprovalMatrix ResponsibilityMatrix
	EscalationMatrix []EscalationLevel
	SLA           SLA
}

// ChangeType represents a type of change request
type ChangeType string

const (
	ChangeStandard ChangeType = "standard"
	ChangeNormal   ChangeType = "normal"
	ChangeEmergency ChangeType = "emergency"
)

// Performance represents performance management components
type Performance struct {
	SupportProcess      SupportProcess
	IncidentManagement  IncidentManagement
	EscalationProcess   []EscalationLevel
	ApplicationSecurity SecurityProvisions
	BusinessContinuity  BusinessContinuity
}

// SupportProcess represents the application support process
type SupportProcess struct {
	Level1Support []string
	Level2Support []string
	Level3Support []string
	SLA          SLA
}

// IncidentManagement represents incident management processes
type IncidentManagement struct {
	ClassificationMatrix []IncidentClass
	PrioritizationMatrix []IncidentPriority
	ResponseMatrix      []IncidentResponse
}

// IncidentClass represents an incident classification
type IncidentClass struct {
	Severity    int
	Name        string
	Description string
	ResponseTime time.Duration
}

// IncidentPriority represents incident prioritization
type IncidentPriority struct {
	Priority     int
	Name         string
	Description  string
	SLA         time.Duration
}

// IncidentResponse represents an incident response action
type IncidentResponse struct {
	IncidentClass string
	Action        string
	Responsible   string
	Timeframe     time.Duration
}

// Conformance represents conformance to standards and regulations
type Conformance struct {
	LegalRequirements    []LegalRequirement
	ContractualRequirements []ContractualRequirement
	IndustryStandards    []IndustryStandard
	ComplianceMonitoring ComplianceMonitoring
}

// LegalRequirement represents a legal requirement
type LegalRequirement struct {
	Name        string
	Description string
	Authority   string
	EffectiveDate time.Time
	Status      ComplianceStatus
}

// ContractualRequirement represents a contractual requirement
type ContractualRequirement struct {
	Name        string
	Description string
	ContractID  string
	Party       string
	Status      ComplianceStatus
}

// IndustryStandard represents an industry standard requirement
type IndustryStandard struct {
	Name        string
	Description string
	Organization string
	Version     string
	Status      ComplianceStatus
}

// ComplianceStatus represents the compliance status
type ComplianceStatus string

const (
	ComplianceCompliant    ComplianceStatus = "compliant"
	ComplianceNonCompliant ComplianceStatus = "non_compliant"
	CompliancePartial      ComplianceStatus = "partial"
	ComplianceUnderReview  ComplianceStatus = "under_review"
)

// ComplianceMonitoring represents compliance monitoring processes
type ComplianceMonitoring struct {
	MonitoringFrequency string
	ResponsibleParties  []string
	ReportingSchedule   string
	AuditRequirements   []AuditRequirement
}

// AuditRequirement represents an audit requirement
type AuditRequirement struct {
	Name         string
	Description  string
	Frequency     string
	Responsible   string
	LastAudit     time.Time
	NextAudit     time.Time
}

// Implementation represents implementation and deployment processes
type Implementation struct {
	ImplementationProcess ImplementationProcess
	ReleaseManagement     ReleaseManagement
	DeploymentStrategy    DeploymentStrategy
}

// ImplementationProcess represents the application implementation process
type ImplementationProcess struct {
	Phases          []ImplementationPhase
	Roles           ResponsibilityMatrix
	QualityGates    []QualityGate
	RollbackPlan    string
}

// ImplementationPhase represents a phase in implementation
type ImplementationPhase struct {
	PhaseNumber int
	Name        string
	Description string
	Duration    time.Duration
	Responsible string
}

// QualityGate represents a quality gate in the implementation process
type QualityGate struct {
	Name        string
	Description string
	Criteria    string
	Responsible string
}

// ReleaseManagement represents release management processes
type ReleaseManagement struct {
	ReleaseTypes     []ReleaseType
	ApprovalProcess  []ApprovalStep
	TestingRequirements []TestingRequirement
	DeploymentWindows []DeploymentWindow
}

// ReleaseType represents a type of release
type ReleaseType string

const (
	ReleaseMajor    ReleaseType = "major"
	ReleaseMinor    ReleaseType = "minor"
	ReleasePatch    ReleaseType = "patch"
	ReleaseEmergency ReleaseType = "emergency"
)

// TestingRequirement represents a testing requirement for releases
type TestingRequirement struct {
	Type        string
	Description string
	Responsible string
	Duration    time.Duration
}

// DeploymentWindow represents a deployment time window
type DeploymentWindow struct {
	Environment string
	StartTime   string
	EndTime     string
	Days        []string
}

// DeploymentStrategy represents the deployment strategy
type DeploymentStrategy struct {
	Type           DeploymentType
	AutomationLevel string
	RollbackCapability bool
	Monitoring     string
}

// DeploymentType represents the type of deployment
type DeploymentType string

const (
	DeploymentBigBang DeploymentType = "big_bang"
	DeploymentPhased  DeploymentType = "phased"
	DeploymentBlueGreen DeploymentType = "blue_green"
	DeploymentCanary  DeploymentType = "canary"
)
