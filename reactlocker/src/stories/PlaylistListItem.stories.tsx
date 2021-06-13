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
		youtubeId: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
		thumbnailUrl:
			"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
		title: "DogeLog",
		description: "Videos showing Ben Awad as he builds dogehouse.",
		// url: "https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK",
		created: new Date(),
		channels: [],
		videos: [],
		color: "red-1",
	},
];

const Mocked: Story<{}> = ({ ...props }) => {
	let url = "https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK";

	let onClick = () => {};

	return (
		<Card>
			<div></div>
			<div>
				<PlaylistListItem url={url} playlist={playlists[0]} onClick={onClick}></PlaylistListItem>
				<PlaylistListItem
					url={url}
					playlist={playlists[0]}
					onClick={onClick}
					className="mt-3"
				></PlaylistListItem>
				<PlaylistListItem
					url={url}
					playlist={playlists[0]}
					onClick={onClick}
					className="mt-3"
				></PlaylistListItem>
			</div>
		</Card>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
