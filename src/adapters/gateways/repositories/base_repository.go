package repositories

import "go_gift_list_api/src/entities"

type BaseRepository interface {
	StartTransaction() error
	CommitTransaction() error
	CancelTransaction() error
	NextEntityID() entities.UniqueEntityID
}
