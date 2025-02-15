package handler

import (
	"github.com/Axontik/comin-leave-management-service/internal/service"
	"github.com/gin-gonic/gin"
)

type HolidayHandler struct {
	leaveService service.LeaveService
}

func NewHolidayHandler(leaveService service.LeaveService) *HolidayHandler {
	return &HolidayHandler{
		leaveService: leaveService,
	}
}

func (h *HolidayHandler) Create(c *gin.Context) {
	// Implementation
}

func (h *HolidayHandler) List(c *gin.Context) {
	// Implementation
}

func (h *HolidayHandler) Update(c *gin.Context) {
	// Implementation
}

func (h *HolidayHandler) Delete(c *gin.Context) {
	// Implementation
}

func (h *HolidayHandler) GetCalendarView(c *gin.Context) {
	// Implementation
}
