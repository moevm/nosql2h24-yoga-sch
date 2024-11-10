db = new Mongo().getDB("fitness_aggregator");

db.createCollection('studios', { capped: false });
db.createCollection('clients', { capped: false });
db.createCollection('trainers', { capped: false });
db.createCollection('classes', { capped: false });

db.clients.insertOne(
    {
        "name": "Elizaveta Andreeva",
        "phone": "+7(999)99-9999",
        "gender": "FEMALE",
        "birth_date": "2001-10-28T23:58:18Z",
        "created_at": "2040-10-28T23:58:18Z",
        "updated_at": "2040-10-29T23:58:18Z",
        "password": "11111111",
        "picture_uri": "https://cdn.example.com",
        "classes": []
    }
);

db.trainers.insertOne(
    {
        "name": "Olga Ivanova",
        "phone": "+7(922)99-1111",
        "gender": "FEMALE",
        "birth_date": "2001-10-28T23:58:18Z",
        "created_at": "2040-10-28T23:58:18Z",
        "updated_at": "2040-10-29T23:58:18Z",
        "picture_uri": "https://cdn.example.com",
        "class_ids": []
    }
);
