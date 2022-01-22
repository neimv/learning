
variable "ami_id" {
    default=""
    description="Ami id"
    type=string | map | list
}

variable "instance_type" {}

variable "tags" {
    type=map
}
