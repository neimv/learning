provider "aws" {
    region = "us-east-2"
}

resource "vpc" "test" {
    cidr_block = "10.0.0.0/16"
}
