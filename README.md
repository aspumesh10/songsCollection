# Songs Collection App

## This application provides endpoints to interact with songs information fetched from spotify account


### Built With
Go 1.18

Mysql

Postman

## Getting Started

### Prerequisite

    Install Mysql on your system by using following link
    https://dev.mysql.com/downloads/mysql/

    Install Golang support in your system
    https://go.dev/doc/install

    Download postman
    https://www.postman.com/downloads/

### Setup
   clone the repo to local directory
    
    git clone https://github.com/aspumesh10/songsCollection.git .
    

   add .env File in the main repo directory 
    with content as below including mysql db details 
    and spotify account details
   ```sh
    db_name=artistDb
    db_host=127.0.0.1
    db_port=3306
    db_username=<username>
    db_password=<password
    client_id=<spotify account clientId>
    secret_key=<Spotiffy account secret key>
   ```

   Build the project
    
    go build main.go
     
   run the project
    
    ./main
    
   The server will run on the port 3000

### Api end point details

1. This endpoint retrieves track information from spotify and stores its basic metadata
in local database for further processing

    Request:
    
        Method: POST
        URL: http://localhost:3000/v1/addTrack
        Body Variable: (x-www-form-urlencoded)
            isrc (string, required): The ISRC of the track to be retrieved.

    Response:
    ```sh
        {
            "data": {
                "TrackId": 1
            }
        }
    ```
2.  This endpoint retrieves track information based on the provided ISRC (International Standard Recording Code).

    Request:

        Method: GET
        URL: http://localhost:3000/v1/getTracks/:isrc
        Path Variable: isrc (string, required): The ISRC of the track to be retrieved.

    Response:
    ```sh
        {
            "data": {
                "id": 5,
                "title": "Summer Paradise (feat. Sean Paul) - Single Version",
                "isrc": "USAT21200343",
                "popularity": 47,
                "albumName": "Summer Paradise (feat. Sean Paul)",
                "image": "https://i.scdn.co/image/ab67616d0000b27308ca7af4356772a6518ef8af",
                "releaseDate": "2012-02-21T05:30:00+05:30",
                "isActive": 1,
                "isDeleted": 0,
                "createdById": 1,
                "updatedById": {
                    "Int32": 0,
                    "Valid": false
                },
                "createdAt": "2024-01-07T17:22:17+05:30",
                "updatedAt": "2024-01-07T17:22:17+05:30"
            },
            "message": ""
        }
    ```

3.  This endpoint retrieves tracks by artist from the server. 
    
    Request:

        Method: GET
        URL: http://localhost:3000/v1/getTracks/?artist=paul

    Response:
    ```sh
        {
            "data": [
                {
                    "id": 4,
                    "title": "Cheap Thrills (feat. Sean Paul)",
                    "isrc": "USRC11600201",
                    "popularity": 69,
                    "albumName": "Cheap Thrills (feat. Sean Paul)",
                    "image": "https://i.scdn.co/image/ab67616d0000b2739f5a5f3d50cd3939ba8e465c",
                    "releaseDate": "2016-02-11T05:30:00+05:30",
                    "isActive": 1,
                    "isDeleted": 0,
                    "createdById": 1,
                    "updatedById": {
                        "Int32": 0,
                        "Valid": false
                    },
                    "createdAt": "2024-01-07T17:20:20+05:30",
                    "updatedAt": "2024-01-07T17:20:20+05:30"
                },
                {
                    "id": 5,
                    "title": "Summer Paradise (feat. Sean Paul) - Single Version",
                    "isrc": "USAT21200343",
                    "popularity": 47,
                    "albumName": "Summer Paradise (feat. Sean Paul)",
                    "image": "https://i.scdn.co/image/ab67616d0000b27308ca7af4356772a6518ef8af",
                    "releaseDate": "2012-02-21T05:30:00+05:30",
                    "isActive": 1,
                    "isDeleted": 0,
                    "createdById": 1,
                    "updatedById": {
                        "Int32": 0,
                        "Valid": false
                    },
                    "createdAt": "2024-01-07T17:22:17+05:30",
                    "updatedAt": "2024-01-07T17:22:17+05:30"
                }
            ],
            "error": "",
            "totalCount": 2
        }
        ```
