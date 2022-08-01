import { PayloadAction } from "@reduxjs/toolkit";
import { call, put, takeEvery } from "redux-saga/effects";
import { getMe, signin, signout, signup } from "../../api/user";
import { userSagaActionsType, userSagaActions, userActions, SignUpPayload, SignInPayload } from "./modules";

export function* fetchUser() {
    try {
        const { data } = yield call(getMe);
        yield put(userSagaActions.getmeDone());
        yield put(userActions.storeUser(data));
    } catch (e) {
        yield put(userSagaActions.getmeFail());
    }
}

function* signupUser(action: PayloadAction<SignUpPayload>) {
    try {
        const { data } = yield call(signup, action.payload);
        yield put(userSagaActions.signupDone());
        yield put(userActions.storeUser(data));
    } catch (e) {
        yield put(userSagaActions.signupFail());
    }
}

function* signinUser(action: PayloadAction<SignInPayload>) {
    try {
        const { data } = yield call(signin, action.payload);
        yield put(userSagaActions.signinDone());
        yield put(userActions.storeUser(data));
    } catch (e) {
        yield put(userSagaActions.signinFail());
    }
}

function* signoutUser() {
    try {
        yield call(signout);
        yield put(userSagaActions.signoutDone());
        yield put(userActions.clearUser());
    } catch (e) {
        yield put(userSagaActions.signoutFail());
    }
}

export function* watchUser() {
    yield takeEvery(userSagaActionsType.getme, fetchUser);
    yield takeEvery(userSagaActionsType.signup, signupUser);
    yield takeEvery(userSagaActionsType.signin, signinUser);
    yield takeEvery(userSagaActionsType.signout, signoutUser);
}
