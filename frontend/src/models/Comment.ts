export interface Comment {
    threadKey: string;
    commentKey: string;
    contributer: string;
    comment: string;
    updateDate: string;
}

export interface Comments {
    comments: Array<Comment>;
}
