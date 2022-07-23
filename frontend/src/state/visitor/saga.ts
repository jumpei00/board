import { call, put, takeEvery } from "redux-saga/effects";
import { visitorCountup, visitorStat } from "../../api/visitor";
import { visitorSagaActionsType, visitorSagaActions, visitorActions } from "./modules";

function* getVisitorStat() {
    try {
        const { data } = yield call(visitorStat);
        yield put(visitorSagaActions.getStatDone());
        yield put(visitorActions.storeVisitor(data));
    } catch (e) {
        yield put(visitorSagaActions.getStatFail());
    }
}

function* countupTodayVisitor() {
    try {
        const { data } = yield call(visitorCountup);
        yield put(visitorSagaActions.coutupDone());
        yield put(visitorActions.storeVisitor(data));
    } catch (e) {
        console.log(e);
        yield put(visitorSagaActions.coutupFail());
    }
}

function* watchVisitor() {
    yield takeEvery(visitorSagaActionsType.getStat, getVisitorStat);
    yield takeEvery(visitorSagaActionsType.coutup, countupTodayVisitor);
}

export default watchVisitor;
