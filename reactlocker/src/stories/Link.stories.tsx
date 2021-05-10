import { Story } from "@storybook/react";
import { LinkProps } from "react-router-dom";
import { Link } from "../components/Link";
import { sString } from "./utils/utils";

export default {
	title: "Link",
	component: Link,
};

const Mocked: Story<LinkProps & { message: string }> = ({
	children,
	message,
	...props
}) => {
	return <Link {...props}>{message || "Simple link to change state"}</Link>;
};

export const Primary = Mocked.bind({});
Primary.argTypes = {
	className: sString(),
};

const Mocked2: Story<LinkProps & { message: string }> = ({
	children,
	message,
	...props
}) => {
	return (
		<Link {...props} target="_blank">
			{message || "Link to external website"}
		</Link>
	);
};

export const External = Mocked2.bind({});
External.argTypes = {
	className: sString(),
	href: sString(),
	message: sString(),
};
