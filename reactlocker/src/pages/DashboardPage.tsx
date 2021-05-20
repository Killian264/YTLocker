import React, { useEffect, useState } from "react";
import { RouteComponentProps } from "react-router-dom";
import { Card } from "../components/Card";
import { PlaylistList } from "../components/ListCards";
import { UserInfoBar } from "../components/UserInfoBar";
import { VideoListItemController } from "../controllers/VideoListItemController";
import { ChannelListItemController } from "../controllers/ChannelListItemController";
import { API } from "../shared/api/api";
import { useBearer } from "../shared/hooks/useBearer";
import { Playlist, StatCard, User } from "../shared/types";

export const DashboardPage: React.FC<RouteComponentProps> = ({ history }) => {
	const [bearer] = useBearer("");
	const [playlists, setPlaylists] = useState<Playlist[]>([]);
	const [channels, setChannels] = useState<number[]>([]);
	const [stats, setStats] = useState<StatCard[]>([]);
	const [user, setUser] = useState<User | null>(null);

	if (bearer === "") {
		history.push("/login");
	}

	useEffect(() => {
		(async () => {
			let res = await API.UserInformation(bearer);

			if (!res.success) {
				history.push("/login");
			}

			setUser(res.user);
		})();
	}, [bearer, history]);

	useEffect(() => {
		(async () => {
			let res = await API.PlaylistList(bearer);

			if (!res.success) {
				history.push("/login");
			}

			let playlists = res.playlists;
			let channels2: number[] = [];

			playlists.forEach((playlist) => {
				channels2 = [...channels2, ...playlist.channels];
			});

			let stats = ParseStats(playlists);

			setPlaylists(playlists);
			setChannels(channels2);
			setStats(stats);
		})();
	}, [bearer, history]);

	return (
		<div>
			<div className="p-4 mx-auto max-w-7xl">
				{user !== null && <UserInfoBar className="flex-grow" user={user} stats={stats}></UserInfoBar>}
			</div>
			<div className="px-4 mx-auto max-w-7xl grid grid-cols-12 gap-4">
				<PlaylistList
					className="xl:col-span-7 lg:col-span-6 col-span-12"
					playlists={playlists}
					limit={5}
				></PlaylistList>
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

const ParseStats = (playlists: Playlist[]): StatCard[] => {
	let videoCount = 0;
	let channelCount = 0;

	playlists.forEach((playlist) => {
		videoCount += playlist.videos.length;
		channelCount += playlist.channels.length;
	});

	return [
		{
			header: "Playlists",
			count: playlists.length,
			measurement: "total",
		},
		{
			header: "Videos",
			count: videoCount,
			measurement: "total",
		},
		{
			header: "Subscriptions",
			count: channelCount,
			measurement: "total",
		},
		{
			header: "Updated",
			count: 12,
			measurement: "hours ago",
		},
	];
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
