# Development container

Development container that can be used with VSCode.

It works on Linux, Windows and OSX.

## Requirements

- [VS code](https://code.visualstudio.com/download) installed
- [VS code remote containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) installed
- [Docker](https://www.docker.com/products/docker-desktop) installed and running
  - If you don't use Linux or WSL 2, share your home directory `~/` and the directory of your project with Docker Desktop
- [Docker Compose](https://docs.docker.com/compose/install/) installed
- Ensure your host has the following and that they are accessible by Docker:
  - `~/.ssh` directory
  - `~/.gitconfig` file (can be empty)

## Build the container image

You first need to build the container image and have it locally cached. This prevents VSCode from rebuilding the container from scratch every time you launch the workspace.

To build the container run the following command from the .devcontainer folder:

``` bash
docker build -f ./dev.Dockerfile --tag  local/tfctl-dev .
```

## Setup

1. Open the command palette in Visual Studio Code (CTRL+SHIFT+P).
1. Select `Remote-Containers: Open Folder in Container...` and choose the project directory.
