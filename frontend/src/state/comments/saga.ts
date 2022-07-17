import { PayloadAction } from "@reduxjs/toolkit";
import { call, put, takeEvery } from "redux-saga/effects";
import { createCommentAPI, deleteCommentAPI, getAllCommentsAPI, updateCommentAPI } from "../../api/comment";
import {
    commentSagaActionsType,
    commentSagaActions,
    CreateCommentPayload,
    UpdateCommentPayload,
    DeleteCommentPayload,
} from "./modules";

function* fetchComments(action: PayloadAction<string>) {
    try {
        const { res } = yield call(getAllCommentsAPI, action.payload);
        yield put(commentSagaActions.getAllDone(res));
    } catch (e) {
        console.log(e);
        yield put(commentSagaActions.getAllFail());
    }
}

function* createComment(action: PayloadAction<CreateCommentPayload>) {
    try {
        const { res } = yield call(createCommentAPI, action.payload);
        yield put(commentSagaActions.createDone(res));
    } catch (e) {
        console.log(e);
        yield put(commentSagaActions.createFail());
    }
}

function* updateComment(action: PayloadAction<UpdateCommentPayload>) {
    try {
        const { res } = yield call(updateCommentAPI, action.payload);
        yield put(commentSagaActions.updateDone(res));
    } catch (e) {
        console.log(e);
        yield put(commentSagaActions.updateFail());
    }
}

function* deleteComment(action: PayloadAction<DeleteCommentPayload>) {
    try {
        yield call(deleteCommentAPI, action.payload);
        yield put(commentSagaActions.deleteDone(action.payload.commentKey));
    } catch (e) {
        console.log(e);
        yield put(commentSagaActions.deleteFail());
    }
}

function* watchComment() {
    yield takeEvery(commentSagaActionsType.getAll, fetchComments);
    yield takeEvery(commentSagaActionsType.create, createComment);
    yield takeEvery(commentSagaActionsType.update, updateComment);
    yield takeEvery(commentSagaActionsType.delete, deleteComment);
}

export default watchComment;
