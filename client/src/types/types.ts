
export interface Client {
    _id: string,
    name: string,
    phone: string,
    gender: string,
    birth_date: string,
    created_at: string,
    updated_at: string,
    password: string,
    picture_uri: string,
    classes: string[]
}

export interface Trainer {
    _id: string,
    name: string,
    phone: string,
    studio_id: string,
    gender: string,
    birth_date: string,
    created_at: string,
    updated_at: string,
    picture_uri: string,
    classes: string[]
}

export interface Studio {
    _id: string,
    address: string,
    classes: string[]
    trainers: string[]
}

export interface Class {
    _id: string,
    class_name: string,
    time: string,
    studio_id: string,
    trainer_id: string,
    clients: string[]
}
