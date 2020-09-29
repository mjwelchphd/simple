# starter

A starter web server with authentication, sessions, postgresql, static and dynamic pages, testing, email, logging, and PayPal.

This is a new project which is intended to be a starter project for a beginner.

I will be adding parts to it as I figure out how to implement each one. As I add each part, I'll document it here.

I'm a newbie to Go, but not a new programmer. I started my career in programming in 1966, so I have 55 years of experience. The sad truth is, however, that anything you did more than a few years ago is outdated and useless information in the field of programming. So, I'm learning Go!

I do have a use for it, though. I'm a professional administrator with a Ph.D., and I'm the COO of a 13 physician medical group. We use Ruby apps now, and I'd like to replace Ruby with Go.

This code is a "Proof of Concept" app, and should be useful to newbies in the Go world when it is fully implemented.

## Current Status

I figured out how to put middleware between the router and the methods the router will call. At this point, I can test the IP of the requester against a database of known hackers.

A hack is any misuse of the server. For example, if the request is for some 'wp' (word press) data, data that would be returned if this were a WP website, I log the IP and count it. After 3 tries, I block the IP. Hackers are quickly discouraged when they get a 'forbidden' screen back! (This code isn't implemented yet in this starter, but has been very successful in a Ruby website I run.)

I also figured out how to pass global data, like the database, to web pages by making them methods of 'Application.'

I figured out how to use templates, but I don't like the Go template functions because they are too limited in their functionality. I prefer something like the 'Demeler' gem on github.com/mjwelchphd/demeler.

I also figured out how to do static pages, css, and images.

My next task is to make testing for the code I have now.