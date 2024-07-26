# GitHub CLI

## Overview

This project is a simple but very useful tool for automating repository creation and deletion processes on GitHub. It was built in Golang and is constantly improving. It is already very useful in my daily life.

## Requirements

- Go (Golang) installed on your machine.
- A GitHub access token with appropriate permissions.

## Feature

With this CLI you can automate the repository creation process, increasing the project development, as it clones the boilerplates, creates the repository, sets the main branch, and makes the first commit.
      
## How to use this project

### 1. Install Go

If you don't already have Go installed, follow the installation instructions on the [official Go website](https://golang.org/doc/install).

2 Install and compile the Project

```bash
git clone https://github.com/hebertsanto/cli.git
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

```shell
export GITHUB_TOKEN=your_token_here
````

5 Check if the Environment Variable is Set:

```shell
echo $GITHUB_TOKEN
```
6 Change variables

Change the following variables with your github data

![Captura de tela de 2024-07-26 09-00-23](https://github.com/user-attachments/assets/b996cbf7-62ce-4937-92a7-32e5eff6a38c)


7 If everything is ok you can start using it directly in your terminal

Create a repo

```shell
 cli create-repo [repo-name] --type node
```

Delete a repo

```shell
 cli delete-repo [username] [repo-name]
```
