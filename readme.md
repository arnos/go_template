# Yet another Go(lang) template for microservices or cli (and using just)

## Why Just

> Why not using Make?  
> 
> Erich N.

I can point to the technical reasons and design some one else [posted](https://github.com/casey/just), or go on with simple vs easy.

But in the end I found just, it simply had **the right balance** for what I intended.

And now we can go on to the 

##  The Goal

The Goal is to make a simple template, inspired by clean Architecture and screaming Architecture.

1. You shouldn't need to open file to get what they do.
1. Files should have a standard format.   
    > as such I propose the following standard:  
    > 
    > [domain]_[purpose].[language extension]
    >
    > In the golang context I define purposes as:
    > * _main (where everything starts and ends)
    > * _logic (how everything works)
    > * _interfaces (how to abstract implementation from the logic) 
    > * _implementation (where everything happens)
    > * _test (because it inspired the template).
1. Having the above purposes allows one to quickly find the proper file for the piece of code 
1. Then comes the domain, use a domain that draws the attention of what your code is doing
    > I.E. the file domain should tell what the code in it is doing take api and you will have the following
    > * api_logic.go
    > * api_interfaces.go
    > * api_implementation.go
    > * api_test.go
1. The domains should be micro-domains and only be used when needed to seperate code into smaller more managable junks.

With the above file name format I submit that we are using screaming architecture

Clean Architecture is about seperating your core entities (interfaces) and logic from the messy implementations and I submit that using the purpose in the file names provides the seperation.

## The Journey

This must be my 4th revision of a template I started a few years ago based on clean architecture, while the template structure withstood the names of the file (and methods) became generic creating more code debt than functionality.

As such I wanted to evolve it:
* Remove the need for the need for folder as structure (using _purpose)
* Bringing the practice of descriptive function names to file names (as domains) 

I'm fairly certain that this template will withstand time again, and hopefully it will communicate the intent of the code for our fellow maintainers.

## Notes

I did explore going one degree further by abstracting the main intake using [watermill](https://watermill.io/) so that I could abstract the different lambda main call (we are working between GCP and Azure) and/or be able to re-design a given binary by compiling sub-components together. 

But as with just felt it had the right balance, I felt that changing this didn't belong here in the template but rather as an implementation example. 
