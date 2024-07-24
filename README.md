# GitHub CLI

## Overview

This project is a simple but very useful tool for automating repository creation and deletion processes on GitHub. It was built in Golang and is constantly improving. It is already very useful in my daily life.

## Requirements

- Go (Golang) installed on your machine.
- A GitHub access token with appropriate permissions.

## How to use this project

### 1. Install Go

If you don't already have Go installed, follow the installation instructions on the [official Go website](https://golang.org/doc/install).

2 Install and compile the Project

```bash
git clone https://github.com/seu_usuario/github-cli.git
cd cli
go build -o cli
```

3 Move the Binary to a Directory in the PATH:
```shell
sudo mv cli /usr/local/bin/
```

### 4. Configure GitHub Token

You need a GitHub access token to authenticate your requests. Create a token on [GitHub](https://github.com/settings/tokens) and set the permissions required to create and delete repositories.

Set the `GITHUB_TOKEN` environment variable to your token:

```bash
export GITHUB_TOKEN=your_token_here
````

5 Check if the Environment Variable is Set:

```shell
echo $GITHUB_TOKEN
```

6 If everything is ok you can start using it directly in your terminal

Create a repo

```shell
 cli create-repo [repo-name]
```

Delete a repo

```shell
 cli delete-repo [username] [repo-name]
```