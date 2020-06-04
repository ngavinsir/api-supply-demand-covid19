# Create Allocation

* Endpoint: `/api/v1/allocations`
* HTTP Method: `POST`
* Request Header:
    * Accept: `application/json`
    * Content-type: `application/json`
    * Authorization: `Bearer token`
* Request Body:
    ```
    {
        "requestID": "1aMhktyXsy1gTLZOAtc2ObhOwnI",
        "date": "2020-04-14T08:55:30Z",
        "items": [
            {
                "item_id": "1",
                "unit_id": "1",
                "quantity": "0.001"
            }	
        ]
    }
    ```
* Response Body:
    ```
    {
        "id": "1crwQARJ0wWC6HfDaMJUXJ9aj5j",
        "date": "2020-04-14T08:55:30Z",
        "request": {
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
                        "id": "1",
                        "name": "Masker"
                    },
                    "unit": {
                        "id": "1",
                        "name": "Buah"
                    },
                    "quantity": "3.00"
                }
            ]
        },
        "allocator": {
            "id": "1a7vQqjAeg3rCBryTNfaWpZnczh",
            "email": "admin@admin.com",
            "name": "admin@admin.com",
            "contact_person": null,
            "contact_number": null,
            "role": "ADMIN"
        },
        "photoUrl": "",
        "items": [
            {
                "id": "1crwQBCPN121OHwW1DExk01NY2F",
                "item": {
                    "id": "1",
                    "name": "Masker"
                },
                "unit": {
                    "id": "1",
                    "name": "Buah"
                },
                "quantity": "0.001"
            }
        ]
    }
    ```