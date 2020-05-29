# goydlidar
this is a lidar driver in go

I wouldn't use go-get for this.  I mean you could do that but I wouldn't do it. Just use git clone.  

I figured since I worked on this and someone else might be interested in using it I will make it available.  

Warning, getting this to work is not intuitive.  I was going to try to get it to work with the new Github repository found at https://github.com/YDLIDAR/YDLidar-SDK . It would probably require a rewrite of some of the functions  It would also require a bunch of workaround that I don't want to deal with since how the #include paths are different.  I forked the other sdk to https://github.com/dereklstinson/ydlidar, and used that for these bindings.  The problem with the last repository is that it has several branches.  Since, I have a X4 model I made bindings using the X4 branch.

A lot of the API is written in Chinese.  I did a google translate on it, but it didn't help much. 

Instructions.  
Clone this library https://github.com/dereklstinson/goydlidar
Then clone https://github.com/dereklstinson/ydlidar 

Since, I used the x4 branch here are the instructions on building that (I changed the place where it is cloned):
If you use a different branch checkout a different one.  

How to build YDLIDAR SDK samples
---------------
    $ git clone https://github.com/dereklstinson/ydlidar #updated url
    $ cd sdk
    $ git checkout X4
    $ cd ..
    $ mkdir build
    $ cd build
    $ cmake ../sdk
    $ make			###linux

These instructions build the library outside of the cloned folder.

Then several files need to be copied.

In the build folder copy libydlidar_driver.a
  ```
  cp libydlidar_driver.a GOYDLIDARFOLDERLOCATION/libydlidar_driver.a
  ```
  
Then in the sdk folder 

  ```
   cp include/angles.h GOYDLIDARFOLDERLOCATION/angles.h
   cp include/CYdLidar.h GOYDLIDARFOLDERLOCATION/CYdLidar.h
   cp include/lock.h GOYDLIDARFOLDERLOCATION/lock.h
   cp include/locker.h GOYDLIDARFOLDERLOCATION/locker.h
   cp include/serial.h GOYDLIDARFOLDERLOCATION/serial.h
   cp include/thread.h GOYDLIDARFOLDERLOCATION/thread.h
   cp include/timer.h GOYDLIDARFOLDERLOCATION/timer.h
   cp include/utils.h GOYDLIDARFOLDERLOCATION/utils.h
   cp include/v8stdint.h GOYDLIDARFOLDERLOCATION/v8stdint.h
   cp include/ydlidar_driver.h GOYDLIDARFOLDERLOCATION/ydlidar_driver.h
   cp include/ydlidar_protocol.h GOYDLIDARFOLDERLOCATION/ydlidar_protocol.h
   
  ```

Also you need to give non root access to the usb port that the ydlidar is connected.  If it is connected to usb0 then use
```
sudo chmod 666 /dev/ttyUSB0
```

This doesn't come without faults.  First, this package was created for the X4.  It may or may not work with any other device.  You will have to do your own hacking.  Second, type YDdriver methods don't all work.  They some of the methods return errors, and since I could get what I wanted to get done using type Lidar I didn't put much more effort into it. 
