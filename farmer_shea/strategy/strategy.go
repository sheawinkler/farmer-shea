package strategy

import (
	"farmer_shea/oracle"
	"sort"
)

// DetermineBestLendingMarket determines the best lending market on Solend based on APY.
func DetermineBestLendingMarket(markets []oracle.SolendMarket) (*oracle.SolendMarket, error) {
	if len(markets) == 0 {
		return nil, nil
	}

	sort.Slice(markets, func(i, j int) bool {
		return markets[i].SupplyApy > markets[j].SupplyApy
	})

	return &markets[0], nil
}
