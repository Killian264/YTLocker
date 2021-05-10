import { Story } from "@storybook/react";
import { Card, CardProps } from "../components/Card";
import { PlaylistListItem } from "../components/PlaylistListItem";
import { PlusButton } from "../components/PlusButton";

export default {
	title: "PlaylistListItem",
	component: PlaylistListItem,
};

const Mocked: Story<{}> = ({ ...props }) => {
	return (
		<Card>
			<div className="flex justify-between -mb-1">
				<div className="text-2xl">
					<span className="font-bold leading-none -mt-0.5">
						Playlists
					</span>
				</div>
				<PlusButton color="primary" disabled={false}></PlusButton>
			</div>
			<div>
				<PlaylistListItem></PlaylistListItem>
				<PlaylistListItem className="mt-2"></PlaylistListItem>
				<PlaylistListItem className="mt-2"></PlaylistListItem>
			</div>
		</Card>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
