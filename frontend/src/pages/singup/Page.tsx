import React from "react";
import { useDispatch } from "react-redux";
import { AuthForm } from "../../components/organisms/form/AuthForm";
import { signup } from "../../state/user/redux";
import { userPayload } from "../../state/user/redux/type";

export const SingUp: React.FC = () => {
    const dispatch = useDispatch();

    const buttonClickBySignUp = (user: userPayload) => {
        dispatch(signup(user));
    };

    return <AuthForm formName="ユーザー登録" buttonName="登録" OnClick={buttonClickBySignUp}></AuthForm>;
};
