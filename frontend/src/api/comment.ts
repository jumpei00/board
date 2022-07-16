import { CreateCommentPayload, DeleteCommentPayload, UpdateCommentPayload } from "../state/comments/modules"
import { boardApi } from "./init"

export const getAllCommentsAPI = (payload: string) => {
    return boardApi.get(`/api/threads/${payload}/comments`)
}

export const createCommentAPI = (payload: CreateCommentPayload) => {
    return boardApi.post(`/api/threads/${payload.threadKey}/comments`, payload.body)
}

export const updateCommentAPI = (payload: UpdateCommentPayload) => {
    return boardApi.put(`/api/threads/${payload.threadKey}/comments/${payload.commentKey}`, payload.body)
}

export const deleteCommentAPI = (payload: DeleteCommentPayload) => {
    return boardApi.delete(`/api/threads/${payload.threadKey}/comments/${payload.commentKey}`);
}