data "aws_iam_policy" "lambda_execution" {
  arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
data "aws_iam_policy" "ssm_readonly" {
  arn = "arn:aws:iam::aws:policy/AmazonSSMReadOnlyAccess"
}