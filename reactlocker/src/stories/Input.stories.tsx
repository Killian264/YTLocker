import { Story } from "@storybook/react";
import { Input, InputProps } from "../components/Input";
import { sString } from "./utils/utils";

export default {
	title: "Input",
	component: Input,
};

const Mocked: Story<InputProps> = ({ children, ...props }) => {
	return <Input {...props} placeholder={"Placeholder"}></Input>;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
	placeholder: sString(),
	value: sString(),
};
