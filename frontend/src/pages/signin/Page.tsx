import React from "react";
import { useDispatch } from "react-redux";
import { AuthForm } from "../../components/organisms/form/AuthForm";
import { userSagaActions, SignInPayload } from "../../state/user/modules";

export const SignIn: React.FC = () => {
    const dispatch = useDispatch();

    const buttonClickBySignIn = (payload: SignInPayload) => {
        dispatch(userSagaActions.signin(payload));
    };

    return <AuthForm formName="ログイン" buttonName="ログイン" OnClick={buttonClickBySignIn}></AuthForm>;
};
