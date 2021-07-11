#!/usr/bin/env bash

BUCKET="emma-gymshark-challenge-terraform-state"
DYNAMODB_TABLE="terraform-state-lock"
REGION="eu-west-2"
PROFILE="default"
PROFILE_INFO="--profile ${PROFILE}"

BUCKET_EXISTS=$(aws s3api head-bucket --bucket "${BUCKET}" ${PROFILE_INFO} 2>&1)
if [[ "${BUCKET_EXISTS}" == '' ]]; then
    echo 'Bucket already exists'
else
    aws s3api create-bucket --bucket ${BUCKET} --create-bucket-configuration LocationConstraint=${REGION} ${PROFILE_INFO}
    #Wait for bucket creation
    sleep 5
    echo "S3 ${BUCKET} is created"
fi
echo "Configuring bucket security and other policies"
aws s3api put-bucket-versioning --bucket ${BUCKET} --versioning-configuration Status=Enabled  ${PROFILE_INFO}
aws s3api put-bucket-encryption --bucket ${BUCKET} \
    --server-side-encryption-configuration '{ "Rules": [{ "ApplyServerSideEncryptionByDefault": { "SSEAlgorithm": "AES256" }}]}'  ${PROFILE_INFO}
aws s3api put-public-access-block --bucket ${BUCKET} \
  --public-access-block-configuration BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=true,RestrictPublicBuckets=true

DYNAMO_DB_EXISTS=$(aws --profile "${PROFILE}" dynamodb describe-table --table-name "${DYNAMODB_TABLE}" --output text --query 'Table.TableName' --region "${REGION}" 2>&1)
if [[ "${DYNAMO_DB_EXISTS}" == "${DYNAMODB_TABLE}" ]]; then
    echo 'DynamoDB Table already exists'
else
    echo "create dynamodb_table ${DYNAMODB_TABLE}"
    aws dynamodb create-table --table-name ${DYNAMODB_TABLE} \
          --attribute-definitions AttributeName=LockID,AttributeType=S \
          --key-schema AttributeName=LockID,KeyType=HASH \
          --provisioned-throughput ReadCapacityUnits=20,WriteCapacityUnits=20 --region $REGION  ${PROFILE_INFO} --output text --query 'TableDescription.TableName'

    sleep 25
    STATUS=$(aws dynamodb describe-table --table-name ${DYNAMODB_TABLE} --output text --query 'Table.TableStatus' --region $REGION  ${PROFILE_INFO})
    echo "DynamoDB table status: $STATUS"
fi