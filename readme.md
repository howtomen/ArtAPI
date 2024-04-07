<h1>Art API</h1>

<h3>Goals: Create an API that can be deployed via Docker. Art API is intended to be able to process data from the Metropolitan Museum of Art Open Access CSV found here https://github.com/metmuseum/openaccess</h3>

<h3>API Design</h3>

- Made up of HTTP Layer, Service Layer, and Repository Layer
- - Service Layer hold business logic
- - HTTP layer takes incoming requests and responds appropriately
- - Repo Layer interacts with DB
- Comes with Dockerfile, docker compose file and Tasker File
- - Dockerfile creates docker image and runs container after build
- - docker compose file sets up network, declares base DB image and volumes , Env variables needed for DB and App
- - Tasker file creates shortcut commanda to do things such as run, build, test, run-app, run-db amongst others. Will expand on these at a later time.  

<h3>API EndPoints</h3>

- /api/v3/art       GET retrieves all art objects in DB
- /api/v3/art/{id}  GET retrieves specific Art object
- /api/v3/art       POST Adds art object  to DB
- /api/v3/art/{id}  PUT Updates specific Art object
- /api/v3/art/{id}  DELETE deletes specific Art Object

<h3>TODO</h3 
    <text>           
- Create stored procedures to address Deletion creating gap in key<br>
- Create search based off any attribute.<br>
- Fix update to only update fields that are provided. <br>
    </text>
    <br>


**Art Object  JSON Structure:**
```
{
    "object_id": 1,                                 //int   
    "is_highlight": "False",                        //string
    "accession_year": "2024",                       //string
    "department": "The American Wing",              //string
    "title": "One-dollar Liberty Head Coin",        //string
    "object_name": "Coin",                          //string
    "culture": "American",                          //string
    "period": "Modern",                             //string
    "artist_display_name": "James Barton Longacre", //string
    "city": "New York",                             //string 
    "country": "Unites States"                      //string 
}
```



