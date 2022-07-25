import pulumi_aws as aws

# Dynamically fetch AZs so we can spread across them.
availability_zones = aws.get_availability_zones()

# Set up a Virtual Private Cloud to deploy our EC2 instance and RDS datbase into.
prod_vpc = aws.ec2.Vpc(
    'ztpl-deployment-name-prod-vpc',
    cidr_block='10.192.0.0/16',
    enable_dns_support=True, # gives you an internal domain name
    enable_dns_hostnames=True, # gives yoiu an internal host name
    enable_classiclink=False,
    instance_tenancy='default'
)

# Create public subnets for the EC2 instance.
prod_subnet_public1 = aws.ec2.Subnet(
    'ztpl-deployment-name-prod-subnet-public-1',
    vpc_id=prod_vpc.id,
    cidr_block='10.192.0.0/24',
    map_public_ip_on_launch=True,
    availability_zone=availability_zones.names[0]
)

# Create private subnets for RDS:
prod_subnet_private1 = aws.ec2.Subnet(
    'ztpl-deployment-name-prod-subnet-private-1',
    vpc_id=prod_vpc.id,
    cidr_block='10.192.20.0/24',
    map_public_ip_on_launch=False,
    availability_zone=availability_zones.names[1]
)
prod_subnet_private2 = aws.ec2.Subnet(
    'ztpl-deployment-name-prod-subnet-private-2',
    vpc_id=prod_vpc.id,
    cidr_block='10.192.21.0/24',
    map_public_ip_on_launch=False,
    availability_zone=availability_zones.names[2]
)

# Create a gateway for internet connectivity:
prod_igw = aws.ec2.InternetGateway('ztpl-deployment-name-prod-igw', vpc_id=prod_vpc.id)

# Create a route table:
prod_public_rt = aws.ec2.RouteTable('ztpl-deployment-name-prod-public-rt',
    vpc_id=prod_vpc.id,
    routes=[
        aws.ec2.RouteTableRouteArgs(
            # associated subnets can reach anywhere:
            cidr_block='0.0.0.0/0',
            # use this IGW to reach the internet:
            gateway_id=prod_igw.id,
        )
    ]
)
prod_rta_public_subnet1 = aws.ec2.RouteTableAssociation(
    'ztpl-deployment-name-prod-rta-public-subnet-1',
    subnet_id=prod_subnet_public1.id,
    route_table_id=prod_public_rt.id
)

# Security group for EC2:
ec2_allow_rule = aws.ec2.SecurityGroup(
    'ztpl-deployment-name-ec2-allow-rule',
    vpc_id=prod_vpc.id,
    ingress=[
        aws.ec2.SecurityGroupIngressArgs(
            description='HTTPS',
            from_port=443,
            to_port=443,
            protocol='tcp',
            cidr_blocks=['0.0.0.0/0'],
        ),
        aws.ec2.SecurityGroupIngressArgs(
            description='HTTP',
            from_port=80,
            to_port=80,
            protocol='tcp',
            cidr_blocks=['0.0.0.0/0'],
        ),
        aws.ec2.SecurityGroupIngressArgs(
            description='SSH',
            from_port=22,
            to_port=22,
            protocol='tcp',
            cidr_blocks=['0.0.0.0/0'],
        ),
    ],
    egress=[aws.ec2.SecurityGroupEgressArgs(
        from_port=0,
        to_port=0,
        protocol='-1',
        cidr_blocks=['0.0.0.0/0'],
    )],
    tags={
        'Name': 'allow ssh,http,https',
    }
)

# Security group for RDS:
rds_allow_rule = aws.ec2.SecurityGroup(
    'ztpl-deployment-name-rds-allow-rule',
    vpc_id=prod_vpc.id,
    ingress=[aws.ec2.SecurityGroupIngressArgs(
        description='MySQL',
        from_port=3306,
        to_port=3306,
        protocol='tcp',
        security_groups=[ec2_allow_rule.id],
    )],
    # allow all outbound traffic.
    egress=[aws.ec2.SecurityGroupEgressArgs(
        from_port=0,
        to_port=0,
        protocol='-1',
        cidr_blocks=['0.0.0.0/0'],
    )],
    tags={
        'Name': 'allow ec2',
    }
)

# Place the RDS instance into private subnets:
rds_subnet_grp = aws.rds.SubnetGroup(
    'ztpl-deployment-name-rds-subnet-grp',
    subnet_ids=[
        prod_subnet_private1.id,
        prod_subnet_private2.id,
    ]
)