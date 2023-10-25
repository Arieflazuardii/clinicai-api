package request

import (
	"clinicai-api/models/domain"
	"clinicai-api/models/schema"
	"clinicai-api/models/web"
)

func ScheduleDomainToScheduleSchema(request domain.Schedule) *schema.Schedule {
	return &schema.Schedule{
		DoctorID: request.DoctorID,
		Date: request.Date,
		Quota: request.Quota,
	}
}

func ScheduleCreateRequestToScheduleDomain(request web.ScheduleCreateRequest) *domain.Schedule{
	return &domain.Schedule{
		DoctorID: request.DoctorID,
		Date: request.Date,
		Quota: request.Quota,
	}
}

func ScheduleUpdateRequestToScheduleDomain(request web.ScheduleUpdateRequest) *domain.Schedule{
	return &domain.Schedule{
		DoctorID: request.DoctorID,
		Date: request.Date,
		Quota: request.Quota,
	}
}

