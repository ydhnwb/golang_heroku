# golang heroku deploy
This is a simple restful api project that deployed to heroku. You can access it with this url:  
```
https://golang-heroku.herokuapp.com
```

Find the tutorial here:
```
https://www.youtube.com/watch?v=_EAkLIoMCNM
```

# Documentation
## health 
to check current server is alive:

<b>GET</b>
```
https://golang-heroku.herokuapp.com/api/check/health
```
  
Response (Status: 200)
```
{
   message: "OK!"
}
```

## register
Registering a new user

<b>POST</b>
```
https://golang-heroku.herokuapp.com/api/auth/register
```

Request Body
```
{
    "name": "Prieyudha Akadita S",
    "email": "ydhnwb@gmail.com",
    "password": "yudhanewbie"
}
```

Response success (Status: 201)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 2,
        "name": "Prieyudha Akadita S",
        "email": "ydhnwb@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsImV4cCI6MTY1MTgyMDAwMCwiaWF0IjoxNjIwMjg0MDAwLCJpc3MiOiJhZG1pbiJ9.HtnuWlBaevEO3fHAI4McH5W8axvw_3Og47RUI3m9IyI"
    }
}
```

Response error (Status : 400) [Depends on what error]
```
{
    "status": false,
    "message": "Failed to process request",
    "errors": [
        "Key: 'RegisterRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
    ],
    "data": {}
}
```

## login
Authenticate by email and password

<b>POST</b>
```
https://golang-heroku.herokuapp.com/api/auth/login
```

Request body
```
{
    "email": "yudhanewbie@gmail.com",
    "password": "yudhanewbie"
}
```

Response Success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Prieyudha Akadita S",
        "email": "yudhanewbie@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSIsImV4cCI6MTY1MTgyMTQ0NiwiaWF0IjoxNjIwMjg1NDQ2LCJpc3MiOiJhZG1pbiJ9.2m-r1qrCXhNkAxzK-sb4hL0Pzat3zwOmzktES_uzwts"
    }
}
```

Response error, wrong credential (Status: 401)
```
{
    "status": false,
    "message": "Failed to login",
    "errors": [
        "failed to login. check your credential"
    ],
    "data": {}
}
```


## Get Profile
Get current info from logged user

<b>GET</b>
```
https://golang-heroku.herokuapp.com/api/user/profile
```

Headers
```
Authorization: yourToken
```

Response success (status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Prieyudha Akadita S",
        "email": "yudhanewbie@gmail.com"
    }
}
```

## Update profile
Update user data who logged in

<b>PUT</b>
```
https://golang-heroku.herokuapp.com/api/user/profile
```

Headers
```
Authorization: yourToken
```

Request Body
```
{
    "name": "Prieyudha Akadita S",
    "email": "ahmadalbar@gmail.com"
}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "name": "Prieyudha Akadita S",
        "email": "ahmadalbar@gmail.com"
    }
}
```


<b>=============================================</b>
## All products (based on user who logged in)
Only shows products by user who logged in

<b>GET</b>
```
https://golang-heroku.herokuapp.com/api/product
```


Headers
```
Authorization: yourToken
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": [
        {
            "id": 2,
            "product_name": "Xiaomi Redmi 10",
            "price": 3000,
            "user": {
                "id": 1,
                "name": "Prieyudha Akadita S",
                "email": "ahmadalbar@gmail.com"
            }
        },
        {
            "id": 3,
            "product_name": "Indomie Goreng",
            "price": 2500,
            "user": {
                "id": 1,
                "name": "Prieyudha Akadita S",
                "email": "ahmadalbar@gmail.com"
            }
        }
    ]
}
```

## Create product
Create a product with owner is the user who logged in

<b>POST</b>
```
https://golang-heroku.herokuapp.com/api/product
```

Headers
```
Authorization: yourToken
```

Request body
```
{
    "name": "Xiaomi Redmi 5",
    "price": 3000
}
```

Response success (Status: 201)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "product_name": "Xiaomi Redmi 5",
        "price": 3000,
        "user": {
            "id": 1,
            "name": "Prieyudha Akadita S",
            "email": "ahmadalbar@gmail.com"
        }
    }
}
```

## Find one product by id
Find product by id

<b>GET</b>
```
https://golang-heroku.herokuapp.com/api/product/{id}
```

Headers
```
Authorization: yourToken
```


Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "product_name": "Xiaomi Redmi 5",
        "price": 3000,
        "user": {
            "id": 1,
            "name": "Prieyudha Akadita S",
            "email": "ahmadalbar@gmail.com"
        }
    }
}
```

## Update product
<b>You can only update your own product If you are trying to update someone else product, it will return error. </b>  

<b>PUT</b>
```
https://golang-heroku.herokuapp.com/api/product/{id}
```

Request body
```
{
    "name": "Xiaomi Redmi 5 Plus",
    "price": 5000
}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {
        "id": 1,
        "product_name": "Xiaomi Redmi 5 Plus",
        "price": 5000,
        "user": {
            "id": 1,
            "name": "Prieyudha Akadita S",
            "email": "ahmadalbar@gmail.com"
        }
    }
}
```


## Delete product
You can only delete your own product

<b>DELETE</b>
```
https://golang-heroku.herokuapp.com/api/product/{id}
```

Response success (Status: 200)
```
{
    "status": true,
    "message": "OK!",
    "errors": null,
    "data": {}
}
```
