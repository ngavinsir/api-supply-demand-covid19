# Get All Unit

* Endpoint: `/api/v1/units`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body: `-`
* Response Body:
    ```
    [
        {
            "id": "1",
            "name": "Kg"
        }
    ]
    ```

# Create Unit

* Endpoint: `/api/v1/units`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body:
    ```
    {
        "name": "Meter"
    }
    ```
* Response Body:
    ```
    {
        "id": "1Zim6hXOmUTiH28xubzTz2kA0ed",
        "name": "Liter"
    }
    ```