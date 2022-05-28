import React from "react";
import { Alert, AlertIcon, AlertTitle } from "@chakra-ui/react";

type AnyAlertProps = {
    status: "error" | "success" | "warning" | "info";
    title: string;
};

export const AnyAlert: React.FC<AnyAlertProps> = (props) => {
    return (
        <Alert status={props.status}>
            <AlertIcon></AlertIcon>
            <AlertTitle>{props.title}</AlertTitle>
        </Alert>
    );
};
