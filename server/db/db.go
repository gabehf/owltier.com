package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mnrva-dev/owltier.com/server/config"
)

var (
	handle Handle
)

const (
	GSI_NAME = "gsi1"
	GSI_ATTR = "gsi1pk"
)

type Handle struct {
	client  *dynamodb.Client
	table   string
	gsiName string
	gsiAttr string
}

func init() {
	handle = DBConfig(
		config.Environment(),
		GSI_NAME,
		GSI_ATTR,
	)
}

func CreateLocalClient() *dynamodb.Client {
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion("us-east-1"),
		awsconfig.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
		awsconfig.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

func DBConfig(env, gsiname, gsiattr string) Handle {
	var c *dynamodb.Client

	if env == "local" {
		c = CreateLocalClient()
		log.Println("* Local Environment Detected")
	}
	return Handle{
		client:  c,
		table:   config.DbTable(),
		gsiName: gsiname,
		gsiAttr: gsiattr,
	}
}
