package payment

import (
	"errors"
	models "payment/internal/payment/models"
	appErr "payment/pkg/errors"

	"gorm.io/gorm"
)

type PaymentRepo struct {
	db *gorm.DB
}
// var _ <INTERFACE> = (*<STRUCT>)(nil)
var _ PaymentRepository = (*PaymentRepo)(nil)

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{
		db: db,
	}
}


func (r *PaymentRepo) Create(p * models.Payment) error {
	payment := models.Payment{
		ID: p.ID,
		UserID: p.UserID,
		Amount: p.Amount,
		Status: p.Status,
	}

	err := r.db.Create(&payment).Error
	if err != nil {
		return err  
	}

	return nil 
}

func (r *PaymentRepo) GetAll()([]models.Payment,error){
	var payments []models.Payment

	err := r.db.Find(&payments).Error
	if err != nil{
		return nil,err 
	}
	return payments, nil 
}

func (r *PaymentRepo) Update(id int,status models.PaymentStatus) (models.Payment,error){
	var payment models.Payment
	err := r.db.First(&payment,id).Error 
	if err != nil{
		if errors.Is(err,gorm.ErrRecordNotFound){
			return models.Payment{},appErr.ErrNotFound
		}
	}

	payment.Status = status 
	err = r.db.Save(&payment).Error
	if err != nil{
		return models.Payment{},err
	}
	return models.Payment{},nil
}