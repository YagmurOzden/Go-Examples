PS: The go.sum file contains all the dependency checksums, and is managed by the go tools. We don’t have to worry about it.

configs is for modularizing project configuration files

controllers is for modularizing application logics.

models is for modularizing data and database logics.

responses is for modularizing files describing the response we want our API to give. This will become clearer later on.

routes is for modularizing URL pattern and handler information.

## On postman ##
Add User
POST http://0.0.0.0:6000/user
BODY:
{
    "name": "Yağmur Ozden",
    "location":"Istanbul",
    "title":"Computer Engineer"
}
RESPONSE:
{
    "status": 201,
    "message": "success",
    "data": {
        "data": {
            "InsertedID": "63450b37c9d7b161570815d8"
        }
    }
}
###
Get All Users
GET http://0.0.0.0:6000/users
RESPONSE: 
{
    "status": 200,
    "message": "success",
    "data": {
        "data": [
            {
                "id": "63440ebeacc36f0897fa398c",
                "name": "Duru Ozden",
                "location": "Istanbul",
                "title": "Student"
            },
            {
                "id": "63450b37c9d7b161570815d7",
                "name": "Yağmur Ozden",
                "location": "Istanbul",
                "title": "Computer Engineer"
            }
        ]
    }
}
##
GET USER WITH ID

GET http://0.0.0.0:6000/user/63450b37c9d7b161570815d7
RESPONSE: {
    "status": 200,
    "message": "success",
    "data": {
        "data": {
            "id": "63450b37c9d7b161570815d7",
            "name": "Yağmur Ozden",
            "location": "Istanbul",
            "title": "Computer Engineer"
        }
    }
}

##
MODIFY A USER
PUT http://0.0.0.0:6000/user/63440ebeacc36f0897fa398c
BODY: 
{
                "id": "63440ebeacc36f0897fa398c",
                "name": "Duru Ozden",
                "location": "Artvin",
                "title": "Student"
}
RESPONSE:
{
    "status": 200,
    "message": "success",
    "data": {
        "data": {
            "id": "63440ebeacc36f0897fa398c",
            "name": "Duru Ozden",
            "location": "Artvin",
            "title": "Student"
        }
    }
}
##
DELETE A USER
DELETE  http://0.0.0.0:6000/user/63440ebeacc36f0897fa398c
RESPONSE:
{
    "status": 200,
    "message": "success",
    "data": {
        "data": "User successfully deleted!"
    }
}