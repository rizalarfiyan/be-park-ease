package utils

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func Str(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

func PGTimeStamp(date time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{
		Time:  date,
		Valid: true,
	}
}

func PGText(text string) pgtype.Text {
	return pgtype.Text{
		String: text,
		Valid:  text != "",
	}
}

func StrToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func PGNumericFloat64(number float64) pgtype.Numeric {
	heightFloatVal := big.NewFloat(number)
	heightBigInt := new(big.Int)
	heightFloatVal.Int(heightBigInt)

	return pgtype.Numeric{
		Int:   heightBigInt,
		Valid: true,
	}
}

func PGInt32(number int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: number,
		Valid: true,
	}
}

func StrToBool(str string) (bool, error) {
	switch strings.ToLower(str) {
	case "true", "t", "yes", "y", "1":
		return true, nil
	case "false", "f", "no", "n", "0":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value: %s", str)
	}
}
