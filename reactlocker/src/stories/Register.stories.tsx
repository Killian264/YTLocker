import { Story } from "@storybook/react";
import { Register, RegisterProps } from "../components/Register";

export default {
	title: "Register",
	component: Register,
};

const Mocked: Story<RegisterProps> = ({ ...props }) => {
	return <Register {...props}></Register>;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
