import { configureStore } from "@reduxjs/toolkit";
import { threadsReducer } from "../pages/home/reducks/threads";
import { visitorsReducer } from "../pages/home/reducks/visitors";

export const store = configureStore({
    reducer: {
        visitors: visitorsReducer,
        threads: threadsReducer,
    },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
