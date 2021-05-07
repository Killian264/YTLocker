import React from 'react';
import { LoginPage } from '../components/LoginPage'

export default {
  title: 'LoginPage',
  component: LoginPage,
  argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({
  children,
  ...props
}) => {
	return (
		<LoginPage></LoginPage>
	);
};


export const Primary = Mocked.bind({})

Primary.argTypes = {};