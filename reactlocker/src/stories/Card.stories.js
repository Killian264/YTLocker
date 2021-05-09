import React from 'react';
import { Card } from '../components/Card';
import { toEnum, toBoolean} from "./utils/utils"

export default {
  title: 'Card',
  component: Card,
  argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({
  children,
  ...props
}) => {
	return (
		<Card {...props} >
			{children}
		</Card>
	);
};


export const Primary = Mocked.bind({})

Primary.argTypes = {};