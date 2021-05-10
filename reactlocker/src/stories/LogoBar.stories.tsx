import { Story } from "@storybook/react";
import { LogoBar, LogoBarProps } from "../components/LogoBar";
import { sString } from "./utils/utils";

export default {
	title: "LogoBar",
	component: LogoBar,
};

const Mocked: Story<LogoBarProps> = ({ ...props }) => {
	return <LogoBar {...props}></LogoBar>;
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
};
