FROM ubuntu:jammy

LABEL maintainer="Vanessasaurus <@vsoch>"

# Match the default user id for a single system so we aren't root
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000
ENV USERNAME=${USERNAME}
ENV USER_UID=${USER_UID}
ENV USER_GID=${USER_GID}
USER root

# Install extra buildrequires for flux-sched:
RUN apt-get update && \
    apt-get -qq install -y --no-install-recommends \
        wget \
        curl \ 
        less \
        git \
        build-essential \
        protobuf-compiler \
        ca-certificates

# Assuming installing to /usr/local
ENV LD_LIBRARY_PATH=/usr/lib:/usr/local/lib
ENV PATH=$PATH:/workspaces/fluence-kubectl/bin

# Install Go 19 for TBA bindings (if Go bindings desired)
RUN wget --no-check-certificate https://go.dev/dl/go1.19.10.linux-amd64.tar.gz && tar -xvf go1.19.10.linux-amd64.tar.gz && \
         mv go /usr/local && rm go1.19.10.linux-amd64.tar.gz
      
ENV PATH=$PATH:/usr/local/go/bin:/home/vscode/go/bin

# Add the group and user that match our ids
RUN groupadd -g ${USER_GID} ${USERNAME} && \
    adduser --disabled-password --uid ${USER_UID} --gid ${USER_GID} --gecos "" ${USERNAME} && \
    echo "${USERNAME} ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers

# TODO clone flux-sched here / just remove from build when we can?
