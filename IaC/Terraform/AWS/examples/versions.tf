terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
#  profile = "your profile"
  region = "ap-northeast-2"
}

# Concfigure the vshpher Provider
# resource "vsphere_license" "licenseKey" {
#   license_key = "452CQ-2EK54-K8742-00000-00000"

#   labels {
#     VpxClientLicenseLabel = "Hello World"
#     Workflow              = "Hello World"
#   }

# }