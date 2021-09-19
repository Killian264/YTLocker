import { Alert, AlertProps } from "../components/Alert";
import { sRadio, sString } from "./utils/utils";
import { Story } from "@storybook/react";

export default {
	title: "Alert",
	component: Alert,
};

const Mocked: Story<AlertProps & { message: string }> = ({ ...props }) => {
	return <Alert {...props}></Alert>;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
	type: sRadio(["success", "failure"]),
	message: sString("This is an alert"),
};
