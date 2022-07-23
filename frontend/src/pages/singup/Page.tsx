import React from "react";
import { useDispatch, useSelector } from "react-redux";
import { AuthForm } from "../../components/organisms/form/AuthForm";
import { userSagaActions, SignUpPayload } from "../../state/user/modules";
import { RootState } from "../../store/store";

export const SingUp: React.FC = () => {
    const userSagaState = useSelector((state: RootState) => state.user.userSaga);
    const dispatch = useDispatch();

    const buttonClickBySignUp = (payload: SignUpPayload) => {
        dispatch(userSagaActions.signup(payload));
    };

    return (
        <AuthForm
            formName="ユーザー登録"
            buttonName="登録"
            OnClick={buttonClickBySignUp}
            pending={userSagaState.signupResponse.pending}
            success={userSagaState.signupResponse.success}
        ></AuthForm>
    );
};
