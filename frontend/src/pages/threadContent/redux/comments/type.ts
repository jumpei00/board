export type getAllCommentPayload = {
    threadKey: string;
}

export type createCommentPayload = {
    threadKey: string;
    comment: string;
    contributer: string;
}

export type editCommentPayload = {
    commentKey: string;
    comment: string;
    contributer: string;
}

export type deleteCommentPayload = {
    commentKey: string;
    contributer: string;
}