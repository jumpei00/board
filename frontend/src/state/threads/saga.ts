import { PayloadAction } from "@reduxjs/toolkit";
import { call, put, takeEvery } from "redux-saga/effects";
import { createThreadAPI, deleteThreadAPI, getAllThreadsAPI, updateThreadAPI } from "../../api/thread";
import { CreateThreadPayload, UpdateThreadPayload, threadSagaActionsType, threadSagaActions } from "./modules";

function* fetchThreads() {
    try {
        const { res } = yield call(getAllThreadsAPI);
        yield put(threadSagaActions.getAllDone(res));
    } catch (e) {
        console.log(e);
        yield put(threadSagaActions.getAllFail());
    }
}

function* createThread(action: PayloadAction<CreateThreadPayload>) {
    try {
        const { res } = yield call(createThreadAPI, action.payload);
        yield put(threadSagaActions.createDone(res));
    } catch (e) {
        console.log(e);
        yield put(threadSagaActions.createFail());
    }
}

function* updateThread(action: PayloadAction<UpdateThreadPayload>) {
    try {
        const { res } = yield call(updateThreadAPI, action.payload);
        yield put(threadSagaActions.updateDone(res));
    } catch (e) {
        console.log(e);
        yield put(threadSagaActions.updateFail());
    }
}

function* deleteThread(action: PayloadAction<string>) {
    try {
        yield call(deleteThreadAPI, action.payload);
        yield put(threadSagaActions.deleteDone(action.payload));
    } catch (e) {
        console.log(e);
        yield put(threadSagaActions.deleteFail());
    }
}

function* watchThread() {
    yield takeEvery(threadSagaActionsType.getAll, fetchThreads);
    yield takeEvery(threadSagaActionsType.create, createThread);
    yield takeEvery(threadSagaActionsType.update, updateThread);
    yield takeEvery(threadSagaActionsType.delete, deleteThread);
}

export default watchThread;
