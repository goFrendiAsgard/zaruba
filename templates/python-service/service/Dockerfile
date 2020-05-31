FROM continuumio/miniconda3

ENV PYTHONUNBUFFERED 1

WORKDIR /home/anaconda/

COPY ./environment.yml ./environment.yml
RUN /opt/conda/bin/conda env update -n base -f environment.yml

COPY . /home/anaconda/

CMD /opt/conda/bin/conda init && source /home/anaconda/.bashrc && python main.py
