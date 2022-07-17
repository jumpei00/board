import { fork } from "redux-saga/effects";
import watchComment from "../../state/comments/saga";
import watchThread from "../../state/threads/saga";
import watchUser from "../../state/user/saga";
import watchVisitor from "../../state/visitor/saga";

function* rootSaga() {
    yield fork(watchUser)
    yield fork(watchVisitor)
    yield fork(watchThread)
    yield fork(watchComment)
}

export default rootSaga