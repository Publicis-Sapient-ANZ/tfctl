ARG GO_VERSION=1.18

FROM mcr.microsoft.com/vscode/devcontainers/go:${GO_VERSION}-buster

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000

USER root

#Install software dependencies
RUN apt-get update \
    && apt-get install -y --no-install-recommends sudo git make apt-transport-https gnupg2 curl lsb-release zsh ca-certificates\
    && apt-get purge -y --auto-remove \
    && rm -rf /var/lib/apt/lists/*

# Setup shell for root and ${USERNAME}
RUN usermod --shell /bin/zsh root && \
    usermod --shell /bin/zsh ${USERNAME}
COPY --chown=${USER_UID}:${USER_GID} shell/.zshrc shell/.profile.sh shell/.welcome.sh /home/${USERNAME}/
RUN mkdir -p /home/${USERNAME}/.config
COPY --chown=${USER_UID}:${USER_GID} shell/.config/starship.toml /home/${USERNAME}/.config

# Install starship
RUN sh -c "$(curl -fsSL https://starship.rs/install.sh)" -- --yes

#Install McFly
RUN curl -LSfs https://raw.githubusercontent.com/cantino/mcfly/master/ci/install.sh | sh -s -- --git cantino/mcfly

#Install Terraform
RUN curl -fsSL https://apt.releases.hashicorp.com/gpg | apt-key add -
RUN echo "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | tee -a /etc/apt/sources.list.d/hashicorp.list
RUN apt-get update && apt-get install terraform

# Python and pre-commit
RUN apt-get update \
    && apt-get install -y --no-install-recommends python3 python3-pip\
    && apt-get purge -y --auto-remove \
    && rm -rf /var/lib/apt/lists/*
RUN pip3 install pre-commit

#Install Azure CLI
RUN curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash

#Install golang stuff
RUN go install github.com/cosmtrek/air@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
RUN go install github.com/go-critic/go-critic/cmd/gocritic@latest
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin 

#Fix some file permissions
RUN chown -R $USERNAME:$USERNAME /home/$USERNAME/
# More permissions. Need to find a beteer way.
RUN chown -R $USERNAME:$USERNAME /go/

ENV HOME /home/${USERNAME}
USER ${USERNAME}

CMD [ "/bin/zsh" ]