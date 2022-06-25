FROM python:3.8-slim

# define environments
ENV PYTHONUNBUFFERED 1

# Create app directory
WORKDIR /app

# Install app dependencies
COPY requirements.txt ./
RUN pip install -r requirements.txt

# Bundle app source
COPY . .

EXPOSE 3000
RUN chmod 755 ./start.sh
CMD ./start.sh
