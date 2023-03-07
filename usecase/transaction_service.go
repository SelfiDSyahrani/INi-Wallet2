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
	Donasi(form *dto.DonasiReq) (string, error) // harus dalam rupiah
	Bayar(form *dto.BayarReq) (string, error)
	IsiUlang(form *dto.TopUpReq) (string, error)
}

type transService struct {
	transRepo repository.TransactionRepository
	userRepo  repository.UserRepository
	userUC    UserUseCase
	MoneyUC   MoneyChangerUsecase
}

func (ts *transService) Donasi(form *dto.DonasiReq) (string, error) {
	input := dto.LoginRequestBody{
		Email:    form.Email,
		Password: form.Password,
	}
	user, err := ts.userUC.Login(&input)
	if err != nil {
		return err.Error(), &utils.IncorrectCredentialsError{}
	}
	var buktiTrans *model.Transaction
	receiver, err := ts.userUC.GetUserByID(form.ReceiverId)
	if err != nil {
		buktiTrans = &model.Transaction{
			Transaction_ID:      utils.GenerateId(),
			Userwallet_id:       user.ID,
			Money_Changer_ID:    "MC001",
			Transaction_Type_ID: "TT001",
			Payment_method_id:   "PM001",
			Amount:              float64(form.Amount),
			Status:              false,
			Date_Time:           time.Now(),
		}
		return buktiTrans.Transaction_ID, ts.transRepo.CreateTrans(buktiTrans)
	} else if form.ReceiverName != receiver.Name {
		return err.Error(), &utils.IncorrectCredentialsError{}
	} else {
		if user.Balance < float64(form.Amount) {
			return err.Error(), &utils.NotEnoughBalance{}
		}
		receiver.Balance = receiver.Balance + float64(form.Amount)
		user.Balance = user.Balance - float64(form.Amount)
		//update amount receiver & user
		errU := ts.userRepo.UpdateById(user)
		if errU != nil {
			fmt.Println("gagal update receiver")
			return err.Error(), errU
		}
		errR := ts.userRepo.UpdateById(receiver)
		if errR != nil {
			fmt.Println("gagal update receiver")
			return err.Error(), errR
		}
		buktiTrans = &model.Transaction{
			Transaction_ID:      utils.GenerateId(),
			Userwallet_id:       user.ID,
			Money_Changer_ID:    "MC001",
			Transaction_Type_ID: "TT001",
			Payment_method_id:   "PM001",
			Amount:              float64(form.Amount),
			Status:              true,
			Date_Time:           time.Now(),
		}
	}
	return buktiTrans.Transaction_ID, ts.transRepo.CreateTrans(buktiTrans)
}

func (ts *transService) IsiUlang(form *dto.TopUpReq) (string, error) {
	input := &dto.LoginRequestBody{
		Email:    form.Email,
		Password: form.Password,
	}
	user, err := ts.userUC.Login(input)
	if err != nil {
		return err.Error(), &utils.IncorrectCredentialsError{}
	}
	user.Balance = user.Balance + float64(form.Amount)
	errU := ts.userRepo.UpdateById(user)
	if errU != nil {
		return errU.Error(), errU
	}
	buktiTrans := &model.Transaction{
		Transaction_ID:      utils.GenerateId(),
		Userwallet_id:       user.ID,
		Money_Changer_ID:    "MC001",
		Transaction_Type_ID: "TT002",
		Payment_method_id:   form.Method_id,
		Amount:              float64(form.Amount),
		Status:              true,
		Date_Time:           time.Now(),
	}
	return buktiTrans.Transaction_ID, ts.transRepo.CreateTrans(buktiTrans)
}

// Bayar implements TransService
func (ts *transService) Bayar(form *dto.BayarReq) (string, error) {
	input := dto.LoginRequestBody{
		Email:    form.Email,
		Password: form.Password,
	}
	user, err := ts.userUC.Login(&input)
	if err != nil {
		return err.Error(), &utils.IncorrectCredentialsError{}
	}
	if form.Currency == "IDR" {
		Money, err := ts.MoneyUC.MoneyChangerById("MC001")
		fmt.Println(err)
		if err != nil {
			return err.Error(), err
		}
		user.Balance = user.Balance - float64(form.Amount)*Money.Exchange_rate
		errU := ts.userRepo.UpdateById(user)
		if errU != nil {
			return errU.Error(), errU
		}
		buktiTrans := &model.Transaction{
			Transaction_ID:      utils.GenerateId(),
			Userwallet_id:       user.ID,
			Money_Changer_ID:    "MC001",
			Transaction_Type_ID: "TT003",
			Payment_method_id:   "PM003",
			Amount:              float64(form.Amount),
			Status:              true,
			Date_Time:           time.Now(),
		}
		return buktiTrans.Transaction_ID, ts.transRepo.CreateTrans(buktiTrans)
	} else {
		usd, err := ts.MoneyUC.MoneyChangerById("MC002")
		fmt.Println(err, usd)
		if err != nil {
			return err.Error(), err
		}
		user.Balance = user.Balance - float64(form.Amount)*usd.Exchange_rate
		errU := ts.userRepo.UpdateById(user)
		if errU != nil {
			return errU.Error(), errU
		}
		buktiTrans := &model.Transaction{
			Transaction_ID:      utils.GenerateId(),
			Userwallet_id:       user.ID,
			Money_Changer_ID:    "MC002",
			Transaction_Type_ID: "TT003",
			Payment_method_id:   "PM003",
			Amount:              float64(form.Amount),
			Status:              true,
			Date_Time:           time.Now(),
		}
		return buktiTrans.Transaction_ID, ts.transRepo.CreateTrans(buktiTrans)
	}
}

func NewServiceTrans(userRepo repository.UserRepository, userUC UserUseCase, TransRepo repository.TransactionRepository, MoneyUC MoneyChangerUsecase) TransService {
	return &transService{
		transRepo: TransRepo,
		userRepo:  userRepo,
		userUC:    userUC,
		MoneyUC:   MoneyUC,
	}
}
