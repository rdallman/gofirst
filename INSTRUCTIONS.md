### Windows

The first thing we will need is Go. Navigate to
<http://code.google.com/p/go/downloads/list>. Since you are running Windows, you
will want to download and install the go1.1.2 Windows (x86) MSI installer.
Either 32bit or 64bit depending on your system and version of windows. (If you
are unsure just get 32bit or feel free to ask)

Next you will want to download Git. Navigate to
<http://git-scm.com/download/win>
and download and install Git. The installation will prompt you with several
options, just use the defaults.

Run command prompt as an administrator. To do this hit the windows key, type
cmd, right click command prompt and select run as administrator.

First we need to create a new directory to store our Go projects in. Type
the following into the command prompt. 

```
mkdir C:\GoCode
```
This has created a directory in your C: drive called GoCode. 

Next, we will need to set the `GOPATH`. For a description, see the below section
called "GOPATH"

```
setx /m GOPATH C:\GoCode\ 
``` 

The command prompt should inform you that this was successful.

Next, we need to set the path for git. To do this you will need to type
the following into the command prompt : 

```
setx /m PATH “%PATH%;C:\Program Files (x86)\Git\bin”
```

(the quotes are important, don't forget them). It
should also inform you that it was successful.

Close command prompt and reopen it (either admin or not, it doesn't
matter). To verify the modifications we made were successful first type

```
git
```

It should list a bunch of git commands. Next type 

```
echo %GOPATH%
```

It should return `C:\GoCode\`. If neither of these return what they are
supposed to let one of us know!

### Mac / Linux

Navigate to the google code repository to grab a copy of Go.
Point your browser to <http://code.google.com/p/go/downloads/list>

select `go1.1.2.darwin-amd64` PKG installer and download it

once complete, open the installer and complete the instructions.

Now open up a terminal and enter:

```
$ git --version
```

If git is not installed, for __linux:__

```
sudo apt-get install git
```

__mac:__

go to <http://git-scm.com/download/mac>

now I'd like for you to enter the following commands to set up your $GOPATH
environment variable:

```
$ mkdir ~/gocode
$ echo 'export GOPATH=~/gocode' >> ~/.profile
$ echo 'export PATH=$GOPATH:$PATH' >> ~/.profile
$ source ~/.profile
$ echo $GOPATH
```

these should all work and the last one should print the location of your $GOPATH

__If not__ call somebody over

