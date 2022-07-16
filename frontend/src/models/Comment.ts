import { Thread } from "./thread";

// ---- state ---- //
export interface Comment {
    commentKey: string;
    contributor: string;
    comment: string;
    createDate: string;
    updateDate: string;
}

export interface Comments {
    thread: Thread;
    comments: Array<Comment>;
}

// ---- saga response ---- //
export interface CreateCommentResponse {
    commentKey: string;
    contributor: string;
    comment: string;
    createDate: string;
    updateDate: string;
}

export interface UpdateCommentResponse {
    commentKey: string;
    contributor: string;
    comment: string;
    createDate: string;
    updateDate: string;
}

export interface FetchCommentResponse {
    thread: Thread;
    comments: Array<Comment>;
}

export interface CommentResponse {
    fetchResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: FetchCommentResponse;
    };
    createResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: CreateCommentResponse;
    };
    updateResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: UpdateCommentResponse;
    };
    deleteResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
    }
}
