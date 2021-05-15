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
	youtube: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
	thumbnail:
		"https://i.ytimg.com/vi/1PBNAoKd-70/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLCFnLzV-VCKC28TFfjTi5cQL7zXiA",
	title: "DogeLog",
	description: "Videos showing Ben Awad as he builds dogehouse.",
	url: "https://www.youtube.com/playlist?list=PLN3n1USn4xlkZgqq9SdgUXPmgpoxUM9QK",
	created: new Date(),
	Videos: [],
	Channels: [],
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
