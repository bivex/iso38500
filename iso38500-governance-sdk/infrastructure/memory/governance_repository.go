package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/iso38500/iso38500-governance-sdk/domain"
)

// GovernanceAgreementRepositoryMemory is an in-memory implementation of GovernanceAgreementRepository
type GovernanceAgreementRepositoryMemory struct {
	mu          sync.RWMutex
	agreements  map[domain.GovernanceAgreementID]domain.GovernanceAgreement
	byApplication map[domain.ApplicationID]domain.GovernanceAgreementID
}

// NewGovernanceAgreementRepositoryMemory creates a new in-memory governance agreement repository
func NewGovernanceAgreementRepositoryMemory() *GovernanceAgreementRepositoryMemory {
	return &GovernanceAgreementRepositoryMemory{
		agreements:   make(map[domain.GovernanceAgreementID]domain.GovernanceAgreement),
		byApplication: make(map[domain.ApplicationID]domain.GovernanceAgreementID),
	}
}

// Save saves a governance agreement
func (r *GovernanceAgreementRepositoryMemory) Save(ctx context.Context, agreement domain.GovernanceAgreement) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.agreements[agreement.ID] = agreement
	r.byApplication[agreement.ApplicationID] = agreement.ID
	return nil
}

// FindByID finds a governance agreement by ID
func (r *GovernanceAgreementRepositoryMemory) FindByID(ctx context.Context, id domain.GovernanceAgreementID) (domain.GovernanceAgreement, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	agreement, exists := r.agreements[id]
	if !exists {
		return domain.GovernanceAgreement{}, errors.New("governance agreement not found")
	}
	return agreement, nil
}

// FindByApplicationID finds a governance agreement by application ID
func (r *GovernanceAgreementRepositoryMemory) FindByApplicationID(ctx context.Context, appID domain.ApplicationID) (domain.GovernanceAgreement, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	agreementID, exists := r.byApplication[appID]
	if !exists {
		return domain.GovernanceAgreement{}, errors.New("governance agreement not found for application")
	}

	agreement, exists := r.agreements[agreementID]
	if !exists {
		return domain.GovernanceAgreement{}, errors.New("governance agreement not found")
	}
	return agreement, nil
}

// FindAll finds all governance agreements
func (r *GovernanceAgreementRepositoryMemory) FindAll(ctx context.Context) ([]domain.GovernanceAgreement, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	agreements := make([]domain.GovernanceAgreement, 0, len(r.agreements))
	for _, agreement := range r.agreements {
		agreements = append(agreements, agreement)
	}
	return agreements, nil
}

// FindByStatus finds governance agreements by status
func (r *GovernanceAgreementRepositoryMemory) FindByStatus(ctx context.Context, status domain.AgreementStatus) ([]domain.GovernanceAgreement, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	agreements := make([]domain.GovernanceAgreement, 0)
	for _, agreement := range r.agreements {
		if agreement.Status == status {
			agreements = append(agreements, agreement)
		}
	}
	return agreements, nil
}

// Update updates a governance agreement
func (r *GovernanceAgreementRepositoryMemory) Update(ctx context.Context, agreement domain.GovernanceAgreement) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.agreements[agreement.ID]; !exists {
		return errors.New("governance agreement not found")
	}

	r.agreements[agreement.ID] = agreement
	return nil
}

// Delete deletes a governance agreement
func (r *GovernanceAgreementRepositoryMemory) Delete(ctx context.Context, id domain.GovernanceAgreementID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	agreement, exists := r.agreements[id]
	if !exists {
		return errors.New("governance agreement not found")
	}

	delete(r.agreements, id)
	delete(r.byApplication, agreement.ApplicationID)
	return nil
}

// Exists checks if a governance agreement exists
func (r *GovernanceAgreementRepositoryMemory) Exists(ctx context.Context, id domain.GovernanceAgreementID) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.agreements[id]
	return exists, nil
}
