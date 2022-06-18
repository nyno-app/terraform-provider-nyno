#!/bin/bash
cd terraform-provider-nyno
go build -o terraform-provider-nyno
make install
cd ..
rm -rf ".terraform"
rm -f ".terraform.lock.hcl"
terraform init
terraform destroy --auto-approve