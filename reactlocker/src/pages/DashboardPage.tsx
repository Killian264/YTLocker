import React, { useEffect, useState } from "react";
import { RouteComponentProps } from "react-router-dom";
import { Card } from "../components/Card";
import { VideoListItemController } from "../controllers/VideoListItemController";
import { ChannelListItemController } from "../controllers/ChannelListItemController";
import { UserInfoBarController } from "../controllers/UserInfoBarController";
import { usePlaylists } from "../shared/api/usePlaylists";
import { Playlist } from "../shared/types";
import { PlaylistListItem } from "../components/PlaylistListItem";
import { PlusButton } from "../components/PlusButton";
import { BuildPlaylistUrl } from "../shared/urls";

export const DashboardPage: React.FC<RouteComponentProps> = ({ history }) => {
	const [channels, setChannels] = useState<number[]>([]);

	const [playlistsLoading, playlists] = usePlaylists();

	useEffect(() => {
		let merged: number[] = [];

		playlists.forEach((playlist) => {
			merged = [...merged, ...playlist.channels];
		});

		setChannels(merged);
	}, [playlists]);

	return (
		<div>
			<div className="p-4 mx-auto max-w-7xl">
				<UserInfoBarController></UserInfoBarController>
			</div>
			<div className="px-4 mx-auto max-w-7xl grid grid-cols-12 gap-4">
				{!playlistsLoading && playlists != null && (
					<PlaylistList
						className="xl:col-span-7 lg:col-span-6 col-span-12"
						playlists={playlists}
						limit={5}
					></PlaylistList>
				)}
				{playlists.length == 2 && (
					<VideoList
						className="xl:col-span-5 lg:col-span-6 col-span-12"
						videos={playlists[1].videos}
						limit={5}
					></VideoList>
				)}
				<ChannelsList
					className="col-span-12"
					channels={channels}
					limit={Number.MAX_VALUE}
				></ChannelsList>
			</div>
			<div className="px-4 mx-auto max-w-7xl m-3"></div>
		</div>
	);
};

interface VideoListProps {
	className?: string;
	videos: number[];
	limit: number;
}

export const VideoList: React.FC<VideoListProps> = ({ className, videos, limit }) => {
	let list = videos.map((id, index) => {
		if (index >= limit) {
			return "";
		}

		return <VideoListItemController key={id} videoId={id}></VideoListItemController>;
	});

	return (
		<Card className={className}>
			<div className="flex justify-between -mb-1 -mt-1">
				<div className="text-2xl font-semibold">
					<span className="leading-none -mt-0.5">Videos</span>
				</div>
			</div>
			<div className="grid gap-2">{list}</div>
		</Card>
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

		return <ChannelListItemController key={id} channelId={id}></ChannelListItemController>;
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
