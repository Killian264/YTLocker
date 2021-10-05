import { Story } from "@storybook/react";
import { PlaylistView, PlaylistViewProps } from "../components/PlaylistView";
import { Account, Playlist } from "../shared/types";
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
	accountId: 12,
	created: new Date(),
	videos: [],
	channels: [],
	isActive: true,
	color: "red-1",
};

const account: Account = {
	id: 12,
	username: "Killian's Account",
	email: "killian@ytlocker.com",
	picture: "google.com",
	permissionLevel: "manage",
};

const Mocked: Story<PlaylistViewProps> = ({ ...props }) => {
	return (
		<PlaylistView
			EditClick={() => {
				console.log("EDIT CLICKED");
			}}
			DeleteClick={() => {
				console.log("DELETE CLICKED");
			}}
			BackClick={() => {
				console.log("BACK CLICKED");
			}}
			PauseClick={() => {}}
			CopyClick={() => {}}
			RefreshClick={() => {}}
			playlist={playlist}
			account={account}
		></PlaylistView>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {
	className: sString(),
};
