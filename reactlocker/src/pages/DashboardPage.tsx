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
import { Plus, RightArrow, SvgBox } from "../components/Svg";
import { Card } from "../components/Card";
import { LinkButton } from "../components/LinkButton";
import { Footer } from "../components/Footer";
import { Logo } from "../components/Logo";

export const DashboardPage: React.FC<{}> = () => {
	const [viewPage, setViewPage] = useState<"playlists" | "accounts">("playlists");

	let view = <AccountsView></AccountsView>;

	if (viewPage === "playlists") {
		view = <PlaylistView></PlaylistView>;
	}

	return (
		<div className="h-screen flex flex-col flex-grow justify-between">
			<div className="mb-5 mt-6 pb-4">
				<div className="mb-4 px-4 mx-auto max-w-7xl">
					<UserInfoBarController></UserInfoBarController>
				</div>
				<div className="mb-4 px-4 mx-auto max-w-7xl">
					<div className="flex gap-6 mx-3">
						<LinkButton
							onClick={() => setViewPage("playlists")}
							selected={viewPage === "playlists"}
						>
							Playlists
						</LinkButton>
						<LinkButton
							onClick={() => setViewPage("accounts")}
							selected={viewPage === "accounts"}
						>
							Accounts
						</LinkButton>
					</div>
				</div>
				<div className="grid gap-4 px-4 mx-auto max-w-7xl grid-cols-12">{view}</div>
			</div>
			<Footer></Footer>
		</div>
	);
};

export const AccountsView: React.FC<{}> = () => {
	const [isLoadingPlaylists, playlists] = usePlaylistList();

	if (isLoadingPlaylists) {
		return <div>Loading...</div>;
	}

	const css = `hover:bg-primary-600 rounded-md flex justify-between cursor-pointer overflow-hidden`;
	const imageSize = "md:h-20 md:w-32 h-16 w-24";
	const textSize = "sm:text-md text-md";

	return (
		<div className="col-span-12 gap-4 flex flex-col">
			<Card className={"className"}>
				<div className="flex justify-between -mb-2 -mt-2">
					<div className="text-2xl font-semibold">
						<span className="leading-none -mt-0.5">Linked Accounts</span>
					</div>
					<SvgBox className={`border-green-500 p-0.5 opacity-50 cursor-default`} onClick={() => {}}>
						<Plus className="text-green-500" size={24}></Plus>
					</SvgBox>
				</div>
				<div>
					<div className={css} onClick={() => {}}>
						<div className="flex p-1 overflow-hidden">
							<div
								className={`rounded-lg flex-shrink-0 bg-black ${imageSize} flex justify-center items-center`}
							>
								<Logo></Logo>
							</div>
							<div className="pl-3 flex flex-col">
								<span className={`${textSize} font-semibold whitespace-nowrap`}>
									YTLocker Playlists |
									<span className="text-accent ml-2">{playlists.length}</span>/6
								</span>
								<span className="whitespace-nowrap whitespace-nowrap">
									FREE-tier playlists. Playlists limited.
									<br />
									New videos may take longer to be added.
								</span>
							</div>
						</div>
						<div className="mr-2 my-auto select-none">
							<RightArrow size={24}></RightArrow>
						</div>
					</div>
				</div>
			</Card>
		</div>
	);
};

export const PlaylistView: React.FC<{}> = () => {
	const [playlistId, setPlaylistId] = useState<number | null>(null);

	if (playlistId === null) {
		return (
			<PlaylistListView
				PlaylistClick={(id) => {
					setPlaylistId(id);
				}}
			></PlaylistListView>
		);
	}

	return (
		<PlaylistViewPage
			BackClick={() => {
				setPlaylistId(null);
			}}
			playlistId={playlistId}
		></PlaylistViewPage>
	);
};

export interface PlaylistListViewProps {
	className?: string;
	PlaylistClick: (id: number) => void;
}

export const PlaylistListView: React.FC<PlaylistListViewProps> = ({ PlaylistClick }) => {
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
