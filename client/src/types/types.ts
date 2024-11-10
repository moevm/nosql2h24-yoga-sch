
export interface Client {
    _id: string,
    name: string,
    phone: string,
    gender: string,
    birthDate: string,
    createdAt: string,
    updatedAt: string,
    password: string,
    pictureUri: string,
    classes: string[]
}

export interface Trainer {
    _id: string,
    name: string,
    phone: string,
    studioId: string,
    gender: string,
    birthDate: string,
    createdAt: string,
    updatedAt: string,
    pictureUri: string,
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
    name: string,
    time: string,
    studioId: string,
    trainerId: string,
    clients: string[]
}
