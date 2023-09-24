data "aws_ami" "server-ami"
  most_resent=true
  owners = ["<get aws ami owner>"]
  filter {
    name= "name"
    values = ["ubuntu/images....."]
  }