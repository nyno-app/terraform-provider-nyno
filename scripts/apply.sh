#!/bin/bash
# Run with examples/apply.sh

cd internal/provider
pwd
echo "Building"
go build -o terraform-provider-nyno
echo "Making"
cd ../..
make install
cd scripts
rm -rf ".terraform"
rm -f ".terraform.lock.hcl" "terraform.tfstate"
terraform init
terraform apply --auto-approve