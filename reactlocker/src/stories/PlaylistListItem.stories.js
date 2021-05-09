import React from 'react';
import { Card } from '../components/Card';
import { PlaylistListItem } from '../components/PlaylistListItem';
import { PlusButton } from '../components/PlusButton';
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
		<Card>
			<div className="flex justify-between">
				<div className="text-2xl">
					<span className="font-bold inline-block align-middle leading-none">Playlists</span>
				</div>
				<PlusButton></PlusButton>
			</div>
			<div>
				<PlaylistListItem></PlaylistListItem>
				<PlaylistListItem className="mt-2" ></PlaylistListItem>
				<PlaylistListItem className="mt-2" ></PlaylistListItem>
			</div>
		</Card>
		// <PlaylistListItem {...props} >
		// 	{children}
		// </PlaylistListItem>
	);
};


export const Primary = Mocked.bind({})

Primary.argTypes = {};