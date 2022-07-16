import { combineReducers, createSlice, PayloadAction } from "@reduxjs/toolkit";
import { User, UserResponse } from "../../models/user";

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
        },
        getmeDone: (state, action: PayloadAction<User>) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.success = true;
            userSlice.actions.storeUser(action.payload);
        },
        getmeFail: (state) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.error = true;
        },
        signup: (state, action: PayloadAction<SignUpPayload>) => {
            state.signupResponse.pending = true;
        },
        signupDone: (state, action: PayloadAction<User>) => {
            state.signupResponse.pending = false;
            state.signupResponse.success = true;
            userSlice.actions.storeUser(action.payload);
        },
        signupFail: (state) => {
            state.signupResponse.pending = false;
            state.signupResponse.error = true;
        },
        signin: (state, action: PayloadAction<SignInPayload>) => {
            state.signinResponse.pending = true;
        },
        signinDone: (state, action: PayloadAction<User>) => {
            state.signinResponse.pending = false;
            state.signinResponse.success = true;
            userSlice.actions.storeUser(action.payload);
        },
        signinFail: (state) => {
            state.signinResponse.pending = false;
            state.signinResponse.error = true;
        },
        signout: (state) => {
            state.signoutResponse.pending = true;
        },
        signoutDone: (state) => {
            state.signoutResponse.pending = false;
            state.signoutResponse.success = true;
            userSlice.actions.clearUser();
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
    user: userSlice.reducer,
    userSaga: userSagaSlice.reducer,
});
