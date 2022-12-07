terraform {
    backend "s3" {
        key = "terraform/tfstate.tfstate"
        bucket = "name-bucket"
        region = "us-east-1"
        access_key = "some"
        secret_key = "another"
    }
}