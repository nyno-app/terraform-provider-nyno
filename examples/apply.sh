#!/bin/bash
# Run with examples/apply.sh

cd terraform-provider-nyno
go build -o terraform-provider-nyno
make install
cd ../examples
rm -rf ".terraform"
rm -f ".terraform.lock.hcl" "terraform.tfstate"
terraform init
terraform apply --auto-approve