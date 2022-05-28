import React, { ReactNode } from "react";
import { Button } from "@chakra-ui/react";

type PrimaryButtonProps = {
    children: ReactNode;
    colorScheme: string;
    onClick: () => void;
    isDisabled?: boolean;
};

export const PrimaryButton: React.FC<PrimaryButtonProps> = (props) => {
    return (
        <Button
            colorScheme={props.colorScheme}
            size="md"
            _hover={{ opacity: 0.8 }}
            onClick={props.onClick}
            isDisabled={props.isDisabled}
        >
            {props.children}
        </Button>
    );
};
