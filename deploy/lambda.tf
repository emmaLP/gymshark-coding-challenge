
resource "aws_lambda_function" "pack_calc" {
  depends_on       = [aws_cloudwatch_log_group.pack_calc_lambda]
  function_name    = "gymshark-pack-calc"
  handler          = "main"
  role             = aws_iam_role.lambda.arn
  runtime          = "go1.x"
  filename         = data.archive_file.pack_calc.output_path
  source_code_hash = data.archive_file.pack_calc.output_base64sha256

}
resource "aws_cloudwatch_log_group" "pack_calc_lambda" {
  name              = "/aws/lambda/gymshark-pack-calc"
  retention_in_days = 14
}

data "archive_file" "pack_calc" {
  type        = "zip"
  source_file = "${path.module}/../backend/main"
  output_path = "${path.module}/.terraform/archives/gymshark-pack-calc.zip"
}

data "archive_file" "token_authorizer" {
  type        = "zip"
  source_file = "${path.module}/files/token_authoriser.js"
  output_path = "${path.module}/.terraform/archives/token_authorizer.zip"
}

resource "aws_lambda_permission" "apigateway" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.pack_calc.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.pack_calc.execution_arn}/*/*/*"
}


resource "aws_lambda_function" "token_authoriser" {
  depends_on    = [aws_cloudwatch_log_group.token_authoriser]
  function_name = "gymshark-pack-calc-authorizer"
  handler       = "token_authoriser.handler"
  role          = aws_iam_role.lambda.arn
  runtime       = "nodejs12.x"
  environment {
    variables = {
      SECRET_KEY = aws_ssm_parameter.api_token.name
    }
  }
  filename         = data.archive_file.token_authorizer.output_path
  source_code_hash = data.archive_file.token_authorizer.output_base64sha256

}
resource "aws_cloudwatch_log_group" "token_authoriser" {
  name              = "/aws/lambda/gymshark-pack-calc-authorizer"
  retention_in_days = 14
}

resource "aws_lambda_permission" "apigateway_auth" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.token_authoriser.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.pack_calc.execution_arn}/authorizers/${aws_apigatewayv2_authorizer.authoriser.id}"
}