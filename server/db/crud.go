package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Create(i DbItem) error {

	i.BuildKeys()

	av, err := attributevalue.MarshalMap(i)
	if err != nil {
		return err
	}
	_, err = handle.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(handle.table),
		Item:      av,
	})
	if err != nil {
		return err
	}
	return nil
}

// out must be a non-nil pointer
func Fetch(i DbItem, out interface{}) error {

	i.BuildKeys()

	o, err := handle.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(handle.table),
		Key:       i.GetKeys(),
	})
	if err != nil {
		return err
	} else if o.Item == nil {
		return NotFoundError(errors.New("item not found"))
	}

	err = attributevalue.UnmarshalMap(o.Item, out)
	if err != nil {
		return err
	}
	return nil

}

func FetchByGsi(i DbItem, out interface{}) error {

	i.BuildKeys()

	o, err := handle.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String(handle.table),
		IndexName:                 aws.String(handle.gsiName),
		KeyConditionExpression:    aws.String(fmt.Sprintf("%s = :%s", handle.gsiAttr, handle.gsiAttr)),
		ExpressionAttributeValues: i.GetGsi(),
	})
	if err != nil {
		return err
	}
	if o.Count > 1 {
		return MultipleItemsError(errors.New("multiple items found"))
	} else if o.Count < 1 {
		return NotFoundError(errors.New("item not found"))
	}

	err = attributevalue.UnmarshalMap(o.Items[0], out)
	if err != nil {
		return err
	}
	return nil
}

func Update(i DbItem, key string, value interface{}) error {
	val, err := attributevalue.Marshal(value)
	if err != nil {
		return err
	}
	_, err = handle.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:        aws.String(handle.table),
		Key:              i.GetKeys(),
		UpdateExpression: aws.String(fmt.Sprintf("set %s = :val", key)),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":val": val,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func MultiUpdate(i DbItem, keys []string, values []interface{}) error {
	// TODO implement this
	return errors.New("not implemented")
}

func Delete(i DbItem) error {

	k := i.GetKeys()

	_, err := handle.client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(handle.table),
		Key:       k,
	})
	if err != nil {
		return err
	}
	return nil

}
