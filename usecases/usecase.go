package usecases

import "student-planner/domain"

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}
func (usecase *UserUsecase) Login(email string, password string) (domain.UserModel, error) {
	user, err := usecase.repo.Login(email, password)
	if err != nil {
		return domain.UserModel{}, err
	}
	return user, nil
}
func (usecase *UserUsecase) Register(name string, email string, password string) error {
	err := usecase.repo.Register(name, email, password)
	if err != nil {
		return err
	}
	return nil
}
func (usecase *UserUsecase) GetUser(id int) (domain.UserModel, error) {
	user, err := usecase.repo.GetUser(id)
	if err != nil {
		return domain.UserModel{}, err
	}
	return user, nil
}

func (usecase *UserUsecase) UpdateBadge(id int, badge string) error {
	err := usecase.repo.UpdateBadge(id, badge)
	if err != nil {
		return err
	}
	return nil
}
