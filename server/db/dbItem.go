package db

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DbItem interface {
	BuildKeys()
	GetKeys() map[string]types.AttributeValue
	GetGsi() map[string]types.AttributeValue
}
