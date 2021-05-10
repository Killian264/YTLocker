import { Story } from "@storybook/react";
import { Card, CardProps } from "../components/Card";
import { sString } from "./utils/utils";

export default {
	title: "Card",
	component: Card,
};

const Mocked: Story<CardProps> = ({ ...props }) => {
	return (
		<Card {...props} className="flex-grow">
			<div>ELEMENT 1</div>
			<div>ELEMENT 2</div>
			<div>ELEMENT 3</div>
		</Card>
	);
};

export const UserInformation = Mocked.bind({});

UserInformation.argTypes = {
	className: sString(),
};
