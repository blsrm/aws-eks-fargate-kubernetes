# aws-rust-kubernetes

When it comes to building our EKS cluster we have a few different options.


1 - Using the AWS CLI / eksctl
2 - AWS Console
3 - Using the AWS Cloud Development Kit
4 - Using Cloudformation
5 - Using Terraform

In this tutorial I will be exploring options 1, 2 and 3. I will also provided links for repo examples to guide you if you want to try out 3 and 4.

## 1. AWS CLI / eksctl

For this option you need to have the following tools installed

- [eksctl](https://github.com/weaveworks/eksctl): eksctl is a simple CLI tool for creating clusters on EKS
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/): CLI to interact with the kubernetes API server
- [AWS CLI 2](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-mac.html) + Docker: We will use Docker and the AWS CLI to build and push a Docker image for our application.

Lets use the eksctl tool to create a new EKS cluster with Fargate enabled, we'll call it eks-fargate-demo. The `--fargate` flag creates a Fargate profile, which is used to run Kubernetes pods as Fargate tasks, instead of using EC2 instances as worker nodes. This fargate profile has access to the default and kube-system namespaces.

EKS on Fargate only supports the Application Load Balancer (ALB) so the `--alb-ingress-access` flag helps setup some of the scaffolding required to setup an ALB to work with EKS.

```bash
eksctl create cluster \
--name eks-fargate-demo \
--version 1.16 \
--region us-east-1 \
--fargate \
--alb-ingress-access
```

Creating a new cluster can take almost 15-20 minutes to finish so be patient once you run this. If everything worked ok you should see this result in your terminal

```bash
[✔] EKS cluster "eks-fargate-demo" in "us-east-1" region is ready
```

check that the nodes were created:

```bash
kubectl get nodes

# NAME                                      STATUS   ROLES    AGE    VERSION
# fargate-ip-192-168-121-116.ec2.internal   Ready    <none>   2m4s   v1.16.8-eks-e16311
# fargate-ip-192-168-67-11.ec2.internal     Ready    <none>   2m     v1.16.8-eks-e16311
```

The `eksctl create cluster` command that we ran above also created a vpc called `eksctl-eks-fargate-demo-cluster/VPC` with 2 public subnets and 2 private subnets. We need to get the VPC ID, which we can grab by running the following `eksctl get cluster` command:

```bash
eksctl get cluster --region us-east-1 --name eks-fargate-demo -o yaml
```

This will print out some yaml in the console

```yaml
- Arn: arn:aws:eks:us-east-1:190103567417:cluster/eks-fargate-demo
  CertificateAuthority:
    Data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5RENDQWJDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJd01EY3hOekV5TlRZeU1Gb1hEVE13TURjeE5URXlOVFl5TUZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBS1hQClBhc3ZEVFJFTFVRemRVdUNCSzhsL0ZtaStwS0JTcUhLSVUwTW1ERW56Smd4VUlHWGRtc21xNW1aL0NkYS9qN24KREYvdEVMUEpUR05Fd2pLNWJJT1JkUFFxeElvSGowUmJleFJJTDVVc0NnMnB2YWFadUMySnltL1NlSHFyamptaQpPakVkK2dmdXNCdHgxZnd5d1ZIbGRyWmxLK2xrQjh6M1IxQk1vRG5Od3hCYm5kZkxPeWw0eGtMeUU5aDlTQUdWCjBxZWxDeXNQaHJiTXh0RVY0MFRvL0RmL3VnajFYSlFOUkFuNTljNmdLdWQ0bTBtcnpoREZ5clhqZ1pRSERkTG0KaDFhZENuckk5akdNMWJoWHFFZC82bFNvS0FvSGVqTkk2ZitsRER6VmlxQnhsa3FjRHUxaDRjaTl1WGJWb2pGVQo5MXNuWit6UzlSNlpPVUlmTjhjQ0F3RUFBYU1qTUNFd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFGdmRaWHprZnI0aHZmN3JFQ01KUEtnS1hTMkQKTS8xNjU1UmVsSDY0ZWx0dnBaUVhJbExBd0NteVZvNXpFMlZOcXpYb21NeGFpSzk2djZWSlk0Q2tYbXVwVnJMdwpwTTlJZml6c05BSDFnNVFLaHZUQklLemluZTl1T0J3anRvNUtKUzlsbG1JM0RvYlhXVk1xWTJUSXhZdDRGb1I4CmxXUURiYVEwdG9SMi9HM05EZWtEVVNtM2RqOFRDdERTUWFiVnpkQml5Y3k1cCtHTXZDMU8wMUJqWU5UeVdYSkcKcmp1ZFdhYUQ1bktucENVUkxnR0p3MFRzTXhKWS9BeHQ0YzRod2c2R0FBdlJONXdCclZjSGNmYk5PeHFQZXRkWgoySDhlZGxZK2dzeE1iTnJvY1VPMEk3ekhBWndJejI0aGU0Zk1FL09LQS9mU0VMTTJlMUFhQ2ovZGszOD0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
  ClientRequestToken: null
  CreatedAt: "2020-07-17T12:48:47Z"
  EncryptionConfig: null
  Endpoint: https://766221B7380D4D259F92FB03F9593A74.gr7.us-east-1.eks.amazonaws.com
  Identity:
    Oidc:
      Issuer: https://oidc.eks.us-east-1.amazonaws.com/id/766221B7380D4D259F92FB03F9593A74
  Logging:
    ClusterLogging:
    - Enabled: false
      Types:
      - api
      - audit
      - authenticator
      - controllerManager
      - scheduler
  Name: eks-fargate-demo
  PlatformVersion: eks.2
  ResourcesVpcConfig:
    ClusterSecurityGroupId: sg-084ec8ebe3843535f
    EndpointPrivateAccess: false
    EndpointPublicAccess: true
    PublicAccessCidrs:
    - 0.0.0.0/0
    SecurityGroupIds:
    - sg-044b64c08ca5f3847
    SubnetIds:
    - subnet-09cd6c5db56035cc3
    - subnet-0f71efc72dae05650
    - subnet-097e26d09a1858658
    - subnet-090a2f3ec5e6bd6e4
    VpcId: vpc-0807555a7163ca290
  RoleArn: arn:aws:iam::190103567417:role/eksctl-eks-fargate-demo-cluster-ServiceRole-55SSGKTODNQA
  Status: ACTIVE
  Tags: {}
  Version: "1.16"
```

