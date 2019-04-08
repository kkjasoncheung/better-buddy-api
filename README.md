# Better Buddy API and Back End

This repo contains the API and Back End built with [Golang](https://golang.org/), [Gin](https://gin-gonic.com/docs/introduction/), and [GORM](http://gorm.io/) for the Better Buddy Application.

This serves as the back-end and API for an application who's purpose is to influence positive lifestyle choices for the user.

### Getting Started
To get started with the API, clone the repo to your workspace. In the Terminal or Command Prompt:
```
$ git clone https://github.com/kkjasoncheung/better-buddy-api.git
```
Run the app
```
$ go run main.go
```

Navigate to `localhost:8080`

Available endpoints so far (on `create_user_controller` branch):
- GET user
- GET user/:id
- POST user

... Others are currently being worked on.

_*Reminder to create Swagger to document endpoints.*_
