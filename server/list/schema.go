package list

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

type List struct {
	Pk        string   `json:"-" dynamodbav:"pk"`
	Id        string   `json:"id" dynamodbav:"id"`
	CreatedAt int64    `json:"created_at" dynamodbav:"created_at"`
	CreatedBy string   `json:"created_by" dynamodbav:"created_by"`
	Format    string   `json:"format" dynamodbav:"format"`
	Breaks    []bool   `json:"breaks" dynamodbav:"breaks"`
	NA        []string `json:"na,omitempty" dynamodbav:"na,omitempty"`
	APAC      []string `json:"apac,omitempty" dynamodbav:"apac,omitempty"`
	Combined  []string `json:"combined,omitempty" dynamodbav:"combined,omitempty"`
}

func (u *List) BuildKeys() {
	u.Pk = "list#" + u.Id
}

func (u *List) GetKeys() map[string]types.AttributeValue {
	u.BuildKeys()
	k := make(map[string]types.AttributeValue)
	k["pk"] = &types.AttributeValueMemberS{Value: u.Pk}
	return k
}

func (u *List) GetGsi() map[string]types.AttributeValue {
	return nil
}
