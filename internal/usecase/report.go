package usecase

// ReportUseCase -.
type ReportUseCase struct {
	repo TimeEntryRepo
}

// NewReportUseCase -.
func NewReportUseCase(t TimeEntryRepo) *ReportUseCase {
	return &ReportUseCase{
		repo: t,
	}
}
