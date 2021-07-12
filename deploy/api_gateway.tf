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
  api_id             = aws_apigatewayv2_api.pack_calc.id
  authorization_type = "CUSTOM"
  authorizer_id      = aws_apigatewayv2_authorizer.authoriser.id
  route_key          = "POST /calculate-packs"
  target             = "integrations/${aws_apigatewayv2_integration.pack_calc.id}"
}

resource "aws_apigatewayv2_integration" "pack_calc" {
  api_id                 = aws_apigatewayv2_api.pack_calc.id
  integration_type       = "AWS_PROXY"
  connection_type        = "INTERNET"
  integration_uri        = aws_lambda_function.pack_calc.invoke_arn
  integration_method     = "POST"
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_authorizer" "authoriser" {
  api_id          = aws_apigatewayv2_api.pack_calc.id
  authorizer_type = "REQUEST"
  authorizer_uri  = aws_lambda_function.token_authoriser.invoke_arn

  identity_sources                  = ["$request.header.Authorization"]
  authorizer_payload_format_version = "2.0"
  name                              = "gymshark-token-authoriser"
  enable_simple_responses           = true
  authorizer_result_ttl_in_seconds  = 0
}