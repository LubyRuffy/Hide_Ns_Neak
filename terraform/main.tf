
	terraform {
		backend "s3" {
		  bucket         = "hidensneak-terraform"
		  key            = "filename.tfstate"
		  dynamodb_table = "terraform-state-lock-dynamo"
		  region         = "us-east-1"
		  encrypt        = true
		}
	  }

	module "ec2deploy1" {
	source          = "modules/ec2-deployment"
	default_sg_name = ""
	aws_sg_id       = ""
  
	#Example of region_count
	region_count         = "${map("us-east-1","2", "us-east-2","1")}"
	custom_ami           = ""
	aws_instance_type    = "t2.micro"
	ec2_default_user     = "ubuntu"
	aws_access_key       = "${var.aws_access_key}"
	aws_secret_key       = "${var.aws_secret_key}"
	aws_keypair_name     = "do_rsa"
	aws_private_key_file = "/Users/mike.hodges/.ssh/do_rsa"
	aws_public_key_file  = "/Users/mike.hodges/.ssh/do_rsa"
  }

