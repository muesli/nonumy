# nonumy

A minimalistic GUI based file hasher written in Go

## About

nonumy contains a self explaining GUI

![nonumy's GUI] (/assets/nonumy.png?raw=true)

nonumy supports the comparison of two hashsums

You can also choose by command line which file you want to hash.
Simply type:

    nonumy /path/to/file
    
And the file is directly choosen on startup

## Current supported methods:

* **MD5**
* **SHA1**
* **SHA256**
* **SHA512**

## Installation

Make sure you have a working Go environment. See the [install instructions](http://golang.org/doc/install.html).

First of all you need to checkout the source code:
    
    git clone https://github.com/PenguWin/nonumy.git
    cd nonumy
    
Now you need to get the required dependencies:

    go get -v

Let's build nonumy:

    go build

Thats all!

To make nonumy global executeable just copy it to /usr/local/bin 

    cp nonumy /usr/local/bin/nonumy
    
## Development

I'm currently working on some improvements for nonumy.
It's not a complete, finished and tested programm yet!
