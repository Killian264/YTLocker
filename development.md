# Requirements

Docker Desktop -- https://www.docker.com/products/docker-desktop  
WSL 2 -- https://docs.microsoft.com/en-us/windows/wsl/install-win10

-   Ubuntu -- Microsoft Store
-   Windows Terminal -- Microsoft Store
-   Golang -- https://golang.org/doc/install
-   NodeJS -- https://github.com/nodesource/distributions/blob/master/README.md

# Development Setup (Windows)

## WSL 2

Install and Open WSL 2 Ubuntu

-   Create a user remember the `USER_NAME`

## Windows Terminal

-   Install and Open Windows Terminal
-   Dropdown > Settings > Open JSON File
-   Paste below inside "profiles" > "list" replacing `USER_NAME` with the Ubuntu username chosen

```
{
    "guid": "{2c4de342-38b7-51cf-b940-2309a097f518}",
    "hidden": false,
    "name": "Ubuntu",
    "source": "Windows.Terminal.Wsl",
    "startingDirectory": "//wsl$/Ubuntu/home/USER_NAME"
}
```

-   Ensure that the dropdown in Windows Terminal now shows Ubuntu and it opens properly.

## Docker Desktop

-   Install Docker Destop

## Inside Ubuntu

-   Install Golang
-   Install NodeJS
-   Ensure these commands all work `docker -v` `node -v` `go version`
-   Follow cloning and running steps in `README.md`

# VSCode Extensions

-   golang.go -- golang tooling, also install required packages
-   cweijan.vscode-mysql-client2 -- create connection with secrets from `docker-compose.yml`
-   ms-vscode-remote.remote-wsl -- used to connect to WSL 2

# Other Commands

-   `go version` -- check version, use after install
-   `go get` -- gets packages from go.mod
-   `node -v` -- check version, use after install
-   `yarn install` -- installs packages for node from package.json
-   `yarn add` -- adds a dependency
