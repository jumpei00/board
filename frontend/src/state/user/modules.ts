import { combineReducers, createSlice, PayloadAction } from "@reduxjs/toolkit";
import { User, UserResponse } from "../../models/User";

// ---- Payload ---- //
export type SignUpPayload = {
    username: string;
    password: string;
};

export type SignInPayload = {
    username: string;
    password: string;
};

// ---- state ---- //
const initialState: User = {
    id: "",
    username: "",
};

const initialSagaResponse: UserResponse = {
    fetchResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            id: "",
            username: "",
        },
    },
    signupResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            id: "",
            username: "",
        },
    },
    signinResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            id: "",
            username: "",
        },
    },
    signoutResponse: {
        pending: false,
        success: false,
        error: false,
    },
};

// ---- reducer ---- //
export const userSlice = createSlice({
    name: "user",
    initialState,
    reducers: {
        storeUser: (state, action: PayloadAction<User>) => {
            return {
                ...state,
                ...action.payload,
            };
        },
        clearUser: () => initialState,
    },
});

// ---- saga reducer ---- //
const sagaSliceName = "userSaga";

export const userSagaActionsType = {
    getme: `${sagaSliceName}/getme`,
    signup: `${sagaSliceName}/signup`,
    signin: `${sagaSliceName}/signin`,
    signout: `${sagaSliceName}/signout`,
};

export const userSagaSlice = createSlice({
    name: sagaSliceName,
    initialState: initialSagaResponse,
    reducers: {
        getme: (state) => {
            state.fetchResponse.pending = true;
            state.fetchResponse.success = false;
            state.fetchResponse.error = false;
        },
        getmeDone: (state) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.success = true;
        },
        getmeFail: (state) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.error = true;
        },
        signup: (state, action: PayloadAction<SignUpPayload>) => {
            state.signupResponse.pending = true;
            state.signupResponse.success = false;
            state.signupResponse.error = false;
        },
        signupDone: (state) => {
            state.signupResponse.pending = false;
            state.signupResponse.success = true;
        },
        signupFail: (state) => {
            state.signupResponse.pending = false;
            state.signupResponse.error = true;
        },
        signin: (state, action: PayloadAction<SignInPayload>) => {
            state.signinResponse.pending = true;
            state.signinResponse.success = false;
            state.signinResponse.error = false;
        },
        signinDone: (state) => {
            state.signinResponse.pending = false;
            state.signinResponse.success = true;
        },
        signinFail: (state) => {
            state.signinResponse.pending = false;
            state.signinResponse.error = true;
        },
        signout: (state) => {
            state.signoutResponse.pending = true;
            state.signoutResponse.success = false;
            state.signoutResponse.error = false;
        },
        signoutDone: (state) => {
            state.signoutResponse.pending = false;
            state.signoutResponse.success = true;
            state.signinResponse.success = false;
            state.signinResponse.error = false;
            state.signupResponse.success = false;
            state.signupResponse.error = false;
        },
        signoutFail: (state) => {
            state.signoutResponse.pending = false;
            state.signoutResponse.error = true;
        },
    },
});

export const userActions = userSlice.actions;
export const userSagaActions = userSagaSlice.actions;
export const userReducer = combineReducers({
    userState: userSlice.reducer,
    userSaga: userSagaSlice.reducer,
});
