# quizapi
TO INSTALL GO
If you have a previous version of Go installed, be sure to remove it before installing another.

Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go.
For example, run the following as root or through sudo:

tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
Add /usr/local/go/bin to the PATH environment variable.
You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

export PATH=$PATH:/usr/local/go/bin
Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

Verify that you've installed Go by opening a command prompt and typing the following command:
$ go version
Confirm that the command prints the installed version of Go.

TO INSTALL BEEGO
Beego contains sample applications to help you learn and use the Beego app framework.

You will need a Go 1.1+ installation for this to work.

You will need to install or upgrade Beego and the Bee dev tool:

go get -u github.com/astaxie/beego
go get -u github.com/beego/bee
For convenience, you should add $GOPATH/bin to your $PATH environment variable. Please make sure you have already set the $GOPATH environment variable.

If you haven’t set $GOPATH add it to the shell you’re using (~/.profile, ~/.zshrc, ~/.cshrc or any other).

For example ~/.zsh
echo 'export GOPATH="$HOME/go"' >> ~/.zsh

If you have already set $GOPATH
echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.profile # or ~/.zshrc, ~/.cshrc, whatever shell you use exec $SHELL

Want to quickly see how it works? Then just set things up like this:
cd $GOPATH/src bee new hello cd hello bee run

Windows users：
cd %GOPATH%/src bee new hello cd hello bee run

CLONE THE CODE THEN COMMAND
bee run -downdoc=true -gendoc=true

see the swagger page at localhost:8080/swagger
