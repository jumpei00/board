import { PayloadAction } from "@reduxjs/toolkit";
import { call, put, takeEvery } from "redux-saga/effects";
import { getMe, signin, signout, signup } from "../../api/user";
import { userSagaActionsType, userSagaActions, SignUpPayload, SignInPayload } from "./modules";

function* fetchUser() {
    try {
        const { res } = yield call(getMe);
        yield put(userSagaActions.getmeDone(res));
    } catch (e) {
        console.log(e);
        yield put(userSagaActions.getmeFail());
    }
}

function* signupUser(action: PayloadAction<SignUpPayload>) {
    try {
        const { res } = yield call(signup, action.payload);
        yield put(userSagaActions.signupDone(res));
    } catch (e) {
        console.log(e);
        yield put(userSagaActions.signupFail());
    }
}

function* signinUser(action: PayloadAction<SignInPayload>) {
    try {
        const { res } = yield call(signin, action.payload);
        yield put(userSagaActions.signinDone(res));
    } catch (e) {
        console.log(e);
        yield put(userSagaActions.signinFail());
    }
}

function* signoutUser() {
    try {
        yield call(signout);
        yield put(userSagaActions.signoutDone());
    } catch (e) {
        console.log(e);
        yield put(userSagaActions.signoutFail());
    }
}

function* watchUser() {
    yield takeEvery(userSagaActionsType.getme, fetchUser);
    yield takeEvery(userSagaActionsType.signup, signupUser);
    yield takeEvery(userSagaActionsType.signin, signinUser);
    yield takeEvery(userSagaActionsType.signout, signoutUser);
}

export default watchUser;
