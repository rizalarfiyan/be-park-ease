package utils

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

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
