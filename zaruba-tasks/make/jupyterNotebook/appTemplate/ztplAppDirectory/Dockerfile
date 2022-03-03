FROM docker.io/bitnami/jupyter-base-notebook:latest

USER 0
RUN conda install --quiet --yes \
    'matplotlib-base' \
    'scipy' && \
    conda clean --all -f -y
USER 1001