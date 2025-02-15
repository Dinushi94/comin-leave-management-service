package handler

import (
	"github.com/Axontik/comin-leave-management-service/internal/service"
	"github.com/gin-gonic/gin"
)

type LeaveBalanceHandler struct {
	leaveService service.LeaveService
}

func NewLeaveBalanceHandler(leaveService service.LeaveService) *LeaveBalanceHandler {
	return &LeaveBalanceHandler{
		leaveService: leaveService,
	}
}

func (h *LeaveBalanceHandler) List(c *gin.Context) {
	// Implementation
}

func (h *LeaveBalanceHandler) GetByEmployee(c *gin.Context) {
	// Implementation
}

func (h *LeaveBalanceHandler) AdjustBalance(c *gin.Context) {
	// Implementation
}

func (h *LeaveBalanceHandler) GetBalanceHistory(c *gin.Context) {
	// Implementation
}

func (h *LeaveBalanceHandler) YearlyReset(c *gin.Context) {
	// Implementation
}
