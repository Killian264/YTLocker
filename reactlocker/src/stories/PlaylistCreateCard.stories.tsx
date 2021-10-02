import { Story } from "@storybook/react";
import { PlaylistCreateCard } from "../components/PlaylistCreateCard";
import { Playlist } from "../shared/types";

export default {
	title: "PlaylistCreateCard",
	component: PlaylistCreateCard,
};

const Mocked: Story<{}> = () => {
	let editPlaylist: Playlist = {
		id: 932423423,
		youtubeId: "PLamdXAekZPYiqLDNQXQTbm4N_cPBmLPyr",
		thumbnailUrl:
			"https://yt3.ggpht.com/ytc/AAUvwngwzt6aURebMaKQpNkT9WpY3A3z0rtocMufWQwLxA=s800-c-k-c0x00ffffff-no-rj-mo",
		title: "DogeLog",
		description: "Videos showing Ben Awad as he builds dogehouse.",
		accountId: 12,
		channels: [],
		videos: [],
		color: "red-1",
		created: new Date(),
	};

	return (
		<div>
			<PlaylistCreateCard
				playlists={[]}
				accounts={[]}
				CreateClick={() => console.log("created")}
				BackClick={() => console.log("go back")}
			></PlaylistCreateCard>
			<PlaylistCreateCard
				className="mt-3"
				editPlaylist={editPlaylist}
				playlists={[]}
				accounts={[]}
				CreateClick={() => console.log("created")}
				BackClick={() => console.log("go back")}
			></PlaylistCreateCard>
		</div>
	);
};

export const Primary = Mocked.bind({});

Primary.argTypes = {};
