module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = "2.21.0"

  name = var.vpc_name
  cidr = var.vpc_cidr

  azs = var.vpc_azs
  private_subnets = var.vpc_private_subnets
  public_subnets = var.vpc_public_subnets
  database_subnets = var.vpc_db_subnets

  create_database_subnet_group = true

  enable_nat_gateway = true
  enable_vpn_gateway = true

  default_vpc_enable_dns_hostnames = true
  enable_dns_hostnames = true
}
