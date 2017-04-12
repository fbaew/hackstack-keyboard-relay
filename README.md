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

## Installing Client

### Windows
* Put kbclient.exe somewhere on your filesystem 
* Create a shortcut to kbclient.exe
* Configure the shortcut path to be `"X:\path\to\kbclient.exe" -detach -key=X:\path\to\private.key`
* Configure a hotkey combo for the shortcut (I like `Ctrl+Alt+Scroll Lock` as it also makes
  the scroll lock inicator somewhat useful)

### Linux
* Put kbclient somewhere on your filesystem
* Configure a hotkey combo to run `/path/to/kbclient -attach -key=/path/to/private.key`
  (How you do this depends on your choice of DE but I assume if you've made it as far as
  needing something like this you've got that shit under control)
    * I suggest using the same hotkey combo for Windows and Linux both

### config.json

Specify the hostname/IP and port number where `kbserver` is listening. `kbclient` will attempt to load `config.json`
from the current working directory, or it will try to load the file specified by `-config=/path/to/config.json`

## Configuring QEMU
You must add a new monitor to your qemu invocation:

`qemu ... -chardev socket,id=mon2,*host=localhost,port=4445*,server,nowait -mon chardev=mon2,mode=readline`

Note: unless you want any joker to be able to connect to and control your guest, don't bind this
to a different interface. In fact, it should really be a local socket, but I don't support that yet.

## Installing Linux Server

Nothing much to this yet; just run `./kbserver`. This will attempt to load `config.json` in the current
working directory. If it isn't there, the server won't start. If you want to specify a different config file, do so
with `./kbserver -config=/path/to/config.json`.

### config.json

* *KeyboardName* - A string that uniquely distinguishes the keyboard to be manipulated from all other lines
   of output from `info usb` on the qemu monitor.

* *VendorID* - A 4-digit hex identifier corresponding to your hardware vendor. Find this in the output of `lsusb`.
* *ProductID* - A 4-digit hex identifier representing your specific product model. Find this in the output of `lsusb`.
* *ManagementPort* - The TCP on which to listen for connections from `kbclient`. Always binds to `localhost`.

#### Sample:
```
{
    "KeyboardName":"Corsair K65 Gaming Keyboard",
    "VendorID":"1b1c",
    "ProductID":"1b07"
}
```

# TODO

* Auto-generate config file
* (?) Develop Windows wrapper to listen for key events?
