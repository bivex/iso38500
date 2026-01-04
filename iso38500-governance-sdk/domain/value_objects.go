package domain

import (
	"errors"
	"time"
)

// ResponsibilityMatrix represents the RACI matrix for stakeholders
type ResponsibilityMatrix struct {
	Entries []RACIEntry
}

// RACIEntry represents a single entry in the RACI matrix
type RACIEntry struct {
	Activity    string
	Responsible string // Who does the work
	Accountable string // Who is ultimately accountable
	Consulted   string // Who needs to be consulted
	Informed    string // Who needs to be informed
}

// Validate ensures the RACI entry has valid data
func (r *RACIEntry) Validate() error {
	if r.Activity == "" {
		return errors.New("activity cannot be empty")
	}
	if r.Responsible == "" {
		return errors.New("responsible party cannot be empty")
	}
	if r.Accountable == "" {
		return errors.New("accountable party cannot be empty")
	}
	return nil
}

// AddEntry adds a RACI entry to the matrix
func (rm *ResponsibilityMatrix) AddEntry(entry RACIEntry) error {
	if err := entry.Validate(); err != nil {
		return err
	}
	rm.Entries = append(rm.Entries, entry)
	return nil
}

// KPI represents a Key Performance Indicator
type KPI struct {
	ID          string
	Name        string
	Description string
	Target      float64
	Unit        string
	Category    string
	Frequency   string // daily, weekly, monthly, quarterly
	Status      KPIStatus
}

// KPIStatus represents the status of a KPI measurement
type KPIStatus string

const (
	KPIStatusOnTrack    KPIStatus = "on_track"
	KPIStatusAtRisk     KPIStatus = "at_risk"
	KPIStatusOffTrack   KPIStatus = "off_track"
	KPIStatusNotMeasured KPIStatus = "not_measured"
)

// Validate ensures the KPI has valid data
func (k *KPI) Validate() error {
	if k.ID == "" {
		return errors.New("KPI ID cannot be empty")
	}
	if k.Name == "" {
		return errors.New("KPI name cannot be empty")
	}
	return nil
}

// ApplicationCatalogue represents the business functionality of an application
type ApplicationCatalogue struct {
	Functionality []Functionality
	LastUpdated   time.Time
}

// Functionality represents a specific business function provided by the application
type Functionality struct {
	ID          string
	Name        string
	Description string
	Category    string
	Priority    Priority
	Status      FunctionalityStatus
}

// Priority represents the business priority of functionality
type Priority string

const (
	PriorityCritical Priority = "critical"
	PriorityHigh     Priority = "high"
	PriorityMedium   Priority = "medium"
	PriorityLow      Priority = "low"
)

// FunctionalityStatus represents the status of functionality
type FunctionalityStatus string

const (
	FunctionalityAvailable   FunctionalityStatus = "available"
	FunctionalityPlanned     FunctionalityStatus = "planned"
	FunctionalityDeprecated  FunctionalityStatus = "deprecated"
	FunctionalityUnavailable FunctionalityStatus = "unavailable"
)

// ApplicationInterface represents technical interfaces of an application
type ApplicationInterface struct {
	ID          string
	Name        string
	Type        InterfaceType
	Description string
	Protocol    string
	Endpoint    string
	Status      InterfaceStatus
}

// InterfaceType represents the type of interface
type InterfaceType string

const (
	InterfaceAPI      InterfaceType = "api"
	InterfaceDatabase InterfaceType = "database"
	InterfaceFile     InterfaceType = "file"
	InterfaceMessage  InterfaceType = "message"
	InterfaceUI       InterfaceType = "ui"
)

// InterfaceStatus represents the status of an interface
type InterfaceStatus string

const (
	InterfaceActive   InterfaceStatus = "active"
	InterfaceInactive InterfaceStatus = "inactive"
	InterfaceTesting  InterfaceStatus = "testing"
	InterfaceFailed   InterfaceStatus = "failed"
)

// ConfigurationStandard represents the configuration requirements for an application
type ConfigurationStandard struct {
	EnvironmentVariables []EnvironmentVariable
	ConfigurationFiles   []ConfigurationFile
	SecuritySettings     []SecuritySetting
	LastUpdated          time.Time
}

// EnvironmentVariable represents a required environment variable
type EnvironmentVariable struct {
	Name        string
	Value       string
	Description string
	Required    bool
	Sensitive   bool
}

// ConfigurationFile represents a configuration file requirement
type ConfigurationFile struct {
	Path        string
	Format      string
	Description string
	Required    bool
}

// SecuritySetting represents a security configuration requirement
type SecuritySetting struct {
	Name        string
	Value       string
	Description string
	Category    string
}

// SecurityProvisions represents security measures for an application
type SecurityProvisions struct {
	DataConfidentiality   []SecurityMeasure
	DataIntegrity        []SecurityMeasure
	ApplicationAvailability SLA
	ApplicationAuthenticity []SecurityMeasure
	RolesAndPermissions   []RolePermission
}

// SecurityMeasure represents a specific security measure
type SecurityMeasure struct {
	Name        string
	Description string
	Category    string
	Status      SecurityStatus
}

// SecurityStatus represents the implementation status of a security measure
type SecurityStatus string

const (
	SecurityImplemented SecurityStatus = "implemented"
	SecurityPlanned     SecurityStatus = "planned"
	SecurityPartial     SecurityStatus = "partial"
	SecurityNotStarted  SecurityStatus = "not_started"
)

// RolePermission represents a role-based permission
type RolePermission struct {
	Role        string
	Permissions []string
	Resource    string
}

// SLA represents a Service Level Agreement
type SLA struct {
	ServiceName      string
	ResponseTime     time.Duration
	Availability     float64 // percentage (e.g., 99.9)
	Uptime           string
	SupportHours     string
	EscalationMatrix []EscalationLevel
}

// EscalationLevel represents a level in the escalation matrix
type EscalationLevel struct {
	Level       int
	Description string
	ResponseTime time.Duration
	Contacts    []string
}

// BusinessContinuity represents business continuity provisions
type BusinessContinuity struct {
	RecoveryTimeObjective time.Duration
	RecoveryPointObjective time.Duration
	BusinessImpactAnalysis string
	ContinuityPlans       []ContinuityPlan
	TestingSchedule       string
}

// ContinuityPlan represents a specific continuity plan
type ContinuityPlan struct {
	Name        string
	Description string
	Type        ContinuityType
	Status      PlanStatus
}

// ContinuityType represents the type of continuity plan
type ContinuityType string

const (
	ContinuityDisaster  ContinuityType = "disaster_recovery"
	ContinuityBackup    ContinuityType = "backup"
	ContinuityFailover  ContinuityType = "failover"
	ContinuityRedundant ContinuityType = "redundant_systems"
)

// PlanStatus represents the status of a continuity plan
type PlanStatus string

const (
	PlanDocumented PlanStatus = "documented"
	PlanTested     PlanStatus = "tested"
	PlanActive     PlanStatus = "active"
	PlanOutdated   PlanStatus = "outdated"
)
