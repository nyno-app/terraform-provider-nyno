#!/bin/bash
# Run with examples/apply.sh

cd internal/provider
go build -o terraform-provider-nyno
make install
cd ../../scripts
rm -rf ".terraform"
rm -f ".terraform.lock.hcl" "terraform.tfstate"
terraform init
terraform apply --auto-approve