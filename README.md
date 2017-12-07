# GoLang notes

**Go is an open source programming language that make it easy to build simple, reliable, and efficient software.**



## Install


### MACOS

Install Mercurial from https://www.mercurial-scm.org/downloads

```bash
# or bash whatever
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

source /Users/james/.gvm/scripts/gvm

# A Note on Compiling Go 1.5+
gvm install go1.8 -B
gvm use go1.8
export GOROOT_BOOTSTRAP=$GOROOT

# show release version
gvm listall

# install 
gvm install <version>

# list
gvm list

```

