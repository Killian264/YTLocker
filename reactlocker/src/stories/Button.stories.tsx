import { Story } from "@storybook/react";
import { Button, ButtonProps } from "../components/Button";
import { sRadio, sBoolean, sString } from "./utils/utils";

export default {
	title: "Button",
	component: Button,
};

const Mocked: Story<ButtonProps & { message: string }> = ({ ...props }) => {
	return <Button {...props}>{"New Playlist"}</Button>;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	size: sRadio(["small", "medium", "large"]),
	color: sRadio(["primary"]),
	disabled: sBoolean(),
	loading: sBoolean(),
	message: sString(),
};
