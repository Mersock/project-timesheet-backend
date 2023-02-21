package usecase

// ProjectsUseCase -.
type ProjectsUseCase struct {
	repo ProjectRepo
}

// NewProjectsUseCase -.
func NewProjectsUseCase(r ProjectRepo) *ProjectsUseCase {
	return &ProjectsUseCase{repo: r}
}
