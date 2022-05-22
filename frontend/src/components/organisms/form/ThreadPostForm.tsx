import React, { useState, ChangeEvent } from "react";
import { useDispatch } from "react-redux";
import { Text, Input, Box } from "@chakra-ui/react";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { createThreadPayload } from "../../../pages/home/redux/threads/type";
import { createThread } from "../../../pages/home/redux/threads";

type ThreadPostFormProps = {
    loginUsername: string;
}

export const ThreadPostForm: React.FC<ThreadPostFormProps> = (props) => {
    const dispatch = useDispatch();
    const [value, setValue] = useState("");
    const [buttonDisabled, setButtonDisabled] = useState(true);

    const createThreadPayload: createThreadPayload = {
        title: value,
        contributer: props.loginUsername,
    };

    const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
        const currentValue = event.target.value;
        currentValue === "" ? setButtonDisabled(true) : setButtonDisabled(false);
        setValue(currentValue);
    };

    const threadPostByButtonClick = () => {
        dispatch(createThread(createThreadPayload));
        setValue("");
    };

    return (
        <Box w="70%" m="auto">
            <Text mb="10px">スレッドタイトル</Text>
            <Input
                type="text"
                size="lg"
                mb="10px"
                variant="flushed"
                placeholder="話題を投稿しましょう"
                value={value}
                onChange={handleChange}
            ></Input>
            <PrimaryButton colorScheme="teal" onClick={threadPostByButtonClick} isDisabled={buttonDisabled}>
                スレッドを投稿
            </PrimaryButton>
        </Box>
    );
};
