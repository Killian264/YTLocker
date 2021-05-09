import React from 'react';
import { Alert } from '../components/Alert';
import { toEnum, toStr} from "./utils/utils"

export default {
  title: 'Alert',
  component: Alert,
  argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({
  children,
  ...props
}) => {
	return (
		<Alert {...props} >
			{children || "Successfully created user account."}
		</Alert>
	);
};


export const Primary = Mocked.bind({})

Primary.argTypes = {
	type: toEnum(["success", "failure"]),
	className: toStr(),
};