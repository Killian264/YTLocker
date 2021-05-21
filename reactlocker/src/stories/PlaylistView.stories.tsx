import { Story } from "@storybook/react";
import { PlaylistView, PlaylistViewProps } from "../components/PlaylistView";
import { Playlist } from "../shared/types";
import { sString } from "./utils/utils";

export default {
	title: "PlaylistView",
	component: PlaylistView,
};

const playlist: Playlist = {
	id: 932423423,
	youtubeId: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
	thumbnailUrl:
		"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
	title: "DogeLog",
	description: "Videos showing Ben Awad as he builds dogehouse.",
	created: new Date(),
	videos: [],
	channels: [],
	color: "red-1",
};

const Mocked: Story<PlaylistViewProps> = ({ ...props }) => {
	return (
		<PlaylistView
			DeleteClick={() => {
				console.log("hello");
			}}
			BackClick={() => {
				console.log("hello");
			}}
			playlist={playlist}
		></PlaylistView>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
};
