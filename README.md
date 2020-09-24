<div align="center">

# Questioner

![GitHub repo size](https://img.shields.io/github/repo-size/sarthakpranesh/Questioner)
[![GitHub issues](https://img.shields.io/github/issues/sarthakpranesh/Questioner)](https://github.com/sarthakpranesh/Questioner/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/sarthakpranesh/Questioner)](https://github.com/sarthakpranesh/Questioner/pulls)
[![GitHub Repo stars](https://img.shields.io/github/stars/sarthakpranesh/Questioner)](https://github.com/sarthakpranesh/Questioner/stargazers)
![GitHub](https://img.shields.io/github/license/sarthakpranesh/Questioner)

</div>

<br />

## Introduction
Many online games and competitions like various CTFs or events like [Enigma](https://github.com/IEEE-VIT/enigma6) have something in common and that is a simple Question and Answer model that the main backend service is made to handle. Questioner is this common Question and Answer backend service that can be easily used to host such an online event. This would help the team focus more on structuring questions, investing more time in frontend and other aspects of the event.

<br />

## Technologies Used
- Go Lang
- Mongo Atlas Database
- Gorilla Mux
- godotenv
- jwt-go

<br />

## For Developers
Make sure you have go installed

For local development
1. `git clone https://github.com/sarthakpranesh/Questioner`
2. `cd Questioner`
3. create a `.env` file and add the following content
    ```
    MONGO_URL = "<your mongo cluster connection URL>"
    ADMIN_PASSWORD = "<your admin password>"
    PORT = "<your port number>"
    ```
3. `go mod tidy`
4. `go run main.go` - you'll have to restart the server each time you make a change to see its affect

For hosting their is a docker file included in the project that can be easily used to build and host a docker image of the project.

<br />

## Documentation
The End Points are tested and documented using Postman and the collection can be accessed from link below

[![Run in Postman](https://run.pstmn.io/button.svg)](https://documenter.getpostman.com/view/7649159/TVKFzFn6)

<br />

## Found Something Broken
If you find any bug, vulnerability, or have any feature suggestion please feel free to open an Issue [here](https://github.com/sarthakpranesh/Questioner/issues)


<br />

<div align="center">

##### Made with ❤️

</div>