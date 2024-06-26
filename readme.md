<h1>Art API</h1>

**Goals:** Create a production ready containerized API and db that can be deployed via Kubernetes to store and access art records. Art API is intended to be able to process data from the Metropolitan Museum of Art Open Access CSV found here https://github.com/metmuseum/openaccess.

<h3>API Design & Features</h3>

- Made up of HTTP Layer, Service Layer, and Repository Layer
    - Service Layer hold business logic
        - initiates service and routes HTTP request to correct repo function.
    - HTTP layer takes incoming requests and responds appropriately
        - Implements Graceful Shutdown 
        - Implements **JWT Auth** for POST, UPDATE and DELETE endpoints
        - **ZeroLog** for logging with Request ID, elapsed time, method accessed, status code, and user agent. Also set up to provide debugging info and write to console and log file. Log file would be used in production and saved to be processed by another service (datadog etc...)
            - ZeroLog is passed within context throughout the application.
        - Includes Middleware for logging(Zerolog), setting timeout for requests, setting JSON content type
        - input is validated using **validatorv10** 
    - Repo Layer interacts with DB
- Dockerfile, docker compose file and Tasker File
    - Dockerfile has Build Env and Deployment set up using different base images
        - build env uses full golang image and deployment uses lighter alpine image
    - docker compose file 
        - Two services: DB & App.
        - Sets up health check for DB to make sure App doesnt make requests before DB is ready
        - Sets up environment Variables for both DB & App
        - Sets network, exposes port etc...
    - Tasker file creates shortcut commands to do things such as run, build, test, run-app, run-db amongst others. Will expand on these at a later time.  

<h3>API EndPoints</h3>

- /api/v3/art       GET retrieves all art objects in DB
- /api/v3/art/{id}  GET retrieves specific Art object
- /api/v3/art       POST Adds art object  to DB
- /api/v3/art/{id}  PUT Updates specific Art object
- /api/v3/art/{id}  DELETE deletes specific Art Object

<h3>TODO</h3>

- ~~Implement Logrus Logging levels and more logging aside from HTTP middleware logging~~
    - ~~implement log exporting to file which can be extracted from container~~
- Implement a safer way to handle config, ports, server info, passwords and other things
- ~~Implement Request validation~~
- Implement Tests
- Implement CI
- Implement k8s 
- Finish implementing Search endpoint


<h3>Art Object  JSON Structure:</h3>

```
{
    "ID": "b8b907d0-331c-48df-88ea-580355ba6dc5",   //uuid
    "object_id": 1,                                 //int; numeric; required
    "is_highlight": false,                          //bool
    "accession_year": "2024",                       //string; numeric
    "department": "The American Wing",              //string
    "title": "One-dollar Liberty Head Coin",        //string; required
    "object_name": "Coin",                          //string; required
    "culture": "American",                          //string
    "period": "Modern",                             //string
    "artist_display_name": "James Barton Longacre", //string; required
    "city": "New York",                             //string 
    "country": "Unites States"                      //string 
}
```



