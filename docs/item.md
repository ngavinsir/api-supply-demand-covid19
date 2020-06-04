# Get All Item

* Endpoint: `/api/v1/items`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body: `-`
* Response Body:
    ```JSON
    [
        {
            "id": "1",
            "name": "Beras"
        },
        {
            "id": "1Zio8iFvaASndYT8wlcSEW5m91e",
            "name": "Masker"
        }
    ]
    ```

# Create Item

* Endpoint: `/api/v1/items`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body:
    ```JSON
    {
        "name": "Masker"
    }
    ```
* Response Body:
    ```JSON
    {
        "id": "1Zio8iFvaASndYT8wlcSEW5m91e",
        "name": "Masker"
    }
    ```

# Delete Item

* Endpoint: `/api/v1/items/{itemID}`
* HTTP Method: `DELETE`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body: `-`
* Response Body: `-`