// internal/domain/stats.go
package domain

import (
	"time"

	"github.com/google/uuid"
)

// LeaveStats represents overall leave statistics
type LeaveStats struct {
	TotalRequests  int64           `json:"total_requests"`
	TotalDaysTaken float64         `json:"total_days_taken"`
	LeaveByType    []LeaveByType   `json:"leave_by_type"`
	LeaveByStatus  []LeaveByStatus `json:"leave_by_status"`
	MonthlyStats   []MonthlyStats  `json:"monthly_stats"`
}

// LeaveByType represents leave statistics grouped by leave type
type LeaveByType struct {
	LeaveType string  `json:"leave_type"`
	Count     int64   `json:"count"`
	TotalDays float64 `json:"total_days"`
}

// LeaveByStatus represents leave statistics grouped by status
type LeaveByStatus struct {
	Status    string  `json:"status"`
	Count     int64   `json:"count"`
	TotalDays float64 `json:"total_days"`
}

// MonthlyStats represents leave statistics by month
type MonthlyStats struct {
	Month     time.Time `json:"month"`
	Count     int64     `json:"count"`
	TotalDays float64   `json:"total_days"`
}

// DepartmentLeaveStats represents leave statistics for a department
type DepartmentLeaveStats struct {
	DepartmentID   uuid.UUID     `json:"department_id"`
	DepartmentName string        `json:"department_name"`
	TotalRequests  int64         `json:"total_requests"`
	TotalDaysTaken float64       `json:"total_days_taken"`
	LeaveByType    []LeaveByType `json:"leave_by_type"`
}

// EmployeeLeaveStats represents leave statistics for an employee
type EmployeeLeaveStats struct {
	EmployeeID     uuid.UUID           `json:"employee_id"`
	EmployeeName   string              `json:"employee_name"`
	TotalRequests  int64               `json:"total_requests"`
	TotalDaysTaken float64             `json:"total_days_taken"`
	LeaveBalances  []LeaveBalanceStats `json:"leave_balances"`
}

// LeaveBalanceStats represents leave balance statistics
type LeaveBalanceStats struct {
	LeaveType     string  `json:"leave_type"`
	TotalDays     float64 `json:"total_days"`
	UsedDays      float64 `json:"used_days"`
	PendingDays   float64 `json:"pending_days"`
	RemainingDays float64 `json:"remaining_days"`
}

// StatsRequest represents the request parameters for statistics
type StatsRequest struct {
	OrganizationID uuid.UUID  `json:"organization_id"`
	DepartmentID   *uuid.UUID `json:"department_id,omitempty"`
	EmployeeID     *uuid.UUID `json:"employee_id,omitempty"`
	StartDate      time.Time  `json:"start_date"`
	EndDate        time.Time  `json:"end_date"`
	GroupBy        []string   `json:"group_by,omitempty"` // month, type, status
}

// LeaveAnalytics represents advanced leave analytics
type LeaveAnalytics struct {
	AverageLeaveLength    float64     `json:"average_leave_length"`
	MostCommonLeaveType   string      `json:"most_common_leave_type"`
	LeastUsedLeaveType    string      `json:"least_used_leave_type"`
	PeakLeaveMonth        string      `json:"peak_leave_month"`
	ApprovalRate          float64     `json:"approval_rate"`
	AverageProcessingTime float64     `json:"average_processing_time"` // in hours
	TrendAnalysis         []TrendData `json:"trend_analysis"`
}

// TrendData represents trend analysis data points
type TrendData struct {
	Period time.Time `json:"period"`
	Value  float64   `json:"value"`
	Trend  string    `json:"trend"` // increasing, decreasing, stable
}

// Methods for LeaveStats
func (s *LeaveStats) GetAverageLeaveLength() float64 {
	if s.TotalRequests == 0 {
		return 0
	}
	return s.TotalDaysTaken / float64(s.TotalRequests)
}

func (s *LeaveStats) GetMostUsedLeaveType() *LeaveByType {
	if len(s.LeaveByType) == 0 {
		return nil
	}

	mostUsed := &s.LeaveByType[0]
	for _, lt := range s.LeaveByType {
		if lt.TotalDays > mostUsed.TotalDays {
			mostUsed = &lt
		}
	}
	return mostUsed
}

func (s *LeaveStats) GetMonthlyAverage() float64 {
	if len(s.MonthlyStats) == 0 {
		return 0
	}

	total := 0.0
	for _, ms := range s.MonthlyStats {
		total += ms.TotalDays
	}
	return total / float64(len(s.MonthlyStats))
}
