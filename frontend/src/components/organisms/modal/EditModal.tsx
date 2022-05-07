import React from "react";
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalCloseButton,
    ModalBody,
    ModalFooter,
    Input
} from "@chakra-ui/react";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";

type EditModalProps = {
    content: string;
    isOpen: boolean;
    onClose: () => void;
}

export const EditModal: React.FC<EditModalProps> = (props) => {
    const { content, isOpen, onClose } = props

    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <ModalOverlay></ModalOverlay>
            <ModalContent>
                <ModalHeader>編集</ModalHeader>
                <ModalCloseButton></ModalCloseButton>
                <ModalBody>
                    <Input value={content}></Input>
                </ModalBody>
                <ModalFooter>
                    <PrimaryButton colorScheme="teal">更新</PrimaryButton>
                    <PrimaryButton colorScheme="blue">閉じる</PrimaryButton>
                </ModalFooter>
            </ModalContent>
        </Modal>
    )
}
