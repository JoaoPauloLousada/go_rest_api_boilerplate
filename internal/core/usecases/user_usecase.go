package usecases

import (
	domain "github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/core/entity"
	"github.com/JoaoPauloLousada/go_rest_api_boilerplate/internal/core/interfaces"
)

type UserUseCase struct {
	repository interfaces.Repository[domain.User]
}

func NewUserUseCase(repository interfaces.Repository[domain.User]) *UserUseCase {
	return &UserUseCase{repository: repository}
}

func (useCase *UserUseCase) GetAll() ([]domain.User, error) {
	return useCase.repository.GetAll()
}

func (useCase *UserUseCase) Create(user domain.User) error {
	return useCase.repository.Create(user)
}
