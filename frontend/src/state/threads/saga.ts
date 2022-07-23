import { PayloadAction } from "@reduxjs/toolkit";
import { call, put, takeEvery } from "redux-saga/effects";
import { createThreadAPI, deleteThreadAPI, getAllThreadsAPI, updateThreadAPI } from "../../api/thread";
import { commentActions } from "../comments/modules";
import {
    CreateThreadPayload,
    UpdateThreadPayload,
    threadSagaActionsType,
    threadSagaActions,
    threadActions,
} from "./modules";

function* fetchThreads() {
    try {
        const { data } = yield call(getAllThreadsAPI);
        yield put(threadSagaActions.getAllDone());
        yield put(threadActions.storeThreads(data));
    } catch (e) {
        yield put(threadSagaActions.getAllFail());
    }
}

function* createThread(action: PayloadAction<CreateThreadPayload>) {
    try {
        const { data } = yield call(createThreadAPI, action.payload);
        yield put(threadSagaActions.createDone());
        yield put(threadActions.addThread(data));
    } catch (e) {
        yield put(threadSagaActions.createFail());
    }
}

function* updateThread(action: PayloadAction<UpdateThreadPayload>) {
    try {
        const { data } = yield call(updateThreadAPI, action.payload);
        yield put(threadSagaActions.updateDone());
        yield put(threadActions.editThread(data));
    } catch (e) {
        yield put(threadSagaActions.updateFail());
    }
}

function* deleteThread(action: PayloadAction<string>) {
    try {
        yield call(deleteThreadAPI, action.payload);
        yield put(threadSagaActions.deleteDone());
        yield put(threadActions.deleteThread(action.payload));
        yield put(commentActions.clearComment());
    } catch (e) {
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
