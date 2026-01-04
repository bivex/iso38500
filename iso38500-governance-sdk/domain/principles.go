package domain

import (
	"time"
)

// EvaluatePrinciple represents the Evaluate principle from ISO 38500
type EvaluatePrinciple struct {
	CurrentSituation CurrentSituationAssessment
	NeedsAssessment  NeedsAssessment
	RiskAssessment   RiskAssessment
	PerformanceMetrics []KPIMeasurement
	LastEvaluated    time.Time
}

// CurrentSituationAssessment represents assessment of current situation
type CurrentSituationAssessment struct {
	ApplicationInventory []ApplicationAssessment
	PortfolioHealth      PortfolioHealthAssessment
	GovernanceMaturity   GovernanceMaturityAssessment
}

// ApplicationAssessment represents assessment of a specific application
type ApplicationAssessment struct {
	ApplicationID   ApplicationID
	TechnicalHealth TechnicalHealth
	BusinessValue   BusinessValueAssessment
	RiskLevel       RiskLevel
	Recommendations []Recommendation
}

// TechnicalHealth represents the technical health of an application
type TechnicalHealth struct {
	CodeQuality       int // 1-5 scale
	Documentation     int // 1-5 scale
	TestCoverage      float64
	SecurityScore     int // 1-5 scale
	PerformanceScore  int // 1-5 scale
}

// BusinessValueAssessment represents business value assessment
type BusinessValueAssessment struct {
	UsageMetrics      UsageMetrics
	BusinessAlignment float64 // percentage
	CostEfficiency    float64 // percentage
	UserSatisfaction  float64 // percentage
}

// UsageMetrics represents application usage metrics
type UsageMetrics struct {
	ActiveUsers       int
	TransactionVolume int
	UptimePercentage  float64
	ResponseTime      time.Duration
}

// RiskLevel represents the risk level
type RiskLevel string

const (
	RiskLow      RiskLevel = "low"
	RiskMedium   RiskLevel = "medium"
	RiskHigh     RiskLevel = "high"
	RiskCritical RiskLevel = "critical"
)

// Recommendation represents a recommendation from assessment
type Recommendation struct {
	ID          string
	Type        RecommendationType
	Description string
	Priority    Priority
	EstimatedEffort time.Duration
	BusinessImpact   string
}

// RecommendationType represents the type of recommendation
type RecommendationType string

const (
	RecModernize     RecommendationType = "modernize"
	RecReplace       RecommendationType = "replace"
	RecEnhance       RecommendationType = "enhance"
	RecRetire        RecommendationType = "retire"
	RecMaintain      RecommendationType = "maintain"
)

// PortfolioHealthAssessment represents overall portfolio health
type PortfolioHealthAssessment struct {
	TotalApplications     int
	ActiveApplications    int
	DeprecatedApplications int
	RedundantApplications int
	TotalCost            float64
	AverageApplicationAge time.Duration
	RiskDistribution     map[RiskLevel]int
}

// GovernanceMaturityAssessment represents governance maturity level
type GovernanceMaturityAssessment struct {
	MaturityLevel      int // 1-5 scale
	Strengths         []string
	Weaknesses        []string
	ImprovementAreas  []string
}

// NeedsAssessment represents assessment of organizational needs
type NeedsAssessment struct {
	BusinessObjectives []BusinessObjective
	TechnologyNeeds    []TechnologyNeed
	ResourceRequirements []ResourceRequirement
	Timeline          time.Duration
}

// BusinessObjective represents a business objective
type BusinessObjective struct {
	ID          string
	Name        string
	Description string
	Priority    Priority
	Deadline    time.Time
}

// TechnologyNeed represents a technology requirement
type TechnologyNeed struct {
	ID          string
	Name        string
	Description string
	Category    string
	Priority    Priority
}

// ResourceRequirement represents a resource requirement
type ResourceRequirement struct {
	Type        string
	Description string
	Quantity    int
	Timeframe   time.Duration
}

// RiskAssessment represents risk assessment
type RiskAssessment struct {
	Risks           []Risk
	MitigationPlans []MitigationPlan
	OverallRiskLevel RiskLevel
}

// Risk represents an identified risk
type Risk struct {
	ID          string
	Name        string
	Description string
	Category    string
	Probability float64 // 0-1
	Impact      RiskImpact
	Level       RiskLevel
}

// RiskImpact represents the impact of a risk
type RiskImpact string

const (
	ImpactLow      RiskImpact = "low"
	ImpactMedium   RiskImpact = "medium"
	ImpactHigh     RiskImpact = "high"
	ImpactCritical RiskImpact = "critical"
)

