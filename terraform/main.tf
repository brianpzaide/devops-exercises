resource "aws_vpc" "mtc_vpc"{
  cidr_block = "192.168.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support = true
  tags{
    name="dev"
  }
}

resource "aws_subnet" "mtc_subnet"{
  vpc_id = aws_vpc.mtc_vpc.id
  cidr_block = "192.168.1.0/16"
  map_public_ip_on_launch = true
  availability_zone = "us-west-2a"
  tags{
    name="dev-public"
  }
}

resource "aws_internet_gateway" "mtc_igw"{
  vpc_id = aws_vpc.mtc_vpc.id
  tags{
    name="dev-igw"
  }
}

resource "aws_route_table" "mtc_rt"{
  vpc_id = aws_vpc.mtc_vpc.id
  tags{
    name="dev-rt"
  }
}

resource "aws_route" "mtc_route"{
  route_table_id = aws_route_table.mtc_rt.id
  destination_id = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.mtc_igw.id
}

resource "aws_route_table_association" "mtc_public_assoc"{
  subnet_id = aws_subnet.mtc_subnet.id
  route_table_id = aws_route_table.mtc_rt.id
}

resource "aws_key_pair" "mtc_auth"{
  key_pair = "mtckey"
  public_key = file("~/.ssh/mtckey.pub")
}

resource "aws_security_group" "mtc_asg"{
  name = "dev_sg"
  description = "dev security group"
  vpc_id = aws_vpc.mtc_vpc.id
  ingress {
    from_port = 0
    to_port = 80
    protocol = "TCP"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port = 22
    to_port = 22
    protocol = "TCP"
    cidr_blocks = ["<my dev machine ip>"]
  }
  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "mtc_node"{
  ami = ""
  instance_type = ""
  tags{
    name = "dev-node"
  }
  key_name = aws_key_pair.mtc_auth.id
  vpc_security_group_ids = [aws_security_group.mtc_sg.id]
  subnet_id = aws_subnet_id.mtc_subnet.id
  root_block_device{
    volume_size = 10
  }  
}






