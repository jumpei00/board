import { combineReducers, createSlice, PayloadAction } from "@reduxjs/toolkit";
import {
    Threads,
    ThreadResponse,
    FetchThreadResponse,
    CreateThreadResponse,
    UpdateThreadResponse,
} from "../../models/Thread";

// ---- Payload ---- //
export type CreateThreadPayload = {
    title: string;
};

export type UpdateThreadPayload = {
    threadKey: string;
    body: {
        title: string;
    };
};

// ---- state ---- //
const initialState: Threads = {
    threads: [],
};

const initialSagaResponse: ThreadResponse = {
    fetchResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            threads: [],
        },
    },
    createResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            threadKey: "",
            title: "",
            contributor: "",
            views: 0,
            commentSum: 0,
            createDate: "",
            updateDate: "",
        },
    },
    updateResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            threadKey: "",
            title: "",
            contributor: "",
            views: 0,
            commentSum: 0,
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

// ---- reducer ---- //
export const threadSlice = createSlice({
    name: "thread",
    initialState,
    reducers: {
        storeThreads: (state, action: PayloadAction<FetchThreadResponse>) => {
            if (action.payload.threads === null) {
                return state
            }
            return action.payload;
        },
        addThread: (state, action: PayloadAction<CreateThreadResponse>) => {
            state.threads.unshift(action.payload);
        },
        editThread: (state, action: PayloadAction<UpdateThreadResponse>) => {
            const targetThread = state.threads.find((thread) => thread.threadKey === action.payload.threadKey);
            if (targetThread) {
                targetThread.title = action.payload.title
                targetThread.updateDate = action.payload.updateDate
            }
        },
        deleteThread: (state, action: PayloadAction<string>) => {
            state.threads = state.threads.filter((thread) => thread.threadKey !== action.payload);
        },
        clearThread: () => initialState,
    },
});

// ---- saga ---- //
const sagaSliceName = "threadSaga";

export const threadSagaActionsType = {
    getAll: `${sagaSliceName}/getAll`,
    create: `${sagaSliceName}/create`,
    update: `${sagaSliceName}/update`,
    delete: `${sagaSliceName}/delete`,
};

export const threadSagaSlice = createSlice({
    name: sagaSliceName,
    initialState: initialSagaResponse,
    reducers: {
        getAll: (state) => {
            state.updateResponse.pending = true;
        },
        getAllDone: (state) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.success = true;
        },
        getAllFail: (state) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.error = true;
        },
        create: (state, action: PayloadAction<CreateThreadPayload>) => {
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
        update: (state, action: PayloadAction<UpdateThreadPayload>) => {
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
        delete: (state, action: PayloadAction<string>) => {
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

export const threadActions = threadSlice.actions;
export const threadSagaActions = threadSagaSlice.actions;
export const threadReducer = combineReducers({
    threadState: threadSlice.reducer,
    threadSaga: threadSagaSlice.reducer,
});
