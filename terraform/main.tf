resource "aws_vpc" "devopsexercises_vpc"{
  cidr_block = "192.168.0.0/16"
}

resource "aws_subnet" "devopsexercises_subnet"{
  vpc_id = aws_vpc.devopsexercises_vpc.id
  cidr_block = "192.168.0.0/24"
  availability_zone = "us-west-2a"
}

resource "aws_internet_gateway" "devopsexercises_igw"{
  vpc_id = aws_vpc.devopsexercises_vpc.id
}

resource "aws_route_table" "devopsexercises_rt"{
  vpc_id = aws_vpc.devopsexercises_vpc.id
}

resource "aws_route" "devopsexercises_route"{
  route_table_id = aws_route_table.devopsexercises_rt.id
  destination_id = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.devopsexercises_igw.id
}

resource "aws_route_table_association" "devopsexercises_public_assoc"{
  subnet_id = aws_subnet.devopsexercises_subnet.id
  route_table_id = aws_route_table.devopsexercises_rt.id
}

resource "aws_key_pair" "devopsexercises_auth"{
  key_pair = "devopsexercises"
  public_key = file("~/.ssh/devopsexercises.pub")
}

resource "aws_security_group" "devopsexercises_sg"{
  name = "dev_sg"
  description = "dev security group"
  vpc_id = aws_vpc.devopsexercises_vpc.id
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

resource "aws_instance" "devopsexercises_node"{
  ami = data.aws_ami.server-ami.id
  instance_type = "t2.micro"
  key_name = aws_key_pair.devopsexercises_auth.id
  vpc_security_group_ids = [aws_security_group.devopsexercises_sg.id]
  subnet_id = aws_subnet_id.devopsexercises_subnet.id
  root_block_device{
    volume_size = 10
  }  
}






