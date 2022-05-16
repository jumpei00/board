export interface Comment {
    commentKey: string;
    contributer: string;
    comment: string;
    pictureURL: string;
    updateDate: string;
}

export interface Comments {
    comments: Array<Comment>;
}
