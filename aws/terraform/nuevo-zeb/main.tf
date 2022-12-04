
provider "aws" {
    region = "us-east-1"
}

variable "ingressrules" {
    type = list(number)
    default = [80, 443]
}

variable "egressrules" {
    type = list(number)
    default = [80, 443, 25, 3306, 8080]
}

resource "aws_instance" "ec2" {
    ami = "ami-032598fcc7e9d1c7a"
    instance_type = "t2.micro"
    security_groups = [aws_security_group.webtraffic.name]
}

resource "aws_eip" "elasticip" {
    instance = aws_instance.ec2.id
}

resource "aws_security_group" "webtraffic" {
    name = "allow HTTPS"

    dynamic ingress {
        iterator = port
        for_each = var.ingressrules
        content {
            from_port = port.value
            to_port = port.value
            protocol = "TCP"
            cidr_blocks = [
                "0.0.0.0/0"
            ]
        }
    }

    egress {
        iterator = port
        for_each = var.egressrules
        content {
            from_port = port.value
            to_port = port.value
            protocol = "TCP"
            cidr_blocks = [
                "0.0.0.0/0"
            ]
        }
    }
}

output "EIP" {
    value = aws_eip.elasticip.public_ip
}