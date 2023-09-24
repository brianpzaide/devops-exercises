data "aws_ami" "server-ami"{
  most_resent=true
  owners = ["amazon"]
  filter {
    name= "name"
    values = ["amzn2-ami-hvm-*"]
  }
}