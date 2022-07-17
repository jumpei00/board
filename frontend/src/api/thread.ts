import { CreateThreadPayload, UpdateThreadPayload } from "../state/threads/modules";
import { boardApi } from "./init";

export const getAllThreadsAPI = () => {
    return boardApi.get(`/api/threads`);
};

export const createThreadAPI = (payload: CreateThreadPayload) => {
    return boardApi.post(`/api/threads`, payload);
};

export const updateThreadAPI = (payload: UpdateThreadPayload) => {
    return boardApi.put(`/api/threads/${payload.threadKey}`, payload.body);
};

export const deleteThreadAPI = (payload: string) => {
    return boardApi.delete(`/api/threads/${payload}`);
};
