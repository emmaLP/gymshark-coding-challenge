
resource "aws_ssm_parameter" "api_token" {
  name   = "gmyshark-api-token"
  type   = "SecureString"
  value  = random_password.api_token.result
}

resource "random_password" "api_token" {
  length           = 16
  special          = true
  upper            = true
  lower            = true
  override_special = "!@#$_-"
}