import { Story } from "@storybook/react";
import { useState } from "react";
import { Input, InputProps } from "../components/Input";
import { sString } from "./utils/utils";

export default {
	title: "Input",
	component: Input,
};

const Mocked: Story<InputProps> = ({ children, ...props }) => {
	let [value, setValue] = useState("");

	return (
		<Input
			{...props}
			value={value}
			onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
				setValue(e.target.value)
			}
			placeholder={"Placeholder"}
		></Input>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
	placeholder: sString(),
	value: sString(),
};
