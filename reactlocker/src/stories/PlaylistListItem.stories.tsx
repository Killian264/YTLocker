import { Story } from "@storybook/react";
import { Card, CardProps } from "../components/Card";
import { PlaylistListItem } from "../components/PlaylistListItem";
import { PlusButton } from "../components/PlusButton";
import { Playlist } from "../shared/types";

export default {
	title: "PlaylistListItem",
	component: PlaylistListItem,
};

const playlists: Playlist[] = [
	{
		id: 932423423,
		youtube: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
		thumbnail:
			"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
		title: "DogeLog",
		description: "Videos showing Ben Awad as he builds dogehouse.",
		url:
			"https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK",
		created: new Date(),
	},
];

const Mocked: Story<{}> = ({ ...props }) => {
	return (
		<Card>
			<div></div>
			<div>
				<PlaylistListItem playlist={playlists[0]}></PlaylistListItem>
				<PlaylistListItem
					playlist={playlists[0]}
					className="mt-3"
				></PlaylistListItem>
				<PlaylistListItem
					playlist={playlists[0]}
					className="mt-3"
				></PlaylistListItem>
			</div>
		</Card>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
