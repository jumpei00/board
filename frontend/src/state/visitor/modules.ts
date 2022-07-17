import { combineReducers, createSlice, PayloadAction } from "@reduxjs/toolkit";
import { VisitorResponse, Visitor } from "../../models/Visitors";

// ---- state ---- //
const initialState: Visitor = {
    yesterday: 0,
    today: 0,
    sum: 0,
};

const initialSagaResponse: VisitorResponse = {
    fetchResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            yesterday: 0,
            today: 0,
            sum: 0,
        },
    },
    countupResponse: {
        pending: false,
        success: false,
        error: false,
        body: {
            yesterday: 0,
            today: 0,
            sum: 0,
        },
    },
};

// ---- resucer ----- //
export const visitorSlice = createSlice({
    name: "visitors",
    initialState,
    reducers: {
        storeVisitor: (state, action: PayloadAction<Visitor>) => {
            return {
                ...state,
                ...action.payload,
            };
        },
        clearVisitor: () => initialState,
    },
});

// ---- saga reducer ---- //
const sagaSliceName = "visitorSaga";

export const visitorSagaActionsType = {
    getStat: `${sagaSliceName}/getStat`,
    coutup: `${sagaSliceName}/coutup`,
};

export const visitorSagaSlice = createSlice({
    name: sagaSliceName,
    initialState: initialSagaResponse,
    reducers: {
        getStat: (state) => {
            state.fetchResponse.pending = true;
        },
        getStatDone: (state, action: PayloadAction<Visitor>) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.success = true;
            visitorSlice.actions.storeVisitor(action.payload);
        },
        getStatFail: (state) => {
            state.fetchResponse.pending = false;
            state.fetchResponse.error = true;
        },
        coutup: (state) => {
            state.countupResponse.pending = true;
        },
        coutupDone: (state, action: PayloadAction<Visitor>) => {
            state.countupResponse.pending = false;
            state.countupResponse.success = true;
            visitorSlice.actions.storeVisitor(action.payload);
        },
        coutupFail: (state) => {
            state.countupResponse.pending = false;
            state.countupResponse.error = true;
        },
    },
});

export const visitorActions = visitorSlice.actions;
export const visitorSagaActions = visitorSagaSlice.actions;
export const visitorsReducer = combineReducers({
    visitorState: visitorSlice.reducer,
    visitorSaga: visitorSagaSlice.reducer,
});
