package db

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserSchema struct {
	Pk              string   `dynamodbav:"pk"`
	Gsi1pk          string   `dynamodbav:"gsi1pk"`
	Id              string   `dynamodbav:"id"`
	Username        string   `dynamodbav:"username,omitempty"`
	Email           string   `dynamodbav:"email"`
	EmailIsVerified bool     `dynamodbav:"email_is_verified"`
	Password        string   `dynamodbav:"password"`
	Refresh         string   `dynamodbav:"refresh_token"`
	Scope           string   `dynamodbav:"scope"`
	Policies        []string `dynamodbav:"policies"`
	CreatedAt       int64    `dynamodbav:"created_at"`
	LastLoginAt     int64    `dynamodbav:"last_login_at"`
}

func (u *UserSchema) buildKeys() {
	u.Pk = "user#" + u.Id
	u.Gsi1pk = "email#" + u.Email
}

func (u *UserSchema) getKey() map[string]types.AttributeValue {
	u.buildKeys()
	k := make(map[string]types.AttributeValue)
	k["pk"] = &types.AttributeValueMemberS{Value: u.Pk}
	return k
}

func (u *UserSchema) getGsi() map[string]types.AttributeValue {
	u.buildKeys()
	k := make(map[string]types.AttributeValue)
	k[":gsi1pk"] = &types.AttributeValueMemberS{Value: u.Gsi1pk}
	return k
}
