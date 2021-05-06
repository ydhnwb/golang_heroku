# golang heroku deploy
This is a simple restful api project that deployed to heroku. You can access it with this url:  
```
https://golang-heroku.herokuapp.com
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
