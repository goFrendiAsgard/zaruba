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

```
.
├── Dockerfile
├── README.md
├── __pycache__
│   └── main.cpython-38.pyc
├── auth
│   ├── __init__.py
│   ├── __pycache__
│   ├── authModel.py
│   ├── event.py
│   ├── route.py
│   ├── rpc.py
│   ├── tokenModel.py
│   ├── userModel.py
│   ├── userRoute.py
│   ├── userRpc.py
│   └── userSeederModel.py
├── blog
│   ├── __init__.py
│   ├── __pycache__
│   ├── event.py
│   ├── route.py
│   └── rpc.py
├── database.db
├── helpers
│   ├── __init__.py
│   ├── __pycache__
│   ├── app
│   └── transport
├── library
│   ├── __init__.py
│   ├── __pycache__
│   ├── bookModel.py
│   ├── bookRoute.py
│   ├── bookRpc.py
│   ├── event.py
│   ├── route.py
│   └── rpc.py
├── main.py
├── repos
│   ├── __init__.py
│   ├── __pycache__
│   ├── book.py
│   ├── dbBook.py
│   ├── dbUser.py
│   └── user.py
├── requirements.txt
├── schemas
│   ├── __init__.py
│   ├── __pycache__
│   ├── book.py
│   └── user.py
├── start.sh
└── template.env
```

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