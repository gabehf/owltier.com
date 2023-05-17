package db

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserSchema struct {
	Pk          string `dynamodbav:"pk"`
	Gsi1pk      string `dynamodbav:"gsi1pk"`
	Session     string `dynamodbav:"session_key"`
	Username    string `dynamodbav:"username"`
	Password    string `dynamodbav:"password"`
	CreatedAt   int64  `dynamodbav:"created_at"`
	LastLoginAt int64  `dynamodbav:"last_login_at"`
}

func (u *UserSchema) BuildKeys() {
	u.Pk = "user#" + u.Username
	u.Gsi1pk = "session_key#" + u.Session
}

func (u *UserSchema) GetKeys() map[string]types.AttributeValue {
	u.BuildKeys()
	k := make(map[string]types.AttributeValue)
	k["pk"] = &types.AttributeValueMemberS{Value: u.Pk}
	return k
}

func (u *UserSchema) GetGsi() map[string]types.AttributeValue {
	u.BuildKeys()
	k := make(map[string]types.AttributeValue)
	k[":gsi1pk"] = &types.AttributeValueMemberS{Value: u.Gsi1pk}
	return k
}
