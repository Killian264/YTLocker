import { Story } from "@storybook/react";
import { useState } from "react";
import { Modal, ModalProps } from "../components/Modal";
import { sString } from "./utils/utils";

export default {
	title: "Modal",
	component: Modal,
};

const Mocked: Story<ModalProps> = ({ ...props }) => {
	let [isOpen, setIsOpen] = useState(true);

	setTimeout(() => {
		setIsOpen(true);
	}, 5000);

	return (
		<div>
			{isOpen && (
				<Modal
					{...props}
					AcceptClick={() => {
						setIsOpen(false);
					}}
					RejectClick={() => {
						setIsOpen(false);
					}}
				/>
			)}
		</div>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	header: sString("Are you sure?"),
	body: sString(
		"Playlist information will be deleted, but the playlist will remain accessible on youtube."
	),
	confirmMessage: sString("I am sure"),
};
