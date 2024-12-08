const { ObjectId } = require('bson');


db = new Mongo().getDB("fitness_aggregator");

db.createCollection('studios', { capped: false });
db.createCollection('clients', { capped: false });
db.createCollection('trainers', { capped: false });
db.createCollection('classes', { capped: false });


// STUDIOS

db.studios.insertOne(
    {
        "_id": ObjectId("507f1f77bcf86cd799439011"),
        "address": "ul. Popova 1",
        "created_at": new Date("2024-10-28T23:58:18Z"),
        "updated_at": new Date("2024-10-29T23:58:18Z"),
        "class_ids": [
            ObjectId("009f1f77bcf86cd799439011"),
            ObjectId("019f1f77bcf86cd799439011"),
            ObjectId("029f1f77bcf86cd799439011"),
            ObjectId("039f1f77bcf86cd799439011"),
        ],
        "trainer_ids": [
            ObjectId("407f1f77bcf86cd799439011"),
            ObjectId("417f1f77bcf86cd799439011")
        ]
    }
);

db.studios.insertOne(
    {
        "_id": ObjectId("507f1f77bcf86cd799439012"),
        "address": "ul. Lesnaya 2",
        "created_at": new Date("2024-10-28T23:58:18Z"),
        "updated_at": new Date("2024-10-29T23:58:18Z"),
        "class_ids": [
            ObjectId("049f1f77bcf86cd799439011"),
            ObjectId("059f1f77bcf86cd799439011"),
            ObjectId("069f1f77bcf86cd799439011"),
            ObjectId("079f1f77bcf86cd799439011"),
        ],
        "trainer_ids": [
            ObjectId("427f1f77bcf86cd799439011"),
            ObjectId("437f1f77bcf86cd799439011")
        ]
    }
);

db.studios.insertOne(
    {
        "_id": ObjectId("507f1f77bcf86cd799439013"),
        "address": "ul. Kosmonavtov 3",
        "created_at": new Date("2024-10-28T23:58:18Z"),
        "updated_at": new Date("2024-10-29T23:58:18Z"),
        "class_ids": [
            ObjectId("089f1f77bcf86cd799439011"),
            ObjectId("099f1f77bcf86cd799439011"),
            ObjectId("109f1f77bcf86cd799439011"),
            ObjectId("119f1f77bcf86cd799439011"),
        ],
        "trainer_ids": [
            ObjectId("447f1f77bcf86cd799439011"),
            ObjectId("457f1f77bcf86cd799439011")
        ]
    }
);


// TRAINERS

db.trainers.insertOne(
    {
        "_id": ObjectId("407f1f77bcf86cd799439011"),
        "name": "Boris Va",
        "phone": "+7(999)444-4444",
        "gender": "MALE",
        "birth_date": new Date("2001-10-28T23:58:18Z"),
        "created_at": new Date("2024-10-28T23:58:18Z"),
        "updated_at": new Date("2024-10-29T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("009f1f77bcf86cd799439011"),
            ObjectId("019f1f77bcf86cd799439011")
        ],
        "studio_id": ObjectId("507f1f77bcf86cd799439011")
    }
);

db.trainers.insertOne(
    {
        "_id": ObjectId("417f1f77bcf86cd799439011"),
        "name": "Egor Shmatcko",
        "phone": "+7(999)333-3333",
        "gender": "MALE",
        "birth_date": new Date("2002-09-11T23:52:14Z"),
        "created_at": new Date("2024-09-22T23:58:18Z"),
        "updated_at": new Date("2024-09-22T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("029f1f77bcf86cd799439011"),
            ObjectId("039f1f77bcf86cd799439011")
        ],
        "studio_id": ObjectId("507f1f77bcf86cd799439011")
    }
);

db.trainers.insertOne(
    {
        "_id": ObjectId("427f1f77bcf86cd799439011"),
        "name": "Oleg Gi",
        "phone": "+7(999)222-2222",
        "gender": "MALE",
        "birth_date": new Date("2012-04-22T23:52:14Z"),
        "created_at": new Date("2024-01-22T23:58:18Z"),
        "updated_at": new Date("2024-01-22T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("049f1f77bcf86cd799439011"),
            ObjectId("059f1f77bcf86cd799439011")
        ],
        "studio_id": ObjectId("507f1f77bcf86cd799439012")
    }
);

