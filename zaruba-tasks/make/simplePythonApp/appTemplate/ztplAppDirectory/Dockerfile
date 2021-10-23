FROM python:3.8-slim as requirements-maker

# define environments
ENV PYTHONUNBUFFERED 1

# Create app directory
WORKDIR /app

# Install app dependencies
RUN pip install pipenv
COPY Pipfile Pipfile.lock ./
RUN pipenv lock -r > requirements.txt


FROM python:3.8-slim

# define environments
ENV PYTHONUNBUFFERED 1

# Create app directory
WORKDIR /app

# Install app dependencies
COPY --from=requirements-maker /app/requirements.txt ./
RUN pip install -r requirements.txt

# Bundle app source
COPY . .

CMD python main.py
