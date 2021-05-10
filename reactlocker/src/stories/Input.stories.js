import React from "react";
import { Input } from "../components/Input";
import { toStr } from "./utils/utils";

export default {
	title: "Input",
	component: Input,
	argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({ children, ...props }) => {
	return (
		<Input {...props} placeholder={props.placeholder || "Placeholder"}>
			{children}
		</Input>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	placeholder: toStr(),
	value: toStr(),
	className: toStr(),
};
