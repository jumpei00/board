import React from "react";
import { useDispatch } from "react-redux";
import { AuthForm } from "../../components/organisms/form/AuthForm";
import { signin } from "../../state/user";
import { userPayload } from "../../state/user/redux/type";

export const SignIn: React.FC = () => {
    const dispatch = useDispatch();

    const buttonClickBySignIn = (user: userPayload) => {
        dispatch(signin(user));
    };

    return <AuthForm formName="ログイン" buttonName="ログイン" OnClick={buttonClickBySignIn}></AuthForm>;
};
