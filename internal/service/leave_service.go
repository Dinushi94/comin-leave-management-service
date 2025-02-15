package service

import (
	"errors"

	"github.com/Axontik/comin-leave-management-service/internal/domain"
	"github.com/Axontik/comin-leave-management-service/internal/repository"
	"github.com/google/uuid"
)

type LeaveService interface {
	// Leave Type methods
	CreateLeaveType(leaveType *domain.LeaveType) error
	GetLeaveType(orgID, id uuid.UUID) (*domain.LeaveType, error)
	UpdateLeaveType(leaveType *domain.LeaveType) error
	DeleteLeaveType(orgID, id uuid.UUID) error
	ListLeaveTypes(orgID uuid.UUID, params *domain.ListLeaveTypesParams) ([]domain.LeaveType, int64, error)
	CreateLeaveRequest(orgID uuid.UUID, req *domain.CreateLeaveRequestRequest) (*domain.LeaveRequest, error)
}

type leaveService struct {
	leaveRepo repository.LeaveRepository
}

func NewLeaveService(leaveRepo repository.LeaveRepository) LeaveService {
	return &leaveService{
		leaveRepo: leaveRepo,
	}
}

// CreateLeaveType creates a new leave type
func (s *leaveService) CreateLeaveType(leaveType *domain.LeaveType) error {
	// Validate leave type
	if err := validateLeaveType(leaveType); err != nil {
		return err
	}

	// Check for duplicate name in the organization
	existingTypes, _, err := s.ListLeaveTypes(leaveType.OrganizationID, &domain.ListLeaveTypesParams{
		Name: leaveType.Name,
	})
	if err != nil {
		return err
	}
	if len(existingTypes) > 0 {
		return errors.New("leave type with this name already exists")
	}

	// Create leave type
	return s.leaveRepo.CreateLeaveType(leaveType)
}

// GetLeaveType retrieves a leave type by ID
func (s *leaveService) GetLeaveType(orgID, id uuid.UUID) (*domain.LeaveType, error) {
	leaveType, err := s.leaveRepo.GetLeaveType(id)
	if err != nil {
		return nil, err
	}

	// Verify organization ownership
	if leaveType.OrganizationID != orgID {
		return nil, errors.New("leave type not found in organization")
	}

	return leaveType, nil
}

// UpdateLeaveType updates an existing leave type
func (s *leaveService) UpdateLeaveType(leaveType *domain.LeaveType) error {
	// Validate leave type
	if err := validateLeaveType(leaveType); err != nil {
		return err
	}

	// Check if leave type exists
	existing, err := s.GetLeaveType(leaveType.OrganizationID, leaveType.ID)
	if err != nil {
		return err
	}

	// Check for name uniqueness if name is being changed
	if existing.Name != leaveType.Name {
		existingTypes, _, err := s.ListLeaveTypes(leaveType.OrganizationID, &domain.ListLeaveTypesParams{
			Name: leaveType.Name,
		})
		if err != nil {
			return err
		}
		if len(existingTypes) > 0 {
			return errors.New("leave type with this name already exists")
		}
	}

	return s.leaveRepo.UpdateLeaveType(leaveType)
}

// DeleteLeaveType deletes a leave type
func (s *leaveService) DeleteLeaveType(orgID, id uuid.UUID) error {
	// Check if leave type exists and belongs to organization
	existing, err := s.GetLeaveType(orgID, id)
	if err != nil {
		return err
	}

	// Check if there are any active leave requests using this type
	hasActiveRequests, err := s.leaveRepo.HasActiveLeaveRequests(id)
	if err != nil {
		return err
	}
	if hasActiveRequests {
		return errors.New("cannot delete leave type with active leave requests")
	}

	return s.leaveRepo.DeleteLeaveType(existing.ID)
}

// ListLeaveTypes lists leave types with filtering and pagination
func (s *leaveService) ListLeaveTypes(orgID uuid.UUID, params *domain.ListLeaveTypesParams) ([]domain.LeaveType, int64, error) {
	// Validate pagination parameters
	if params != nil {
		if params.Page < 1 {
			params.Page = 1
		}
		if params.PageSize < 1 || params.PageSize > 100 {
			params.PageSize = 10
		}
	}

	return s.leaveRepo.ListLeaveTypesWithOptions(orgID, params)
}

// Helper functions

func validateLeaveType(leaveType *domain.LeaveType) error {
	if leaveType.Name == "" {
		return errors.New("name is required")
	}
	if leaveType.DefaultDays < 0 {
		return errors.New("default days cannot be negative")
	}
	if leaveType.MaxDaysPerRequest < 1 {
		return errors.New("max days per request must be at least 1")
	}
	if leaveType.MinDaysNotice < 0 {
		return errors.New("minimum days notice cannot be negative")
	}
	return nil
}

func (s *leaveService) CreateLeaveRequest(orgID uuid.UUID, req *domain.CreateLeaveRequestRequest) (*domain.LeaveRequest, error) {
	// Validate request
	if req.EmployeeID == uuid.Nil {
		return nil, errors.New("employee ID is required")
	}
	if req.LeaveTypeID == uuid.Nil {
		return nil, errors.New("leave type ID is required")
	}
	if req.StartDate.After(req.EndDate) {
		return nil, errors.New("start date cannot be after end date")
	}

	// Get leave type
	leaveType, err := s.GetLeaveType(orgID, req.LeaveTypeID)
	if err != nil {
		return nil, err
	}

	// Calculate total days
	totalDays := int(req.EndDate.Sub(req.StartDate).Milliseconds() / 86400000)
	if totalDays > leaveType.MaxDaysPerRequest {
		return nil, errors.New("total days exceed maximum allowed")
	}

	// Create leave request
	leaveRequest := &domain.LeaveRequest{
		EmployeeID:  req.EmployeeID,
		LeaveTypeID: req.LeaveTypeID,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Status:      domain.LeaveStatusPending,
		Reason:      req.Reason,
	}

	// Save leave request
	if err := s.leaveRepo.CreateLeaveRequest(leaveRequest); err != nil {
		return nil, err
	}

	return leaveRequest, nil
}



