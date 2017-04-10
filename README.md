# Introduction

Have you ever found yourself thinking "gee I wish I could detach my keyboard
from a qemu guest and use it on the host system, all at the press of a button!"

Sound familiar? Well buckle up, bucko! We're plunging straight into golang and racing towards that exact goal.

# TODO

* Authenticate clients of the keyboard server so that Joe Dirt can't hotplug my keyboard
* Read QEMU control socket details, keyboard name + model from config file
* Auto-generate config file

# Objectives

* Develop a server to run on the qemu host
    * Accept authenticated connections
    * Allow clients to issue commands which will be brokered into the qemu monitor
* Develop a Linux client that issues "attach keyboard to guest" command
    * Comb the output of `lsusb` to find the right vendorID and productID
    * Call the above server with `usb_add host:1234:beef`
    * Develop a configuration helper utility so that the keyboard model is not hard-coded
* Develop a Windows client that issues "detach keyboard from guest" command
    * comb the output of `info usb` in the qemu monitor to find the right device ID
    * Call the server and issue `del_usb <deviceid>`
    * Develop a configuration helper utility so that the keyboard model is not hard-coded

# ...Why?

Like most computery stuff that I do, I am mostly doing this out of curiosity. I
wanted to give Go a try, and I have a use case (VFIO-powered Windows gaming
virtual machine). If someone else out there uses this too, terrific!
