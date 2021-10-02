import { useState } from "react";
import { PlaylistCreateCard } from "../components/PlaylistCreateCard";
import { ChannelListController } from "../controllers/ChannelListController";
import { PlaylistChannelListController } from "../controllers/PlaylistChannelsListController";
import { PlaylistsListController } from "../controllers/PlaylistsListController";
import { PlaylistVideoListController } from "../controllers/PlaylistVideoListController";
import { PlaylistViewController } from "../controllers/PlaylistViewController";
import { VideoListLatestController } from "../controllers/VideosListLatestController";
import { useAccountList } from "../hooks/api/useAccountList";
import { usePlaylist } from "../hooks/api/usePlaylist";
import { usePlaylistChannels } from "../hooks/api/usePlaylistChannels";
import { usePlaylistCreate } from "../hooks/api/usePlaylistCreate";
import { usePlaylistList } from "../hooks/api/usePlaylistList";

export const DashboardPlaylistView: React.FC<{}> = () => {
	const [playlistId, setPlaylistId] = useState<number | null>(null);

	if (playlistId === null) {
		return (
			<div className="grid gap-4 grid-cols-12">
				<PlaylistListView
					PlaylistClick={(id) => {
						setPlaylistId(id);
					}}
				></PlaylistListView>
			</div>
		);
	}

	return (
		<div className="grid gap-4 grid-cols-12">
			<PlaylistViewPage
				BackClick={() => {
					setPlaylistId(null);
				}}
				playlistId={playlistId}
			></PlaylistViewPage>
		</div>
	);
};

export interface PlaylistListViewProps {
	className?: string;
	PlaylistClick: (id: number) => void;
}

export const PlaylistListView: React.FC<PlaylistListViewProps> = ({ PlaylistClick }) => {
	const [isLoadingPlaylists, playlists] = usePlaylistList();
	const [isLoadingAccounts, accounts] = useAccountList();
	const [isLoadingChannels, channels] = usePlaylistChannels();
	const [isCreate, setCreate] = useState(false);
	const createPlaylist = usePlaylistCreate();

	let limit = 5;
	if (!(isLoadingPlaylists && isLoadingChannels && isLoadingAccounts)) {
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
				accounts={accounts}
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
