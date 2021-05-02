
FROM python:3.8-slim

# define environments
ENV PYTHONUNBUFFERED 1

# Create app directory
WORKDIR /app

# Install app dependencies
RUN pip install pipenv
COPY Pipfile Pipfile.lock ./
RUN pipenv install --system --deploy

# Bundle app source
COPY . .

EXPOSE 3000
CMD ./start.sh
