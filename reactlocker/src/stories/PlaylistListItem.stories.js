import React from 'react';
import { PlaylistListItem } from '../components/PlaylistListItem';
import { toEnum, toBoolean} from "./utils/utils"

export default {
  title: 'PlaylistListItem',
  component: PlaylistListItem,
  argTypes: { onClick: { action: "clicked" } },
};

const Mocked = ({
  children,
  ...props
}) => {
	return (
		<PlaylistListItem {...props} >
			{children}
		</PlaylistListItem>
	);
};


export const Primary = Mocked.bind({})

Primary.argTypes = {};