package repository

import (
	"context"
	"errors"

	"github.com/Akshayvij07/ecommerce/pkg/domain"
)

func (c *OrderDataBase) FindWalletByUserID(ctx context.Context, UserId int) (domain.Wallet, error) {
	var wallet domain.Wallet
	Find := `SELECT * FROM wallets WHERE user_id = ?`
	if c.DB.Raw(Find, UserId).Scan(&wallet).Error != nil {
		return wallet, errors.New("no cart found using this user_id")
	}
	return wallet, nil
}

func (c *OrderDataBase) SaveWallet(ctx context.Context, UserId int) (uint, error) {
	var wallet domain.Wallet
	Insert := `INSERT INTO wallets (user_id,total_amount) 
	VALUES ($1,$2) RETURNING Id`
	if c.DB.Raw(Insert, UserId, 0).Scan(&wallet).Error != nil {
		return 0, errors.New("failed to create the wallet")
	}
	return wallet.ID, nil
}

//func (c *OrderDataBase)
