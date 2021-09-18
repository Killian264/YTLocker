import { Story } from "@storybook/react";
import { ColorBadge, ColorBadgeProps } from "../components/ColorBadge";
import { sRadio, sString } from "./utils/utils";

export default {
	title: "ColorBadge",
	component: ColorBadge,
};

const Mocked: Story<ColorBadgeProps & { message: string }> = ({ ...props }) => {
	return (
		<div>
			<ColorBadge {...props}></ColorBadge>
			<div className="mt-3"></div>
		</div>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
	color: sString("red-1"),
	status: sRadio(["none", "disabled", "selected"]),
};
