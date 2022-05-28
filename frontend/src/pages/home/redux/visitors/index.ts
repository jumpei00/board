import { createSlice } from "@reduxjs/toolkit";
import { Visitors } from "../../../../models/Visitors";

const initialStatevVisitors: Visitors = {
    yesterdayVisitor: 0,
    todayVisitor: 0,
    sumVisitor: 0,
};

export const visitorsSlice = createSlice({
    name: "visitors",
    initialState: initialStatevVisitors,
    reducers: {
        getVisitors: (state) => {
            state = initialStatevVisitors;
        },
        visitedUserCountup: (state) => {
            state.todayVisitor += 1;
            state.sumVisitor += 1;
        },
    },
});

export const { getVisitors, visitedUserCountup } = visitorsSlice.actions;
export const visitorsReducer = visitorsSlice.reducer;
