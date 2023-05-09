package db

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DbItem interface {
	buildKeys()
	getKey() map[string]types.AttributeValue
	getGsi() map[string]types.AttributeValue
}
