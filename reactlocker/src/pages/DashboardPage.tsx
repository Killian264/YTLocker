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
import { PlaylistCreateCard } from "../components/PlaylistCreateCard";
import { useCreatePlaylist } from "../hooks/api/useCreatePlaylist";

export const DashboardPage: React.FC<{}> = () => {
	const [playlistId, setPlaylistId] = useState(0);

	const PlaylistClick = (id: number) => {
		setPlaylistId(id);
	};

	const PlaylistBack = () => {
		setPlaylistId(0);
	};

	return (
		<div className="my-5">
			<div className="mb-4 px-4 mx-auto max-w-7xl">
				<UserInfoBarController></UserInfoBarController>
			</div>
			<div className="grid gap-4 px-4 mx-auto max-w-7xl grid-cols-12">
				{playlistId === 0 && (
					<DashboardPlaylistsView PlaylistClick={PlaylistClick}></DashboardPlaylistsView>
				)}
				{playlistId !== 0 && (
					<DashboardPlaylistView
						BackClick={PlaylistBack}
						playlistId={playlistId}
					></DashboardPlaylistView>
				)}
			</div>
		</div>
	);
};

export interface DashboardPlaylistsViewProps {
	className?: string;
	PlaylistClick: (id: number) => void;
}

export const DashboardPlaylistsView: React.FC<DashboardPlaylistsViewProps> = ({ PlaylistClick }) => {
	const [loadingP, playlists] = usePlaylists();
	const [loadingC, channels] = usePlaylistChannels();
	const [isCreate, setCreate] = useState(false);
	const createPlaylist = useCreatePlaylist();

	let limit = 5;
	if (!(loadingP && loadingC)) {
		limit = playlists.length + channels.length + 2;
	}

	let view = (
		<PlaylistCreateCard
			CreateClick={createPlaylist}
			BackClick={() => {
				setCreate(false);
			}}
		></PlaylistCreateCard>
	);

	if (!isCreate) {
		view = (
			<PlaylistsListController
				CreatePlaylistClick={() => {
					setCreate(true);
				}}
				PlaylistClick={PlaylistClick}
			></PlaylistsListController>
		);
	}

	return (
		<>
			<div className="grid gap-4 xl:col-span-7 col-span-12">
				{view}
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
	playlistId: number;
	BackClick: () => void;
}
export const DashboardPlaylistView: React.FC<DashboardPlaylistViewProps> = ({ playlistId, BackClick }) => {
	const [loadingP, playlist] = usePlaylist(playlistId);
	const [loadingC, channels] = usePlaylistChannels();

	if (loadingP || loadingC || playlist === null) {
		return <div>Loading...</div>;
	}

	let deletePlaylist = () => {
		console.log("delete playlist click", playlistId);
		BackClick();
	};

	return (
		<>
			<div className="grid gap-4 xl:col-span-7 col-span-12">
				<PlaylistView
					playlist={playlist}
					DeleteClick={deletePlaylist}
					BackClick={BackClick}
				></PlaylistView>
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