db.trainers.insertOne(
    {
        "_id": ObjectId("437f1f77bcf86cd799439011"),
        "name": "Vera Hans",
        "phone": "+7(999)111-1111",
        "gender": "FEMALE",
        "birth_date": new Date("2008-03-03T23:52:14Z"),
        "created_at": new Date("2023-02-11T23:58:18Z"),
        "updated_at": new Date("2023-02-11T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("069f1f77bcf86cd799439011"),
            ObjectId("079f1f77bcf86cd799439011")
        ],
        "studio_id": ObjectId("507f1f77bcf86cd799439012")
    }
);

db.trainers.insertOne(
    {
        "_id": ObjectId("447f1f77bcf86cd799439011"),
        "name": "Lilya Kio",
        "phone": "+7(999)000-0000",
        "gender": "FEMALE",
        "birth_date": new Date("1998-11-12T23:52:14Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("089f1f77bcf86cd799439011"),
            ObjectId("099f1f77bcf86cd799439011")
        ],
        "studio_id": ObjectId("507f1f77bcf86cd799439013")
    }
);

db.trainers.insertOne(
    {
        "_id": ObjectId("457f1f77bcf86cd799439011"),
        "name": "Nastya Vecher",
        "phone": "+7(998)999-9999",
        "gender": "MALE",
        "birth_date": new Date("1998-11-12T23:52:14Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("109f1f77bcf86cd799439011"),
            ObjectId("119f1f77bcf86cd799439011")
        ],
        "studio_id": ObjectId("507f1f77bcf86cd799439013")
    }
);


// Classes

