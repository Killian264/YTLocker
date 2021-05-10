import React from "react";
import { Button } from "../components/Button";
import { toEnum, toBoolean } from "./utils/utils";

export default {
	title: "Button",
	component: Button,
	argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({ children, ...props }) => {
	return <Button {...props}>{children || "New Playlist"}</Button>;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	size: toEnum(["small", "medium", "large"]),
	color: toEnum(["primary", "secondary"]),
	disabled: toBoolean(),
	loading: toBoolean(),
};
