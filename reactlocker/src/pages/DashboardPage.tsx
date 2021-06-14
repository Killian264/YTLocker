import React, { useState } from "react";
import { UserInfoBarController } from "../controllers/UserInfoBarController";
import { usePlaylists } from "../hooks/api/usePlaylists";
import { VideoListLatestController } from "../controllers/VideosListLatestController";
import { ChannelListController } from "../controllers/ChannelListController";
import { usePlaylistChannels } from "../hooks/api/usePlaylistChannels";
import { PlaylistsListController } from "../controllers/PlaylistsListController";
import { usePlaylist } from "../hooks/api/usePlaylist";
import { PlaylistView } from "../components/PlaylistView";
import { PlaylistVideoListController } from "../controllers/PlaylistVideoListController";
import { PlaylistChannelListController } from "../controllers/PlaylistChannelsListController";

export const DashboardPage: React.FC<{}> = () => {
	const [id, setId] = useState(0);

	const playlistClick = (id: number) => {
		setId(id);
	};

	const playlistBack = () => {
		setId(0);
	};

	return (
		<div className="my-5">
			<div className="mb-4 px-4 mx-auto max-w-7xl">
				<UserInfoBarController></UserInfoBarController>
			</div>
			<div className="grid gap-4 px-4 mx-auto max-w-7xl grid-cols-12">
				{id === 0 && (
					<DashboardPlaylistsView onPlaylistClick={playlistClick}></DashboardPlaylistsView>
				)}
				{id !== 0 && <DashboardPlaylistView onBack={playlistBack} id={id}></DashboardPlaylistView>}
			</div>
		</div>
	);
};

export interface DashboardPlaylistsViewProps {
	className?: string;
	onPlaylistClick: (id: number) => void;
}

export const DashboardPlaylistsView: React.FC<DashboardPlaylistsViewProps> = ({ onPlaylistClick }) => {
	const [loadingP, playlists] = usePlaylists();
	const [loadingC, channels] = usePlaylistChannels();

	let limit = 5;
	if (!(loadingP && loadingC)) {
		limit = playlists.length + channels.length + 2;
	}

	return (
		<>
			<div className="grid gap-4 xl:col-span-7 col-span-12">
				<PlaylistsListController onPlaylistClick={onPlaylistClick}></PlaylistsListController>
				<ChannelListController
					className="xl:block hidden"
					limit={Number.MAX_VALUE}
				></ChannelListController>
			</div>
			<div className="xl:col-span-5 col-span-12">
				<VideoListLatestController limit={limit}></VideoListLatestController>
			</div>
			<div className="col-span-12 xl:hidden block">
				<ChannelListController limit={Number.MAX_VALUE}></ChannelListController>
			</div>
		</>
	);
};

export interface DashboardPlaylistViewProps {
	className?: string;
	id: number;
	onBack: () => void;
}
export const DashboardPlaylistView: React.FC<DashboardPlaylistViewProps> = ({ id, onBack }) => {
	const [loadingP, playlist] = usePlaylist(id);
	const [loadingC, channels] = usePlaylistChannels();

	if (loadingP || loadingC || playlist === null) {
		return <div>Loading...</div>;
	}

	return (
		<>
			<div className="grid gap-4 xl:col-span-7 col-span-12">
				<PlaylistView playlist={playlist} DeleteClick={() => {}} BackClick={onBack}></PlaylistView>
				<PlaylistChannelListController
					className="xl:block hidden"
					limit={Number.MAX_VALUE}
					playlist={playlist}
				></PlaylistChannelListController>
			</div>
			<div className="xl:col-span-5 col-span-12">
				{/* <VideoListLatestController limit={limit}></VideoListLatestController> */}
				<PlaylistVideoListController
					limit={channels.length + 4}
					playlist={playlist}
				></PlaylistVideoListController>
			</div>
			<div className="col-span-12 xl:hidden block">
				<PlaylistChannelListController
					limit={Number.MAX_VALUE}
					playlist={playlist}
				></PlaylistChannelListController>
			</div>
		</>
	);
};
