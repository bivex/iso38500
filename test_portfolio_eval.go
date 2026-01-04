package main

import (
	"fmt"
	"log"
	domain "github.com/iso38500/iso38500-governance-sdk/domain"
	memory "github.com/iso38500/iso38500-governance-sdk/infrastructure/memory"
)

func main() {
	// Initialize repositories
	appRepo := memory.NewApplicationRepositoryMemory()
	govRepo := memory.NewGovernanceAgreementRepositoryMemory()
	portfolioRepo := memory.NewApplicationPortfolioRepositoryMemory()

	// Create test data similar to the example
	app := domain.Application{
		ID:          domain.ApplicationID("test-app-001"),
		Name:        "Test Application",
		Description: "Test app for portfolio evaluation",
		Version:     "1.0.0",
		Status:      domain.StatusActive,
	}
	appRepo.Save(nil, app)

	portfolio := domain.ApplicationPortfolio{
		ID:           domain.PortfolioID("test-portfolio-001"),
		Name:         "Test Portfolio",
		Description:  "Test portfolio",
		Owner:        "test",
		Applications: []domain.Application{app},
	}
	portfolioRepo.Save(nil, portfolio)

	// Test portfolio evaluation
	evalService := domain.NewEvaluationService(appRepo, govRepo, portfolioRepo, nil, nil)
	assessment, err := evalService.EvaluatePortfolio(nil, domain.PortfolioID("test-portfolio-001"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("✅ Portfolio Evaluation Test Passed!\n")
	fmt.Printf("Total Applications: %d\n", assessment.TotalApplications)
	fmt.Printf("Active Applications: %d\n", assessment.ActiveApplications)
	fmt.Printf("App Status: %s\n", app.Status)
	fmt.Printf("StatusActive constant: %s\n", domain.StatusActive)
}
