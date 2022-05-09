import React from "react";
import { Modal, ModalOverlay, ModalContent, ModalHeader, ModalCloseButton, ModalBody } from "@chakra-ui/react";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { AnyAlert } from "../../atoms/alert/AnyAlert";

type DeleteModalProps = {
    isOpen: boolean;
    onClose: () => void;
};

export const DeleteModal: React.FC<DeleteModalProps> = (props) => {
    const { isOpen, onClose } = props;

    return (
        <Modal isOpen={isOpen} onClose={onClose}>
            <ModalOverlay></ModalOverlay>
            <ModalContent>
                <ModalHeader>
                    <AnyAlert status="error" title="本当に削除しますか？"></AnyAlert>
                </ModalHeader>
                <ModalCloseButton></ModalCloseButton>
                <ModalBody>
                    <PrimaryButton colorScheme="red">削除</PrimaryButton>
                </ModalBody>
            </ModalContent>
        </Modal>
    );
};