// MitigationPlan represents a risk mitigation plan
type MitigationPlan struct {
	RiskID       string
	Actions      []string
	Responsible  string
	Timeline     time.Duration
	Budget       float64
	Effectiveness float64 // 0-1
}

// KPIMeasurement represents a KPI measurement
type KPIMeasurement struct {
	KPIID       string
	Value       float64
	Target      float64
	Achieved    bool
	MeasuredAt  time.Time
	Notes       string
}

// DirectPrinciple represents the Direct principle from ISO 38500
type DirectPrinciple struct {
	StrategicDirection StrategicDirection
	ResourceAllocation ResourceAllocation
	PolicyFramework    PolicyFramework
	ActionPlans        []ActionPlan
	LastDirected       time.Time
}

// StrategicDirection represents strategic direction setting
type StrategicDirection struct {
	Vision        string
	Mission       string
	Objectives    []StrategicObjective
	Initiatives   []StrategicInitiative
	Timeframe     time.Duration
}

// StrategicObjective represents a strategic objective
type StrategicObjective struct {
	ID          string
	Name        string
	Description string
	KPIs        []KPI
	Deadline    time.Time
}

// StrategicInitiative represents a strategic initiative
type StrategicInitiative struct {
	ID          string
	Name        string
	Description string
	Owner       string
	Budget      float64
	Deadline    time.Time
}

// ResourceAllocation represents resource allocation decisions
type ResourceAllocation struct {
	BudgetAllocations  []BudgetAllocation
	PersonnelAllocations []PersonnelAllocation
	TechnologyAllocations []TechnologyAllocation
}

// BudgetAllocation represents budget allocation
type BudgetAllocation struct {
	Category    string
	Amount      float64
	Timeframe   string
	Justification string
}

// PersonnelAllocation represents personnel allocation
type PersonnelAllocation struct {
	Role        string
	Count       int
	SkillLevel  string
	Timeframe   string
}

// TechnologyAllocation represents technology allocation
type TechnologyAllocation struct {
	Technology  string
	Purpose     string
	Budget      float64
	Timeframe   string
}

// PolicyFramework represents the policy framework
type PolicyFramework struct {
	Policies     []Policy
	Standards    []Standard
	Procedures   []Procedure
	Guidelines   []Guideline
}

// Policy represents a governance policy
type Policy struct {
	ID          string
	Name        string
	Description string
	Scope       string
	Owner       string
	Status      PolicyStatus
}

// PolicyStatus represents the status of a policy
type PolicyStatus string

const (
	PolicyDraft     PolicyStatus = "draft"
	PolicyApproved  PolicyStatus = "approved"
	PolicyPublished PolicyStatus = "published"
	PolicyRetired   PolicyStatus = "retired"
)

// Standard represents a governance standard
type Standard struct {
	ID          string
	Name        string
	Description string
	Category    string
	Mandatory   bool
}

// Procedure represents a governance procedure
type Procedure struct {
	ID          string
	Name        string
	Description string
	Steps       []ProcedureStep
}

// ProcedureStep represents a step in a procedure
type ProcedureStep struct {
	StepNumber  int
	Description string
	Responsible string
}

// Guideline represents a governance guideline
type Guideline struct {
	ID          string
	Name        string
	Description string
	Category    string
}

// ActionPlan represents an action plan
type ActionPlan struct {
	ID          string
	Name        string
	Description string
	Actions     []Action
	Owner       string
	Deadline    time.Time
	Status      ActionStatus
}

// Action represents a specific action in an action plan
type Action struct {
	ID          string
	Description string
	Responsible string
	Deadline    time.Time
	Status      ActionStatus
}

// ActionStatus represents the status of an action
type ActionStatus string

const (
	ActionPending   ActionStatus = "pending"
	ActionInProgress ActionStatus = "in_progress"
	ActionCompleted ActionStatus = "completed"
	ActionCancelled ActionStatus = "cancelled"
)

// MonitorPrinciple represents the Monitor principle from ISO 38500
type MonitorPrinciple struct {
	PerformanceMonitoring PerformanceMonitoring
	ComplianceMonitoring  ComplianceMonitoring
	RiskMonitoring        RiskMonitoring
	StakeholderFeedback   StakeholderFeedback
	Reporting            GovernanceReporting
	LastMonitored        time.Time
}

// PerformanceMonitoring represents performance monitoring
type PerformanceMonitoring struct {
	KPIMonitoring      []KPIMonitoring
	ServiceLevelMonitoring []ServiceLevelMonitoring
	UserExperienceMonitoring UserExperienceMonitoring
}

// KPIMonitoring represents KPI monitoring configuration
type KPIMonitoring struct {
	KPIID       string
	Frequency   string
	Responsible string
	Thresholds  []Threshold
	Alerts      []Alert
}

