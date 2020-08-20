FROM python:latest

ENV PYTHONUNBUFFERED 1

WORKDIR /app

COPY . .

RUN pip install pipenv && pipenv sync

CMD pipenv run python main.py
