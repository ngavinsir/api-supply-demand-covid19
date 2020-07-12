# Create Request

* Endpoint: `/api/v1/requests`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body:
    ```JSON
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
    ```JSON
    {
        "id": "1Zw60g0nM99QJIJjN37SpIKlYF9",
        "date": "2020-04-01T17:39:47.940516264+07:00",
        "isFulfilled": false,
        "requestItems": [
            {
                "id": "1Zw60dckGOR6v7nIA0UofRWrWMl",
                "item": {
                    "id":"1",
                    "name":"Masker"
                },
                "unit": {
                    "id":"1",
                    "name":"Buah"
                },
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
* Request Body: `-`
* Response Body:
    ```JSON
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
                        "item": {
                            "id":"2",
                            "name":"Hand Sanitizer"
                        },
                        "unit": {
                            "id":"2",
                            "name":"Liter"
                        },
                        "quantity": "22.13"
                    },
                    {
                        "id": "1Zw6eKbga0OzptlUPa8UgcXIGfN",
                        "item": {
                            "id":"1",
                            "name":"Masker"
                        },
                        "unit": {
                            "id":"1",
                            "name":"Buah"
                        },
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

# Update Request

* Endpoint: `/api/v1/requests/{requestID}`
* HTTP Method: `PUT`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body:
    ```JSON
    {
        "requestItems": [
            {
                "id": "1aMhkuANqCoTPHj7EkA5NFrkkna",
                "unit_id": "1",
                "item_id": "1",
                "quantity": 3
            }
        ]
    }
    ```
* Response Body:
    ```JSON
    {
        "id": "1aMhktyXsy1gTLZOAtc2ObhOwnI",
        "date": "2020-04-10T20:45:16.952605Z",
        "isFulfilled": false,
        "requestItems": [
            {
                "id": "1aMhkuANqCoTPHj7EkA5NFrkkna",
                "item": {
                    "id":"1",
                    "name":"Masker"
                },
                "unit": {
                    "id":"1",
                    "name":"Buah"
                },
                "quantity": "3",
                "request_id": "1aMhktyXsy1gTLZOAtc2ObhOwnI"
            }
        ]
    }
    ```

# Get Request Detail

* Endpoint: `/api/v1/requests/{requestID}`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body: `-`
* Response Body:
    ```JSON
    {
        "data": {
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
                    "item": {
                        "id":"2",
                        "name":"Hand Sanitizer"
                    },
                    "unit": {
                        "id":"2",
                        "name":"Liter"
                    },
                    "quantity": "22.13"
                },
                {
                    "id": "1Zw6eKbga0OzptlUPa8UgcXIGfN",
                    "item": {
                        "id":"1",
                        "name":"Masker"
                    },
                    "unit": {
                        "id":"1",
                        "name":"Buah"
                    },
                    "quantity": "11.00"
                }
            ]
        }
    }
    ```

# Get User Requests

* Endpoint: `/api/v1/requests/user/{userID}`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body: `-`
* Response Body:
    ```JSON
    {
        "data": [
            {
                "id": "1aMhktyXsy1gTLZOAtc2ObhOwnI",
                "date": "2020-04-10T20:45:16.952605Z",
                "isFulfilled": false,
                "donationApplicant": {
                    "id": "1a7vQqjAeg3rCBryTNfaWpZnczh",
                    "email": "admin@admin.com",
                    "name": "admin@admin.com",
                    "contact_person": null,
                    "contact_number": null,
                    "role": "ADMIN"
                },
                "requestItems": [
                    {
                        "id": "1aMhkuANqCoTPHj7EkA5NFrkkna",
                        "item": {
                            "id":"1",
                            "name":"Masker"
                        },
                        "unit": {
                            "id":"1",
                            "name":"Buah"
                        },
                        "quantity": "2.50"
                    }
                ]
            }
        ],
        "pages": {
            "current": 1,
            "total": 1,
            "first": true,
            "last": true
        }
    }
    ```

# Delete Request

* Endpoint: `/api/v1/requests/{requestID}`
* HTTP Method: `DELETE`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body: `-`
* Response Body: `-`

# Create Request Item Allocation

* Endpoint: `/api/v1/requests/items/{requestItemID}/allocation`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body:
    ```JSON
    {
        "description": "sudah",
        "date": "2020-04-14T08:55:30Z"
    }
    ```
* Response Body:
    ```JSON
    {
        "id": "1eaKh2FOCKy5c5qiZiLWM6HnSC5",
        "request_item_id": "1dL1l6fxEwp4rdplE3eM9JQZdpT",
        "allocation_date": "2020-07-12T18:46:18.888407556+07:00",
        "description": "sudah"
    }
    ```

# Edit Request Item Allocation

* Endpoint: `/api/v1/requests/items/{requestItemID}/allocation/{requestItemAllocationID}`
* HTTP Method: `PUT`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body:
    ```JSON
    {
        "description": "sudahselesai2" 
    }
    ```
* Response Body:
    ```JSON
    {
        "id": "1eaKh2FOCKy5c5qiZiLWM6HnSC5",
        "request_item_id": "1dL1l6fxEwp4rdplE3eM9JQZdpT",
        "allocation_date": "2020-07-12T11:55:41.50752Z",
        "description": "sudahselesai2"
    }
    ```