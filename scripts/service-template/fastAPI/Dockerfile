
FROM python:latest

# define environments
ENV PYTHONUNBUFFERED 1

# Create app directory
WORKDIR /app

# Install app dependencies
COPY Pipfile Pipfile.lock ./
RUN pip install pipenv && pipenv install

# Bundle app source
COPY . .

EXPOSE 3000
CMD ./start.sh
