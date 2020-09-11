<<<<<<< HEAD
## golang-http-server
An example of a HTTP Server written in Golang to test features.

---
### What you are NOT going to find here (yet)
  
* Unit testing
* Lint strategy
* CI/CD
* Any kind of automation
* Go mods or any dependency management

### What you are going to find:

* App deployed on heroku (free plan);
* Simple endpoint just for testing the deploy on heroku (so I can make sure it is running);
* A full http server to test a few things in Go

When I build this app my goal was to write a HTTP server in Go to test a few basic features and
also, deploy it on Heroku w/o any dependency management. The reason of that is because 
Heroku usually detects my apps as Go apps based on `/vendor/vendor.json` file, which on this
case I don't have. So, for this app, it looks for one or more go files inside **subfolders** of 
`src` directory.

It is extremely important to have a subfolder inside `/src`. It is with the name of this subfolder 
that executable file is created inside the `/bin` folder. This gets more clear if you take a look
in the project folder structure AND the `Procfile` file.

---
## Accessing the app

1. Open the url (localhost:8080)
2. Type your name and submit the form;
3. You will see the Form content parsed and rendered in the page;

**IMPORTANT:** The first access to any rout is going to be very slow because the app is probably sleeping on Heroku.
