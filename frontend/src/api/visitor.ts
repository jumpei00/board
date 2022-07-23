import boardApi from "./init"

export const visitorStat = () => {
    return boardApi.get(`/api/visitor`)
}

export const visitorCountup = () => {
    return boardApi.put(`/api/visitor/coutup`)
}