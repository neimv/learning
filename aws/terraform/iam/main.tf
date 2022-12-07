provider "aws" {
    region="us-east-2"
}

resource "aws_iam_user" "myUser" {
    name = "TJ"
}

resource "aws_iam_policy" "customPolicy" {
    name = "GlaciesEFSEC2"

    policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "id",
            "Effect": "Allow",
            "Action": [s3:*]
            "Resource": *
        }
    ]
}
    EOF
}

resource "aws_iam_policy_attachment" "policyBind" {
    name = "attachment"
    user = [aws_iam_user.myUser.name]
    policy_arn = aws_iam_policy.customPolicy
}