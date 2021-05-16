import { Story } from "@storybook/react";
import { useState } from "react";
import { TextBox, TextBoxProps } from "../components/TextBox";
import { sString } from "./utils/utils";

export default {
	title: "TextBox",
	component: TextBox,
};

const Mocked: Story<TextBoxProps> = ({ children, ...props }) => {
	let [value, setValue] = useState("");

	return (
		<TextBox
			{...props}
			value={value}
			onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => setValue(e.target.value)}
		></TextBox>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
	placeholder: sString(),
	value: sString(),
};
