import React, { useEffect, useState } from "react";
import { RouteComponentProps } from "react-router-dom";
import { Card } from "../components/Card";
import { ChannelListItemController } from "../controllers/ChannelListItemController";
import { UserInfoBarController } from "../controllers/UserInfoBarController";
import { usePlaylists } from "../shared/api/usePlaylists";
import { Playlist } from "../shared/types";
import { PlaylistListItem } from "../components/PlaylistListItem";
import { PlusButton } from "../components/PlusButton";
import { BuildPlaylistUrl } from "../shared/urls";
import { VideoListLatestController } from "../controllers/VideosListLatestController";

export const DashboardPage: React.FC<RouteComponentProps> = ({ history }) => {
	const [playlistsLoading, playlists] = usePlaylists();
	const [channels, setChannels] = useState<number[]>([]);

	useEffect(() => {
		let merged: number[] = [];

		playlists.forEach((playlist) => {
			merged = [...merged, ...playlist.channels];
		});

		setChannels(merged);
	}, [playlists]);

	let videoLimit = 5;

	if (!playlistsLoading && playlists !== null) {
		videoLimit = playlists.length + channels.length + 2;
	}

	return (
		<div className="mx-5">
			<div className="p-4 mx-auto max-w-7xl">
				<UserInfoBarController></UserInfoBarController>
			</div>
			<div className="grid gap-4 px-4 mx-auto max-w-7xl grid-cols-12">
				<div className="grid gap-4 xl:col-span-7 col-span-12">
					{!playlistsLoading && playlists !== null && (
						<PlaylistList playlists={playlists} limit={5}></PlaylistList>
					)}
					<ChannelsList
						className="xl:block hidden"
						channels={channels}
						limit={Number.MAX_VALUE}
					></ChannelsList>
				</div>
				<div className="xl:col-span-5 col-span-12">
					{playlists.length === 2 && (
						<VideoListLatestController
							className="xl:col-span-5 lg:col-span-6 col-span-12"
							limit={videoLimit}
						></VideoListLatestController>
					)}
				</div>
				<div className="col-span-12 xl:hidden block">
					<ChannelsList channels={channels} limit={Number.MAX_VALUE}></ChannelsList>
				</div>
			</div>
		</div>
	);
};

interface ChannelsListProp {
	className?: string;
	channels: number[];
	limit: number;
}

export const ChannelsList: React.FC<ChannelsListProp> = ({ className, channels, limit }) => {
	let list = channels.map((id, index) => {
		if (index >= limit) {
			return "";
		}

		return <ChannelListItemController key={index} channelId={id}></ChannelListItemController>;
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Channels</span>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};

interface PlaylistListProps {
	className?: string;
	playlists: Playlist[];
	limit: number;
}

export const PlaylistList: React.FC<PlaylistListProps> = ({ className, playlists, limit }) => {
	let list = playlists.map((playlist, index) => {
		if (index >= limit) {
			return "";
		}

		return (
			<PlaylistListItem
				url={BuildPlaylistUrl(playlist.youtubeId)}
				key={index}
				playlist={playlist}
			></PlaylistListItem>
		);
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Playlists</span>
				</div>
				<PlusButton color="secondary" disabled={false}></PlusButton>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
	);
};
