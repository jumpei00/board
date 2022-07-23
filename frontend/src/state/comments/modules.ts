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
    };
};

export type UpdateCommentPayload = {
    threadKey: string;
    commentKey: string;
    body: {
        comment: string;
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
            if (action.payload.thread === null) {
                // エラーハンドリングするべき
                return state
            }
            if (action.payload.comments === null) {
                return {
                    ...state,
                    thread: action.payload.thread,
                }
            }
            return action.payload;
        },
        addComment: (state, action: PayloadAction<CreateCommentResponse>) => {
            state.comments.unshift(action.payload);
            state.thread.commentSum = state.thread.commentSum + 1
            state.thread.updateDate = action.payload.updateDate
        },
        editComment: (state, action: PayloadAction<UpdateCommentResponse>) => {
            const targetComment = state.comments.find((comment) => comment.commentKey === action.payload.commentKey);
            if (targetComment) {
                targetComment.comment = action.payload.comment
                targetComment.updateDate = action.payload.updateDate
            }
        },
        deleteComment: (state, action: PayloadAction<string>) => {
            state.comments = state.comments.filter((comment) => comment.commentKey !== action.payload);
            state.thread.commentSum = state.thread.commentSum - 1
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
        getAllDone: (state) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.success = true;
        },
        getAllFail: (state) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.error = true;
        },
        create: (state, action: PayloadAction<CreateCommentPayload>) => {
            state.createResponse.pending = true;
        },
        createDone: (state) => {
            state.createResponse.pending = false;
            state.createResponse.success = true;
        },
        createFail: (state) => {
            state.createResponse.pending = false;
            state.createResponse.error = true;
        },
        update: (state, action: PayloadAction<UpdateCommentPayload>) => {
            state.updateResponse.pending = true;
        },
        updateDone: (state) => {
            state.updateResponse.pending = false;
            state.updateResponse.success = true;
        },
        updateFail: (state) => {
            state.updateResponse.pending = false;
            state.updateResponse.error = true;
        },
        delete: (state, action: PayloadAction<DeleteCommentPayload>) => {
            state.deleteResponse.pending = true;
        },
        deleteDone: (state) => {
            state.deleteResponse.pending = false;
            state.deleteResponse.success = true;
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
