package items_keeper

import (
	"fmt"
)

func (s *ItemsKeeperService) StartBot() error {
	err := s.tgClient.Start()
	if err != nil {
		return fmt.Errorf("failed to starting telegram client: %s", err.Error())
	}

	return nil
}
