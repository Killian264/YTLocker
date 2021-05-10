import { Story } from "@storybook/react";
import { PlusButton, PlusButtonProps } from "../components/PlusButton";
import { sRadio, sBoolean, sString } from "./utils/utils";

export default {
	title: "PlusButton",
	component: PlusButton,
};

const Mocked: Story<PlusButtonProps> = ({ ...props }) => {
	return <PlusButton {...props} />;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
	color: sRadio(["primary", "secondary"]),
	disabled: sBoolean(),
};
