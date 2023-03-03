package usecase

import (
	"INi-Wallet2/dto"
	"INi-Wallet2/model"
	"INi-Wallet2/repository"
	"INi-Wallet2/utils"
	"fmt"
	"time"
)

type TransService interface {
	Donasi(form *dto.DonasiReq) error // harus dalam rupiah
	// Bayar
}

type transService struct {
	transRepo repository.TransactionRepository
	userRepo  repository.UserRepository
	userUC    UserUseCase
}

func (ts *transService) Donasi(form *dto.DonasiReq) error {
	fmt.Println(form)
	input := dto.LoginRequestBody{
		Email:    form.Email,
		Password: form.Password,
	}
	user, err := ts.userUC.Login(&input)
	if err != nil {
		return &utils.IncorrectCredentialsError{}
	}
	var buktiTrans *model.Transaction
	receiver, err := ts.userUC.GetUserByID(form.ReceiverId)
	if err != nil {
		buktiTrans = &model.Transaction{
			Transaction_ID:      utils.GenerateId(),
			Userwallet_id:       user.ID,
			Money_Changer_ID:    "MC-001",
			Transaction_Type_ID: "TT-001",
			Payment_method_id:   "PM-001",
			Amount:              float64(form.Amount),
			Status:              false,
			Date_Time:           time.Time{},
		}
		return &utils.UserNotFoundError{}
	} else if form.ReceiverName != receiver.Name {
		return &utils.IncorrectCredentialsError{}
	} else {
		if user.Balance < float64(form.Amount){
			return &utils.NotEnoughBalance{}
		}
		receiver.Balance = receiver.Balance + float64(form.Amount)
		user.Balance = user.Balance - float64(form.Amount)
		//update amount receiver & user
		errU := ts.userRepo.UpdateById(user)
		if errU != nil {
			fmt.Println("gagal update receiver")
			return err
		}
		errR := ts.userRepo.UpdateById(receiver)
		if errR != nil {
			fmt.Println("gagal update receiver")
			return errR
		}
		ts.userRepo.UpdateById(receiver)
		buktiTrans = &model.Transaction{
			Transaction_ID:      utils.GenerateId(),
			Userwallet_id:       user.ID,
			Money_Changer_ID:    "MC-001",
			Transaction_Type_ID: "TT-001",
			Payment_method_id:   "PM-001",
			Amount:              float64(form.Amount),
			Status:              true,
			Date_Time:           time.Time{},
		}
	}
	return ts.transRepo.CreateTrans(buktiTrans)
}

func NewServiceTrans(userRepo repository.UserRepository, userUC UserUseCase, transactionRepo repository.TransactionRepository) TransService {
	return &transService{
		userRepo:  userRepo,
		userUC:    userUC,
		transRepo: transactionRepo,
	}
}
