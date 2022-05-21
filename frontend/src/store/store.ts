import { configureStore } from "@reduxjs/toolkit";
import { threadsReducer } from "../pages/home/redux/threads";
import { visitorsReducer } from "../pages/home/redux/visitors";
import { commentsReducer } from "../pages/threadContent/redux/comments";
import { threadReducer } from "../pages/threadContent/redux/thread";

export const store = configureStore({
    reducer: {
        visitors: visitorsReducer,
        threads: threadsReducer,
        thread: threadReducer,
        comments: commentsReducer
    },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
