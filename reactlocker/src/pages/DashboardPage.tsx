import React, { useEffect, useState } from "react";
import { RouteComponentProps } from "react-router-dom";
import { Card } from "../components/Card";
import { ChannelListItem } from "../components/ChannelListItem";
import { ChannelsList, PlaylistList, VideoList } from "../components/ListCards";
import { PlaylistListItem } from "../components/PlaylistListItem";
import { PlusButton } from "../components/PlusButton";
import { UserInfoBar } from "../components/UserInfoBar";
import { VideoListItem } from "../components/VideosListItem";
import { API, PlaylistListResponse } from "../shared/api";
import { useBearer } from "../shared/hooks/useBearer";
import { Channel, Playlist, StatCard, User, Video } from "../shared/types";

export const DashboardPage: React.FC<RouteComponentProps> = ({ history }) => {
	const [bearer] = useBearer("");
	const [playlists, setPlaylists] = useState<Playlist[]>([]);
	const [channels, setChannels] = useState<Channel[]>([]);
	const [videos, setVideos] = useState<Video[]>([]);
	const [stats, setStats] = useState<StatCard[]>([]);
	const [user, setUser] = useState<User | null>(null);

	if (bearer === "") {
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
	}, [bearer, history]);

	useEffect(() => {
		API.UserInformation(bearer).then((res) => {
			if (!res.success) {
				history.push("/login");
			}

			setUser(res.user);
		});
	}, []);

	return (
		<>
			<div className="p-4 mx-auto max-w-7xl">
				{user !== null && <UserInfoBar className="flex-grow" user={user} stats={stats}></UserInfoBar>}
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
		return b.created.getTime() - a.created.getTime();
	});

	videos = videos.sort((a: Video, b: Video) => {
		return b.created.getTime() - a.created.getTime();
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
			count: today.getHours() - date.getHours(),
			measurement: "hours ago",
		},
	];
};
