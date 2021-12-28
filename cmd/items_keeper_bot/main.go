package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/olezhek28/items-keeper/internal/app/config"
	"github.com/olezhek28/items-keeper/internal/app/repository"
	"github.com/olezhek28/items-keeper/internal/app/service/auth"
	"github.com/olezhek28/items-keeper/internal/app/service/items_keeper"
	"github.com/olezhek28/items-keeper/internal/pkg/api/pocket"
	"github.com/olezhek28/items-keeper/internal/pkg/api/telegram"
	pocketSDK "github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	tgClient, err := getTgClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	pocketClient, err := getPocketClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	tokensRepository, err := getDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	itemsKeeperService := items_keeper.NewItemsKeeperService(tgClient, pocketClient, tokensRepository)
	authService := auth.NewAuthService(pocketClient, tokensRepository, cfg.TelegramBotURL)

	go func() {
		err = itemsKeeperService.Process()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = authService.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func getTgClient(cfg *config.Config) (telegram.ITelegramClient, error) {
	tgBot, err := tgBotAPI.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, fmt.Errorf("failed to creating new tg client: %v", err)
	}

	return telegram.NewClient(tgBot), nil
}

func getPocketClient(cfg *config.Config) (pocket.IPocketClient, error) {
	pocketSDKClient, err := pocketSDK.NewClient(cfg.PocketConsumerKey)
	if err != nil {
		return nil, fmt.Errorf("failed to creating new pocket client: %v", err)
	}

	return pocket.NewClient(pocketSDKClient, cfg.AuthServerURL), nil
}

func getDB(cfg *config.Config) (repository.ITokensRepository, error) {
	db, err := initDB(cfg)
	if err != nil {
		return nil, err
	}

	return repository.NewTokensRepository(db), nil
}

func initDB(cfg *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cfg.DBPath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to opening db: %s", err.Error())
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, errBucket := tx.CreateBucketIfNotExists([]byte(repository.BucketAccessTokens))
		if errBucket != nil {
			return errBucket
		}

		_, errBucket = tx.CreateBucketIfNotExists([]byte(repository.BucketRequestTokens))
		if errBucket != nil {
			return errBucket
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to creating buckets in db: %s", err.Error())
	}

	return db, nil
}
