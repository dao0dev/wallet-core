package coin

import (
	"database/sql/driver"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"nexus-wallet/pkg/repository"
)

type Repository struct {
	baseRepository *repository.BaseRepository
}

func NewRepository(baseRepository *repository.BaseRepository) *Repository {
	return &Repository{baseRepository: baseRepository}
}

func (r Repository) FindAllMappedByIds(tx driver.Tx) (map[int64]Coin, error) {
	items, err := repository.FindManyBy(r.baseRepository, repository.FindManyByOptions{
		Expression: goqu.Ex{},
		Limit:      50000,
		Offset:     0,
	}, &Coin{}, tx)
	if err != nil {
		return nil, fmt.Errorf("error find coins %s", err)
	}
	coinsById := make(map[int64]Coin)
	for _, coinData := range items {
		coinsById[coinData.Id] = *coinData
	}

	return coinsById, nil
}
