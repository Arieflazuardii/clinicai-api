package response

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)

func ScheduleSchemaToScheduleDomain(schedule *schema.Schedule) *domain.Schedule {
	return &domain.Schedule{
		DoctorID: schedule.DoctorID,
		Date:     schedule.Date,
		Quota:    schedule.Quota,
	}
}

func ScheduleDomainToScheduleResponse(schedule *domain.Schedule) web.ScheduleResponse {
	return web.ScheduleResponse{
		ID: schedule.ID,
		DoctorID: schedule.DoctorID,
		DoctorName: schedule.DoctorName,
		Date:     schedule.Date,
		Quota:    schedule.Quota,
	}
}

func CreateScheduleDomainToScheduleResponse(schedule *domain.Schedule) web.ScheduleResponse {
	return web.ScheduleResponse{
		ID: 	  schedule.ID,
		DoctorID: schedule.DoctorID,
		DoctorName: schedule.DoctorName,
		Date:     schedule.Date,
		Quota:    schedule.Quota,
	}
}

func UpdateScheduleDomainToScheduleResponse(schedule *domain.Schedule) web.ScheduleUpdateResponse {
	return web.ScheduleUpdateResponse{
		ID: schedule.ID,
		DoctorID: schedule.DoctorID,
		DoctorName: schedule.DoctorName,
		Date:     schedule.Date,
		Quota:    schedule.Quota,
	}
}

func ConvertScheduleResponse(schedules []domain.Schedule) []web.ConvertScheduleResponse {
	var results []web.ConvertScheduleResponse
	for _, schedule := range schedules {
		scheduleResponse := web.ConvertScheduleResponse {
			ID: schedule.ID,
			DoctorID: schedule.DoctorID,
			DoctorName: schedule.DoctorName,
			Date:     schedule.Date,
			Quota:    schedule.Quota,
		}
		results = append(results, scheduleResponse)
	}
	return results
}
