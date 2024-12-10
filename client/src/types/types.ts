
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
    classesInfo: any
}

export interface Trainer {
    id: string,
    name: string,
    phone: string,
    studioInfo: any,
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
    classesInfo: any,
    trainersInfo: any
}

export interface Class {
    id: string,
    name: string,
    time: string,
    studioInfo: any,
    trainerInfo: any,
    clientsInfo: any
}
