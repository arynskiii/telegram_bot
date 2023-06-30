package boltdb

import (
	"errors"
	"strconv"
	"telbot/pkg/repository"

	"github.com/boltdb/bolt"
)

type TokenRepository struct {
	db *bolt.DB
}

func NewTokenRepository(db *bolt.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) Save(chatID int64, token string, bucket repository.Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.Put(intToByte(chatID), []byte(token))
		return nil
	})
}

func (r *TokenRepository) Get(chatID int64, bucket repository.Bucket) (string, error) {
	var token string
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get(intToByte(chatID))
		token = string(data)
		return nil

	})
	if err != nil {
		return "", err
	}
	if token == "" {
		return "", errors.New("Token not found")
	}
	return token, nil
}

func intToByte(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}