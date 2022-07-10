import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { User } from "../../models/User";
import { userPayload } from "./type";

const initialState: User = {
    username: "",
};

export const userSlice = createSlice({
    name: "user",
    initialState,
    reducers: {
        signup: (state, action: PayloadAction<userPayload>) => {
            state.username = action.payload.username;
        },
        signin: (state, action: PayloadAction<userPayload>) => {
            state.username = action.payload.username;
        },
        signout: (state) => {
            return initialState;
        },
    },
});

export const { signup, signin, signout } = userSlice.actions;
export const userReducer = userSlice.reducer;
