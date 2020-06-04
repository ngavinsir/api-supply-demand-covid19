# Login

* Endpoint: `/api/v1/auth/login`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body:
    ```JSON
    {
        "email": "test",
        "password": "test"
    }
    ```
* Response Body:
    ```JSON
    {
        "user": {
            "id": "1ZiESUvzPPP0r9SCTtkt84eAGfP",
            "email": "admin",
            "name": "admin",
            "contact_person": null,
            "contact_number": null,
            "role": "ADMIN"
        },
        "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU3NDYzNjMsInVzZXJfaWQiOiIxWmlFU1V2elBQUDByOVNDVHRrdDg0ZUFHZlAifQ.ffyS1yOhQxuIgbd2l09-Q4tBhVES7BOTOXFt88GkbUc"
    }
    ```


# Register

* Endpoint: `/api/v1/auth/register`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body:
    ```JSON
    {
        "email": "test",
        "password": "test",
        "name": "test",
        "role": "DONATOR",
        "contact_person": "Donator",
        "contact_number": "0846272634"
    }
    ```
* Response Body:
    ```JSON
    {
        "email": "test",
        "name": "test",
        "role": "DONATOR",
        "contact_person": "Donator",
        "contact_number": "0846272634",
        "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJleHAiOjE1ODUyNTgwMzYsInVzZXJfaWQiOiIxWmc0N2RENlVydk8zSkRCY2ZJN0d1Qlo
        2TXoifQ.59-EqcUkQSxzIlND6cCfaI0OVTDd6rcMZnEqttFkLqk"
    }
    ```

# Refresh Token

* Endpoint: `/api/v1/auth/refresh`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body: `-`
* Response Body:
    ```
    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJleHAiOjE1ODUyNTgwMzYsInVzZXJfaWQiOiIxWmc0N2RENlVydk8zSkRCY2ZJN0d1Qlo
    2TXoifQ.59-EqcUkQSxzIlND6cCfaI0OVTDd6rcMZnEqttFkLqk
    ```

# Create Password Reset Request

* Endpoint: `/api/v1/auth/reset`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body:
    ```JSON
    {
        "email": "example@email.com"
    }
* Response Body: `-`

# Confirm Password Reset Request

* Endpoint: `/api/v1/auth/reset/{requestID}/confirm`
* HTTP Method: `PUT`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body:
    ```JSON
    {
        "newPassword": "NEW_PASSWORD"
    }
    ```
* Response Body: `-`
