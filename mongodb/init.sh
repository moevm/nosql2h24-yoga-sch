#!/bin/bash

mongosh --eval "use fitness_aggregator"
mongosh --eval "db.createCollection('clients')"
mongosh --eval "db.createCollection('trainers')"
mongosh --eval "db.clients.insertOne({"name": "Elizaveta Andreeva", "phone": "+7(999)99-9999", "gender": "FEMALE", "birth_date": "2002-10-28T23:58:18Z", "created_at": "2040-10-28T23:58:18Z","updated_at": "2040-10-29T23:58:18Z","password": "11111111","picture_uri": "https://cdn.example.com","classes": []})"
mongosh --eval "db.trainers.insertOne({"name": "Olga Ivanova", "phone": "+7(999)11-1111", "gender": "FEMALE", "birth_date": "2000-10-28T23:58:18Z", "created_at": "2040-10-28T23:58:18Z","updated_at": "2040-10-29T23:58:18Z","studio_id": "","picture_uri": "https://cdn.example.com","classes": []})"