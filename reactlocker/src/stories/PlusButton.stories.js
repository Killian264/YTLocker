import React from 'react';
import { PlusButton } from '../components/PlusButton';
import { toEnum, toBoolean} from "./utils/utils"

export default {
  title: 'PlusButton',
  component: PlusButton,
  argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({...props}) => {
	return (
		<PlusButton {...props} />
	)
};


export const Primary = Mocked.bind({})

Primary.argTypes = {
	color: toEnum(["primary", "secondary"]),
	disabled: toBoolean(),
};