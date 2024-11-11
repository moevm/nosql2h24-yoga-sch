
export interface Client {
    id: string,
    name: string,
    phone: string,
    gender: string,
    birthDate: string,
    createdAt: string,
    updatedAt: string,
    password: string,
    pictureUri: string,
    classIds: string[]
}

export interface Trainer {
    id: string,
    name: string,
    phone: string,
    studioId: string,
    gender: string,
    birthDate: string,
    createdAt: string,
    updatedAt: string,
    pictureUri: string,
    classIds: string[]
}

export interface Studio {
    id: string,
    address: string,
    classes: string[]
    trainers: string[]
}

export interface Class {
    id: string,
    name: string,
    time: string,
    studioId: string,
    trainerId: string,
    clients: string[]
}