// Threshold represents a monitoring threshold
type Threshold struct {
	Level      string // warning, critical
	Value      float64
	Condition  string // >, <, =, etc.
}

// Alert represents an alert configuration
type Alert struct {
	Type        string
	Recipient   string
	Message     string
	Escalation  string
}

// ServiceLevelMonitoring represents service level monitoring
type ServiceLevelMonitoring struct {
	ServiceID   string
	SLAs        []SLA
	Metrics     []string
	Dashboards  []string
}

// UserExperienceMonitoring represents user experience monitoring
type UserExperienceMonitoring struct {
	Surveys         []Survey
	FeedbackChannels []FeedbackChannel
	SatisfactionScores []SatisfactionScore
}

// Survey represents a user survey
type Survey struct {
	ID          string
	Name        string
	Frequency   string
	Questions   []string
}

// FeedbackChannel represents a feedback collection channel
type FeedbackChannel struct {
	Type        string
	Description string
	Frequency   string
}

// SatisfactionScore represents a satisfaction score measurement
type SatisfactionScore struct {
	Metric      string
	Score       float64
	Date        time.Time
	SampleSize  int
}

// RiskMonitoring represents risk monitoring
type RiskMonitoring struct {
	RiskIndicators     []RiskIndicator
	RiskHeatMaps       []RiskHeatMap
	MitigationTracking []MitigationTracking
}

// RiskIndicator represents a risk indicator
type RiskIndicator struct {
	Name        string
	Value       float64
	Threshold   float64
	Status      RiskStatus
}

// RiskStatus represents the status of a risk indicator
type RiskStatus string

const (
	RiskStatusNormal RiskStatus = "normal"
	RiskStatusWarning RiskStatus = "warning"
	RiskStatusCritical RiskStatus = "critical"
)

// RiskHeatMap represents a risk heat map
type RiskHeatMap struct {
	Name        string
	Description string
	Data        map[string]map[string]float64 // risk vs impact matrix
}

// MitigationTracking represents mitigation action tracking
type MitigationTracking struct {
	MitigationID string
	Status       ActionStatus
	Progress     float64 // 0-1
	Notes        string
}

// StakeholderFeedback represents stakeholder feedback collection
type StakeholderFeedback struct {
	FeedbackItems    []FeedbackItem
	SurveyResults    []SurveyResult
	CommunicationLog []CommunicationLogEntry
}

// FeedbackItem represents a piece of stakeholder feedback
type FeedbackItem struct {
	ID          string
	Stakeholder string
	Feedback    string
	Category    string
	Sentiment   string
	Date        time.Time
}

// SurveyResult represents survey results
type SurveyResult struct {
	SurveyID    string
	Responses   []SurveyResponse
	Summary     SurveySummary
}

// SurveyResponse represents an individual survey response
type SurveyResponse struct {
	QuestionID  string
	Response    string
	Score       int
}

// SurveySummary represents survey summary statistics
type SurveySummary struct {
	TotalResponses   int
	AverageScore     float64
	ResponseRate     float64
	KeyInsights      []string
}

// CommunicationLogEntry represents a communication log entry
type CommunicationLogEntry struct {
	Date        time.Time
	Type        string
	Subject     string
	Recipients  []string
	Response    string
}

// GovernanceReporting represents governance reporting
type GovernanceReporting struct {
	Reports          []Report
	Dashboards       []Dashboard
	KPIDashboards    []KPIDashboard
	ExecutiveSummary ExecutiveSummary
}

// Report represents a governance report
type Report struct {
	ID          string
	Name        string
	Type        ReportType
	Frequency   string
	Recipients  []string
	LastGenerated time.Time
}

// ReportType represents the type of report
type ReportType string

const (
	ReportPerformance ReportType = "performance"
	ReportCompliance  ReportType = "compliance"
	ReportRisk        ReportType = "risk"
	ReportExecutive   ReportType = "executive"
)

// Dashboard represents a governance dashboard
type Dashboard struct {
	ID          string
	Name        string
	Description string
	Widgets     []Widget
	AccessRoles []string
}

// Widget represents a dashboard widget
type Widget struct {
	ID       string
	Type     string
	Title    string
	DataSource string
	Config   map[string]interface{}
}

// KPIDashboard represents a KPI dashboard
type KPIDashboard struct {
	ID          string
	Name        string
	KPIs        []string
	TimeRange   string
	RefreshRate string
}

// KeyMetric represents a key metric for executive summary
type KeyMetric struct {
	Name   string
	Value  float64
	Unit   string
	Trend  string
	Status string
}

// ExecutiveSummary represents an executive summary
type ExecutiveSummary struct {
	Period         string
	KeyMetrics     []KeyMetric
	Achievements   []string
	Challenges     []string
	Recommendations []string
}
