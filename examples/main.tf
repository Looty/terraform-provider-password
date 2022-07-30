terraform {
  required_providers {
    password = {
      version = "0.0.1"
      source  = "hashicorp.com/edu/password"
    }
  }
}

provider "password" {

}

data "password_generator" "secret" {
  group = "asd"

  recipe {
    length       = 31
    num_digits   = 2
    num_symbols  = 21
    allow_upper  = true
    allow_repeat = false
  }
}

output "example_secret" {
  value = data.password_generator.secret.value
}