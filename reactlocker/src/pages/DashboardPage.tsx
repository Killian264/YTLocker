import React, { useState } from "react";
import { UserInfoBarController } from "../controllers/UserInfoBarController";
import { usePlaylistList } from "../hooks/api/usePlaylistList";
import { VideoListLatestController } from "../controllers/VideosListLatestController";
import { ChannelListController } from "../controllers/ChannelListController";
import { usePlaylistChannels } from "../hooks/api/usePlaylistChannels";
import { PlaylistsListController } from "../controllers/PlaylistsListController";
import { usePlaylist } from "../hooks/api/usePlaylist";
import { PlaylistVideoListController } from "../controllers/PlaylistVideoListController";
import { PlaylistChannelListController } from "../controllers/PlaylistChannelsListController";
import { PlaylistCreateCard } from "../components/PlaylistCreateCard";
import { usePlaylistCreate } from "../hooks/api/usePlaylistCreate";
import { PlaylistViewController } from "../controllers/PlaylistViewController";

export const DashboardPage: React.FC<{}> = () => {
	const [playlistId, setPlaylistId] = useState<number | null>(null);

	let view = (
		<DashboardPlaylistListView
			PlaylistClick={(id) => {
				setPlaylistId(id);
			}}
		></DashboardPlaylistListView>
	);

	if (playlistId !== null) {
		view = (
			<PlaylistViewPage
				BackClick={() => {
					setPlaylistId(null);
				}}
				playlistId={playlistId}
			></PlaylistViewPage>
		);
	}

	return (
		<div className="my-5 mb-8">
			<div className="mb-4 px-4 mx-auto max-w-7xl">
				<UserInfoBarController></UserInfoBarController>
			</div>
			<div className="grid gap-4 px-4 mx-auto max-w-7xl grid-cols-12">{view}</div>
		</div>
	);
};

export interface DashboardPlaylistListViewProps {
	className?: string;
	PlaylistClick: (id: number) => void;
}

export const DashboardPlaylistListView: React.FC<DashboardPlaylistListViewProps> = ({ PlaylistClick }) => {
	const [isLoadingPlaylists, playlists] = usePlaylistList();
	const [isLoadingChannels, channels] = usePlaylistChannels();
	const [isCreate, setCreate] = useState(false);
	const createPlaylist = usePlaylistCreate();

	let limit = 5;
	if (!(isLoadingPlaylists && isLoadingChannels)) {
		limit = playlists.length + channels.length + 2;
	}

	let view = (
		<PlaylistsListController
			CreatePlaylistClick={() => {
				setCreate(true);
			}}
			PlaylistClick={PlaylistClick}
		></PlaylistsListController>
	);

	if (isCreate) {
		view = (
			<PlaylistCreateCard
				CreateClick={createPlaylist}
				BackClick={() => {
					setCreate(false);
				}}
				playlists={playlists}
			></PlaylistCreateCard>
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

export interface PlaylistViewPageProps {
	className?: string;
	playlistId: number;
	BackClick: () => void;
}
export const PlaylistViewPage: React.FC<PlaylistViewPageProps> = ({ playlistId, BackClick }) => {
	const [isLoadingPlaylist, playlist] = usePlaylist(playlistId);
	const [isLoadingChannels, channels] = usePlaylistChannels();

	if (isLoadingPlaylist || isLoadingChannels || playlist === null) {
		return <div>Loading...</div>;
	}

	return (
		<>
			<div className="grid gap-4 xl:col-span-7 col-span-12">
				<PlaylistViewController
					playlistId={playlistId}
					BackClick={BackClick}
				></PlaylistViewController>
				<PlaylistChannelListController
					className="xl:block hidden"
					limit={Number.MAX_VALUE}
					playlist={playlist}
				></PlaylistChannelListController>
			</div>
			<div className="xl:col-span-5 col-span-12">
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
