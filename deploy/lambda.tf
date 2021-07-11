
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

resource "aws_lambda_permission" "apigateway" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.pack_calc.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.pack_calc.execution_arn}/*/*/*"
}