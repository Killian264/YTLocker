# Requirements #
Golang: https://golang.org/doc/install  
WSL:   
NodeJS: https://github.com/nodesource/distributions/blob/master/README.md
* curl -fsSL https://deb.nodesource.com/setup_15.x | sudo -E bash -
* sudo apt-get install -y nodejs


# Steps #
*Pull down code*  
Git Clone: https://github.com/Killian264/YTLocker

*Make WSL not be dumb about permissions*  
sudo chown -R "user-name" "directory-name"

*Build Services*  
docker-compose up

# Other Commands #
-------
* go version -- check version, use after install
* go get -- gets packages  from go.mod
* go run main.go -- builds and runs main
------
* node -v -- check version, use after install
* npm install -- installs packages for node from package.json
* npm update -- updates packages probably
* npm update "package-name" -- updates a single package
* npm start -- starts a dev server 
------

# Extensions #
* golang.go -- also install required packages 
* ms-azuretools.vscode-docker
* coenraads.bracket-pair-colorizer