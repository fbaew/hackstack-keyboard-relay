# Introduction

Have you ever found yourself thinking "gee I wish I could hotplug my keyboard
between a qemu guest and my host system, all at the press of a button!"

Sound familiar? Well buckle up, bucko! We're plunging straight into golang and racing towards that exact goal.

## ...Why?

Like most computery stuff that I do, I am mostly doing this out of curiosity. I
wanted to give Go a try, and I have a use case (VFIO-powered Windows gaming
virtual machine). If someone else out there uses this too, terrific!

# INSTALLING

None of this is automatic or pretty. You have been warned.

## Generating a private key
Commands issued to kbserver are encrypted with 256-bit AES. `kbclient` checks the directory from
which it was executed for `private.key`, unless you specify `-key=path/to/private.key`. To generate this file,
run `keyutil` with no arguments. This will generate a (pseudo)random key and stick it in `private.key`.

## Installing Windows Client 
* Put kbclient.exe somewhere on your filesystem 
* Create a shortcut to kbclient.exe
* Configure the shortcut path to be `"X:\path\to kbclient.exe" -detach -key=X:\path\to\private.key`
* Configure a hotkey combo for the shortcut (I like `Ctrl+Alt+Scroll Lock` as it also makes
  the scroll lock inicator somewhat useful)

## Installing Linux Client
* Put kbclient somewhere on your filesystem
* Configure a hotkey combo to run `/path/to/kbclient -attach -key=/path/to/private.key`
  (How you do this depends on your choice of DE but I assume if you've made it as far as
  needing something like this you've got that shit under control)
    * I suggest using the same hotkey combo for Windows and Linux both

## Configuring QEMU
You must add a new monitor to your qemu invocation:

`qemu ... -chardev socket,id=mon2,*host=localhost,port=4445*,server,nowait -mon chardev=mon2,mode=readline`

Note: unless you want any joker to be able to connect to and control your guest, don't bind this
to a different interface. In fact, it should really be a local socket.

## Installing Linux Server

Nothing much to this yet; just run `./kbserver`

# TODO

* Read QEMU control socket details, keyboard name + model from config file
* Auto-generate config file
* (?) Develop Windows wrapper to listen for key events?
