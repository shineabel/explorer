package utils

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/explorer/backend/logger"
	"math"
	"math/big"
	"strconv"
)

func ParseInt(text string) (i int64, b bool) {
	i, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		return i, false
	}
	return i, true
}

func ParseIntWithDefault(text string, def int64) (i int64) {
	i, err := strconv.ParseInt(text, 10, 0)
	if err != nil {
		logger.Error("ParseIntWithDefault error", logger.String("str", text))
		return def
	}
	return i
}

func ParseUint(text string) (uint64, bool) {
	i, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return i, false
	}
	return i, true
}

func RoundFloat(num float64, bit int) (i float64, b bool) {
	format := "%" + fmt.Sprintf("0.%d", bit) + "f"
	s := fmt.Sprintf(format, num)
	i, err := strconv.ParseFloat(s, 0)
	if err != nil {
		return i, false
	}
	return i, true
}

func Round(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}

func RoundString(x string) int64 {
	f, err := strconv.ParseFloat(x, 0)
	if err != nil {
		logger.Error("RoundString error", logger.String("str", x))
	}
	return int64(math.Floor(f + 0.5))
}

func RoundToString(decimal string, bit int) (i string) {
	f, err := strconv.ParseFloat(decimal, bit)
	if err != nil {
		logger.Error("RoundFloatString error", logger.String("str", decimal))
	}
	return strconv.FormatFloat(f, 'f', bit, 64)
}

func Copy(src interface{}, dest interface{}) error {
	bz, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bz, dest)
	if err != nil {
		return err
	}
	return nil
}

func NewRatFromFloat64(f float64) *big.Rat {
	return new(big.Rat).SetFloat64(f)
}
