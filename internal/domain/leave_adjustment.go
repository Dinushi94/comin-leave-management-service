package domain

import (
	"time"

	"github.com/google/uuid"
)

type LeaveBalanceAdjustment struct {
	Base
	LeaveBalanceID uuid.UUID     `json:"leave_balance_id" gorm:"type:uuid;not null"`
	Adjustment     float64       `json:"adjustment" gorm:"type:decimal(5,2);not null"`
	Reason         string        `json:"reason" gorm:"not null"`
	PerformedBy    uuid.UUID     `json:"performed_by" gorm:"type:uuid;not null"`
	ApprovedBy     *uuid.UUID    `json:"approved_by,omitempty" gorm:"type:uuid"`
	ApprovedAt     *time.Time    `json:"approved_at,omitempty"`
	Comments       string        `json:"comments"`
	Status         string        `json:"status" gorm:"default:'pending'"`
	LeaveBalance   *LeaveBalance `json:"leave_balance,omitempty" gorm:"foreignKey:LeaveBalanceID"`
}

type CreateBalanceAdjustmentRequest struct {
	LeaveBalanceID uuid.UUID `json:"leave_balance_id" binding:"required"`
	Adjustment     float64   `json:"adjustment" binding:"required,ne=0"`
	Reason         string    `json:"reason" binding:"required,min=5,max=500"`
	Comments       string    `json:"comments" binding:"max=1000"`
}

type UpdateBalanceAdjustmentRequest struct {
	Status     string    `json:"status" binding:"required,oneof=pending approved rejected"`
	Comments   string    `json:"comments" binding:"max=1000"`
	ApprovedBy uuid.UUID `json:"approved_by" binding:"required_if=Status approved"`
}

// Constants for balance adjustment
const (
	AdjustmentStatusPending  = "pending"
	AdjustmentStatusApproved = "approved"
	AdjustmentStatusRejected = "rejected"
)

// Methods for LeaveBalanceAdjustment
func (a *LeaveBalanceAdjustment) IsPositive() bool {
	return a.Adjustment > 0
}

func (a *LeaveBalanceAdjustment) IsNegative() bool {
	return a.Adjustment < 0
}

func (a *LeaveBalanceAdjustment) CanApprove() bool {
	return a.Status == AdjustmentStatusPending
}

func (a *LeaveBalanceAdjustment) CanReject() bool {
	return a.Status == AdjustmentStatusPending
}
