FROM python:latest

# define environments
ENV SERVICE_NAME servicename
ENV PYTHONUNBUFFERED 1

# Create app directory
WORKDIR /app

# Install app dependencies
COPY Pipfile Pipfile.lock ./
RUN pip install pipenv && pipenv sync

# Bundle app source
COPY . .

EXPOSE 3000
CMD pipenv run python main.py
