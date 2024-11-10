# InvenTech

## Equipment inventory management software solution

Hello! In this readme you will find a brief description of the code and tech stack that was used to create the project.

#### Aim

This software aims to simplify and automate the process of cllecting data on new equipment, and autiding existing equipment in buildings.

#### Tech stack

- <ins>[AngularJS](InventTek/README.md)</ins>:
    - a popular and reliable enterprise-level frontend framework
    - used to create the web application to carry out task of equipment inventory management
    - queries data on equipment from the server and allows the user to scan technical plates using a camera for further automated image processing
    - alternatively, can be replaced by a native mobile app for IOS and Android

- <ins>[Python](image-text-extractor-python/text_extractor_v2.py)</ins>:
    - a script that sends an image via an API call to 
Google Vision, where the image gets processed
    - the API returns a text which was recognized on the image
    - the text then gets processes in python where the key information is picked out (e.g. model number, manufacturer)

- <ins>[Go](junction2024-server-go/main.go)</ins>:
    - a Go server receives the equipment data from the python script
    - data is then parsed and passed onto the database to be saved
    - additionally, the server provides the API endpoints to query, create and delete data from the database