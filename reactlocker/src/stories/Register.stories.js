import React from 'react';
import { Register } from '../components/Register'

export default {
  title: 'Register',
  component: Register,
  argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({
  children,
  ...props
}) => {
	return (
		<Register
			{...props} 
			onSubmit={(user) => console.log("Submitted", user)}
			onClickLogin={() => console.log("Swapped to login page.")}
		>
		</Register>
	);
};


export const Primary = Mocked.bind({})

Primary.argTypes = {};