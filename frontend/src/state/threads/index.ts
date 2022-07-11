import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Thread, Threads } from "../../models/thread";
import { createThreadPayload, deleteThreadPayload, editThreadPayload } from "./type";

const initialState: Threads = {
    threads: [],
};

export const threadsSlice = createSlice({
    name: "threads",
    initialState,
    reducers: {
        getAllThread: (state) => {
            state = initialState;
        },
        createThread: (state, action: PayloadAction<createThreadPayload>) => {
            const now = new Date();
            const thread: Thread = {
                threadKey: Math.random().toString(32).substring(2),
                title: action.payload.title,
                contributer: action.payload.contributer,
                postDate: now.toLocaleString(),
                updateDate: now.toLocaleString(),
                views: 0,
                sumComment: 0,
            };
            state.threads.push(thread);
        },
        editThreadTitle: (state, action: PayloadAction<editThreadPayload>) => {
            const now = new Date();
            state.threads.forEach((thread) => {
                if (thread.threadKey === action.payload.threadKey) {
                    thread.title = action.payload.title;
                    thread.updateDate = now.toLocaleString();
                }
            });
        },
        deleteThread: (state, action: PayloadAction<deleteThreadPayload>) => {
            state.threads = state.threads.filter((thread) => thread.threadKey !== action.payload.threadKey);
        },
    },
});

export const { getAllThread, createThread, editThreadTitle, deleteThread } = threadsSlice.actions;
export const threadsReducer = threadsSlice.reducer;
