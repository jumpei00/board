import { combineReducers, createSlice, PayloadAction } from "@reduxjs/toolkit";
import {
    Comments,
    CommentResponse,
    FetchCommentResponse,
    CreateCommentResponse,
    UpdateCommentResponse,
} from "../../models/Comment";

// ---- Payload ---- //
export type CreateCommentPayload = {
    threadKey: string;
    body: {
        comment: string;
        contributor: string;
    };
};

export type UpdateCommentPayload = {
    threadKey: string;
    commentKey: string;
    body: {
        comment: string;
        contributor: string;
    };
};

export type DeleteCommentPayload = {
    threadKey: string;
    commentKey: string;
};

// ---- state ---- //
const initialState: Comments = {
    thread: {
        threadKey: "",
        title: "",
        contributor: "",
        views: 0,
        commentSum: 0,
        createDate: "",
        updateDate: "",
    },
    comments: [],
};

const initialSagaResponse: CommentResponse = {
    fetchResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            thread: {
                threadKey: "",
                title: "",
                contributor: "",
                views: 0,
                commentSum: 0,
                createDate: "",
                updateDate: "",
            },
            comments: [],
        },
    },
    createResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            commentKey: "",
            contributor: "",
            comment: "",
            createDate: "",
            updateDate: "",
        },
    },
    updateResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            commentKey: "",
            contributor: "",
            comment: "",
            createDate: "",
            updateDate: "",
        },
    },
    deleteResponse: {
        pending: false,
        success: false,
        error: false,
    },
};

// ---- comment reducer ---- //
export const commentSlice = createSlice({
    name: "comment",
    initialState,
    reducers: {
        storeComments: (state, action: PayloadAction<FetchCommentResponse>) => {
            return action.payload;
        },
        addComment: (state, action: PayloadAction<CreateCommentResponse>) => {
            state.comments.unshift(action.payload);
            return {
                ...state,
                thread: {
                    ...state.thread,
                    commentSum: state.thread.commentSum + 1,
                    updateDate: action.payload.updateDate,
                },
            };
        },
        editComment: (state, action: PayloadAction<UpdateCommentResponse>) => {
            let taegetComment = state.comments.find((comment) => comment.commentKey === action.payload.commentKey);
            if (taegetComment) {
                taegetComment = {
                    ...taegetComment,
                    ...action.payload,
                };
            }
        },
        deleteComment: (state, action: PayloadAction<string>) => {
            state.comments = state.comments.filter((comment) => comment.commentKey !== action.payload);
        },
        clearComment: () => initialState,
    },
});

// ---- comment saga reducer ---- //
const sagaSliceName = "commentSaga";

export const commentSagaActionsType = {
    getAll: `${sagaSliceName}/getAll`,
    create: `${sagaSliceName}/create`,
    update: `${sagaSliceName}/update`,
    delete: `${sagaSliceName}/delete`,
};

export const commentSagaSlice = createSlice({
    name: sagaSliceName,
    initialState: initialSagaResponse,
    reducers: {
        getAll: (state, action: PayloadAction<string>) => {
            state.fetchResponse.pending = true;
        },
        getAllDone: (state, action: PayloadAction<FetchCommentResponse>) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.success = true;
            commentSlice.actions.storeComments(action.payload);
        },
        getAllFail: (state) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.error = true;
        },
        create: (state, action: PayloadAction<CreateCommentPayload>) => {
            state.createResponse.pending = true;
        },
        createDone: (state, action: PayloadAction<CreateCommentResponse>) => {
            state.createResponse.pending = false;
            state.createResponse.success = true;
            commentSlice.actions.addComment(action.payload);
        },
        createFail: (state) => {
            state.createResponse.pending = false;
            state.createResponse.error = true;
        },
        update: (state, action: PayloadAction<UpdateCommentPayload>) => {
            state.updateResponse.pending = true;
        },
        updateDone: (state, action: PayloadAction<UpdateCommentResponse>) => {
            state.updateResponse.pending = false;
            state.updateResponse.success = true;
            commentSlice.actions.editComment(action.payload);
        },
        updateFail: (state) => {
            state.updateResponse.pending = false;
            state.updateResponse.error = true;
        },
        delete: (state, action: PayloadAction<DeleteCommentPayload>) => {
            state.deleteResponse.pending = true;
        },
        deleteDone: (state, action: PayloadAction<string>) => {
            state.deleteResponse.pending = false;
            state.deleteResponse.success = true;
            commentSlice.actions.deleteComment(action.payload);
        },
        deleteFail: (state) => {
            state.deleteResponse.pending = false;
            state.deleteResponse.error = true;
        },
    },
});

export const commentActions = commentSlice.actions;
export const commentSagaActions = commentSagaSlice.actions;
export const commentReducer = combineReducers({
    commentState: commentSlice.reducer,
    commentSaga: commentSagaSlice.reducer,
});
