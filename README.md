# go-rest
A simple REST based application built in GOLANG to perform basic CRUD operations

This also includes a dockerfile which can be used to build docker image and push that image to dockerhub.

I have uploaded separate terraform code which will preconfigure DB for app in a pod and also will build/deploy application to separate pod.

#### **Dependencies -**

This app will connect to a mysql database and following are default DB configs -

DB_SERVER_NAME=localhost
DB_NAME=user
DB_USER=admin
DB_PASSWORD=Test1234

These values will be automatically overridden if these are set as ENV Variables.


#### **Table structure  -**

The database should have a table names "person" with following columns :


mysql> SHOW COLUMNS from person;

| Field     | Type         | Null | Key | Default | Extra |
|----------:|-------------:|-----:|----:|--------:|------:|
| id        | int          | NO   | PRI | NULL    |       |
| lastname  | varchar(255) | YES  |     | NULL    |       |
| firstname | varchar(255) | YES  |     | NULL    |       |
| city      | varchar(255) | YES  |     | NULL    |       |

4 rows in set (0.00 sec)

mysql>

#### **Usage -**

This app supports below REST calls :

GET - http://<host-name>:8090/get --- This will get all the persons in Database in JSON format

GET - http://<host-name>:8090/get/{id} --- Get details of person with ID 

PUT - http://<host-name>:8090/update/{id} -- Update the details about person with ID, it expects payload as below

            {
                "id": 1,
                "lastname": "last",
                "firstname": "first",
                "city": "Toronto"
            }

POST - http://<host-name>:8090/create -- This will add entry to database and expects below payload 

        {
            "id": 1,
            "lastname": "last",
            "firstname": "first",
            "city": "Toronto"
        }

DELETE - http://<host-name>:8090/delete/{id} --- This will delete the record of person with ID.

