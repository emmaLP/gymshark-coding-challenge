resource "aws_apigatewayv2_api" "pack_calc" {
  name          = "gymshark-pack-calc"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "default" {
  api_id      = aws_apigatewayv2_api.pack_calc.id
  name        = "$default"
  auto_deploy = true
}


resource "aws_apigatewayv2_route" "pack_calc" {
  api_id    = aws_apigatewayv2_api.pack_calc.id
  route_key = "POST /calculate-packs"
  target    = "integrations/${aws_apigatewayv2_integration.pack_calc.id}"
}

resource "aws_apigatewayv2_integration" "pack_calc" {
  api_id                 = aws_apigatewayv2_api.pack_calc.id
  integration_type       = "AWS_PROXY"
  connection_type        = "INTERNET"
  integration_uri        = aws_lambda_function.pack_calc.invoke_arn
  integration_method     = "POST"
  payload_format_version = "2.0"
}