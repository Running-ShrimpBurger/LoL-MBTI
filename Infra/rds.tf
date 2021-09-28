module "security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4"

  name        = "mbti-db"
  description = "MySQL security group"
  vpc_id      = module.vpc.vpc_id

  # ingress
  ingress_with_cidr_blocks = [
    {
      from_port   = 3306
      to_port     = 3306
      protocol    = "tcp"
      description = "MySQL access from within VPC"
      cidr_blocks = module.vpc.vpc_cidr_block
    },
  ]
}

module "db" {
  source  = "terraform-aws-modules/rds/aws"
  version = "~> 3.0"

  identifier = "mbtidb"

  engine            = "mysql"
  engine_version    = "5.7.19"
  instance_class    = "db.t3.small"
  allocated_storage = 5

  name     = "mbtidb"
  username = var.db_username
  password = var.db_username
  port     = "3306"

  iam_database_authentication_enabled = true

  vpc_security_group_ids = [module.security_group.security_group_id]
  subnet_ids = [module.vpc.database_subnets[0], module.vpc.database_subnets[1]]

  create_db_option_group = false
  create_db_parameter_group = false
}