import React from 'react';
import { Login } from '../components/Login'
import { toEnum, toBoolean} from "./utils/utils"

export default {
  title: 'Login',
  component: Login,
  argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({
  children,
  ...props
}) => {
	return (
		<Login
			{...props} 
			onSubmit={(user) => console.log("Submitted", user)}
			onClickRegister={() => console.log("Swapped to register page.")}
		>
		</Login>
	);
};


export const Primary = Mocked.bind({})

Primary.argTypes = {};