import React, { ChangeEvent, useState } from "react";
import { Modal, ModalOverlay, ModalContent, ModalHeader, ModalBody, ModalFooter, Input } from "@chakra-ui/react";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import { AnyAlert } from "../../atoms/alert/AnyAlert";

type GeneralModalProps = {
    content: string;
    isEdit: boolean;
    isOpen: boolean;
    onClose: () => void;
    updateOnClick: (title: string) => void;
    deleteOnClick: () => void;
};

export const GeneralModal: React.FC<GeneralModalProps> = (props) => {
    const [value, setValue] = useState(props.content);

    const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
        setValue(event.target.value);
    };

    const modalClose = () => {
        setValue(props.content);
        props.onClose();
    };

    return (
        <Modal isOpen={props.isOpen} onClose={modalClose}>
            <ModalOverlay></ModalOverlay>
            <ModalContent>
                {props.isEdit ? (
                    <>
                        <ModalHeader>タイトル：{value}</ModalHeader>
                        <ModalBody>
                            <Input value={value} onChange={handleChange}></Input>
                        </ModalBody>
                        <ModalFooter>
                            <PrimaryButton colorScheme="teal" onClick={() => props.updateOnClick(value)}>
                                更新
                            </PrimaryButton>
                            <PrimaryButton colorScheme="blue" onClick={modalClose}>
                                閉じる
                            </PrimaryButton>
                        </ModalFooter>
                    </>
                ) : (
                    <>
                        <ModalHeader>
                            <AnyAlert status="error" title="本当に削除しますか？"></AnyAlert>
                        </ModalHeader>
                        <ModalFooter>
                            <PrimaryButton colorScheme="red" onClick={() => props.deleteOnClick()}>
                                削除
                            </PrimaryButton>
                            <PrimaryButton colorScheme="blue" onClick={modalClose}>
                                閉じる
                            </PrimaryButton>
                        </ModalFooter>
                    </>
                )}
            </ModalContent>
        </Modal>
    );
};
