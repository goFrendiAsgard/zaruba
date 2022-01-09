# ZtplAppDirectory

An opinionated Fast API application. Built with ❤️ based on past mistakes/experiences.

# How to run

```bash
# create virtual environment if not exist
if [ ! -d ./venv ]; then python -m venv ./venv; fi

# activate virtual environment
source venv/bin/activate

# install pip packages
pip install -r requirements.txt

# load environments
source template.env

# run the application
./start.sh
```

# Application structure

# Logical layers

* Repo
* Schema
* Model
* RPC and Messagebus
* Routes

# Dependency Injection

# Message bus

## Local

## Rabbitmq

## Kafka

## Your custom message bus

# RPC

## Local

## Rabbitmq

## Your custom RPC