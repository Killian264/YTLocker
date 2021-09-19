import { Story } from "@storybook/react";
import { Badge, BadgeProps } from "../components/Badge";
import { sRadio, sString } from "./utils/utils";

export default {
	title: "Badge",
	component: Badge,
};

const Mocked: Story<BadgeProps & { message: string }> = ({ ...props }) => {
	return <Badge {...props}>{props.message}</Badge>;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
	color: sRadio(["primary", "secondary"]),
	message: sString("PRO"),
};
