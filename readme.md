<h1>Art API</h1>

<h3>Goals: Create an API that can be deployed via Docker. Art API is intended to be able to process data from the Metropolitan Museum of Art Open Access CSV found here https://github.com/metmuseum/openaccess</h3>


<h3>API Features</h3>

- /objects GET retrieves all art objects in DB
- /objects/{id} GET retrieves specific Art object
- /objects POST Adds art object  to DB
- /objects/{id} PUT Updates specific Art object
- /objects/{id} DELETE deletes specific Art Object

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



