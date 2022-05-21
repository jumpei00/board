import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Comment, Comments } from "../../../../models/Comment";
import { getAllCommentPayload, createCommentPayload, editCommentPayload, deleteCommentPayload } from "./type";

const initialState: Comments = {
    comments: [],
};

export const commentsSlice = createSlice({
    name: "comments",
    initialState,
    reducers: {
        getAllCommentByThreadKey: (state, action: PayloadAction<getAllCommentPayload>) => {
            state.comments = state.comments.filter((comment) => comment.threadKey === action.payload.threadKey)
        },
        createComment: (state, action: PayloadAction<createCommentPayload>) => {
            const now = new Date();
            const comment: Comment = {
                threadKey: action.payload.threadKey,
                commentKey: Math.random().toString(32).substring(2),
                contributer: action.payload.contributer,
                comment: action.payload.comment,
                updateDate: now.toLocaleDateString(),
            };
            state.comments.push(comment);
        },
        editComment: (state, action: PayloadAction<editCommentPayload>) => {
            const now = new Date();
            state.comments.forEach((comment) => {
                if (comment.commentKey === action.payload.commentKey) {
                    comment.comment = action.payload.comment;
                    comment.updateDate = now.toLocaleString();
                }
            });
        },
        deleteComment: (state, action: PayloadAction<deleteCommentPayload>) => {
            state.comments = state.comments.filter((comment) => comment.commentKey !== action.payload.commentKey);
        },
    },
});

export const { getAllCommentByThreadKey, createComment, editComment, deleteComment } = commentsSlice.actions;
export const commentsReducer = commentsSlice.reducer;
