# Create Request

* Endpoint: `/api/v1/requests`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body:
    ```
    {
        "requestItems": [
            {
                "item_id": "2",
                "unit_id": "3",
                "quantity": 11
            },
            {
                "item_id": "3",
                "unit_id": "1Zim6hXOmUTiH28xubzTz2kA0ed",
                "quantity": 22.13
            }
        ]
    }
    ```
* Response Body:
    ```
    {
        "id": "1Zw60g0nM99QJIJjN37SpIKlYF9",
        "date": "2020-04-01T17:39:47.940516264+07:00",
        "isFulfilled": false,
        "requestItems": [
            {
                "id": "1Zw60dckGOR6v7nIA0UofRWrWMl",
                "item": "Masker",
                "unit": "Buah",
                "quantity": "11"
            },
            {
                "id": "1Zw60kCOybHjxtQdXRTk6xf2cAw",
                "item": "Hand Sanitizer",
                "unit": "Liter",
                "quantity": "22.13"
            }
        ]
    }
    ```

# Get All Request

* Endpoint: `/api/v1/requests?page=10&size=1`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body: `-`
* Response Body:
    ```
    {
        "data": [
            {
                "id": "1Zw6eIqGb6CXoxw3WL2JofbBjcA",
                "date": "2020-04-01T10:45:02.489632Z",
                "isFulfilled": false,
                "donationApplicant": {
                    "id": "1ZiESUvzPPP0r9SCTtkt84eAGfP",
                    "email": "admin",
                    "name": "admin",
                    "contact_person": null,
                    "contact_number": null,
                    "role": "ADMIN"
                },
                "requestItems": [
                    {
                        "id": "1Zw6eJgNwKNuLu3rH3WqbeYoq6n",
                        "item": "Hand Sanitizer",
                        "unit": "Liter",
                        "quantity": "22.13"
                    },
                    {
                        "id": "1Zw6eKbga0OzptlUPa8UgcXIGfN",
                        "item": "Masker",
                        "unit": "Buah",
                        "quantity": "11.00"
                    }
                ]
            }
        ],
        "pages": {
            "current": 10,
            "total": 11,
            "first": false,
            "last": false
        }
    }
    ```