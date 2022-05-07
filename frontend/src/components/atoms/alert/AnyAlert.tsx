import React from "react";
import { Alert, AlertIcon, AlertTitle } from "@chakra-ui/react";
import { title } from "process";

type AnyAlertProps = {
    status: "error" | "success" | "warning" | "info";
    title: string;
};

export const AnyAlert: React.FC<AnyAlertProps> = (props) => {
    const { status } = props;

    return (
        <Alert status={status}>
            <AlertIcon></AlertIcon>
            <AlertTitle>{title}</AlertTitle>
        </Alert>
    );
};
