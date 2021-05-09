import React from 'react';
import { LogoBar } from '../components/LogoBar';
import { toStr}  from "./utils/utils"

export default {
  title: 'LogoBar',
  component: LogoBar,
  argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({
  children,
  ...props
}) => {
	return (
		<LogoBar {...props} ></LogoBar>
	);
};


export const Primary = Mocked.bind({})

Primary.argTypes = {
	className: toStr()
};