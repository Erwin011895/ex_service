package user

import (
	"log"

	dto "ex_service/src/app/dto/user"

	repo "ex_service/src/infra/persistence/postgres/user"

	helper "ex_service/src/infra/helper"

	"golang.org/x/crypto/bcrypt"
)

type UserUCInterface interface {
	Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error)
	Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error)
}

type userUseCase struct {
	Repo repo.UserRepository
}

func NewUserUseCase(repo repo.UserRepository) UserUCInterface {
	return &userUseCase{
		Repo: repo,
	}
}

func (uc *userUseCase) Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error) {
	crypted, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	data.Password = string(crypted)

	result, err := uc.Repo.Register(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result.Token, err = helper.GenerateToken(result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *userUseCase) Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error) {

	result, err := uc.Repo.Login(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(data.Password))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result.Token, err = helper.GenerateToken(result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
