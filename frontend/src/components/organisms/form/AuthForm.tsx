import React, { ChangeEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Box, Stack, FormControl, FormLabel, Input, Heading } from "@chakra-ui/react";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { userPayload } from "../../../state/user/redux/type";

type AuthFormProps = {
    formName: string;
    buttonName: string;
    OnClick: (user: userPayload) => void;
};

export const AuthForm: React.FC<AuthFormProps> = (props) => {
    const navigate = useNavigate();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleUsernameChange = (event: ChangeEvent<HTMLInputElement>) => {
        setUsername(event.target.value);
    };

    const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value);
    };

    const buttonOnClick = () => {
        const user: userPayload = {
            username,
            password,
        };
        props.OnClick(user);
        navigate("/");
    };

    return (
        <Box w="50%" m="200px auto" bg="gray.300" boxShadow="dark-lg" p="50px">
            <FormControl>
                <Stack spacing={8}>
                    <Heading>{props.formName}</Heading>
                    <Box>
                        <FormLabel>ユーザー名</FormLabel>
                        <Input
                            type="text"
                            variant="flushed"
                            placeholder="username"
                            value={username}
                            onChange={handleUsernameChange}
                        ></Input>
                    </Box>
                    <Box>
                        <FormLabel>パスワード</FormLabel>
                        <Input
                            type="password"
                            variant="flushed"
                            placeholder="password"
                            value={password}
                            onChange={handlePasswordChange}
                        ></Input>
                    </Box>
                    <Box textAlign="right">
                        <PrimaryButton colorScheme="teal" onClick={buttonOnClick}>
                            {props.buttonName}
                        </PrimaryButton>
                    </Box>
                </Stack>
            </FormControl>
        </Box>
    );
};
