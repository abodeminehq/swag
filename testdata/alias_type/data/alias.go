package data

import (
	"time"

	"github.com/abodeminehq/swag/v2/testdata/alias_type/types"
)

type TimeContainer struct {
	Name      types.StringAlias `json:"name"`
	Timestamp time.Time         `json:"timestamp"`
	CreatedAt types.DateOnly    `json:"created_at"`
}
