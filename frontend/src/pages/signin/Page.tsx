import React from "react";
import { useDispatch, useSelector } from "react-redux";
import { AuthForm } from "../../components/organisms/form/AuthForm";
import { userSagaActions, SignInPayload } from "../../state/user/modules";
import { RootState } from "../../store/store";

export const SignIn: React.FC = () => {
    const userSagaState = useSelector((state: RootState) => state.user.userSaga);
    const dispatch = useDispatch();

    const buttonClickBySignIn = (payload: SignInPayload) => {
        dispatch(userSagaActions.signin(payload));
    };

    return (
        <AuthForm
            formName="ログイン"
            buttonName="ログイン"
            OnClick={buttonClickBySignIn}
            pending={userSagaState.signinResponse.pending}
            success={userSagaState.signinResponse.success}
        ></AuthForm>
    );
};
