#!/bin/bash
aws dynamodb delete-table --table-name 'owltier-local' --endpoint-url http://localhost:8000 --output json