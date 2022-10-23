package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BetChartKeyPrefix is the prefix to retrieve all BetChart
	BetChartKeyPrefix = "BetChart/value/"
)

// BetChartKey returns the store key to retrieve a BetChart from the index fields
func BetChartKey(
	accountName string,
) []byte {
	var key []byte

	accountNameBytes := []byte(accountName)
	key = append(key, accountNameBytes...)
	key = append(key, []byte("/")...)

	return key
}
