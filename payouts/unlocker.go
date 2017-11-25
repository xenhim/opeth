package payouts

import (
	"strings"

	"github.com/EMoneyPools/opeth/payouts/unlocker"
	"github.com/EMoneyPools/opeth/payouts/unlocker/pplns"
	"github.com/EMoneyPools/opeth/payouts/unlocker/pps"
	"github.com/EMoneyPools/opeth/storage"
)

type BlockUnlocker struct {
	Config  *unlocker.UnlockerConfig
	Backend *storage.RedisClient
}

func NewBlockUnlocker(cfg *unlocker.UnlockerConfig, backend *storage.RedisClient) *BlockUnlocker {
	return &BlockUnlocker{
		Config:  cfg,
		Backend: backend,
	}
}
func (u *BlockUnlocker) Start() {
	switch strings.ToLower(u.Config.PayoutModel) {
	case "pps":
		BlockUnlock_pps := pps.NewBlockUnlocker(u.Config, u.Backend)
		BlockUnlock_pps.Start()
	case "pplns":
		BlockUnlock_pplns := pplns.NewBlockUnlocker(u.Config, u.Backend)
		BlockUnlock_pplns.Start()
	}
}