db.classes.insertOne(
    {
        "_id": ObjectId("009f1f77bcf86cd799439011"),
        "name": "Vodnaya yoga",
        "time": new Date("2025-01-09T13:00:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("407f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439011"),
        "client_ids": [
            ObjectId("307f1f77bcf86cd799439011"),
            ObjectId("317f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("019f1f77bcf86cd799439011"),
        "name": "Goryachaya yoga",
        "time": new Date("2025-01-09T13:30:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("407f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439011"),
        "client_ids": [
            ObjectId("307f1f77bcf86cd799439011"),
            ObjectId("317f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("029f1f77bcf86cd799439011"),
        "name": "Holodnaya yoga",
        "time": new Date("2025-01-09T14:00:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("417f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439011"),
        "client_ids": [
            ObjectId("317f1f77bcf86cd799439011"),
            ObjectId("327f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("039f1f77bcf86cd799439011"),
        "name": "Coal yoga",
        "time": new Date("2025-01-09T14:30:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("417f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439011"),
        "client_ids": [
            ObjectId("317f1f77bcf86cd799439011"),
            ObjectId("327f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("049f1f77bcf86cd799439011"),
        "name": "Vodnaya yoga",
        "time": new Date("2025-01-09T13:00:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("427f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439012"),
        "client_ids": [
            ObjectId("327f1f77bcf86cd799439011"),
            ObjectId("337f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("059f1f77bcf86cd799439011"),
        "name": "Goryachaya yoga",
        "time": new Date("2025-01-09T13:30:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("427f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439012"),
        "client_ids": [
            ObjectId("327f1f77bcf86cd799439011"),
            ObjectId("337f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("069f1f77bcf86cd799439011"),
        "name": "Holodnaya yoga",
        "time": new Date("2025-01-09T14:00:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("437f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439012"),
        "client_ids": [
            ObjectId("337f1f77bcf86cd799439011"),
            ObjectId("347f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("079f1f77bcf86cd799439011"),
        "name": "Coal yoga",
        "time": new Date("2025-01-09T14:30:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("437f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439012"),
        "client_ids": [
            ObjectId("337f1f77bcf86cd799439011"),
            ObjectId("347f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("089f1f77bcf86cd799439011"),
        "name": "Vodnaya yoga",
        "time": new Date("2025-01-09T13:00:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("447f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439013"),
        "client_ids": [
            ObjectId("347f1f77bcf86cd799439011"),
            ObjectId("357f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("099f1f77bcf86cd799439011"),
        "name": "Goryachaya yoga",
        "time": new Date("2025-01-09T13:30:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("447f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439013"),
        "client_ids": [
            ObjectId("347f1f77bcf86cd799439011"),
            ObjectId("357f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("109f1f77bcf86cd799439011"),
        "name": "Holodnaya yoga",
        "time": new Date("2025-01-09T14:00:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("457f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439013"),
        "client_ids": [
            ObjectId("357f1f77bcf86cd799439011"),
            ObjectId("307f1f77bcf86cd799439011")
        ]
    }
);

db.classes.insertOne(
    {
        "_id": ObjectId("119f1f77bcf86cd799439011"),
        "name": "Coal yoga",
        "time": new Date("2025-01-09T14:30:00Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "trainer_id": ObjectId("457f1f77bcf86cd799439011"),
        "studio_id": ObjectId("507f1f77bcf86cd799439013"),
        "client_ids": [
            ObjectId("357f1f77bcf86cd799439011"),
            ObjectId("307f1f77bcf86cd799439011")
        ]
    }
);


// CLIENTS

db.clients.insertOne(
    {
        "_id": ObjectId("307f1f77bcf86cd799439011"),
        "name": "Elizaveta Andreeva",
        "phone": "+7(999)999-9999",
        "gender": "FEMALE",
        "birth_date": new Date("2001-10-28T23:58:18Z"),
        "created_at": new Date("2024-10-28T23:58:18Z"),
        "updated_at": new Date("2024-10-29T23:58:18Z"),
        "password": "1",
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("009f1f77bcf86cd799439011"),
            ObjectId("019f1f77bcf86cd799439011"),
            ObjectId("109f1f77bcf86cd799439011"),
            ObjectId("119f1f77bcf86cd799439011")
        ]
    }
);

db.clients.insertOne(
    {
        "_id": ObjectId("317f1f77bcf86cd799439011"),
        "name": "Egor Butylo",
        "phone": "+7(999)888-8888",
        "gender": "MALE",
        "birth_date": new Date("2002-09-11T23:52:14Z"),
        "created_at": new Date("2024-09-22T23:58:18Z"),
        "updated_at": new Date("2024-09-22T23:58:18Z"),
        "password": "1",
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("009f1f77bcf86cd799439011"),
            ObjectId("019f1f77bcf86cd799439011"),
            ObjectId("029f1f77bcf86cd799439011"),
            ObjectId("039f1f77bcf86cd799439011")
        ]
    }
);

db.clients.insertOne(
    {
        "_id": ObjectId("327f1f77bcf86cd799439011"),
        "name": "Oleg Mongol",
        "phone": "+7(999)777-7777",
        "gender": "MALE",
        "birth_date": new Date("2012-04-22T23:52:14Z"),
        "created_at": new Date("2024-01-22T23:58:18Z"),
        "updated_at": new Date("2024-01-22T23:58:18Z"),
        "password": "1",
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("029f1f77bcf86cd799439011"),
            ObjectId("039f1f77bcf86cd799439011"),
            ObjectId("049f1f77bcf86cd799439011"),
            ObjectId("059f1f77bcf86cd799439011")
        ]
    }
);

db.clients.insertOne(
    {
        "_id": ObjectId("337f1f77bcf86cd799439011"),
        "name": "Irina Chikipiki",
        "phone": "+7(999)666-6666",
        "gender": "FEMALE",
        "birth_date": new Date("2008-03-03T23:52:14Z"),
        "created_at": new Date("2023-02-11T23:58:18Z"),
        "updated_at": new Date("2023-02-11T23:58:18Z"),
        "password": "1",
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("049f1f77bcf86cd799439011"),
            ObjectId("059f1f77bcf86cd799439011"),
            ObjectId("069f1f77bcf86cd799439011"),
            ObjectId("079f1f77bcf86cd799439011")
        ]
    }
);

db.clients.insertOne(
    {
        "_id": ObjectId("347f1f77bcf86cd799439011"),
        "name": "Vladislav Frolov",
        "phone": "+7(999)555-5555",
        "gender": "MALE",
        "birth_date": new Date("1998-11-12T23:52:14Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "password": "1",
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("069f1f77bcf86cd799439011"),
            ObjectId("079f1f77bcf86cd799439011"),
            ObjectId("089f1f77bcf86cd799439011"),
            ObjectId("099f1f77bcf86cd799439011")
        ]
    }
);

db.clients.insertOne(
    {
        "_id": ObjectId("357f1f77bcf86cd799439011"),
        "name": "Vova Shustiy",
        "phone": "+7(998)888-8888",
        "gender": "MALE",
        "birth_date": new Date("2011-01-11T23:52:14Z"),
        "created_at": new Date("2022-04-22T23:58:18Z"),
        "updated_at": new Date("2022-04-22T23:58:18Z"),
        "password": "1",
        "picture_uri": "cdn.example.com",
        "class_ids": [
            ObjectId("089f1f77bcf86cd799439011"),
            ObjectId("099f1f77bcf86cd799439011"),
            ObjectId("109f1f77bcf86cd799439011"),
            ObjectId("119f1f77bcf86cd799439011")
        ]
    }
);
