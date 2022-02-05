FROM docker.io/bitnami/jupyter-base-notebook:latest

RUN conda install --quiet --yes \
    'matplotlib-base' \
    'scipy' && \
    conda clean --all -f -y