import { call, put, takeEvery } from "redux-saga/effects";
import { visitorCountup, visitorStat } from "../../api/visitor";
import { visitorSagaActionsType, visitorSagaActions } from "./modules";

function* getVisitorStat() {
    try {
        const { res } = yield call(visitorStat);
        yield put(visitorSagaActions.getStatDone(res));
    } catch (e) {
        console.log(e);
        yield put(visitorSagaActions.getStatFail());
    }
}

function* countupTodayVisitor() {
    try {
        const { res } = yield call(visitorCountup);
        yield put(visitorSagaActions.coutupDone(res));
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
