package repository

//go:generate mockgen -destination=mocks/mock_tokens_repository.go -package=mocks . ITokensRepository

import (
	"errors"
	"strconv"

	"github.com/boltdb/bolt"
)

type Bucket string

const (
	BucketAccessTokens  Bucket = "access_tokens"
	BucketRequestTokens Bucket = "request_tokens"
)

type ITokensRepository interface {
	Save(chatID int64, token string, bucket Bucket) error
	Get(chatID int64, bucket Bucket) (string, error)
}

type tokensRepository struct {
	db *bolt.DB
}

func NewTokensRepository(db *bolt.DB) ITokensRepository {
	return &tokensRepository{db: db}
}

func (r *tokensRepository) Save(chatID int64, token string, bucket Bucket) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put([]byte(strconv.FormatInt(chatID, 10)), []byte(token))
	})
}

func (r *tokensRepository) Get(chatID int64, bucket Bucket) (string, error) {
	var token string
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		data := b.Get([]byte(strconv.FormatInt(chatID, 10)))

		token = string(data)

		return nil
	})

	if err != nil {
		return "", err
	}
	if token == "" {
		return "", errors.New("token not found")
	}

	return token, nil
}
