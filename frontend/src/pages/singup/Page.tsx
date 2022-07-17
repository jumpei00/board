import React from "react";
import { useDispatch } from "react-redux";
import { AuthForm } from "../../components/organisms/form/AuthForm";
import { userSagaActions, SignUpPayload } from "../../state/user/modules";

export const SingUp: React.FC = () => {
    const dispatch = useDispatch();

    const buttonClickBySignUp = (payload: SignUpPayload) => {
        dispatch(userSagaActions.signup(payload));
    };

    return <AuthForm formName="ユーザー登録" buttonName="登録" OnClick={buttonClickBySignUp}></AuthForm>;
};
