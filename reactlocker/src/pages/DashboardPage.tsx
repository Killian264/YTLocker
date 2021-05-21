import React from "react";
import { UserInfoBarController } from "../controllers/UserInfoBarController";
import { usePlaylists } from "../shared/api/usePlaylists";
import { VideoListLatestController } from "../controllers/VideosListLatestController";
import { ChannelListController } from "../controllers/ChannelListController";
import { usePlaylistChannels } from "../shared/api/usePlaylistChannels";
import { PlaylistsListController } from "../controllers/PlaylistsListController";

export const DashboardPage: React.FC<{}> = () => {
	const [loadingP, playlists] = usePlaylists();
	const [loadingC, channels] = usePlaylistChannels();

	let limit = 5;
	if (!(loadingP && loadingC)) {
		limit = playlists.length + channels.length + 2;
	}

	return (
		<div className="my-5">
			<div className="mb-4 px-4 mx-auto max-w-7xl">
				<UserInfoBarController></UserInfoBarController>
			</div>
			<div className="grid gap-4 px-4 mx-auto max-w-7xl grid-cols-12">
				<div className="grid gap-4 xl:col-span-7 col-span-12">
					<PlaylistsListController></PlaylistsListController>
					<ChannelListController limit={Number.MAX_VALUE}></ChannelListController>
				</div>
				<div className="xl:col-span-5 col-span-12">
					<VideoListLatestController
						className="xl:col-span-5 lg:col-span-6 col-span-12"
						limit={limit}
					></VideoListLatestController>
				</div>
				<div className="col-span-12 xl:hidden block">
					<ChannelListController limit={Number.MAX_VALUE}></ChannelListController>
				</div>
			</div>
		</div>
	);
};
