Art API Project:

Goals: Create an API that can be deployed via Docker. Art API is intended to be able to process data from the Metropolitan Museum of Art Open Access CSV found here https://github.com/metmuseum/openaccess


API v1 Features:
    - /objects GET retrieves all art objects in DB
    - /objects/{id} GET retrieves specific Art object
    - /objects POST Adds art object  to DB
    - /objects/{id} PUT Updates specific Art object
    - /objects/{id} DELETE deletes specific Art Object

Art Object  JSON Structure: 
{
    "object_id": 1,  //int   
    "is_highlight": "False", //string
    "accession_year": "2024", //string
    "department": "The American Wing", //string
    "title": "One-dollar Liberty Head Coin",  //string
    "object_name": "Coin",  //string
    "culture": "American",  /string
    "period": "Modern", //string
    "artist_display_name": "James Barton Longacre", //string
    "city": "New York", //string 
    "country": "Unites States" //string 
}


TODO: 
Create stored procedures to address Deletion creating gap in key
Create search based off any attribute.
Fix update to only update fields that are provided.