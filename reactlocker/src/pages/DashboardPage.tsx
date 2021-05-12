import React, { useEffect, useState } from "react";
import { RouteComponentProps } from "react-router-dom";
import { Card } from "../components/Card";
import { ChannelListItem } from "../components/ChannelListItem";
import { PlaylistListItem } from "../components/PlaylistListItem";
import { PlusButton } from "../components/PlusButton";
import { UserInfoBar } from "../components/UserInfoBar";
import { VideoListItem } from "../components/VideosListItem";
import { API, PlaylistListResponse } from "../shared/api";
import { useBearer } from "../shared/hooks/useBearer";
import { Channel, Playlist, StatCard, Video } from "../shared/types";

const user = {
	username: "Killian",
	email: "killiandebacker@gmail.com",
	joined: "Mar 13 2021",
};

const stats: StatCard[] = [
	{
		header: "Playlists",
		count: 454,
		measurement: "total",
	},
	{
		header: "Videos",
		count: 357,
		measurement: "total",
	},
	{
		header: "Subscriptions",
		count: 17,
		measurement: "total",
	},
	{
		header: "Updated",
		count: 13,
		measurement: "seconds ago",
	},
];

export const DashboardPage: React.FC<RouteComponentProps> = ({ history }) => {
	const [bearer, setBearer] = useBearer("");
	const [playlists, setPlaylists] = useState<Playlist[]>([]);
	const [channels, setChannels] = useState<Channel[]>([]);
	const [videos, setVideos] = useState<Video[]>([]);
	const [stats, setStats] = useState<StatCard[]>([]);

	if (bearer == "") {
		history.push("/login");
	}

	useEffect(() => {
		API.PlaylistList(bearer).then((res) => {
			if (!res.success) {
				history.push("/login");
			}

			let [playlistsList, channelsList, videosList] = ParsePlaylistListIntoLists(res);
			let stats = ParseStats(playlistsList, channelsList, videosList);

			setPlaylists(playlistsList);
			setChannels(channelsList);
			setVideos(videosList);
			setStats(stats);
		});
	}, []);

	return (
		<>
			<div className="p-4 mx-auto max-w-7xl">
				<UserInfoBar className="flex-grow" user={user} stats={stats}></UserInfoBar>
			</div>
			<div className="px-4 mx-auto max-w-7xl grid grid-cols-12 gap-4">
				<PlaylistList
					className="xl:col-span-7 lg:col-span-6 col-span-12"
					playlists={playlists}
				></PlaylistList>
				<VideoList className="xl:col-span-5 lg:col-span-6 col-span-12" videos={videos}></VideoList>
				<ChannelsList className="col-span-12" channels={channels}></ChannelsList>
			</div>
			<div className="px-4 mx-auto max-w-7xl m-3"></div>
		</>
	);
};

const ParsePlaylistListIntoLists = (res: PlaylistListResponse): [Playlist[], Channel[], Video[]] => {
	let items = res.items;

	let channels: Channel[] = [];
	let videos: Video[] = [];

	items.forEach((playlist) => {
		channels.push(...playlist.Channels);
	});

	items.forEach((playlist) => {
		videos.push(...playlist.Videos);
	});

	channels = channels.sort((a: Channel, b: Channel) => {
		return a.created.getTime() - b.created.getTime();
	});

	videos = videos.sort((a: Video, b: Video) => {
		return a.created.getTime() - b.created.getTime();
	});

	return [items, channels, videos];
};

const ParseStats = (playlists: Playlist[], channels: Channel[], videos: Video[]): StatCard[] => {
	let today = new Date();

	let date = videos.length > 0 ? videos[0].created : today;

	return [
		{
			header: "Playlists",
			count: playlists.length,
			measurement: "total",
		},
		{
			header: "Videos",
			count: videos.length,
			measurement: "total",
		},
		{
			header: "Subscriptions",
			count: channels.length,
			measurement: "total",
		},
		{
			header: "Updated",
			count: date.getHours() - today.getHours(),
			measurement: "hours ago",
		},
	];
};

interface PlaylistListProps {
	className?: string;
	playlists: Playlist[];
}

export const PlaylistList: React.FC<PlaylistListProps> = ({ className, playlists }) => {
	let list = playlists.map((playlist, index) => {
		if (index >= 5) {
			return;
		}

		return <PlaylistListItem key={index} playlist={playlist}></PlaylistListItem>;
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

interface VideoListProps {
	className?: string;
	videos: Video[];
}

export const VideoList: React.FC<VideoListProps> = ({ className, videos }) => {
	let list = videos.map((video, index) => {
		if (index >= 5) {
			return;
		}

		return <VideoListItem key={index} video={video}></VideoListItem>;
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
	channels: Channel[];
}

export const ChannelsList: React.FC<ChannelsListProp> = ({ className, channels }) => {
	let list = channels.map((channel, index) => {
		if (index >= 5) {
			return;
		}

		return <ChannelListItem key={index} channel={channel}></ChannelListItem>;
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

DashboardPage.propTypes = {};
