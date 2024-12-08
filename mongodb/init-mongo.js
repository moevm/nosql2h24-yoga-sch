
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

db.trainers.insertOne(
    {
        "name": "Egor Shmatcko",
        "phone": "+7(999)333-3333",
        "gender": "MALE",
        "birth_date": new Date("2002-09-11T23:52:14Z"),
        "created_at": new Date("2024-09-22T23:58:18Z"),
        "updated_at": new Date("2024-09-22T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [],
        "studio_id": ObjectId("507f1f77bcf86cd799439011")
    }
);

db.trainers.insertOne(
    {
        "name": "Oleg Gi",
        "phone": "+7(999)222-2222",
        "gender": "MALE",
        "birth_date": new Date("2012-04-22T23:52:14Z"),
        "created_at": new Date("2024-01-22T23:58:18Z"),
        "updated_at": new Date("2024-01-22T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [],
        "studio_id": ObjectId("507f1f77bcf86cd799439012")
    }
);

db.trainers.insertOne(
    {
        "name": "Vera Hans",
        "phone": "+7(999)111-1111",
        "gender": "FEMALE",
        "birth_date": new Date("2008-03-03T23:52:14Z"),
        "created_at": new Date("2023-02-11T23:58:18Z"),
        "updated_at": new Date("2023-02-11T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [],
        "studio_id": ObjectId("507f1f77bcf86cd799439012")
    }
);

db.trainers.insertOne(
    {
        "name": "Lilya Kio",
        "phone": "+7(999)000-0000",
        "gender": "FEMALE",
        "birth_date": new Date("1998-11-12T23:52:14Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [],
        "studio_id": ObjectId("507f1f77bcf86cd799439013")
    }
);

db.trainers.insertOne(
    {
        "name": "Nastya Vecher",
        "phone": "+7(998)999-9999",
        "gender": "MALE",
        "birth_date": new Date("1998-11-12T23:52:14Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [],
        "studio_id": ObjectId("507f1f77bcf86cd799439013")
    }
);

// STUDIOS

db.studios.insertOne(
    {
        "_id": ObjectId("507f1f77bcf86cd799439011"),
        "address": "ul. Popova 1",
        "created_at": new Date("2024-10-28T23:58:18Z"),
        "updated_at": new Date("2024-10-29T23:58:18Z"),
        "class_ids": [],
        "trainer_ids": [ObjectId("")]
    }
);

db.studios.insertOne(
    {
        "_id": ObjectId("507f1f77bcf86cd799439012"),
        "address": "ul. Lesnaya 2",
        "created_at": new Date("2024-10-28T23:58:18Z"),
        "updated_at": new Date("2024-10-29T23:58:18Z"),
        "class_ids": [],
        "trainer_ids": []
    }
);