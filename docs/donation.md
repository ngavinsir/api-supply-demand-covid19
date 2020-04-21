# Create Donation

* Endpoint: `/api/v1/donations`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer jwt`
* Request Body:
    ```
    {
        "donationItems": [
            {
                "item_id": "1",
                "unit_id": "1",
                "quantity: 10
            },
            {
                "item_id": "2",
                "unit_id": "2",
                "quantity": 10
            },
        ]
    }
    ```
* Response Body:
    ```
    {
        "donation": {
            "id": "1ZlfXrepQpHoq6e4YYfhpdAZ4HK",
            "date": "2020-03-29T01:04:03.458704516+07:00",
            "is_accepted": false,
            "is_donated": false,
            "donator_id": "1ZiESUvzPPP0r9SCTtkt84eAGfP"
        },
        "items": [
            {
                "id": "1ZlfXrepQpHoq6e4YYfhpdAZ4HK",
                "donation_id": "1ZlfXsUn4fYgvWc7Ke6SZSNuDPn",
                "item_id": "1",
                "unit_id": "1",
                "quantity": "10"
            },
        ]
    }
    ```

# Get All Donation

* Endpoint: `/api/v1/donations?page=1&size=1`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body: `-`
* Response Body:
    ```
    {
        "data": [
            {
                "id": "1aMeoehZkXbki1xJo47X6FwBtcU",
                "date": "2020-04-10T20:21:05.053957Z",
                "isAccepted": false,
                "isDonated": false,
                "donator": {
                    "id": "1a871e37D2acUEUl9ssOZTKsjsy",
                    "email": "ngavinsir@gmail.com",
                    "name": "ngavinsir",
                    "contact_person": null,
                    "contact_number": null,
                    "role": "DONATOR"
                },
                "donationItems": [
                    {
                        "id": "1aMeoatjtEKeCdisz1sbCTnsiQe",
                        "item": {
                            "id":"1",
                            "name":"Masker"
                        },
                        "unit": {
                            "id":"1",
                            "name":"Buah"
                        },
                        "quantity": "10.00"
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

# Update Donation

* Endpoint: `/api/v1/donations/{donationID}`
* HTTP Method: `PUT`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer jwt`
* Request Body:
    ```
    {
        "donationItems": [
            {
                "id": "1aMeoatjtEKeCdisz1sbCTnsiQe",
                "unit_id": "1",
                "item_id": "1",
                "quantity": 10.1
            }
        ]
    }
    ```
* Response Body:
    ```
    {
        "id": "1aMeoehZkXbki1xJo47X6FwBtcU",
        "date": "2020-04-10T20:21:05.053957Z",
        "isAccepted": false,
        "isDonated": false,
        "donationItems": [
            {
                "id": "1aMeoatjtEKeCdisz1sbCTnsiQe",
                "item": "Masker",
                "unit": "Buah",
                "quantity": "10.1",
                "donation_id": "1aMeoehZkXbki1xJo47X6FwBtcU"
            }
        ]
    }
    ```

# Accept Donation

* Endpoint: `/api/v1/donations/{donationID}/accept`
* HTTP Method: `PUT`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer jwt`
* Request Body: `-`
* Response Body: `-`

# Get Donation Detail

* Endpoint: `/api/v1/donations/{donationID}`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body: `-`
* Response Body:
    ```
    {
        "id": "1aMeoehZkXbki1xJo47X6FwBtcU",
        "date": "2020-04-10T20:21:05.053957Z",
        "isAccepted": false,
        "isDonated": false,
        "donator": {
            "id": "1a871e37D2acUEUl9ssOZTKsjsy",
            "email": "ngavinsir@gmail.com",
            "name": "ngavinsir",
            "contact_person": null,
            "contact_number": null,
            "role": "DONATOR"
        },
        "donationItems": [
            {
                "id": "1aMeoatjtEKeCdisz1sbCTnsiQe",
                "item": {
                    "id":"1",
                    "name":"Masker"
                },
                "unit": {
                    "id":"1",
                    "name":"Buah"
                },
                "quantity": "10.00"
            }
        ]
    }
    ```

# Get User Donations

* Endpoint: `/api/v1/donations/user/{userID}`
* HTTP Method: `GET`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
* Request Body: `-`
* Response Body:
    ```
    {
        "data": [
            {
                "id": "1aMeoehZkXbki1xJo47X6FwBtcU",
                "date": "2020-04-10T20:21:05.053957Z",
                "isAccepted": false,
                "isDonated": false,
                "donator": {
                    "id": "1a871e37D2acUEUl9ssOZTKsjsy",
                    "email": "ngavinsir@gmail.com",
                    "name": "ngavinsir",
                    "contact_person": null,
                    "contact_number": null,
                    "role": "DONATOR"
                },
                "donationItems": [
                    {
                        "id": "1aMeoatjtEKeCdisz1sbCTnsiQe",
                        "item": {
                            "id":"1",
                            "name":"Masker"
                        },
                        "unit": {
                            "id":"1",
                            "name":"Buah"
                        },
                        "quantity": "10.00"
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