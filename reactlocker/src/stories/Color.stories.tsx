import { Story } from "@storybook/react";
import { ColorBadge, ColorBadgeProps } from "../components/ColorBadge";
import { sRadio, sString } from "./utils/utils";

export default {
	title: "ColorBadge",
	component: ColorBadge,
};

const Mocked: Story<ColorBadgeProps & { message: string }> = ({ ...props }) => {
	return <ColorBadge {...props}></ColorBadge>;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
	color: sString(),
	message: sString(),
};
