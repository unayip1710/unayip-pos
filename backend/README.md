# POS go 

In this repo you will find the POS backend 

### Prerequisites

- [Go > 1.21](https://go.dev/doc/install)
- [make](https://linuxhint.com/install-use-make-windows/)

## MakeFile

There is a make file there you can add commands and add their shortcut, so be free to add your commands 

## Running Code

Be sure to add your database connection string in app.env file, then

just run ``` make start ``` 

your server will start

## Run the follow endpoint 

```json
/ping
```

Happy Hacking :sunglasses:

## Running the coverage tests

Add your test files, once you can run ```make test``` and then you can see your coverage

## Running Linter

Once you finish your commit you can run ```make lint``` it will run a linter and you can be sure about your good code :wink:

## .github folder

There are github actions workflows there you have a linter, builder, and test runner 80% minimum that runs in github


