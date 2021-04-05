# Requirements #
WSL 2: https://docs.microsoft.com/en-us/windows/wsl/install-win10  
Get Ubuntu From microsoft Store  
Get Windows Terminal From Microsoft Store  
Docker Desktop: https://www.docker.com/products/docker-desktop  
Golang: https://golang.org/doc/install  
NodeJS: https://github.com/nodesource/distributions/blob/master/README.md  
* curl -fsSL https://deb.nodesource.com/setup_15.x | sudo -E bash -  
* sudo apt-get install -y nodejs  


# Steps #
Git Clone: https://github.com/Killian264/YTLocker

sudo chown -R "user-name" "directory-name"

docker-compose up

# Extensions #
* golang.go -- also install required packages 
* ms-azuretools.vscode-docker
* coenraads.bracket-pair-colorizer
* cweijan.vscode-mysql-client2
* ms-vscode-remote.remote-wsl

# Other Commands #
* go version -- check version, use after install
* go get -- gets packages  from go.mod
* go run main.go -- builds and runs main
* node -v -- check version, use after install
* npm install -- installs packages for node from package.json
* npm update -- updates packages probably
* npm update "package-name" -- updates a single package
* npm start -- starts a dev server 

* ~USER_NAME/go/bin/mockery --all

# Full Steps #
Get WSL 2: https://docs.microsoft.com/en-us/windows/wsl/install-win10   
Get Ubuntu From microsoft Store  
Open Ubuntu
Create a user with a USER_NAME
Get Windows Terminal From Microsoft Store  
* Click Dropdown
* Click Settings
* inside "list" paste the below code block and replace USER_NAME with chosen user name

            {
                "guid": "{2c4de342-38b7-51cf-b940-2309a097f518}",
                "hidden": false,
                "name": "Ubuntu",
                "source": "Windows.Terminal.Wsl",
                "startingDirectory": "//wsl$/Ubuntu/home/USER_NAME"
            }
Open Windows Terminal 
* If when you click the dropdown Ubuntu is shown the previos steps worked.  

-------------------------------

Download Docker Desktop for Windows: https://www.docker.com/products/docker-desktop  

INSIDE Ubuntu follow Golang Install steps for linux: https://golang.org/doc/install  
* Instead of the first step run this sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz  

INSIDE Ubuntu follow NodeJS Install steps: https://github.com/nodesource/distributions/blob/master/README.md  
* sudo curl -fsSL https://deb.nodesource.com/setup_15.x | sudo -E bash -  
* sudo apt-get install -y nodejs  

If any of these commands don't work you messed up some previous install step
* docker -v
* node -v
* go version

-------------------  

RUN git clone https://github.com/Killian264/YTLocker

sudo chown -R "USER_NAME" "YTLocker"

cd YTLocker

docker-compose up

Install WSL 2 on vs code

"code ." in YTLocker directory

Install Extensions listed in VSCode Extensions section