package payment

import (
	"testing"

	"log"

	models "payment/internal/payment/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	dsn := "postgresql://postgres.xgkgtnwfcpnwtelkkjmo:golang9890562214@aws-1-ap-southeast-1.pooler.supabase.com:6543/postgres?pgbouncer=true"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect test db: %v", err)
	}

	// Clean slate each run
	err = db.Migrator().DropTable(&models.Payment{})
	if err != nil {
		log.Fatalf("failed to drop table: %v", err)
	}

	err = db.AutoMigrate(&models.Payment{})
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return db
}

func TestPaymentRepo_Create(t *testing.T) {
	db := setupTestDB() // your helper
	repo := NewPaymentRepo(db)

	p := &models.Payment{
		UserID: 1,
		Amount: 100,
		Status: models.StatusCreated,
	}

	res, err := repo.Create(p)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res.ID == 0 {
		t.Error("expected ID to be generated")
	}
}
