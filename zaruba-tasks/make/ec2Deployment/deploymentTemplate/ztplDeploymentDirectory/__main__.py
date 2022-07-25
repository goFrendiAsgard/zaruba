import pulumi
import pulumi_aws as aws
import vpc
import os

config = pulumi.Config()
public_key_path = os.getenv('PUBLIC_KEY_PATH', 'app-keypair.pub')
private_key_path = os.getenv('PRIVATE_KEY_PATH', 'app-keypair')
db_instance_size = os.getenv('DB_INSTANCE_SIZE', 'db.t3.small')
db_name = os.getenv('DB_NAME', 'appDb')
db_username = os.getenv('DB_USERNAME', 'root')
db_password = os.getenv('DB_PASSWORD', 'Alch3mist')
ec2_instance_size = os.getenv('EC2_INSTANCE_SIZE', 't3.small')

# Dynamically query for the AMI in this region.
ami = aws.ec2.get_ami(
    owners=['099720109477'],
    filters=[
            aws.ec2.GetAmiFilterArgs(
            name='name',
            values=['ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*']
        )
    ],
    most_recent=True
)


# Read in the public key for easy use below.
public_key = open(public_key_path).read()
# Read in the private key for easy use below (and to ensure it's marked a secret!)
private_key = pulumi.Output.secret(open(private_key_path).read())


# Create the RDS instance:
app_rds = aws.rds.Instance(
    'ztpl-deployment-name-app-rds',
    allocated_storage=10,
    engine='mysql',
    engine_version='5.7',
    instance_class=db_instance_size,
    db_subnet_group_name=vpc.rds_subnet_grp.id,
    vpc_security_group_ids=[vpc.rds_allow_rule.id],
    db_name=db_name,
    username=db_username,
    password=db_password,
    skip_final_snapshot=True,
)

# Create a keypair to access the EC2 instance:
app_keypair = aws.ec2.KeyPair('ztpl-deployment-name-app-keypair', public_key=public_key)

# Note: user_data is executed by root user, and the default directory is /
app_ec2_user_data = '''#!/bin/bash
apt-get update -y
apt-get install golang -y
sudo su - ubuntu
sh -c "$(curl -fsSL https://raw.githubusercontent.com/state-alchemists/zaruba/master/install.sh)"
'''

# Create an EC2 instance to run app (after RDS is ready).
app_ec2 = aws.ec2.Instance(
    'ztpl-deployment-name-app-instance',
    ami=ami.id,
    instance_type=ec2_instance_size,
    subnet_id=vpc.prod_subnet_public1.id,
    vpc_security_group_ids=[vpc.ec2_allow_rule.id],
    key_name=app_keypair.id,
    tags={
        'Name': 'app.web',
    },
    # Only create after RDS is provisioned.
    opts=pulumi.ResourceOptions(depends_on=[app_rds]),
    # Define what to do once created
    user_data=app_ec2_user_data,
    user_data_replace_on_change=False
)

# Give our EC2 instance an elastic IP address.
app_eip = aws.ec2.Eip('ztpl-deployment-name-app-eip', instance=app_ec2.id)

pulumi.export('app-ec2-public-dns', app_ec2.public_dns)
pulumi.export('app-eip-public-dns', app_eip.public_dns)
pulumi.export('app-rds-address', app_rds.address)
