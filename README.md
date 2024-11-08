# Go-fiber post API

### Introduction

This is a RESTful API built using Golang, Fiber, MongoDB Driver, MongoDB, and Reflex. This API provides a simple CRUD implementation to get, create, update, and delete posts.

</br>

### Setup

Clone the repository to your local machine.

```bash
git clone https://github.com/dipo0x/golang-post-api
```

Ensure that you have Golang and MongoDB installed on your machine. Alternatively, you can use MongoDB Cloud Atlas

Navigate to the root directory of the project in a terminal.

```bash
cd golang-post-api
```

Run the following command to install the necessary dependencies

```bash
go install
```

After that, run this command to create a .env file with which youcan get started with.

```bash
bash setup.sh
```

</br>

### Running Server

#### Locally

Run the following command to start the server:

```bash
reflex -c .reflex
```

<img width="385" alt="Screenshot 2024-11-08 at 4 50 46 PM" src="https://github.com/user-attachments/assets/a47934cc-e802-4e57-8eb7-1b4986f64556">

The server will run on http://localhost:8080 by default

</br>

## Available Endpoints

Base URL[dev]: 0.0.0.0:8080/\

When your server is running, call the base endpoint to ensure it is up, and you will receive a response like this:

<img width="1054" alt="Screenshot 2024-11-08 at 4 25 34 PM" src="https://github.com/user-attachments/assets/432a5af6-aa3a-466a-8adb-043a4d543564">


### Conclusion

You can find additional documentation for this API, including request and response signatures, by visiting https://documenter.getpostman.com/view/17975360/2sAY52beNV in your web browser.
