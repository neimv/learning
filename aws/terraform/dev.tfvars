ami_id="ami id"
instance="t2.micro"
tags={Name="practica1", Environment="Dev"}
sg_name="platzi-rules"
ingress_rules = [
    {
        from_port = "22"
        to_port = "22"
    },
]
