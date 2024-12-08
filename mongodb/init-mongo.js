db = new Mongo().getDB("fitness_aggregator");

db.createCollection('studios', { capped: false });
db.createCollection('clients', { capped: false });
db.createCollection('trainers', { capped: false });
db.createCollection('classes', { capped: false });

db.clients.insertOne(
    {
        "name": "Elizaveta Andreeva",
        "phone": "+7(999)999-9999",
        "gender": "FEMALE",
        "birth_date": "2001-10-28T23:58:18Z",
        "created_at": "2024-10-28T23:58:18Z",
        "updated_at": "2024-10-29T23:58:18Z",
        "password": "1",
        "picture_uri": "cdn.example.com",
        "classes": []
    }
);
