import React from 'react';
import { Link } from '../components/Link';
import {toStr} from "./utils/utils"

export default {
  title: 'Link',
  component: Link,
  argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({
  children,
  ...props
}) => {
	return (
		<Link {...props}>
			Simple link to change state
		</Link>
	);
};

export const Primary = Mocked.bind({})
Primary.argTypes = {
	className: toStr(),
};

const Mocked2 = ({
	children,
	...props
}) => {
	return (
		<Link {...props} target="_blank">
			Link to external website
		</Link>
	);
};

export const External = Mocked2.bind({})
External.argTypes = {
	className: toStr(),
	href: toStr(),
};