It returns a lot of useful information about our cluster, in my case the VPC ID is `vpc-0807555a7163ca290`.

Next, we need to allow external traffic to access our application, to do this we have to setup the AWS [ALB Ingress controller](https://github.com/kubernetes-sigs/aws-alb-ingress-controller). There are 2 files within the `./ingress-controller`, the first is `alb-ingress-controller.yaml` creates an Ingress Controller which uses the AWS Application Load Balancer (ALB), the second file `rbac-role.yaml` gives the right permissions to the ALB ingress controller to communicate with the EKS cluster we created earlier.

Before we can apply these manifests, we need to uncomment and edit the following fields in the `alb-ingress-controller.yaml`:

`cluster-name`: The name that you gave to your cluster in the eksctl command above.
`vpc-id`: The VPC Id that you grabbed in the command above.
`aws-region`: The region for your EKS cluster.
`AWS_ACCESS_KEY_ID`: The AWS access key id that ALB controller can use to communicate with AWS. For this tutorial, we will add the access key in plaintext in the file. However, for a production setup, it is recommended to use a project like kube2iam for providing IAM Access.
`AWS_SECRET_ACCESS_KEY`: The AWS secret access key id that ALB controller can use to communicate with AWS. For this tutorial, we will add the access key in plaintext in the file. However, for a production setup, it is recommended to use a project like kube2iam for providing IAM Access.

Once you have finished editing then we can deploy the 2 files using kubectl

```bash
kubectl apply -f rbac-role.yaml
# clusterrole.rbac.authorization.k8s.io/alb-ingress-controller created
# clusterrolebinding.rbac.authorization.k8s.io/alb-ingress-controller created
# serviceaccount/alb-ingress-controller created
```

```bash
kubectl apply -f alb-ingress-controller.yaml
# deployment.apps/alb-ingress-controller created
```

Verify the deployment was successful and the controller started.
```bash
kubectl logs -n kube-system $(kubectl get po -n kube-system | egrep -o "alb-ingress[a-zA-Z0-9-]+")
# -------------------------------------------------------------------------------
# AWS ALB Ingress controller
#   Release:    v1.1.6
#   Build:      git-95ee2ac8
#   Repository: https://github.com/kubernetes-sigs/aws-alb-ingress-controller.git
# -------------------------------------------------------------------------------
```

Next, we will create a new repository on AWS Elastic Container Registry (ECR) to store our Docker images.

```bash
aws ecr create-repository --repository-name kubernetes-tutorial

# {
#     "repository": {
#         "repositoryArn": "arn:aws:ecr:us-east-1:190103567417:repository/kubernetes-tutorial",
#         "registryId": "190103567417",
#         "repositoryName": "kubernetes-tutorial",
#         "repositoryUri": "190103567417.dkr.ecr.us-east-1.amazonaws.com/kubernetes-tutorial",
#         "createdAt": "2020-07-17T14:17:01+01:00",
#         "imageTagMutability": "MUTABLE",
#         "imageScanningConfiguration": {
#             "scanOnPush": false
#         }
#     }
# }

```

Make a note of the `repositoryUri` from json reponse you get back as you'll need to build our docker image. To be able to deploy our application onto Kubernetes, we need to create a Docker image for our application. Use the commands in the makefile to build and tag your docker image.

```bash
make build
```
```bash
make tag
```

Next, we need to push the Docker image to ECR so that it can be accessed by the EKS cluster. Before we can push the image to ECR, we need to authenticate the Docker CLI:

```bash
aws ecr get-login-password --region <region> | docker login --username AWS --password-stdin <aws_account_id>.dkr.ecr.<region>.amazonaws.com
# Login Succeeded
```

Make sure the account id does not have the dashes so in mu case this was `aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 190103567417.dkr.ecr.us-east-1.amazonaws.com`

Once you have logged in successfully, you should now be able to push your Docker image using:

```bash
docker push 190103567417.dkr.ecr.us-east-1.amazonaws.com/kubernetes-tutorial:1

# The push refers to repository [190103567417.dkr.ecr.us-east-1.amazonaws.com/kubernetes-tutorial]
# d2118dec3e87: Pushed
# 82a15a1dbd15: Pushed
# 35d1ab51a96b: Pushed
# 1ba1431fe2ba: Pushed
# 0f7493e3a35b: Pushed
# 50644c29ef5a: Pushed
# 1: digest: sha256:043e6338b6e46ab98fa62566bbb12adeac4f234e4a71e0c7dc3e71f31dabffc1 size: 1576
```

There are 4 manifest files under the `./infra/k8s` folder. These files are:

`Namespace`: Creates a new Namespace for our application
`Deployment`: Creates a Deployment object for our application
`Service`: Creates a Service object for our application
`Ingress`: Sets up ingress for the application so that it is accessible externally

First, we need to create a new Fargate profile since we will be using a new namespace for our application. EKS on Fargate by default only supports pods running in the default and kube-system namespaces.

```bash
eksctl create fargateprofile --namespace go-web --cluster eks-fargate-demo --region us-east-1

# [ℹ]  creating Fargate profile "fp-ed2af90b" on EKS cluster "eks-fargate-demo"
# [ℹ]  created Fargate profile "fp-ed2af90b" on EKS cluster "eks-fargate-demo"
```

Apply the namespace manifest to create a new namespace for the app

```bash
cd infra/k8s && kubectl apply -f namespace.yaml

# namespace/go-web created
```

Then, we create a service for our application:

```bash
kubectl apply -f service.yaml
```

Check that the service has been created

```bash
kubectl get svc -n go-web

# NAME     TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
# go-web   NodePort   10.100.196.3   <none>        80:32285/TCP   36s
```

Next, we will create the deployment to deploy the pod to fargate:

```bash
kubectl apply -f deployment.yaml
```

Check that the pod is up and running (it will take a minute or two for the pod to be ready):

```bash
kubectl get pods -n go-web

# NAME                      READY   STATUS    RESTARTS   AGE
# go-web-6bd5f865c5-krp2k   1/1     Running   0          107s
```

Lastly, we will apply ingress manifest to expose our application to the public internet.

```bash
kubectl apply -f ingress.yaml

```

To check the status of the ingress, we can run the following command. Note the Address in the output is public URL for our application.

```bash
kubectl describe ing -n go-web go-web

# Name:             go-web
# Namespace:        go-web
# Address:          5fb65de9-goweb-goweb-a520-731462364.us-east-1.elb.amazonaws.com
# Default backend:  default-http-backend:80 (<none>)

```

The ALB can take a while to start working. Once it is functional, we should be able to access it by hitting the endpoint:

```bash
curl -v 5fb65de9-goweb-goweb-a520-731462364.us-east-1.elb.amazonaws.com/api/users/ping

# pong
```


## 2. AWS Console
Documentation to follow

## 3. Using the AWS Cloud Development Kit (CDK)
Documentation to follow

## 4. and 5. Using Cloudformation and Terraform
Links to follow
