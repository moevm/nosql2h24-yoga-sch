
const { ObjectId } = require('bson');


db = new Mongo().getDB("fitness_aggregator");

db.createCollection('studios', { capped: false });
db.createCollection('clients', { capped: false });
db.createCollection('trainers', { capped: false });
db.createCollection('classes', { capped: false });


// TRAINERS

db.trainers.insertOne(
    {
        "name": "Boris Va",
        "phone": "+7(999)444-4444",
        "gender": "MALE",
        "birth_date": new Date("2001-10-28T23:58:18Z"),
        "created_at": new Date("2024-10-28T23:58:18Z"),
        "updated_at": new Date("2024-10-29T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [],
        "studio_id": ObjectId("507f1f77bcf86cd799439011")
    }
);

db.clients.insertOne(
    {
        "name": "Elizaveta Andreeva",
        "phone": "+7(999)99-9999",
        "gender": "FEMALE",
        "birth_date": "2001-10-28T23:58:18Z",
        "created_at": "2024-10-28T23:58:18Z",
        "updated_at": "2024-10-29T23:58:18Z",
        "password": "1",
        "picture_uri": "cdn.example.com",
        "classes": []
    }
);
