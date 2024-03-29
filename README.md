# devops-exercises

This repo contains my solutions for [Devops-Exercises](https://github.com/DivvyPayHQ/devops-exercise).
I enjoyed working on this project. This project involved working with the following technologies:
* Python
* Golang
* Docker
* Kubernetes (k3s)
* Helm
* Ansible
* Terraform

[!NOTE]
For the exercises:
* I opted for K3s instead of Minikube due to its lightweight and resource-efficient nature, among its numerous other advantages.
* I have not yet tested my Terraform configuration, specifically the creation of an AWS EC2 instance. All the exercises were conducted on two virtual machines (a master and a slave) using Vagrant.

## Description (reproducing the description here solely for the sake of convenience.)
### Exercise: Infrastructure as Code
#### Objective
- [x] Create terraform file to deploy a T2.Micro EC2 instance, including VPC Networking and Security Groups
- [x] Security Group must allow access via port 22 for SSH
- [x] EC2 instance must be accessible via the Internet
- [x] All necessary configurations must be captured in Terraform so the instance can be spun down when not in use
- [x] Recommend using the Amazon Linux 2 AMI

#### Bonus Objective
Configure your fresh EC2 Instance with a complete installation of Minikube using Ansible.
- [x] Require no user input beyond launching the Ansible Playbook
- [x] Ansible Playbook should include all the steps necessary to install Minikube.

### Exercise: Bears in the Forest (Golang)

You are an avid hiker but are terrified of bears. Because of this you often find yourself wandering through the deep forest on high-alert. Given your devops background, you decide to automate the process of choosing a route through the forest that avoids bear encounters, so you decide to write an algorithm that will calculate this for you.

#### Objective

Given the input data your goal is to find a path from the starting position to the ending position (the forest uses zero-based indexing) that avoids a bear encounter.

#### Input
You will need to implement a simple HTTP server which accepts a `POST` request. The POST body should be JSON with the following format:

```
{
	"forestSize": 4,
	"forest": "....\n.B..\n..B.\n....",
	"start": [0, 0],
	"end": [3, 3]
}
```

Where _forestSize_ contains an integer n denoting the size of the forest grid, _forest_ contains the grid itself, _start_ indicates the starting position and _end_ indicates the ending position. Any given spot on the grid will be marked with either a `.` representating a safe tile or `B` representing a bear.

#### Result

- [x] Your HTTP server should respond to the `POST` request directly with your result.

- [x] Your result must be a print out of the path's steps in a human readable way or "no path" if no path is found. If there are multiple valid paths choose only one.

For example, given the example input above, a valid output would be "(0,1), (0,2), (0,3), (1,3), (2,3), (3,3)" following the left then the bottom edge of the grid.

#### Requirements

- [x] You must parse the input data.
- [x] You cannot move diagonally and you must move one tile at a time.
- [x] Your code must be split into multiple packages. e.g. a main package and a shortestpath package.
- [x] You should test thoroughly using your own input data.

#### Extra Credit

- [x] You decide to write your algorithm to find the shortest possible path instead of any path from the starting tile to the ending tile that avoids all bears, and you've heard of the legendary explorer Dijkstra who might have some much-needed insight.


### Exercise: Star Wars Cron Job (python)

- [x] Your company has tasked you with writing an background cron job. Input data has been provided as `input.yaml`. Write a script in Python that will call out to the [Star Wars API](https://swapi.co/) and gather information on each object in `input.yaml`. The output should be in JSON format so that it can be ingested by another service, to be created in the future. Write the output to a volume which survives the exit of the cronjob.

Some criteria for the project:
- [x] a `.json` file should be created for each object in `input.yaml`.
- [x] the naming convention should match `<object.name>.json` for each object.

### Exercise: Containerization
#### Preperation
You should have docker, helm and minikube installed for this exercises.

#### Objective 1
Containerize the golang or python challenge in preparation to deploying it to Minikube
- [x] Dockerfile must be produced which will be used to build this container.
- [x] Container must run the executable from the Bears in the Forest challenge, or execute the python script in the Star Wars API challenge
- [x] For Bears in the Forest: Container must expose the HTTP service created in the challenge
- [x] For Star Wars API: Container must write the output json files to a volume which survives the exit of the container

#### Objective 2 (Helm)
- [x] Create a helm chart and use it to deploy the container (from part 1) to kubernetes (locally or to the instance made in the IaC exercise above)
- [x] Must include all necessary manifests to support the challenge container.
- [x] For Bears in the Forest: Must expose the challenge container HTTP service
- [x] For Star Wars API: Must me implemented as a cronjob which is set to run every hour
