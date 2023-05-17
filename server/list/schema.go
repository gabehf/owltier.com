package list

import "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

type List struct {
	Pk        string   `dynamodbav:"pk"`
	Id        string   `json:"id"`
	CreatedAt int64    `json:"created_at"`
	CreatedBy string   `json:"created_by"`
	Format    string   `json:"format"`
	Breaks    []bool   `json:"breaks"`
	NA        []string `json:"na"`
	APAC      []string `json:"apac"`
	Combined  []string `json:"combined"`
}

func (u *List) buildKeys() {
	u.Pk = "list#" + u.Id
}

func (u *List) getKey() map[string]types.AttributeValue {
	u.buildKeys()
	k := make(map[string]types.AttributeValue)
	k["pk"] = &types.AttributeValueMemberS{Value: u.Pk}
	return k
}

func (u *List) getGsi() map[string]types.AttributeValue {
	return nil
}
