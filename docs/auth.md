# Login

* Endpoint: `/api/v1/auth/login`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body:
    ```
    {
        "email": "test",
        "password": "test"
    }
    ```
* Response Body:
    ```
    {
        "email": "test",
        "name": "test",
        "role": "donator",
        "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJleHAiOjE1ODUyNTg0NzAsInVzZXJfaWQiOiIxWmc0N2RENlVydk8zSkRCY2ZJN0d1Qlo
        2TXoifQ.V2oAHjqRt6ekSf1ULamPD87mqZp5ZsWOZqDrCchiBsg"
    }
    ```


# Register

* Endpoint: `/api/v1/auth/register`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body:
    ```
    {
        "email": "test",
        "password": "test",
        "name": "test",
        "role": "donator"
    }
    ```
* Response Body:
    ```
    {
        "email": "test",
        "name": "test",
        "role": "donator",
        "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJleHAiOjE1ODUyNTgwMzYsInVzZXJfaWQiOiIxWmc0N2RENlVydk8zSkRCY2ZJN0d1Qlo
        2TXoifQ.59-EqcUkQSxzIlND6cCfaI0OVTDd6rcMZnEqttFkLqk"
    }
    ```

# Refresh

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